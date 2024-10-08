package database

import (
	model "Backend/Model"
	showe "Backend/Model/Showe"
	util "Backend/Util"
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
var client *mongo.Client = Connect()
var gridBucket *gridfs.Bucket

/*Util handlers*/
func UploadToGridFS(file io.Reader, fileName string) (string, error) {
	log.Println("Invoked uploadToGridFs")

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
	log.Println(client.Database(envFile["DATABASE_NAME"]))

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
			log.Println(envFile["DATABASE_NAME"])
			if client == nil {
				log.Println("mongo client is nil")
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
				return nil, fmt.Errorf("something went wrong while connecting to DB")
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
func SaveNewEventData(collectionName string, event showe.Eventshow) (interface{}, error) {
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

		result, err := collection.InsertOne(ctx, event)

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

func SaveNewLiveshowData(collectionName string, liveshow showe.Liveshow) (interface{}, error) {
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

		result, err := collection.InsertOne(ctx, liveshow)

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

func SaveNewActivityData(collectionName string, activity showe.ActivityShow) (interface{}, error) {
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

		result, err := collection.InsertOne(ctx, activity)

		if err != nil {
			log.Println("Could not save user in mongodb ->SaveNeweventData")
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

func SaveNewMovieData(collectionName string, movie showe.Movie) (interface{}, error) {
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

func FetchShoweByFilter(collectionName string, startIndex int, recordPerPage int, filter string) (interface{}, error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	//match stage
	matchStage := bson.D{{"$match", bson.D{{}}}}

	//group stage
	groupStage := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: bson.D{{"_id", "null"}}},
		{Key: "total_count", Value: bson.D{{"$sum", "1"}}},
		{Key: "data", Value: bson.D{{"$push", "$$ROOT"}}},
	}}}

	//project stage
	projectStage := bson.D{
		{
			"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"data", bson.D{
					{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			},
		},
	}

	switch filter {

	case string(util.Movie):
		collectionName := util.GetCollectionNameByShoweType(string(util.Movie))
		collection := GetCollectionByName(collectionName)
		result, err := collection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})

		defer cancel()

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		var allMovies []bson.M

		if err = result.All(ctx, &allMovies); err != nil {
			log.Fatal(err)
			return nil, err
		}

		return allMovies, nil

		break

	case string(util.Event):
		collectionName := util.GetCollectionNameByShoweType(string(util.Movie))
		collection := GetCollectionByName(collectionName)
		result, err := collection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})

		defer cancel()

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		var allEvents []bson.M

		if err = result.All(ctx, &allEvents); err != nil {
			log.Fatal(err)
			return nil, err
		}

		return allEvents, nil
		break

	case string(util.Activity):
		collectionName := util.GetCollectionNameByShoweType(string(util.Movie))
		collection := GetCollectionByName(collectionName)
		result, err := collection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})

		defer cancel()

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		var allActivities []bson.M

		if err = result.All(ctx, &allActivities); err != nil {
			log.Fatal(err)
			return nil, err
		}

		return allActivities, nil
		break

	case string(util.LiveShow):
		collectionName := util.GetCollectionNameByShoweType(string(util.Movie))
		collection := GetCollectionByName(collectionName)
		result, err := collection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})

		defer cancel()

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		var allLiveShows []bson.M

		if err = result.All(ctx, &allLiveShows); err != nil {
			log.Fatal(err)
			return nil, err
		}

		return allLiveShows, nil
		break

	}

	return nil, fmt.Errorf("the movie type does not exist")
}

func CreateBookingByshowId(showeType string,
		registeredShoweId string,
		seat_number string,
		modelBooking *model.BookingModel)(interface{},error){
	
	var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second);
	collection := GetCollectionByName(util.GetCollectionNameByShoweType(showeType));

	result,err := collection.InsertOne(ctx,modelBooking);
	defer cancel();

	if err!=nil{
		return nil,err;
	}

	return result,nil;
}

func FetchAllMovieShowe(collectionName string, startIndex int, recordPerPage int) (interface{}, error) {
	var allMovieList []bson.M

	collection := GetCollectionByName(collectionName)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	//defer cancel();

	matchStage := bson.D{{"$match", bson.D{{}}}}

	groupStage := bson.D{{Key: "$group",
		Value: bson.D{{Key: "_id",
			Value: bson.D{{"_id", "null"}}},
			{Key: "total_count", Value: bson.D{{"$sum", 1}}},
			{Key: "data", Value: bson.D{{"$push", "$$ROOT"}}},
		}}}

	projectStage := bson.D{
		{
			"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"data", bson.D{
					{"$slice", []interface{}{"$data", startIndex, recordPerPage}},
				}},
			},
		},
	}

	result, err := collection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage,
	})

	defer cancel()

	if err != nil {
		log.Println("Could not fetch movie items")
		log.Fatal(err)
		return nil, err
	}

	err = result.All(ctx, &allMovieList)
	if err != nil {
		log.Println("Could not parse allmovie list!")
		log.Fatal(err)
		return nil, err
	}

	return allMovieList[0], nil
}

func FetchAllActivityShowe(collectionName string, startIndex int, recordPerPage int) (interface{}, error) {
	var allActivityList []bson.M

	collection := GetCollectionByName(collectionName)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	//filter by activity
	matchStage := bson.D{{"$match", bson.D{{}}}}

	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}},
		{"total_count", bson.D{{"$sum", 1}}},
		{"data", bson.D{{"$push", "$$ROOT"}}},
	}}}

	projectStage := bson.D{
		{
			"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"data", bson.D{
					{"$slice", []interface{}{"$data", startIndex, recordPerPage}},
				}},
			},
		},
	}

	result, err := collection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage,
	})
	defer cancel()

	if err != nil {
		log.Println("could not fetch activity items")
		log.Fatal(err)
		return nil, err
	}

	err = result.All(ctx, &allActivityList)
	if err != nil {
		log.Println("could not parse all activity list")
		log.Fatal(err)
		return nil, err
	}
	log.Println("hi2")

	return allActivityList, nil
}

func FetchAllEventShowe(collectionName string, startIndex int, recordPerPage int) (interface{}, error) {
	var allEventList []bson.M

	collection := GetCollectionByName(collectionName)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	//filter by activity
	matchStage := bson.D{{"$match", bson.D{{}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}},
		{"total_count", bson.D{{"$sum", 1}}},
		{"data", bson.D{{"$push", "$$ROOT"}}},
	}}}

	projectStage := bson.D{
		{
			"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"data", bson.D{
					{"$slice", []interface{}{"$data", startIndex, recordPerPage}},
				}},
			},
		},
	}

	result, err := collection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage,
	})
	defer cancel()

	if err != nil {
		log.Println("could not fetch activity items")
		log.Fatal(err)
		return nil, err
	}

	err = result.All(ctx, &allEventList)
	if err != nil {
		log.Println("could not parse all activity list")
		log.Fatal(err)
		return nil, err
	}

	return allEventList, nil
}

func FetchAllLiveshows(collectionName string, startIndex int, recordPerPage int) (interface{}, error) {
	var allLiveshowList []bson.M

	collection := GetCollectionByName(collectionName)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	//filter by activity
	matchStage := bson.D{{"$match", bson.D{{}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}},
		{"total_count", bson.D{{"$sum", 1}}},
		{"data", bson.D{{"$push", "$$ROOT"}}},
	}}}

	projectStage := bson.D{
		{
			"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"data", bson.D{
					{"$slice", []interface{}{"$data", startIndex, recordPerPage}},
				}},
			},
		},
	}

	result, err := collection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage,
	})
	defer cancel()

	if err != nil {
		log.Println("could not fetch activity items")
		log.Fatal(err)
		return nil, err
	}

	err = result.All(ctx, &allLiveshowList)
	if err != nil {
		log.Println("could not parse all activity list")
		log.Fatal(err)
		return nil, err
	}

	return allLiveshowList, nil
}

