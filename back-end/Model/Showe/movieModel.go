package showe

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct{
	ID primitive.ObjectID `bson:"_id"`
	Showe_id string `json:"showe_id"`
	BaseshowModel
	Show_release_date string	`json:"show_release_date"`
	Movie_rating int64 `json:"movie_rating"`
	Movie_votes int64	`json:"movie_votes"`
	Vendor_name string `json:"vendor_name"`
	Movie_experience string	`json:"movie_experience"`
}


//getters and setters
func (mv *Movie)SetThumbnailImg(img string){
	mv.ThumbnailImg = img;
}

func (mv *Movie)SetBannerImages(imglist []string){
	mv.BannerImgList = imglist;
}

func (mv *Movie)SetMovieRating(rating int64){
	mv.Movie_rating = rating;
}
func (mv *Movie)SetMovieVotes(votes int64){
	mv.Movie_votes = votes;
}
func (mv *Movie)SetMovieExperience( exp string){
	mv.Movie_experience = exp;
}