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

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	collection := client.Database(databaseName).Collection(collectionName)
	//bson.D{{"wallet_address", sponsor.WalletAddress}}
	if getErr := collection.FindOne(ctx, bson.D{{}}).Decode(&sponsor); getErr != nil {
		log.Println("PASO ERROR:" + getErr.Error())
		return errors.NewInternalServerError(getErr.Error())
	}
	log.Println("PASO 7")
	disconnect(ctx, client)
	log.Println("PASO 8")
	return nil
}

func connect() (context.Context, *mongo.Client) {
	//dataSourceName := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority",
	//	databaseName,
	//	password,
	//	host,
	//)

	ctx := context.Background()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI(dataSourceName))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://GGCGdb:S%40yley23@cluster0.6hrfc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}
	return ctx, client
}

func disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}
