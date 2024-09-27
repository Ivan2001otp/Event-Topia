package showe

import "go.mongodb.org/mongo-driver/bson/primitive"

type Eventshow struct {
	ID       primitive.ObjectID `bson:"_id"`
	Eventshow_id string             `json:"eventshow_id"`
	BaseshowModel
}

func (mv *Eventshow) SetThumbnailImg(img string) {
	mv.ThumbnailImg = img
}

func (mv *Eventshow) SetBannerImages(imglist []string) {
	mv.BannerImgList = imglist
}