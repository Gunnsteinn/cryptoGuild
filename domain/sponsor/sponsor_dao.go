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
)

// Get method implements Sponsor struct and get sponsor from the mongodb.
func (sponsor *Sponsor) Get() *errors.RestErr {
	log.Println("PASO 1")
	ctx, client := connect()

	collection := client.Database(databaseName).Collection(collectionName)
	if getErr := collection.FindOne(ctx, bson.D{{"wallet_address", sponsor.WalletAddress}}).Decode(&sponsor); getErr != nil {
		return errors.NewInternalServerError(getErr.Error())
	}

	disconnect(ctx, client)

	return nil
}

func connect() (context.Context, *mongo.Client) {
	dataSourceName := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		databaseName,
		password,
		host,
		databaseName,
	)

	// Set client options
	clientOptions := options.Client().ApplyURI(dataSourceName)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.TODO()
	log.Println("Connected to MongoDB!")

	return ctx, client
}

func disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}
