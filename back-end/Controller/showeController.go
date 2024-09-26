package controller

import (
	database "Backend/Database"
	"Backend/Model"
	factory "Backend/Util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//use of mutex to handle concurrent operations

var (
	mu sync.Mutex
)

type status map[string]interface{};

func CreateShowModel(wg *sync.WaitGroup) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {


		if(r.Method !=http.MethodPost){
			http.Error(w,"Invalid request method ",http.StatusMethodNotAllowed);
			return;
		}

		var newShoweModel model.BaseShowe
		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second);

		defer cancel();
		//parse form
		err := r.ParseMultipartForm(20<<30)//10 mb limit

		if err!=nil{
			http.Error(w,"Error parsing form",http.StatusBadRequest);
			return;
		}

		//validate inputs
		var imageList []string;
		var majorThumbnail string;

		frontImages := r.MultipartForm.File["images"];
		thumbnailImages := r.MultipartForm.File["thumbnail"];

		if(len(frontImages)==0 || len(thumbnailImages)==0){
			http.Error(w,"Missing media images",http.StatusBadRequest);
			return;
		}

		showeName := r.FormValue("showeName");
		showeRating := r.FormValue("showeRating");
		showeRatingInt,_ := strconv.ParseInt(showeRating,10,64);

		vendorName := r.FormValue("vendorName");
		showeVotes := r.FormValue("showeVotes");
		showeVotesInt,_ := strconv.ParseInt(showeVotes,10,64)

		showeDuration := r.FormValue("showeDuration");
		showeDurationInt,_ := strconv.ParseInt(showeDuration,10,64);

		showeType := factory.SetShowDurationDynamically(r.FormValue("showType"))
		//showe type cannot be empty

		showeGenre := strings.Split(r.FormValue("showeGenre"),",");
		showeFromTime := r.FormValue("showeFromTime");
		showeEndTime := r.FormValue("showeEndTime");
		showeAllowedAges := r.FormValue("showeAllowedAges");
		showeReleasedDate := r.FormValue("showeReleasedDate")

		showeUserExperience := r.FormValue("showeUserExperience");
		showeNote := r.FormValue("showeNote");
		showeAboutUs := strings.Split(r.FormValue("showeAboutUs"),",");
		showeTermsConditions := strings.Split(r.FormValue("showeTermsConditions"),",");


		showeCause := r.FormValue("");
		showeLanguages := strings.Split(r.FormValue("showeLanguages"),",");
		// showeCrewMembers := strings.Split(r.FormValue("showeCrewMembers"),",");

		//fetching crew members
		var showeCrewMembers []model.CrewMember;
		for i:=0;;i++{
			imageUrl_ := r.FormValue(fmt.Sprintf("crew[%d].imageUrl",i));
			aboutCrew_ := r.FormValue(fmt.Sprintf("crew[%d].aboutCrew",i));

			if imageUrl_=="" && aboutCrew_==""{
				break;
			}

			
			showeCrewMembers = append(showeCrewMembers, model.CrewMember{
				CrewImageUrl: imageUrl_,
				AboutCrew:aboutCrew_,
			});

		}

		// showeReviews := strings.Split(r.FormValue("showeReviews"),",");
		//storing reviews
		var showeReviews []model.Review;

		for i:=0;;i++{
			imageUrl := r.FormValue(fmt.Sprintf("review[%d].ReviewImg",i))
			message := r.FormValue(fmt.Sprintf("review[%d].Message",i))

			if(imageUrl=="" && message==""){
				break;
			}

			showeReviews = append(showeReviews, model.Review{
				ReviewImg: &imageUrl,
				Message:message,
			})
		}

		//store files and parsed data.

		//using wait and signal to implement concurrency
		var wg sync.WaitGroup
		resultChan := make(chan struct{
			ID string
			Err error
		},len(frontImages)+1)

		for _,fileHeader := range frontImages{
			wg.Add(1);

			//goroutine
			go func(fileHeader *multipart.FileHeader){
				defer wg.Done();
				file,err := fileHeader.Open();

				if err!=nil{
					resultChan <- struct{ID string; Err error}{"",err}
					return;
				}
				defer file.Close();



				//upload to gridFS
				fileID,err := database.UploadToGridFS(file,fileHeader.Filename)
				
				resultChan <- struct{ID string; Err error}{fileID,err}
				

			}(fileHeader);
		}

		//handling the thumbnail
		for _,thumbnailHeader:=range thumbnailImages{
			wg.Add(1)
			go func (fileHeader *multipart.FileHeader)  {
				defer wg.Done()
				file,err := fileHeader.Open();
				if err!=nil{
					resultChan <- struct{ID string; Err error}{"",err}
					return;
				}
				defer file.Close();

				fileID,err := database.UploadToGridFS(file,fileHeader.Filename)
				
				if err!=nil{
					majorThumbnail=fileID;
				}
				resultChan <- struct{ID string; Err error}{fileID,err}
				

			}(thumbnailHeader)
		}

		//wait for all uploades to finish
		go func(){
			wg.Wait()
			close(resultChan)
		}()

		//collect results from channel
		for result := range resultChan{
			if result.Err!=nil{
				http.Error(w,result.Err.Error(),http.StatusInternalServerError);
				return;
			}
			if result.ID!=""{
				imageList = append(imageList, result.ID);
			}	
		}

		//store in mongodb
		newShoweModel.ID = primitive.NewObjectID()
		newShoweModel.ImagesOfShowe = imageList;
		newShoweModel.ImageThumbnail = majorThumbnail;
		newShoweModel.ShoweName = showeName;
		newShoweModel.ShoweRating = &showeRatingInt;
		newShoweModel.VendorName = &vendorName;
		newShoweModel.ShoweVotes = &showeVotesInt
		newShoweModel.ShoweDuration = &showeDurationInt;
		newShoweModel.ShoweType = showeType;
		newShoweModel.ShoweGenre = showeGenre;
		newShoweModel.ShoweFromTime = &showeFromTime
		newShoweModel.ShoweEndTime = &showeEndTime
		newShoweModel.ShoweAllowedAges=showeAllowedAges
		newShoweModel.ShoweReleaseDate=showeReleasedDate
		newShoweModel.ShowUserExperience = &showeUserExperience
		newShoweModel.ShoweAboutUs = &showeAboutUs;
		newShoweModel.ShoweTermsConditions = &showeTermsConditions;
		newShoweModel.ShoweNote = &showeNote;
		newShoweModel.ShoweCause = &showeCause;
		newShoweModel.ShoweLanguages = &showeLanguages;
		newShoweModel.ShoweLvl1Crew = &showeCrewMembers;
		newShoweModel.ShoweReviews = &showeReviews;

		//store in mngo
		collectionName := factory.GetCollectionNameByShoweType(showeType)

		mongoCollection :=  database.GetCollectionByName(collectionName)

		result,err := mongoCollection.InsertOne(ctx,newShoweModel);

		if err!=nil{
			log.Println(err.Error());
			http.Error(w,"could not insert showe model",http.StatusInternalServerError);
			return;
		}



		log.Println("Created showe model")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(status{"message":"created successfully","data":result});
		
	}

}


