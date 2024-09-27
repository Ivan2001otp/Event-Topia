package showe

import ("go.mongodb.org/mongo-driver/bson/primitive")

type Liveshow struct{
	ID primitive.ObjectID `bson:"_id"`
	Liveshow_id string `json:"liveshow_id"`
	BaseshowModel
}

func (mv *Liveshow)SetThumbnailImg(img string){
	mv.ThumbnailImg = img;
}

func (mv *Liveshow)SetBannerImages(imglist []string){
	mv.BannerImgList = imglist;
}