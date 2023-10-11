package app

import (
	"log"
	"math"
	"net/http"
	"sort"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/database"
	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/database/models"
	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type BrandResult struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Product       string  `json:"product"`
	Category      string  `json:"category"`
	Location      string  `json:"location"`
	ShopifyRating float32 `json:"shopify_rating"`
	MatchScore    float32 `json:"match_score"`
}

func SearchHandler(c *gin.Context) {
	q := c.Query("q")
	cleanQuery := utils.RemoveNonAlphanumeric(q)
	if cleanQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing query parameter or empty/invalid query",
		})
		return
	}

	uniqueWords := getWords(cleanQuery)
	locations := findMatches(uniqueWords, models.AllBrandLocations)
	categories := findMatches(uniqueWords, models.AllBrandCategories)
	unmatchedTokens := getUnmatched(uniqueWords, append(locations, categories...))

	tokens := tokenize(strings.Join(unmatchedTokens, " "))

	log.Println("query:", q)
	log.Println("cleanQuery:", cleanQuery)
	log.Println("uniqueWords:", uniqueWords)
	log.Println("locations:", locations)
	log.Println("categories:", categories)
	log.Println("unmatchedTokens:", unmatchedTokens)
	log.Println("searchTokens:", tokens)

	records, err := search(tokens, categories, locations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	results := getBrandResults(records, uniqueWords)

	c.JSON(http.StatusOK, gin.H{
		"count":   len(results),
		"results": results,
	})
}

func tokenize(q string) []string {
	clean := stopwords.CleanString(q, "en", true)
	words := strings.Fields(clean)
	stems := make([]string, 0, len(words))

	for _, word := range words {
		stem := word
		stems = append(stems, stem)
	}

	return stems
}

func getWords(s string) []string {
	uniqueParts := make(map[string]bool)
	words := make([]string, 0)

	parts := strings.Fields(s)
	for _, part := range parts {
		_, ok := uniqueParts[part]
		if !ok {
			uniqueParts[part] = true
			words = append(words, part)
		}
	}

	return words
}

func search(
	tokens []string,
	categories []string,
	locations []string,
) ([]*models.Brand, error) {
	db := database.Connect()
	var brands []*models.Brand

	doSearch := false

	statement := db

	if len(categories) > 0 {
		statement = statement.Where("category IN ?", categories)
		doSearch = true
	}

	if len(locations) > 0 {
		statement = statement.Where("location IN ?", locations)
		doSearch = true
	}

	if len(tokens) > 0 {
		similarityTreshold := float32(0.9)
		for _, token := range tokens {
			statement = statement.Where(
				statement.
					Where("word_similarity(?, product) > ?", token, similarityTreshold).
					Or("word_similarity(?, name) > ?", token, similarityTreshold),
			)
		}

		doSearch = true
	}

	if !doSearch {
		return []*models.Brand{}, nil
	}

	if result := statement.Find(&brands); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []*models.Brand{}, nil
		}

		return nil, result.Error
	}

	return brands, nil
}

func getBrandResults(brands []*models.Brand, tokens []string) []BrandResult {
	results := make([]BrandResult, 0, len(brands))

	for _, brand := range brands {
		if brand == nil {
			continue
		}

		result := BrandResult{
			ID:            brand.ID,
			Name:          brand.Name,
			Product:       brand.Product,
			Category:      string(brand.Category),
			Location:      string(brand.Location),
			ShopifyRating: brand.ShopifyRating,
			MatchScore:    getMatchScore(*brand, tokens),
		}

		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].MatchScore > results[j].MatchScore
	})

	return results
}

// Arbitrary match scoring for token occurrence
func getMatchScore(brand models.Brand, tokens []string) float32 {
	var score float32
	var categoryMatches float32
	var locationMatches float32

	categoryTokens := strings.Fields(string(brand.Category))
	locationTokens := strings.Fields(string(brand.Location))

	perfectScore := float32(len(tokens))

	for _, token := range tokens {
		if strings.Contains(brand.Name, token) {
			score += 0.03 * perfectScore
		}

		if strings.Contains(brand.Product, token) {
			score += 0.03 * perfectScore
		}

		for _, categoryToken := range categoryTokens {
			if categoryToken == token {
				categoryMatches++
			}
		}

		for _, locationToken := range locationTokens {
			if locationToken == token {
				locationMatches++
			}
		}

	}

	locationScore := (locationMatches / float32(len(locationTokens))) * 0.47 * perfectScore
	categoryScore := (categoryMatches / float32(len(categoryTokens))) * 0.47 * perfectScore

	score += (locationScore + categoryScore)
	score = score / perfectScore
	if math.IsNaN(float64(score)) {
		score = 0
	}

	return score
}

func findMatches[T ~string](tokens []string, params []T) []string {
	foundMatches := make([]string, 0)

	for _, param := range params {
		paramTokens := strings.Fields(string(param))

		missingParamToken := false
		for _, paramToken := range paramTokens {
			paramTokenIndex := slices.IndexFunc(tokens,
				func(token string) bool {
					return token == paramToken
				},
			)
			if paramTokenIndex == -1 {
				missingParamToken = true
				break
			}
		}

		if !missingParamToken {
			foundMatches = append(foundMatches, string(param))
		}
	}

	return foundMatches
}

func getUnmatched(tokens []string, matched []string) []string {
	unmatchedTokens := make([]string, 0)
	matchedWords := strings.Fields(strings.Join(matched, " "))

	for _, token := range tokens {
		foundIndex := slices.IndexFunc(matchedWords, func(matchedWord string) bool {
			return matchedWord == token
		})

		if foundIndex == -1 {
			unmatchedTokens = append(unmatchedTokens, token)
		}
	}

	return unmatchedTokens
}
