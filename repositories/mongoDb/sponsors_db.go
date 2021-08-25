package mongoDb

import (
	"context"
	"fmt"
	"github.com/Gunnsteinn/cryptoGuild/domain/sponsor"
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
	ClientMongo *mongo.Client

	databaseName = os.Getenv(mongoSponsorDatabaseName)
	password     = os.Getenv(mongoSponsorPassword)
	host         = os.Getenv(mongoSponsorHost)
	collection   = os.Getenv(mongoSponsorCollectionName)
)

//func init() {
//	dataSourceName := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority",
//		databaseName,
//		password,
//		host,
//	)
//
//	var err error
//	ClientMongo, err = mongo.NewClient(options.Client().ApplyURI(dataSourceName))
//	if err != nil {
//		log.Fatal(err)
//	}
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	err = ClientMongo.Connect(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = ClientMongo.Ping(ctx, readpref.Primary())
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("database successfully configured.")
//}

type MongoRepository struct {
}

func (m MongoRepository) FindOne(wallet string) *sponsor.Sponsor {
	ctx, client := connect()

	var result *sponsor.Sponsor
	collection := client.Database(databaseName).Collection(collection)
	collection.FindOne(ctx, sponsor.Sponsor{WalletAddress: wallet}).Decode(&result)

	disconnect(ctx, client)

	return result
}

func connect() (context.Context, *mongo.Client) {
	dataSourceName := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority",
		databaseName,
		password,
		host,
	)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dataSourceName))
	if err != nil {
		log.Fatal(err)
	}
	return ctx, client
}

func disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}
