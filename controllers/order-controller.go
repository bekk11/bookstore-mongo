package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongo_go/pkg/common/models"
	"net/http"
)

//---------------- LIST ORDERS --------------------------

func (h handler) ListOrders(ctx *gin.Context) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("bookstore").Collection("orders")

	var orders []models.Order

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	ctx.JSON(http.StatusOK, orders)
}

//=======================================================

//---------------- CREATE ORDER -------------------------

func (h handler) CreateOrder(ctx *gin.Context) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("bookstore").Collection("orders")

	var order models.Order

	if err := ctx.Bind(&order); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := collection.InsertOne(context.Background(), order)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

//=======================================================

//---------------- RETRIEVE ORDER -----------------------

func (h handler) RetrieveOrder(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("bookstore").Collection("orders")

	var order models.Order
	if err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&order); err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	ctx.JSON(http.StatusOK, order)
}

//=======================================================

//---------------- UPDATE ORDER -------------------------

func (h handler) UpdateOrder(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("bookstore").Collection("orders")

	var order models.Order

	if err := ctx.Bind(&order); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"fullName": order.FullName, "phoneNumber": order.PhoneNumber, "bookId": order.BookId}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if result.ModifiedCount == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.Status(http.StatusOK)
}

//=======================================================

//---------------- DESTROY ORDER ------------------------

func (h handler) DestroyOrder(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("bookstore").Collection("orders")

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusNoContent)
}

//=======================================================
