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
	AddColumn(id string, newColumn *model.CustomField) *model.Table
	DeleteTable(id string) *model.Table
	DeleteTableColumn(tableid string, columnid string) *model.Table
	ModifyTableColumn(tableid string, columnid string, newCustomeField *model.NewCustomField) *model.Table
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
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

func (db *database) AddColumn(id string, newColumn *model.CustomField) *model.Table {

	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	var table *model.Table
	cur := collection.FindOne(context.TODO(), bson.M{"tableid": id})
	cur.Decode(&table)
	table.Fields = append(table.Fields, newColumn)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"tableid": id}, bson.M{"$set": bson.M{"fields": table.Fields}})
	if err != nil {
		log.Fatal(err)
	}
	return table

}

func (db *database) DeleteTable(id string) *model.Table {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	var table *model.Table
	cur := collection.FindOne(context.TODO(), bson.M{"tableid": id})
	cur.Decode(&table)
	_, err := collection.DeleteOne(context.TODO(), bson.M{"tableid": id})
	if err != nil {
		log.Fatal(err)
	}
	return table
}

func (db *database) DeleteTableColumn(tableid string, coulmnid string) *model.Table {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	var table *model.Table
	cur := collection.FindOne(context.TODO(), bson.M{"tableid": tableid})
	cur.Decode(&table)

	var tempfields []*model.CustomField

	for i := 0; i < len(table.Fields); i++ {
		if table.Fields[i].FieldID != coulmnid {
			tempfields = append(tempfields, table.Fields[i])
		}
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"tableid": tableid}, bson.M{"$set": bson.M{"fields": tempfields}})
	if err != nil {
		log.Fatal(err)
	}
	table.Fields = tempfields
	return table
}

func (db *database) ModifyTableColumn(tableid string, columnid string, newCustomeField *model.NewCustomField) *model.Table {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	var table *model.Table
	cur := collection.FindOne(context.TODO(), bson.M{"tableid": tableid})
	cur.Decode(&table)

	for i := 0; i < len(table.Fields); i++ {
		if table.Fields[i].FieldID == columnid {
			table.Fields[i].FieldName = newCustomeField.FieldName
			table.Fields[i].DataType = newCustomeField.DataType
			table.Fields[i].Value = newCustomeField.Value
			table.Fields[i].MaxValue = newCustomeField.MaxValue
			table.Fields[i].MaxValue = newCustomeField.MinValue
			table.Fields[i].DefaultValue = newCustomeField.DefaultValue
			table.Fields[i].IsRequired = newCustomeField.IsRequired
			table.Fields[i].Visibility = newCustomeField.Visibility
			break
		}
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"tableid": tableid}, bson.M{"$set": bson.M{"fields": table.Fields}})
	if err != nil {
		log.Fatal(err)
	}
	return table
}
