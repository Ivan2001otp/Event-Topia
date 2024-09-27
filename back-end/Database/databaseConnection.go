package database

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var envFile map[string]string = *ReadEnvFile();


//database operations
var client *mongo.Client;

var gridBucket *gridfs.Bucket;
//image handling ops

//upload to gridfs
func UploadToGridFS(file io.Reader,fileName string)(string,error){
	bucket,err := BucketProvider();
	
	if err!=nil{
		log.Println("BucketProvider failed")
		log.Fatal(err)
		return "",err;
	}

	fileID,err := bucket.UploadFromStream(fileName,file);

	if err!=nil{
		log.Println("could not upload UploadFromStream()")
		log.Fatal(err)

		return "",err;
	}

	return fileID.Hex(),nil;
}



func ReadEnvFile()(*map[string]string){
	envFile,err:= godotenv.Read(".env")

	if err!=nil{
		log.Fatal("Error loading .env file");
	}else{
		//log.Println(envFile)
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

var lockForBucket =  &sync.Mutex{};
func BucketProvider() (*gridfs.Bucket, error){
	if gridBucket==nil{
		lockForBucket.Lock();

		defer lockVariable.Unlock();

		if gridBucket==nil{
			gridBucket,err := gridfs.NewBucket(client.Database(envFile["DATABASE_NAME"]),);

			if err!=nil{
				log.Fatal("Something wrong with grid Fs in mongodb");
				return nil,err;
			}
		log.Println("Singleton instance of mongo provided")


			return gridBucket,nil;

		}
	}else{
		log.Println("Singleton instance of mongo provided")
	}

	return gridBucket,nil;
}

//singleton instance implementation for database
var lockVariable = &sync.Mutex{}
func MongoDbProvider() (*mongo.Client,error){
	if client==nil{
		lockVariable.Lock()

		defer lockVariable.Unlock()

		if client==nil{
			err:= Connect();
			if err!=nil{
				return nil,fmt.Errorf("Something went wrong while connecting to DB!");
			}
		}else{
		log.Println("Singleton already provided.");
		}
	}else{
		log.Println("Singleton already provided.");
	}

	return client,nil;
}