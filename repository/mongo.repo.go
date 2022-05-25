package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deyr02/crm_mongo_graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerRepository interface {
	Save(Customer *model.Customer)
	FindAll() []*model.Customer
}

const (
	DATABASE   = "BNZL_CRM"
	COLLECTION = "Customers"
)

type database struct {
	client *mongo.Client
}

func New() CustomerRepository {
	//mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	//MONGODB := os.Getenv("MONGODB")
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.4.2")
	clientOptions = clientOptions.SetMaxPoolSize(50)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	dbclient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to database")
	return &database{
		client: dbclient,
	}

}

func (db *database) FindAll() []*model.Customer {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())
	var result []*model.Customer
	for cursor.Next(context.TODO()) {
		var cus *model.Customer
		err := cursor.Decode(&cus)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, cus)
	}
	return result
}

func (db *database) Save(customer *model.Customer) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), customer)

	if err != nil {
		log.Fatal(err)
	}
}
