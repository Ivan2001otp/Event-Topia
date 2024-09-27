package showe

import ()

type Liveshow struct{
	BaseshowModel
}

func (mv *Liveshow)SetThumbnailImg(img string){
	mv.ThumbnailImg = img;
}

func (mv *Liveshow)SetBannerImages(imglist []string){
	mv.BannerImgList = imglist;
}