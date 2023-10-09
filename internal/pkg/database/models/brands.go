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
	BrandLocationLosAngeles BrandLocation = "los_angeles"
	BrandLocationNewYork    BrandLocation = "new_york"
	BrandLocationSeattle    BrandLocation = "seattle"
	BrandLocationDallas     BrandLocation = "dallas"
	BrandLocationMiami      BrandLocation = "miami"
)

type Brand struct {
	gorm.Model
	ID       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name     string `gorm:"unique"`
	Product  string
	Category BrandCategory
	Location BrandLocation
}

func SeedBrands(tx *gorm.DB) error {
	brands := []Brand{
		{Name: "Samsung", Product: "Refrigerator", Category: BrandCategoryAppliances, Location: BrandLocationLosAngeles},
		{Name: "Starbucks", Product: "Coffee Beans", Category: BrandCategoryCoffee, Location: BrandLocationNewYork},
		{Name: "Kraft", Product: "Macaroni and Cheese", Category: BrandCategoryFood, Location: BrandLocationSeattle},
		{Name: "LEGO", Product: "Building Blocks", Category: BrandCategoryToys, Location: BrandLocationDallas},
		{Name: "IKEA", Product: "Sofa", Category: BrandCategoryFurniture, Location: BrandLocationMiami},
		{Name: "Whirlpool", Product: "Washing Machine", Category: BrandCategoryAppliances, Location: BrandLocationLosAngeles},
		{Name: "Dunkin' Donuts", Product: "Coffee", Category: BrandCategoryCoffee, Location: BrandLocationNewYork},
		{Name: "Kellogg's", Product: "Cereal", Category: BrandCategoryFood, Location: BrandLocationSeattle},
		{Name: "Hasbro", Product: "Action Figures", Category: BrandCategoryToys, Location: BrandLocationDallas},
		{Name: "Ashley Furniture", Product: "Dining Table", Category: BrandCategoryFurniture, Location: BrandLocationMiami},
		{Name: "LG", Product: "TV", Category: BrandCategoryAppliances, Location: BrandLocationLosAngeles},
		{Name: "Peet's Coffee", Product: "Coffee Beans", Category: BrandCategoryCoffee, Location: BrandLocationNewYork},
		{Name: "Nestlé", Product: "Chocolate", Category: BrandCategoryFood, Location: BrandLocationSeattle},
		{Name: "Mattel", Product: "Dolls", Category: BrandCategoryToys, Location: BrandLocationDallas},
		{Name: "Rooms To Go", Product: "Bed", Category: BrandCategoryFurniture, Location: BrandLocationMiami},
		{Name: "Bosch", Product: "Dishwasher", Category: BrandCategoryAppliances, Location: BrandLocationLosAngeles},
		{Name: "Caribou Coffee", Product: "Coffee", Category: BrandCategoryCoffee, Location: BrandLocationNewYork},
		{Name: "Campbell's", Product: "Soup", Category: BrandCategoryFood, Location: BrandLocationSeattle},
		{Name: "Disney", Product: "Plush Toys", Category: BrandCategoryToys, Location: BrandLocationDallas},
		{Name: "Crate & Barrel", Product: "Bookshelf", Category: BrandCategoryFurniture, Location: BrandLocationMiami},
	}

	if result := tx.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		},
	).
		Create(&brands); result.Error != nil {
		return result.Error
	}

	return nil

}
