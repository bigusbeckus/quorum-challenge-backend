package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BrandCategory string
type BrandLocation string

const (
	BrandCategoryAppliances BrandCategory = "appliances"
	BrandCategoryCoffee     BrandCategory = "coffee"
	BrandCategoryFood       BrandCategory = "food"
	BrandCategoryToys       BrandCategory = "toys"
	BrandCategoryFurniture  BrandCategory = "furniture"
)

const (
	BrandLocationLosAngeles BrandLocation = "los angeles"
	BrandLocationNewYork    BrandLocation = "new york"
	BrandLocationSeattle    BrandLocation = "seattle"
	BrandLocationDallas     BrandLocation = "dallas"
	BrandLocationMiami      BrandLocation = "miami"
)

type Brand struct {
	gorm.Model
	ID            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name          string `gorm:"unique"`
	Product       string
	Category      BrandCategory
	Location      BrandLocation
	ShopifyRating float32
}

func SeedBrands(tx *gorm.DB) error {
	if result := tx.
		Clauses(
			clause.OnConflict{
				DoNothing: true,
			},
		).
		Create(&brands); result.Error != nil {
		return result.Error
	}

	return nil

}

var brands = []Brand{
	{
		Name:          "Samsung",
		Product:       "Refrigerator",
		Category:      BrandCategoryAppliances,
		Location:      BrandLocationLosAngeles,
		ShopifyRating: 4.3,
	},
	{
		Name:          "Starbucks",
		Product:       "Coffee Beans",
		Category:      BrandCategoryCoffee,
		Location:      BrandLocationNewYork,
		ShopifyRating: 4.7,
	},
	{
		Name:          "Kraft",
		Product:       "Macaroni and Cheese",
		Category:      BrandCategoryFood,
		Location:      BrandLocationSeattle,
		ShopifyRating: 3.8,
	},
	{
		Name:          "LEGO",
		Product:       "Building Blocks",
		Category:      BrandCategoryToys,
		Location:      BrandLocationDallas,
		ShopifyRating: 4.5,
	},
	{
		Name:          "IKEA",
		Product:       "Sofa",
		Category:      BrandCategoryFurniture,
		Location:      BrandLocationMiami,
		ShopifyRating: 4.1,
	},
	{
		Name:          "Whirlpool",
		Product:       "Washing Machine",
		Category:      BrandCategoryAppliances,
		Location:      BrandLocationLosAngeles,
		ShopifyRating: 4.0,
	},
	{
		Name:          "Dunkin' Donuts",
		Product:       "Coffee",
		Category:      BrandCategoryCoffee,
		Location:      BrandLocationNewYork,
		ShopifyRating: 4.6,
	},
	{
		Name:          "Kellogg's",
		Product:       "Cereal",
		Category:      BrandCategoryFood,
		Location:      BrandLocationSeattle,
		ShopifyRating: 3.9,
	},
	{
		Name:          "Hasbro",
		Product:       "Action Figures",
		Category:      BrandCategoryToys,
		Location:      BrandLocationDallas,
		ShopifyRating: 4.2,
	},
	{
		Name:          "Ashley Furniture",
		Product:       "Dining Table",
		Category:      BrandCategoryFurniture,
		Location:      BrandLocationMiami,
		ShopifyRating: 4.0,
	},
	{
		Name:          "LG",
		Product:       "TV",
		Category:      BrandCategoryAppliances,
		Location:      BrandLocationLosAngeles,
		ShopifyRating: 4.4,
	},
	{
		Name:          "Peet's Coffee",
		Product:       "Coffee Beans",
		Category:      BrandCategoryCoffee,
		Location:      BrandLocationNewYork,
		ShopifyRating: 4.7,
	},
	{
		Name:          "Nestlé",
		Product:       "Chocolate",
		Category:      BrandCategoryFood,
		Location:      BrandLocationSeattle,
		ShopifyRating: 4.0,
	},
	{
		Name:          "Mattel",
		Product:       "Dolls",
		Category:      BrandCategoryToys,
		Location:      BrandLocationDallas,
		ShopifyRating: 4.3,
	},
	{
		Name:          "Rooms To Go",
		Product:       "Bed",
		Category:      BrandCategoryFurniture,
		Location:      BrandLocationMiami,
		ShopifyRating: 3.8,
	},
	{
		Name:          "Bosch",
		Product:       "Dishwasher",
		Category:      BrandCategoryAppliances,
		Location:      BrandLocationLosAngeles,
		ShopifyRating: 4.1,
	},
	{
		Name:          "Caribou Coffee",
		Product:       "Coffee",
		Category:      BrandCategoryCoffee,
		Location:      BrandLocationNewYork,
		ShopifyRating: 4.6,
	},
	{
		Name:          "Campbell's",
		Product:       "Soup",
		Category:      BrandCategoryFood,
		Location:      BrandLocationSeattle,
		ShopifyRating: 3.9,
	},
	{
		Name:          "Disney",
		Product:       "Plush Toys",
		Category:      BrandCategoryToys,
		Location:      BrandLocationDallas,
		ShopifyRating: 4.5,
	},
	{
		Name:          "Crate & Barrel",
		Product:       "Bookshelf",
		Category:      BrandCategoryFurniture,
		Location:      BrandLocationMiami,
		ShopifyRating: 4.2,
	},
}
