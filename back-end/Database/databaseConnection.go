package database

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var envFile map[string]string = *ReadEnvFile();

//image handling ops

//upload to gridfs
func UploadToGridFS(file io.Reader,fileName string)(string,error){
	bucket,err := gridfs.NewBucket(client.Database(envFile["DATABASE_NAME"]))

	if err!=nil{
		return "",err;
	}

	fileID,err := bucket.UploadFromStream(fileName,file);

	if err!=nil{
		return "",err;
	}

	return fileID.Hex(),nil;
}


//database operations
var client *mongo.Client = Connect();

func ReadEnvFile()(*map[string]string){
	envFile,err:= godotenv.Read(".env")

	if err!=nil{
		log.Fatal("Error loading .env file");
	}else{
		log.Println(envFile)
	}

	return &envFile;
}



func GetCollectionByName(collectionName string) *mongo.Collection{
	if client==nil{
		log.Println("GetCollectionByName->client not disconnected");
		log.Fatal("mongo client not connected")
	}

	var dbName string = envFile["DATABASE_NAME"]

	if dbName==""{
		log.Fatal("Database name not found")
		return nil;
	}

	return client.Database(dbName).Collection(collectionName);
}

func Connect() *mongo.Client{
	connectionUrl := envFile["MONGO_URL"];
	fmt.Println("Connecting ",connectionUrl)

	if(connectionUrl==""){
		log.Fatal("Url not found !")
		return nil;
	}

	client,err := mongo.NewClient(options.Client().ApplyURI(connectionUrl));

	if err!=nil{
		log.Panic("Failed to connect database!")
		log.Fatal(err.Error());
	}

	ctx,cancel := 	context.WithTimeout(context.Background(),10*time.Second);

	defer cancel();
	

	err = client.Connect(ctx)

	if err!=nil{
		log.Fatal("Error while connecting database ",err.Error());

	}

	log.Println("Mongodb connected successfully !");
	return client;
}

func Close() error{
	if client == nil{
		return nil;
	}

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)

	err := client.Disconnect(ctx);

	defer cancel();

	if err!=nil{
		return err;
	}

	client = nil;

	return nil;
}