package showe

import ("go.mongodb.org/mongo-driver/bson/primitive")

type Liveshow struct{
	ID primitive.ObjectID `bson:"_id"`
	Showe_id string `json:"showe_id"`
	Vendor_name string `json:"vendor_name"`
	BaseshowModel
}

func (mv *Liveshow)SetThumbnailImg(img string){
	mv.ThumbnailImg = img;
}

func (mv *Liveshow)SetBannerImages(imglist []string){
	mv.BannerImgList = imglist;
}