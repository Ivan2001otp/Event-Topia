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
