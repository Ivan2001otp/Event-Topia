package showe

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct{
	BaseshowModel
	ID primitive.ObjectID `bson:"_id"`
	Movie_id string `json:"movie_id"`
	

	MovieRating int64 `json:"movie_rating"`
	MovieVotes int64	`json:"movie_votes"`
	MovieExperience string	`json:"movie_experience"`
}


//getters and setters
func (mv *Movie)SetThumbnailImg(img string){
	mv.ThumbnailImg = img;
}

func (mv *Movie)SetBannerImages(imglist []string){
	mv.BannerImgList = imglist;
}

func (mv *Movie)SetMovieRating(rating int64){
	mv.MovieRating = rating;
}
func (mv *Movie)SetMovieVotes(votes int64){
	mv.MovieVotes = votes;
}
func (mv *Movie)SetMovieExperience( exp string){
	mv.MovieExperience = exp;
}