package showe 

import ("time")
type Crew struct {
	ImgUrl string;
	AboutCrewInfo string;
};



type BaseshowInterface interface{

	SetBannerImgList([]string)
	SetThumbnailImg(string)
	SetshowName(string)
	SetshowDuration(int64)
	SetshowGenre(string)
	SetshowReleaseDate(string)
	SetstartTime(string)
	SetendTime(string);
	Setvenue(string);
	SetaboutShow([]string);
	SetCrewMembers([]Crew);

	GetBannerImgList() []string;
	GetThumbnailImg() string;
	GetshowName() string
	GetshowDuration() int64
	GetshowGenre() string
	GetshowReleaseDate() string
	GetshowStartTime()string
	GetshowEndtime()string
	Getvenue()string
	GetaboutShow()[]string
	GetCrewMembers()[]Crew
}

type BaseshowModel struct{
	BannerImgList []string `json:"banner_img_list"`
	ThumbnailImg string		`json:"thumbnail_img"`
	ShowName string	`json:"show_name"`;
	ShowDuration int64	`json:"show_duration"`;
	ShowGenre string	`json:"show_genre"`
	ShowReleaseDate string	`json:"show_release_date"`
	ShowStartTime string	`json:"show_start_time"`
	ShowEndTime string	`json:"show_end_time"`
	ShowVenue string	`json:"show_venue"`
	ShowAboutUs []string	`json:"show_about_us"`
	ShowCrewMembers []Crew	`json:"show_crew_members"`
	Created_date time.Time `json:"created_at"`
	Updated_date time.Time `json:"updated_at"`
}