package utils

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"fmt"
	"github.com/ShijuPJohn/quizapp-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

var Mg MongoInstance

var Secret string

const dbName = "test"

//	func MongoDBConnect() (error, func()) {
//		mongoURI := getMongoURLandPopulateSecretString()
//		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
//		deferFunc := func() {
//			if err := client.Disconnect(context.TODO()); err != nil {
//				panic(err)
//			}
//		}
//		if err != nil {
//			fmt.Println(err)
//			return err, deferFunc
//		}
//
//		db := client.Database(dbName)
//
//		if err != nil {
//			return err, deferFunc
//		}
//
//		Mg = MongoInstance{
//			Client: client,
//			Db:     db,
//		}
//		indexModel := mongo.IndexModel{
//			Keys:    bson.D{{"email", 1}},
//			Options: options.Index().SetUnique(true),
//		}
//
//		// Create the index
//		_, err = Mg.Db.Collection("users").Indexes().CreateOne(context.Background(), indexModel)
//		if err != nil {
//			log.Fatal("Error creating the index")
//		}
//
//		fmt.Println("Connected to MongoDB cloud")
//		return nil, deferFunc
//	}
func getMongoURLandPopulateSecretString() []string {
	name := "projects/1037996227658/secrets/quizapp_s/versions/2"
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatal("failed to create secretmanager client: %w", err)

	}
	defer client.Close()
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatal("failed to access secret version: %w", err)
	}
	stringVal := string(result.Payload.Data)
	words := strings.Fields(stringVal)
	Secret = words[len(words)-1]
	return words

}
func ConnectDb() {
	sValues := getMongoURLandPopulateSecretString()
	dbName := sValues[0]
	host := sValues[1]
	port := sValues[2]
	username := sValues[3]
	password := sValues[4]
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, username, password, dbName, port)

	//TimeZone=Asia/Shanghai

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.User{})
	DB = Dbinstance{
		Db: db,
	}
}
