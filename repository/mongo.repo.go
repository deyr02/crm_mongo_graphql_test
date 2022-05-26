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

type TableRepository interface {
	Save(Customer *model.Table)
	FindAll() []*model.Table
	FindTableById(id string) *model.Table
}

const (
	DATABASE   = "BNZL_CRM"
	COLLECTION = "Tables"
)

type database struct {
	client *mongo.Client
}

func New() TableRepository {
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

func (db *database) FindAll() []*model.Table {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())
	var result []*model.Table
	for cursor.Next(context.TODO()) {
		var cus *model.Table
		err := cursor.Decode(&cus)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, cus)
	}
	return result
}

func (db *database) Save(table *model.Table) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), table)

	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) FindTableById(id string) *model.Table {

	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())
	var result *model.Table
	for cursor.Next(context.TODO()) {
		var cus *model.Table
		err := cursor.Decode(&cus)
		if err != nil {
			log.Fatal(err)
		}
		if cus.TableID == id {
			result = cus
			break
		}

	}
	return result
}
