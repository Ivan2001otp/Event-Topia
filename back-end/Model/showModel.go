package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// "Backend/factory"

type Review struct{
	ReviewImg *string `json:"review_img"`;
	Message string	`json:"message"`;
}

type CrewMember struct{
	CrewImageUrl string `json:"image_url"`;
	AboutCrew string	`json:"about_crew"`;
}


type BaseShowe struct{

	ID primitive.ObjectID 	`bson:"_id"`
	ImagesOfShowe []string  `json:"images_of_showe"`;
	ImageThumbnail string	`json:"image_of_thumbnail`;
	ShoweName string		`json:"showe_name"`;	
	ShoweRating *int64		`json:"showe_rating"`;

	VendorName *string		`json:"vendor_name"`;

	ShoweVotes *int64		`json:"showe_votes`;
	ShoweDuration *int64	`json:"showe_duration"`;
	ShoweType string		`json:"showe_type"`
	ShoweGenre []string		`json:"showe_genre"`

	ShoweFromTime *string	`json:"showe_from_time`
	ShoweEndTime *string	`json:"showe_end_time"`
	ShoweAllowedAges string	`json:"showe_allowed_ages"`

	ShoweReleaseDate string	`json:"showe_release_date"`
	ShowUserExperience *string	`json:"show_user_experience"`

	ShoweNote *string	`json:"showe_note"`
	ShoweAboutUs *[]string	`json:"showe_about_us"`
	ShoweTermsConditions *[]string	`json:"showe_terms_conditions`

	ShoweCause *string `json:"showe_cause"`;
	ShoweLanguages *[]string `json:"showe_languages`;
	ShoweLvl1Crew *[]CrewMember `json:"crew_members"`;
	ShoweReviews *[]Review `json:"showe_reviews`;	


}

/*
//getters and setters
func(mv *BaseShowe)setVendorName(vendorName string){
	mv.vendorName = &vendorName;
}
func (mv *BaseShowe)getVendorName() string{
	return *mv.vendorName;
}
func (mv *BaseShowe) setListOfShoweImage(imageList []string){
	mv.imagesOfShowe = imageList;
}
func (mv *BaseShowe)setShoweType(showType string){
	mv.showeType = showType;
}
func (mv *BaseShowe) setShoweImageThumbnail(imageThumnail string){
	mv.imageThumbnail = imageThumnail;
}
func (mv *BaseShowe) setShoweName(showName string){
	mv.showeName = showName;
}
func (mv *BaseShowe)setRating(rating int64){
	mv.showeRating = &rating;
}
func (mv *BaseShowe)setVotes(votes int64){
	mv.showeVotes = &votes;
}
func(mv *BaseShowe)setDuration(duration int64){
	mv.showeDuration = &duration;
}
func (mv *BaseShowe)setShoweGenre(genreList []string){
	mv.showeGenre = genreList
}
func (mv *BaseShowe)setShoweStartTime(starTime string){
	mv.showeFromTime = &starTime;
}
func (mv *BaseShowe)setShoweEndTime(endTime string){
	mv.showeEndTime = &endTime;
}
func (mv *BaseShowe)setShoweAllowedAges(ageString string){
	mv.showeAllowedAges = ageString;
}
func (mv *BaseShowe)setShoweReleasedDate(releaseDate string){
	mv.showeReleaseDate = releaseDate;
}
func (mv *BaseShowe)setShoweUX(ux string){
	mv.showUserExperience = &ux;
}
func (mv *BaseShowe)setShoweNote(notes string){
	mv.showeNote = &notes;
}
func (mv *BaseShowe)setShoweAbtUs(aboutUsList []string){
	mv.showeAboutUs = &aboutUsList;
}
func (mv *BaseShowe)setShoweTermsConditions(termsList []string){
	mv.showeTermsConditions = &termsList;
}
func (mv *BaseShowe) setShoweReviews(reviewList []Review){
	mv.showeReviews = &reviewList;
}
func (mv *BaseShowe)setShoweCrewMembers(crewList []CrewMember){
	mv.showeLvl1Crew = &crewList;
}
func (mv *BaseShowe)setShoweLanguages(langList []string){
	mv.showeLanguages = &langList;
}
func (mv *BaseShowe)setShoweCause(cause string){
	mv.showeCause = &cause;
}*/