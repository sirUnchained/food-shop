package models

type FoodRestaurant struct {
	ID           uint `gorm:"primaryKey"`
	RestaurantID uint
	FoodID       uint
	Restaurant   Restaurants `gorm:"foreignKey:RestaurantID"`
	Food         Foods       `gorm:"foreignKey:FoodID"`
}
