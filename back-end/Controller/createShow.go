package controller

import (
	"Backend/Database"
	"Backend/Model/Showe"
	"Backend/Util"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type status map[string]interface{}

func CreateShowController(w http.ResponseWriter, r *http.Request) {
	// return func(w http.ResponseWriter, r *http.Request) {
		
		if(r.Method != http.MethodPost){
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status{"error":"not a post method"})
			return;
		}


		
		//take values from form
	   err :=	r.ParseMultipartForm(20<<30)

	   if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status{"error":"unable to parse form"})
		return;
	}

	//properties
	var movie showe.Movie

	var bannerImgList []string;
	var thumbnailImg string="";

		multipartFormData := r.MultipartForm;
		
		for _,v:= range multipartFormData.File[Util.BANNER_IMAGES]{
			uploadedFile,_ := v.Open();
			downloadUrl,err :=  database.UploadToGridFS(uploadedFile,v.Filename)

			if err!=nil{
				log.Println(err.Error());
				return;
			}

			bannerImgList = append(bannerImgList,downloadUrl);


		}

		//thumnail
		for _,v := range multipartFormData.File[Util.THUMBNAIL]{
			uploadedFile,_ := v.Open();
			downloadUrl,err := database.UploadToGridFS(uploadedFile,v.Filename)
			if err!=nil{
				log.Println(err.Error());
				return;
			}

			// bannerImgList = append(bannerImgList,downloadUrl);
			thumbnailImg = downloadUrl;
			break;
		}

		log.Println("thumbnail image is ",thumbnailImg)
		//make sure u have list of map in r.formvalue(crewmembers)

		var parsedBuff showe.Movie;

		err = json.NewDecoder(r.Body).Decode(&parsedBuff);
		if err!=nil{
			http.Error(w,err.Error(),http.StatusBadRequest);
			return;
		}

		
		movie.ID = primitive.NewObjectID();
		movie.Movie_id = movie.ID.Hex();

		movie.ShowName = r.FormValue(Util.SHOW_NAME)

		duration,_ := strconv.ParseInt(r.FormValue(Util.SHOW_DURATION),10,64)
		movie.ShowDuration = duration

		rating,_ := strconv.ParseInt(r.FormValue(Util.MOVIE_RATING,),10,64);
		movie.MovieRating = rating

		votes,_ := strconv.ParseInt(r.FormValue(Util.MOVIE_VOTING),10,64);
		movie.MovieVotes = votes;

		movie.ShowAboutUs = strings.Split(r.FormValue(Util.SHOW_ABOUT_US),",");

		

		movie.ShowCrewMembers = parsedBuff.ShowCrewMembers;


		movie.ShowGenre = r.FormValue(Util.SHOW_GENRE)
		movie.ShowReleaseDate = r.FormValue(Util.SHOW_RELEASE_DATE)
		
		movie.ShowStartTime = r.FormValue(Util.SHOW_START_TIME)
		movie.ShowEndTime = r.FormValue(Util.SHOW_END_TIME)
		movie.ShowVenue = r.FormValue(Util.SHOW_VENUE)

		movie.SetThumbnailImg(thumbnailImg)
		movie.SetBannerImages(bannerImgList)

		movie.Updated_date,_	= time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		movie.Created_date,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))

		//store in db

		result,err := database.SaveNewShoweData(Util.NEW_SHOWE_COLLECTION,movie)
		log.Println(result);

		if err!=nil{
			log.Println("Something went wrong after saving new created showe!");
			http.Error(w,err.Error(),http.StatusInternalServerError);
			return;
		}
		w.WriteHeader(http.StatusOK);

		json.NewEncoder(w).Encode(status{"message":"success","id":result,"data":movie})

		return;
		
	// }
}