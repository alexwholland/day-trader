package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

// Logger instance
var logger *log.Logger
var client *mongo.Client
var collection *mongo.Collection

func init() {
	// Initialize logger with desired settings, you can modify as per your requirement
	logger = log.New(os.Stdout, "Wallet Logger: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Connect to MongoDB
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Set the collection
	collection = client.Database("logging").Collection("log_collection")
}

// LogBuyOrder logs the details of a buy order
func LogOrder(order Order) {
	logMessage := fmt.Sprintf("Buy Order: StockTxID=%s, StockID=%s, WalletTxID=%s, Quantity=%d, Price=$%.2f, TimeStamp=%s",
		order.StockTxID, order.StockID, order.WalletTxID, order.Quantity, *order.Price, order.TimeStamp)
	_, err := collection.InsertOne(context.TODO(), bson.M{"log": logMessage})
	if err != nil {
		logger.Printf("Failed to save buy order to MongoDB: %v", err)
		// Handle error as required
	}
}
