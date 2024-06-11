package routes

import (
	"context"
	"log"
	"net/http"

	"example/licigo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func addTenderRoutes(rg *gin.RouterGroup) {
	tenders := rg.Group("/tenders")

	tenders.GET(":code", func(c *gin.Context) {
		code := c.Param("code")
		tenderColl := client.Database("licigo").Collection("tender")
		var tender models.Tender
		err := tenderColl.FindOne(context.TODO(), bson.D{{"code", code}}).Decode(&tender)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, tender)
	})

	tenders.GET("", func(c *gin.Context) {
		tenderColl := client.Database("licigo").Collection("tender")
		cursor, err := tenderColl.Find(context.TODO(), bson.D{{}})
		if err != nil {
			panic(err)
		}
		var tenders []models.Tender
		if err = cursor.All(context.TODO(), &tenders); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, tenders)
	})

	tenders.POST("", func(c *gin.Context) {
		var tender models.Tender
		if err := c.ShouldBindBodyWithJSON(&tender); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tenderColl := client.Database("licigo").Collection("tender")

		tenderResult := tenderColl.FindOne(context.TODO(), bson.D{{"code", tender.Code}})

		if tenderResult.Err() == nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println("There's already a tender with that code")
			return
		}

		_, err := tenderColl.InsertOne(context.TODO(), tender)

		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			panic(err)
		}

		c.JSON(http.StatusOK, tender)
	})
}
