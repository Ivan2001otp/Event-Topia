package Util

type Show string

const (
	Movie Show = "movie"
	Event Show = "event"
	LiveShow Show = "liveshow"
	Activity Show = "activity"
)

const (
	MovieCollection Show = "movie-collection"
	EventCollection Show = "event-collection"
	LiveShowCollection Show = "live-show-collection"
	ActivityCollection Show = "activity-collection"
)

const (
	SHOW_NAME  string = "show_name"
	SHOW_DURATION string = "show_duration"
	SHOW_GENRE  string = "show_genre"
	SHOW_RELEASE_DATE  string = "show_release_date"
	SHOW_START_TIME string = "show_start_time"
	SHOW_END_TIME string = "show_end_time"
	SHOW_VENUE string = "show_venue"
	SHOW_ABOUT_US string = "show_about_us"
	SHOW_CREW_MEMBERS string = "show_crew_members"
	MOVIE_RATING string = "movie_rating"
	MOVIE_VOTING string = "movie_votes"
	MOVIE_EXPERIENCE string = "movie_experience"
	THUMBNAIL string = "thumbnail"
	BANNER_IMAGES string = "banner_attachments"
	VENDOR_NAME string = "vendor_name"
	
)

const (
	NEW_SHOWE_COLLECTION = "new-created-showes"
)


func GetCollectionNameByShoweType(showeType string) string{
	switch showeType {
	case string(Movie):
		return string(MovieCollection);
	case string(Event):
		return string(EventCollection);
	case string(LiveShow):
		return string(LiveShowCollection);
	case string(Activity):
		return string(ActivityCollection);
	}
	return "";
}

//utils
func SetShowDurationDynamically(eventType string) string{

	switch(eventType){
	case string(Movie):
		return string(Movie);

	case string(Event):
		return string(Event);

	case string(LiveShow):
		return string(LiveShow)

	case string(Activity):
		return string(Activity)

	default:
		return "";
	}

}
