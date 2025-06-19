package controllers

import (
	"context"
	"fmt"
	"go-rms/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection *mongo.Collection
var validate = validator.New()

func SetMenuCollection(col *mongo.Collection) {
	menuCollection = col
}

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := menuCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching menus"})
			return
		}

		var menus []models.Menu
		if err = cursor.All(ctx, &menus); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding menus"})
			return
		}

		c.JSON(http.StatusOK, menus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		var menu models.Menu

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
			return
		}

		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate struct
		if err := validate.Struct(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		now := time.Now()
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()
		menu.Created_at = now
		menu.Updated_at, _ = time.Parse(time.RFC3339, now.Format(time.RFC3339))

		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu item was not created"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// inTimeSpan checks if check time is between start and end (inclusive)
func inTimeSpan(start, end, check time.Time) bool {
	return (check.After(start) || check.Equal(start)) && (check.Before(end) || check.Equal(end))
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date range"})
				return
			}
			updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
			updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})
		}

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{"name", menu.Name})
		}

		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{"category", menu.Category})
		}

		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

		if len(updateObj) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields to update"})
			return
		}

		upsert := true
		opts := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := menuCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{"$set", updateObj}},
			&opts,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu update failed"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
