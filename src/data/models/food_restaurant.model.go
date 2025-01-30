package models

type FoodRestaurant struct {
	ID           uint `gorm:"primaryKey"`
	Restaurant   Restaurants
	Food         Foods
	RestaurantID uint `gorm:"foreignKey:Restaurant"`
	FoodID       uint `gorm:"foreignKey:Food"`
}
