package services

import (
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderService struct{}

func GetOrderService() *OrderService {
	return &OrderService{}
}

func (rc *OrderService) GetAllAdmin(ctx *gin.Context) *helpers.ResultResponse {
	var err error

	// Get the limit query parameter, default to 10 if not provided or invalid
	limit_str := ctx.Query("limit")
	var limit int
	if limit, err = strconv.Atoi(limit_str); err != nil {
		limit = 10
	}

	// Get the page query parameter, default to 1 if not provided or invalid
	page_str := ctx.Query("page")
	var page int
	if page, err = strconv.Atoi(page_str); err != nil {
		page = 1
	}

	// calculate the offset for pagination
	offset := (page - 1) * limit

	var orders []models.Orders
	db := postgres.GetDb()
	// user, _ := ctx.Get("user")

	// fetch orders from the database with pagination
	if err := db.Model(&models.Orders{}).Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "Orders fetched successfully.", Data: orders}
}

func (rc *OrderService) GetAllUser(ctx *gin.Context) *helpers.ResultResponse {
	var err error
	var orders []models.Orders
	user, _ := ctx.Get("user")
	db := postgres.GetDb()

	// Get the limit query parameter, default to 10 if not provided or invalid
	limit_str := ctx.Query("limit")
	var limit int
	if limit, err = strconv.Atoi(limit_str); err != nil {
		limit = 10
	}

	// Get the page query parameter, default to 1 if not provided or invalid
	page_str := ctx.Query("page")
	var page int
	if page, err = strconv.Atoi(page_str); err != nil {
		page = 1
	}

	// calculate the offset for pagination
	offset := (page - 1) * limit

	// fetch orders from the database with pagination for the specific user
	if err := db.Model(&models.Orders{}).Where("user_id = ?", user.(models.Users).ID).Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "User orders fetched successfully.", Data: orders}
}

func (rc *OrderService) GetOne(ctx *gin.Context) *helpers.ResultResponse {
	var orderID int
	var err error

	orderID_str := ctx.Param("id")
	if orderID, err = strconv.Atoi(orderID_str); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "order not found.", Data: nil}
	}

	var order models.Orders
	user, _ := ctx.Get("user")
	db := postgres.GetDb()

	db.Model(&models.Orders{}).Where("ID = ? AND User = ?", orderID, user.(models.Users).ID).First(&order)
	if order.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "order not found.", Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "here you are.", Data: order}

}

func (rc *OrderService) Create(ctx *gin.Context) *helpers.ResultResponse {
	orderDTO := new(dto.OrderDTO)
	err := ctx.ShouldBindBodyWithJSON(&orderDTO)
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: err.Error(), Data: nil}
	}

	user, _ := ctx.Get("user")

	newOrder := new(models.Orders)
	newOrder.User = user.(models.Users).ID
	newOrder.Restaurant = orderDTO.Restaurant
	newOrder.Address = orderDTO.Address
	newOrder.PostalCode = orderDTO.PostalCode

	db := postgres.GetDb()
	if err := db.Create(newOrder).Error; err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 201, Message: "Order created successfully.", Data: newOrder}
}

func (rc *OrderService) DeliveredStatus(ctx *gin.Context) *helpers.ResultResponse {
	var orderID int
	var err error

	orderID_str := ctx.Param("id")
	if orderID, err = strconv.Atoi(orderID_str); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "order not found.", Data: nil}
	}

	user, _ := ctx.Get("user")
	db := postgres.GetDb()

	var order models.Orders
	db.Model(&models.Orders{}).Where("ID = ?", orderID).First(&order)
	if order.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "order not found.", Data: nil}
	}

	var restaurant models.Restaurants
	db.Model(&models.Restaurants{}).Where("Owner = ?", user.(models.Users).ID).First(&restaurant)
	if restaurant.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: "you should not be here ! somthing went wrong.", Data: nil}
	}

	if restaurant.ID != order.Restaurant {
		return &helpers.ResultResponse{Ok: false, Status: 403, Message: "access denied. you cannot change other restaurant order data.", Data: nil}
	}

	order.IsDelivered = true
	db.Save(order)

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "order delivered status changed !", Data: nil}
}

func (rc *OrderService) AddStars(ctx *gin.Context) *helpers.ResultResponse {}

func (rc *OrderService) Remove(ctx *gin.Context) *helpers.ResultResponse {}
