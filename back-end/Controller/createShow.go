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
		log.Println("Parsing a multipart failed");
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status{"error":"unable to parse form"})
		return;
	}

	//properties
	var movie showe.Movie
	var parsedCrewList []showe.Crew;
	var bannerImgList []string;
	var thumbnailImg string="";

		multipartFormData := r.MultipartForm;

		log.Println("The banner images is ",multipartFormData.File[Util.BANNER_IMAGES])
		log.Println("The thumbnail images ",multipartFormData.File[Util.THUMBNAIL])
		
		//banner images
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
			
		}

		//use r.Form["key"]
		log.Println(r.Form)

		//crafting the list of crew members.
		jsonCrew := r.Form["show_crew_members"]

		//convert json bytes to crew - list struct
		err = json.Unmarshal([]byte(jsonCrew[0]),&parsedCrewList);

		if err!=nil{
			log.Println("wrong while jsonbytes to slice of crew")
			log.Fatal(err);
		}
		
		for _,e := range parsedCrewList{
			log.Println("Img ->",e.ImgUrl);
			log.Println("Info ->",e.AboutCrewInfo)
		}
	
		
		movie.ID = primitive.NewObjectID();
		movie.Movie_id = movie.ID.Hex();

		movie.ShowName = r.FormValue(Util.SHOW_NAME)

		duration,_ := strconv.ParseInt(r.FormValue(Util.SHOW_DURATION),10,64)
		movie.ShowDuration = duration

		rating,_ := strconv.ParseInt(r.FormValue(Util.MOVIE_RATING,),10,64);
		movie.MovieRating = rating

		log.Println("The voting value is ",r.Form[Util.MOVIE_VOTING])
		votes,_ := strconv.ParseInt(r.FormValue(Util.MOVIE_VOTING),10,64);
		movie.MovieVotes = votes;

		movie.ShowAboutUs = strings.Split(r.FormValue(Util.SHOW_ABOUT_US),",");
		movie.ShowType =(string(Util.Movie))
		

		movie.ShowCrewMembers = parsedCrewList;


		movie.ShowGenre = r.FormValue(Util.SHOW_GENRE)
		movie.ShowReleaseDate = r.FormValue(Util.SHOW_RELEASE_DATE)
		movie.VendorName = r.FormValue(Util.VENDOR_NAME)
		
		movie.ShowStartTime = r.FormValue(Util.SHOW_START_TIME)
		movie.ShowEndTime = r.FormValue(Util.SHOW_END_TIME)
		movie.ShowVenue = r.FormValue(Util.SHOW_VENUE)
		movie.MovieExperience = r.FormValue(Util.MOVIE_EXPERIENCE)

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

		json.NewEncoder(w).Encode(status{"message":"success","id":result})		
}