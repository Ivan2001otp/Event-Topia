package database

import (
	showe "Backend/Model/Showe"
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var envFile map[string]string = *ReadEnvFile()
var client *mongo.Client = Connect();
var gridBucket *gridfs.Bucket

/*Util handlers*/
func UploadToGridFS(file io.Reader, fileName string) (string, error) {
	log.Println("Invoked uploadToGridFs");
	
	bucket, err := BucketProvider()
	
	if err != nil {
		log.Println("BucketProvider failed")
		log.Fatal(err)
		return "", err
	}

	fileID, err := bucket.UploadFromStream(fileName, file)
	log.Println("The file id is ", fileID)
	if err != nil {
		log.Println("could not upload UploadFromStream()")
		log.Fatal(err)

		return "", err
	}

	return fileID.Hex(), nil
	// return "https://cdn.pixabay.com/photo/2024/03/04/14/17/ai-generated-8612487_640.jpg",nil;
}

func ReadEnvFile() *map[string]string {
	envFile, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		// log.Println(envFile)
	}

	return &envFile
}


func GetCollectionByName(collectionName string) *mongo.Collection {
	if client == nil {
		log.Println("GetCollectionByName->client not disconnected")
		log.Fatal("mongo client not connected")
	}

	var dbName string = envFile["DATABASE_NAME"]

	if dbName == "" {
		log.Fatal("Database name not found")
		return nil
	}

	return client.Database(dbName).Collection(collectionName)
}

func Connect() *mongo.Client {
	connectionUrl := envFile["MONGO_URL"]
	fmt.Println("Connecting ", connectionUrl)


	client, err := mongo.NewClient(options.Client().ApplyURI(connectionUrl))

	if err != nil {
		log.Panic("Failed to connect database!")
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Error while connecting database ", err.Error())

	}

	log.Println("Mongodb connected successfully !")
	log.Println(client.Database(envFile["DATABASE_NAME"]));

	return client
}

func Close() error {
	if client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err := client.Disconnect(ctx)

	defer cancel()

	if err != nil {
		return err
	}

	client = nil

	return nil
}

var lockForBucket = &sync.Mutex{}

func BucketProvider() (*gridfs.Bucket, error) {
	if gridBucket == nil {
		lockForBucket.Lock()

		defer lockForBucket.Unlock()

		if gridBucket == nil {
			log.Println("Initializing bucket")
			log.Println(envFile["DATABASE_NAME"]);
			if(client==nil){
				log.Println("mongo client is nil");
			}
			gridBucket, err := gridfs.NewBucket(client.Database(envFile["DATABASE_NAME"]))
			
			if err != nil {
				log.Fatal("Something wrong with grid Fs in mongodb")
				return nil, err
			}
			log.Println("Singleton instance of mongo provided")

			return gridBucket, nil

		}
	} else {
		log.Println("Singleton instance of mongo provided")
	}

	return gridBucket, nil
}

// singleton instance implementation for database
var lockVariable = &sync.Mutex{}

func MongoDbProvider() (*mongo.Client, error) {
	if client == nil {
		lockVariable.Lock()

		defer lockVariable.Unlock()

		if client == nil {
			err := Connect()
			if err != nil {
				return nil, fmt.Errorf("Something went wrong while connecting to DB!")
			}
		} else {
			log.Println("Singleton already provided.")
		}
	} else {
		log.Println("Singleton already provided.")
	}

	return client, nil
}


// database operations CRUD
func SaveNewShoweData(collectionName string, movie showe.Movie) (interface{}, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	collection := GetCollectionByName(collectionName)

	//channels
	resultChan := make(chan *mongo.InsertOneResult)
	errChan := make(chan error)

	var wg sync.WaitGroup

	wg.Add(1)

	//goroutine
	go func() {
		defer wg.Done()

		result, err := collection.InsertOne(ctx, movie)

		if err != nil {
			log.Println("Could not save user in mongodb ->SaveNewShoweData")
			errChan <- err
			return
		}

		resultChan <- result
	}()

	//separete goroutine to handle closing channels
	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	select {
	case result := <-resultChan:
		return result.InsertedID, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}


func FetchAllMovieShowe(collectionName string,startIndex int,recordPerPage int)([] showe.Movie,error){
	var allMovieList [] showe.Movie;

	collection := GetCollectionByName(collectionName)
	var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second);

	//defer cancel();	

	matchStage := bson.D{{"$match",bson.D{{}}}}

	groupStage := bson.D{{Key:"$group",
		Value:bson.D{{Key:"_id",
			Value: bson.D{{"_id","null"}}},
		{Key: "total_count",Value: bson.D{{"$sum","1"}}},
		{Key: "data",Value: bson.D{{"$push","$$ROOT"}}},
		
	}}}

	projectStage := bson.D{
		{
			"$project",bson.D{
				{"_id",0},
				{"total_count","1"},
				{"movie_showes",bson.D{
					{"$slice",[]interface{}{"$data",startIndex,recordPerPage}},
				}},
			},
		},
	}

	result ,err := collection.Aggregate(ctx,mongo.Pipeline{
		matchStage,groupStage,projectStage,
	});

	defer cancel();

	if err!=nil{
		log.Println("Could not fetch movie items");
		log.Fatal(err);
		return nil,err;
	}

	err = result.All(ctx,&allMovieList);
	if err!=nil{
		log.Println("Could not parse allmovie list!");
		log.Fatal(err);
		return nil,err;
	}

	return allMovieList,nil;
}