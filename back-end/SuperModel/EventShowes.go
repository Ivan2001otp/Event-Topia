package super_model;

type IShow interface{

	setListOfShoweImage(list []string)
	setShoweImageThumbnail(url string)
	setShoweName(showename string)
	setRating(rating int32)
	setVotes(votes int64)
	setDuration(duration int32)

	setShoweType(showeType string)
	setShoweGenre(genre []string)
	setShoweStartTime( startTime string)
	setShoweEndTime(endTime string)

	setShoweAllowedAges(ages string)//string
	setShoweReleasedDate(date string)
	setShoweUX(ux string);

	setShoweNote(notes string);
	setShoweTermsConditions(t_c []string)
	setShoweAbtUs(aboutUs []string);
	setShoweReviews(reviewList []Review)

	setShoweCrewMembers(memberList []CrewMember);
	setShoweCause(cause string);
	setShoweLanguages(languages []string);
}


type BaseShowe struct{
	imagesOfShowe []string;
	imageThumbnail string;
	showeName string;
	showeRating *int32

	showeVotes *int64;
	showeDuration *int32;
	showeType string
	showeGenre []string

	showeFromTime *string
	showeEndTime *string
	showeAllowedAges string

	showeReleaseDate string
	showUserExperience *string

	showeNote *string
	showeAboutUs *[]string
	showeTermsConditions *[]string

	showeCause *string
	showeLanguages *[]string
	showeLvl1Crew *[]CrewMember
	// showeLvl2Crew *[]CrewMember;
	showeReviews *[]Review;
	
}

//models of lvl1crew,lvl2crew
type Review struct{
	reviewImg *string;
	message string;
}

type CrewMember struct{
	imageUrl string;
	aboutCrew string;
}


//setters and getters.
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
func (mv *BaseShowe)setRating(rating int32){
	mv.showeRating = &rating;
}
func (mv *BaseShowe)setVotes(votes int64){
	mv.showeVotes = &votes;
}
func(mv *BaseShowe)setDuration(duration int32){
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
}