package database

import (
	"context"
	"log"
	"os"

	"github.com/rcsolis/endoflife_client/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	// CollectionName is the name of the collection in the database
	CollectionName = "technologies"
)

type LanguageCycleDb struct {
	Id              bson.ObjectID `json:"id" bson:"_id"`
	Name            string        `json:"name" bson:"name"`
	Version         string        `json:"version" bson:"version"`
	Cycle           string        `json:"cycle" bson:"cycle"`
	ReleaseDate     string        `json:"releaseDate" bson:"releaseDate"`
	Eol             string        `json:"eol" bson:"eol"`
	Latest          string        `json:"latest" bson:"latest"`
	Link            string        `json:"link" bson:"link"`
	Lts             string        `json:"lts" bson:"lts"`
	Support         string        `json:"support" bson:"support"`
	Discontinued    string        `json:"discontinued" bson:"discontinued"`
	ExtendedSupport string        `json:"extendedSupport" bson:"extendedSupport"`
}

func connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@mongo:27017/")
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}
	return client, err
}

/**
* GetDetails is a function to get the details of a technology
*
*@param name    The name of the technology
*@param version The version of the technology
*@return       The details of the technology and a boolean indicating success
 */
func GetDetails(name, version string) (LanguageCycleDb, bool) {
	var result bson.M
	var technologyDetails LanguageCycleDb

	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer func(ctx context.Context) {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}(context.Background())
	log.Println("Getting details from the database")
	// Send a ping to confirm a successful connection
	if err := client.Database(os.Getenv("MONGODB_DBNAME")).RunCommand(context.Background(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatal(err)
	}

	collection := client.Database(os.Getenv("MONGODB_DBNAME")).Collection(CollectionName)

	err = collection.FindOne(context.Background(), bson.M{"name": name, "version": version}).Decode(&technologyDetails)
	if err == mongo.ErrNoDocuments {
		log.Println("Document not found")
		return LanguageCycleDb{}, false
	}
	if err != nil {
		log.Println("Error finding the document")
		return LanguageCycleDb{}, false
	}

	return technologyDetails, true
}

/**
* SaveDetails is a function to save the details of a technology
*
*@param name     The name of the technology
*@param version  The version of the technology
*@param response The response from the rpc server
*@return         An error if any
 */
func SaveDetails(name, version string, response model.LanguageCycle) error {
	var result bson.M
	log.Println("Saving details to the database")
	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	// Send a ping to confirm a successful connection
	if err := client.Database(os.Getenv("MONGODB_DBNAME")).RunCommand(context.Background(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatal(err)
	}

	collection := client.Database(os.Getenv("MONGODB_DBNAME")).Collection(CollectionName)

	_, err = collection.InsertOne(context.Background(), bson.M{
		"name":            name,
		"version":         version,
		"cycle":           response.Cycle,
		"releaseDate":     response.ReleaseDate,
		"eol":             response.Eol,
		"latest":          response.Latest,
		"link":            response.Link,
		"lts":             response.Lts,
		"support":         response.Support,
		"discontinued":    response.Discontinued,
		"extendedSupport": response.ExtendedSupport,
	})
	if err != nil {
		log.Println("Error saving the document:", err)
		return err
	}
	return nil
}
