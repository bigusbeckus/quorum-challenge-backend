package app

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/database"
	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/database/models"
	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/utils"
	"github.com/gin-gonic/gin"
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

	tokens := tokenizeQuery(cleanQuery)
	results, err := search(tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":   len(results),
		"results": results,
	})
}

func tokenizeQuery(q string) []string {
	lowerQuery := strings.ToLower(q)
	words := getWords(lowerQuery)

	return words
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

func search(tokens []string) ([]BrandResult, error) {
	db := database.Connect()
	var brands []*models.Brand

	query := strings.Join(tokens, "|")
	query = fmt.Sprintf("(%s)", query)

	if result := db.
		Where("category ~* ?", query).
		Or("location ~* ?", query).
		Or("product ~* ?", query).
		Or("name ~* ?", query).
		Find(&brands); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []BrandResult{}, nil
		}

		return nil, result.Error
	}

	results := getBrandResults(brands, tokens)
	return results, nil
}

func getBrandResults(brands []*models.Brand, tokens []string) []BrandResult {
	results := make([]BrandResult, 0, len(brands))
	matchScoreTreshold := float32(0.51)

	for _, brand := range brands {
		if brand == nil {
			continue
		}

		matchScore := getMatchScore(*brand, tokens)
		if matchScore < matchScoreTreshold {
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
				score += (0.47 * perfectScore) / float32(len(categoryTokens))
			}
		}

		for _, locationToken := range locationTokens {
			if locationToken == token {
				score += (0.47 * perfectScore) / float32(len(locationTokens))
			}
		}

	}

	score = score / perfectScore

	return score
}
