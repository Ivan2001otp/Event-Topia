package showe

import "go.mongodb.org/mongo-driver/bson/primitive"

type ActivityShow struct {
	ID primitive.ObjectID `bson:"_id"`
	Activity_id string `json:"activity_id"`
	Vendor_name string `json:"vendor_name"`
	BaseshowModel
}

func (mv *ActivityShow)SetThumbnailImg(img string){
	mv.ThumbnailImg = img;
}

func (mv *ActivityShow)SetBannerImages(imglist []string){
	mv.BannerImgList = imglist;
}