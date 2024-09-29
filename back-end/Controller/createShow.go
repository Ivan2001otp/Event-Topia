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

func CreateMovieController(w http.ResponseWriter, r *http.Request) {
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

		result,err := database.SaveNewMovieData(Util.NEW_MOVIE_COLLECTION,movie)
		log.Println(result);

		if err!=nil{
			log.Println("Something went wrong after saving new created showe!");
			http.Error(w,err.Error(),http.StatusInternalServerError);
			return;
		}
		w.WriteHeader(http.StatusOK);

		json.NewEncoder(w).Encode(status{"message":"success","id":result})		
}


func FetchAllMoviesShowe(w http.ResponseWriter, r *http.Request){

	if( r.Method != http.MethodGet){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status{"error":"supposed to be a GET request"});
		return;
	}


	recordPerPage,err := strconv.Atoi(r.URL.Query().Get("recordPerPage"));

	if err!=nil || recordPerPage<1{
		recordPerPage=10;
	}
	page,err := strconv.Atoi(r.URL.Query().Get("page"));

	if err!=nil || page<1{
		page=1;
	}

	startIndex := (page-1) * recordPerPage;

   allMovieList ,err :=	database.FetchAllMovieShowe(Util.NEW_MOVIE_COLLECTION,startIndex,recordPerPage)

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError);
		json.NewEncoder(w).Encode(status{"error":"could not fetch all movie shows"});
		return;
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status{"status":"success","data":allMovieList,});
	return;
}


//create liveshow
func CreateLiveshowController(w http.ResponseWriter,r *http.Request){
	if (r.Method != http.MethodPost){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status{"error":"supposed to be post method"})
		return;
	}

	err := r.ParseMultipartForm(20<<30);

	if err!=nil{
		log.Println("Parsing a multipart failed")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status{"error":"unable to parse the multipart file"})
		return;
	}

	var liveShow showe.Liveshow
	var parsedCrewList []showe.Crew
	var bannerImgList []string
	var thumbnail string="";

	multipartFormData := r.MultipartForm;

	log.Println("The banner image is ",multipartFormData.File[Util.BANNER_IMAGES]);
	log.Println("The thumbnail images is",multipartFormData.File[Util.THUMBNAIL]);

	for _,v:=range multipartFormData.File[Util.BANNER_IMAGES]{
		uploadedFile,_ := v.Open();
		downloadUrl,err := database.UploadToGridFS(uploadedFile,v.Filename);
		
		if err!=nil{
			log.Println(err.Error());
			return;
		}

		bannerImgList = append(bannerImgList, downloadUrl);
	}

	//thumbnail
	for _,v := range multipartFormData.File[Util.THUMBNAIL]{
		uploadedFile,_ := v.Open();

		downloadUrl,err := database.UploadToGridFS(uploadedFile,v.Filename);

		if err!=nil{
			log.Println(err.Error());
			return;
		}

		thumbnail = downloadUrl;
	}

	log.Println(r.Form)

	jsonCrew := r.Form["show_crew_members"];
	err = json.Unmarshal([]byte(jsonCrew[0]),&parsedCrewList);

	if err!=nil{
		log.Println("wrong while jsonbytes to slice of crew");
		log.Fatal(err);
		return;
	}

	for _,e := range parsedCrewList{
		log.Println("img -> ",e.ImgUrl);
		log.Println("info -> ",e.AboutCrewInfo);
	}

	liveShow.ID = primitive.NewObjectID();
	liveShow.Liveshow_id = liveShow.ID.Hex();

	liveShow.ShowName = r.FormValue(Util.SHOW_NAME);
	duration,_ := strconv.ParseInt(r.FormValue(Util.SHOW_DURATION),10,64);
	liveShow.ShowDuration = duration;

	liveShow.ShowAboutUs = strings.Split(r.FormValue(Util.SHOW_ABOUT_US),",");
	liveShow.ShowType = string(Util.Event)
	liveShow.ShowCrewMembers = parsedCrewList;

	liveShow.ShowGenre = r.FormValue(Util.SHOW_GENRE)
	liveShow.ShowReleaseDate = r.FormValue(Util.SHOW_RELEASE_DATE);
	liveShow.VendorName = r.FormValue(Util.VENDOR_NAME);

	liveShow.ShowStartTime = r.FormValue(Util.SHOW_START_TIME);
	liveShow.ShowEndTime = r.FormValue(Util.SHOW_END_TIME);
	liveShow.ShowVenue = r.FormValue(Util.SHOW_VENUE);
	
	liveShow.SetThumbnailImg(thumbnail)
	liveShow.SetBannerImages(bannerImgList)

	liveShow.Updated_date,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339));
	liveShow.Created_date,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339));

	result,err := database.SaveNewLiveshowData(Util.NEW_EVENT_COLLECTION,liveShow)
	log.Println(result);

	if err!=nil{
		log.Println("Something went wrong after saving new created live show");

		http.Error(w,err.Error(),http.StatusInternalServerError);
		return;
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status{"message":"success","id":result});
}

//create activity
func CreateActivityController(w http.ResponseWriter,r *http.Request){
	if (r.Method != http.MethodPost){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status{"error":"supposed to be post method"})
		return;
	}

	err := r.ParseMultipartForm(20<<30);

	if err!=nil{
		log.Println("Parsing a multipart failed")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status{"error":"unable to parse the multipart file"})
		return;
	}

	var activity showe.ActivityShow
	var parsedCrewList []showe.Crew
	var bannerImgList []string
	var thumbnail string="";

	multipartFormData := r.MultipartForm;

	log.Println("The banner image is ",multipartFormData.File[Util.BANNER_IMAGES]);
	log.Println("The thumbnail images is",multipartFormData.File[Util.THUMBNAIL]);

	for _,v:=range multipartFormData.File[Util.BANNER_IMAGES]{
		uploadedFile,_ := v.Open();
		downloadUrl,err := database.UploadToGridFS(uploadedFile,v.Filename);
		
		if err!=nil{
			log.Println(err.Error());
			return;
		}

		bannerImgList = append(bannerImgList, downloadUrl);
	}

	//thumbnail
	for _,v := range multipartFormData.File[Util.THUMBNAIL]{
		uploadedFile,_ := v.Open();

		downloadUrl,err := database.UploadToGridFS(uploadedFile,v.Filename);

		if err!=nil{
			log.Println(err.Error());
			return;
		}

		thumbnail = downloadUrl;
	}

	log.Println(r.Form)

	jsonCrew := r.Form["show_crew_members"];
	err = json.Unmarshal([]byte(jsonCrew[0]),&parsedCrewList);

	if err!=nil{
		log.Println("wrong while jsonbytes to slice of crew");
		log.Fatal(err);
		return;
	}

	for _,e := range parsedCrewList{
		log.Println("img -> ",e.ImgUrl);
		log.Println("info -> ",e.AboutCrewInfo);
	}

	activity.ID = primitive.NewObjectID();
	activity.Activity_id = activity.ID.Hex();

	activity.ShowName = r.FormValue(Util.SHOW_NAME);
	duration,_ := strconv.ParseInt(r.FormValue(Util.SHOW_DURATION),10,64);
	activity.ShowDuration = duration;

	activity.ShowAboutUs = strings.Split(r.FormValue(Util.SHOW_ABOUT_US),",");
	activity.ShowType = string(Util.Event)
	activity.ShowCrewMembers = parsedCrewList;

	activity.ShowGenre = r.FormValue(Util.SHOW_GENRE)
	activity.ShowReleaseDate = r.FormValue(Util.SHOW_RELEASE_DATE);
	activity.VendorName = r.FormValue(Util.VENDOR_NAME);

	activity.ShowStartTime = r.FormValue(Util.SHOW_START_TIME);
	activity.ShowEndTime = r.FormValue(Util.SHOW_END_TIME);
	activity.ShowVenue = r.FormValue(Util.SHOW_VENUE);
	
	activity.SetThumbnailImg(thumbnail)
	activity.SetBannerImages(bannerImgList)

	activity.Updated_date,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339));
	activity.Created_date,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339));

	result,err := database.SaveNewActivityData(Util.NEW_EVENT_COLLECTION,activity)
	log.Println(result);

	if err!=nil{
		log.Println("Something went wrong after saving new created activity");

		http.Error(w,err.Error(),http.StatusInternalServerError);
		return;
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status{"message":"success","id":result});
}

//create event 
func CreateEventController(w http.ResponseWriter,r *http.Request) {
	if (r.Method != http.MethodPost){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status{"error":"supposed to be post method"})
		return;
	}

	err := r.ParseMultipartForm(20<<30);

	if err!=nil{
		log.Println("Parsing a multipart failed")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status{"error":"unable to parse the multipart file"})
		return;
	}

	var event showe.Eventshow
	var parsedCrewList []showe.Crew
	var bannerImgList []string
	var thumbnail string="";

	multipartFormData := r.MultipartForm;

	log.Println("The banner image is ",multipartFormData.File[Util.BANNER_IMAGES]);
	log.Println("The thumbnail images is",multipartFormData.File[Util.THUMBNAIL]);

	for _,v:=range multipartFormData.File[Util.BANNER_IMAGES]{
		uploadedFile,_ := v.Open();
		downloadUrl,err := database.UploadToGridFS(uploadedFile,v.Filename);
		
		if err!=nil{
			log.Println(err.Error());
			return;
		}

		bannerImgList = append(bannerImgList, downloadUrl);
	}

	//thumbnail
	for _,v := range multipartFormData.File[Util.THUMBNAIL]{
		uploadedFile,_ := v.Open();

		downloadUrl,err := database.UploadToGridFS(uploadedFile,v.Filename);

		if err!=nil{
			log.Println(err.Error());
			return;
		}

		thumbnail = downloadUrl;
	}

	log.Println(r.Form)

	jsonCrew := r.Form["show_crew_members"];
	err = json.Unmarshal([]byte(jsonCrew[0]),&parsedCrewList);

	if err!=nil{
		log.Println("wrong while jsonbytes to slice of crew");
		log.Fatal(err);
		return;
	}

	for _,e := range parsedCrewList{
		log.Println("img -> ",e.ImgUrl);
		log.Println("info -> ",e.AboutCrewInfo);
	}

	event.ID = primitive.NewObjectID();
	event.Eventshow_id = event.ID.Hex();

	event.ShowName = r.FormValue(Util.SHOW_NAME);
	duration,_ := strconv.ParseInt(r.FormValue(Util.SHOW_DURATION),10,64);
	event.ShowDuration = duration;

	event.ShowAboutUs = strings.Split(r.FormValue(Util.SHOW_ABOUT_US),",");
	event.ShowType = string(Util.Event)
	event.ShowCrewMembers = parsedCrewList;

	event.ShowGenre = r.FormValue(Util.SHOW_GENRE)
	event.ShowReleaseDate = r.FormValue(Util.SHOW_RELEASE_DATE);
	event.VendorName = r.FormValue(Util.VENDOR_NAME);

	event.ShowStartTime = r.FormValue(Util.SHOW_START_TIME);
	event.ShowEndTime = r.FormValue(Util.SHOW_END_TIME);
	event.ShowVenue = r.FormValue(Util.SHOW_VENUE);
	
	event.SetThumbnailImg(thumbnail)
	event.SetBannerImages(bannerImgList)

	event.Updated_date,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339));
	event.Created_date,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339));

	result,err := database.SaveNewEventData(Util.NEW_EVENT_COLLECTION,event)
	log.Println(result);

	if err!=nil{
		log.Println("Something went wrong after saving new created showe");

		http.Error(w,err.Error(),http.StatusInternalServerError);
		return;
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status{"message":"success","id":result});
	
}