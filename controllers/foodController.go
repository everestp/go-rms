package controllers

import (
	"context"
	"fmt"
	"go-rms/database"
	"go-rms/models"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

// var validate = validator.New()
// GetFoods - Fetch all food items
func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))

		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))
matchStage := bson.D{
    {"$match", bson.D{}} // You can add filters here if needed
}

groupStage := bson.D{
    {"$group", bson.D{
        {"_id", nil},
        {"total_count", bson.D{{"$sum", 1}}},
        {"data", bson.D{{"$push", "$$ROOT"}}}, // push each full document
    }},
}

projectStage := bson.D{
    {"$project", bson.D{
        {"_id", 0},
        {"total_count", 1},
        {"food_items", bson.D{
            {"$slice", bson.A{"$data", startIndex, recordPerPage}},
        }},
    }},
}

 result ,err := foodCollection.Aggregate(ctx ,mongo.Pipeline{
matchStage ,groupStage ,projectStage
})

defer cancel()
if err != nil {
	c.JSON{http.StatusInternalServerError ,gin.H{"error":"error ocurred while listing Food Items"}}
}

var allFoods []bson.M
if err = result.All(ctx ,&allFoods); err != nil {
	log.Fatal(err)
}
c.JSON(http.StatusOK ,allFoods[0])
		
	}
}

// GetFood - Fetch a single food item by ID
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "food not found"})
			return
		}

		c.JSON(http.StatusOK, food)
	}
}

// CreateFood - Create a new food item
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food
		var menu models.Menu

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if Menu exists
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "menu not found"})
			return
		}

		// Set timestamps and IDs
		now := time.Now()
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		food.Created_at = &now
		food.Update_at = &now

		// Fix price precision
		num := toFixed(*food.Price, 2)
		food.Price = &num

		_, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "food creation failed"})
			return
		}

		c.JSON(http.StatusCreated, food)
	}
}

// UpdateFood - Update an existing food item
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var menu models.Menu
		var food models.Food

		foodId := c.Param("food_id")
			if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

 var updateObj primitive.D
 if food.Name !=ni{
updateObj = append(updateObj, bson.E{"name",food.Name})
 }
  if food.Price !=ni{
	updateObj = append(updateObj, bson.E{"price",food.Price})
 }
  if food.Food_image !=ni{
	updateObj = append(updateObj, bson.E{"food_image",food.Food_image})
 }
  if food.Menu_id !=ni{

	err := menuCollection.FindOne(ctx ,bson.M{"menu_id": food.Menu_id}).Decode(&menu)
	defer cancel()
	if err!= nil{
		msg := fmt.Sprintf("message :menu was not found")
		c.JSON(http.StatusInternalServerError ,gin.H{"error":msg})
		return
	}
	updateObj = append(updateObj, bson.E{"menu",food.Price})
 }
 food.Update_at , _ = time.Parse(time.RFC3339, now.Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at",food.Update_at})	
	upsert := true
	filter := bson.M{"food_id":foodId}
opt := options.UpdateOptions{
	Upsert: &upsert,
}
result ,err := foodCollection.UpdateOne(
	ctx,
	filter,
	bson.D {
		{"$set",updateObj}
	},
	&opt
)
if err !=nil {
	msg :=fmt.Sprint("food item update failed")
	c.JSON(http.StatusInternalServerError ,gin.H{"error" :msg})
	return
}
	c.JSON(http.StatusOK , result)
	}
}

// Helper function to round numbers
func round(num float64) int {
	return int(num + math.Copysign(0.5 ,num))
}

// Helper function to fix precision in floating-point numbers
func toFixed(num float64, precision int) float64 {

	
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}
