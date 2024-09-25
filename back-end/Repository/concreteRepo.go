package repository

import (
	"Backend/SuperModel"
	"Backend/Model"
)

func createNewMovieModel(
	imageList []string,
	imageThumbnail string,
	movieShowName string,
	movieShowRating *int32,

	movieShowUpvotes *int64,
	movieShowDuration *int32,
	//set the show type manually
	showTye string,
	movieShowGenre []string,	
	movieShowStartTime *string,
	movieShowEndTime *string,
	
	movieShowWatchAllowedAges string,
	releaseDate string,
	movieShowUx *string,
	
	notes *string,
	aboutUsList *[]string,
	termsList *[]string,

	causeToStream *string,
	availableInLanguages *[]string,
	crewList *[]super_model.CrewMember,
	reviewList *[]super_model.Review,
) super_model.IShow{
	return &model.MovieModel{}
} 

func createEventShowModel(
	imageList []string,
	imageThumbnail string,
	eventShowName string,
	eventShowRating *int32,

	eventShowUpvotes *int64,
	eventShowDuration *int32,
	//set the show type manually
	showTye string,
	eventShowGenre []string,	
	eventShowStartTime *string,
	eventShowEndTime *string,
	
	eventShowWatchAllowedAges string,
	releaseDate string,
	eventShowUx *string,
	
	notes *string,
	aboutUsList *[]string,
	termsList *[]string,

	causeToStream *string,
	availableInLanguages *[]string,
	crewList *[]super_model.CrewMember,
	reviewList *[]super_model.Review,
) super_model.IShow{
	return &model.EventModel{}
}

func createNewLiveShowModel(
		
	imageList []string,
	imageThumbnail string,
	liveShowName string,
	liveShowRating *int32,

	liveShowUpvotes *int64,
	liveShowDuration *int32,
	//set the show type manually
	showTye string,
	liveShowGenre []string,	
	liveShowStartTime *string,
	liveShowEndTime *string,
	
	liveShowWatchAllowedAges string,
	releaseDate string,
	liveShowUx *string,
	
	notes *string,
	aboutUsList *[]string,
	termsList *[]string,

	causeToStream *string,
	availableInLanguages *[]string,
	crewList *[]super_model.CrewMember,
	reviewList *[]super_model.Review,
)super_model.IShow{
	return &model.LiveShowModel{};
}

func createNewActivityModel(
		
	imageList []string,
	imageThumbnail string,
	activityName string,
	activityRating *int32,

	activityUpvotes *int64,
	activityDuration *int32,
	//set the show type manually
	showTye string,
	activityGenre []string,	
	activityStartTime *string,
	activityEndTime *string,
	
	activityWatchAllowedAges string,
	releaseDate string,
	eventUx *string,
	
	notes *string,
	aboutUsList *[]string,
	termsList *[]string,

	causeToStream *string,
	availableInLanguages *[]string,
	crewList *[]super_model.CrewMember,
	reviewList *[]super_model.Review,
)super_model.IShow{
	return &model.ActivityModel{};
}