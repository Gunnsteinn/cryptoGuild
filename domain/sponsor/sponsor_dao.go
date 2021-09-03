package sponsor

import (
	"context"
	"fmt"
	"github.com/Gunnsteinn/cryptoGuild/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const (
	mongoSponsorDatabaseName   = "mongo_sponsor_database_name"
	mongoSponsorPassword       = "mongo_sponsor_password"
	mongoSponsorHost           = "mongo_sponsor_host"
	mongoSponsorCollectionName = "mongo_sponsor_collection_name"
)

// Client is a variable. DB is a database handle representing a pool of zero or more underlying connections.
// databases, password, host, collection are variables that are settings in Environment at the begins like this:
//  >> mongoSponsorDatabaseName   = "mongo_sponsor_database_name";
//  >> mongoSponsorPassword       = "mongo_sponsor_password";
//  >> mongoSponsorHost           = "mongo_sponsor_host";
//  >> mongoSponsorCollectionName = "mongo_sponsor_collection_name";
var (
	databaseName   = os.Getenv(mongoSponsorDatabaseName)
	password       = os.Getenv(mongoSponsorPassword)
	host           = os.Getenv(mongoSponsorHost)
	collectionName = os.Getenv(mongoSponsorCollectionName)

	dataSourceName = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		databaseName,
		password,
		host,
		databaseName,
	)
)

// Get method implements Sponsor struct and get sponsor from the mongodb.
func (sponsor *Sponsor) Get() *errors.RestErr {
	ctx, client := connect()

	collection := client.Database(databaseName).Collection(collectionName)
	if getErr := collection.FindOne(ctx, bson.D{{"wallet_address", sponsor.WalletAddress}}).Decode(&sponsor); getErr != nil {
		return errors.NewInternalServerError(getErr.Error())
	}

	disconnect(ctx, client)
	return nil
}

// GetAll method implements Sponsor struct and get all sponsors from the mongodb.
func (sponsor *Sponsor) GetAll() ([]Sponsor, *errors.RestErr) {
	ctx, client := connect()

	collection := client.Database(databaseName).Collection(collectionName)
	cursor, getErr := collection.Find(ctx, bson.M{})
	if getErr != nil {
		return nil, errors.NewInternalServerError(getErr.Error())
	}

	defer cursor.Close(ctx)
	results := make([]Sponsor, 0)
	for cursor.Next(ctx) {
		var result Sponsor
		cursor.Decode(&result)
		results = append(results, result)
	}

	disconnect(ctx, client)
	return results, nil
}

// Create method implements Sponsor struct and create sponsor from the mongodb.
func (sponsor *Sponsor) Create() *errors.RestErr {
	ctx, client := connect()

	collection := client.Database(databaseName).Collection(collectionName)
	if _, getErr := collection.InsertOne(ctx, sponsor); getErr != nil {
		return errors.NewInternalServerError(getErr.Error())
	}

	disconnect(ctx, client)
	return nil
}

func (sponsor *Sponsor) Update() *errors.RestErr {
	filter := bson.D{{"wallet_address", sponsor.WalletAddress}}
	update := bson.M{"$set": sponsor}
	ctx, client := connect()

	collection := client.Database(databaseName).Collection(collectionName)
	_, errUpdate := collection.UpdateOne(ctx, filter, update)
	disconnect(ctx, client)
	if errUpdate != nil {
		return errors.NewInternalServerError(errUpdate.Error())
	}

	return nil
}

// Delete method implements Sponsor struct and delete sponsor from the mongodb.
func (sponsor *Sponsor) Delete() *errors.RestErr {
	ctx, client := connect()

	collection := client.Database(databaseName).Collection(collectionName)
	_, errDelete := collection.DeleteMany(ctx, bson.D{{"wallet_address", sponsor.WalletAddress}})
	disconnect(ctx, client)

	if errDelete != nil {
		return errors.NewInternalServerError(errDelete.Error())
	}

	return nil
}

// Connect to MongoDB
func connect() (context.Context, *mongo.Client) {
	// Set client options
	clientOptions := options.Client().ApplyURI(dataSourceName)

	// Connect to MongoDB
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return ctx, client
}

// Disconnect to MongoDB
func disconnect(ctx context.Context, client *mongo.Client) {
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, ctx)
}
