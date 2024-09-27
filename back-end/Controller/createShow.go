package controller

import (
	"Backend/Database"
	"Backend/Model"
	"Backend/Model/Showe"
	"Backend/Util"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type status map[string]interface{}

func CreateShowController() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
		if(r.Method != http.MethodPost){
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status{"error":"not a post method"})
			return;
		}

		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second);

		
		//take values from form
	   err :=	r.ParseMultipartForm(20<<30)

	   if err!=nil{
		defer cancel();
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
				defer cancel();
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
				defer cancel();
				log.Println(err.Error());
				return;
			}

			// bannerImgList = append(bannerImgList,downloadUrl);
			thumbnailImg = downloadUrl;
			break;
		}

		log.Println("thumbnail image is ",thumbnailImg)

		var parsedBuff showe.Movie;
		

		movie.ShowName = r.FormValue(Util.SHOW_NAME)

		duration,_ := strconv.ParseInt(r.FormValue(Util.SHOW_DURATION),10,64)
		movie.ShowDuration = duration

		rating,_ := strconv.ParseInt(r.FormValue(Util.MOVIE_RATING,),10,64);
		movie.MovieRating = rating

		votes,_ := strconv.ParseInt(r.FormValue(Util.MOVIE_VOTING),10,64);
		movie.MovieVotes = votes;

		movie.ShowAboutUs = strings.Split(r.FormValue(Util.SHOW_ABOUT_US),",");

		//make sure u have list of map in r.formvalue(crewmembers)

		err = json.NewDecoder(r.Body).Decode(&parsedBuff);
		if err!=nil{
			defer cancel();
			http.Error(w,err.Error(),http.StatusBadRequest);
			return;
		}

		movie.ShowCrewMembers = parsedBuff.ShowCrewMembers;


		movie.ShowGenre = r.FormValue(Util.SHOW_GENRE)
		movie.ShowReleaseDate = r.FormValue(Util.SHOW_RELEASE_DATE)
		
		movie.ShowStartTime = r.FormValue(Util.SHOW_START_TIME)
		movie.ShowEndTime = r.FormValue(Util.SHOW_END_TIME)
		movie.ShowVenue = r.FormValue(Util.SHOW_VENUE)

		movie.SetThumbnailImg(thumbnailImg)
		movie.SetBannerImages(bannerImgList)


		//parse it to golang struct
		//store in db
		defer cancel();

		result,err := SaveNewShowe(Util.NEW_SHOWE_COLLECTION,movie)

		if err!=nil{
			log.Println("Something went wrong after saving new created showe!");
			http.Error(w,err.Error(),http.StatusInternalServerError);
			return;
		}
		w.WriteHeader(http.StatusOK);

		json.NewEncoder(w).Encode(status{"message":"success","data":movie})

		return;
		
	}
}