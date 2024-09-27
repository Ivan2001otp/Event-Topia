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
	SHOW_NAME  string = "show-name"
	SHOW_DURATION string = "show-duration"
	SHOW_GENRE  string = "show-genre"
	SHOW_RELEASE_DATE  string = "show-release-date"
	SHOW_START_TIME string = "show-start-time"
	SHOW_END_TIME string = "show-end-time"
	SHOW_VENUE string = "show-venue"
	SHOW_ABOUT_US string = "show-about-us"
	SHOW_CREW_MEMBERS string = "show-crew-members"
	MOVIE_RATING string = "movie-rating"
	MOVIE_VOTING string = "movie-voting"
	MOVIE_EXPERIENCE string = "movie-experience"
	THUMBNAIL string = "thumb-nail"
	BANNER_IMAGES string = "banner-attachments"
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
