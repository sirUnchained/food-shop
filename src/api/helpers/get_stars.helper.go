package helpers

import (
	"foodshop/data/models"
	"foodshop/data/postgres"
)

func GetCalcRestaurantStars(restaurantID int) int {
	db := postgres.GetDb()
	var result struct {
		Count int
		Sum   int
	}

	db.Model(&models.Orders{}).Where("restaurant_id = ? AND stars != 0", restaurantID).Select("COUNT(*) as count, SUM(stars) as sum").Scan(&result)

	return result.Sum / result.Count
}
