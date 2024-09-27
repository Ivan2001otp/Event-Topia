package showe

import (
)

type ActivityShow struct {
	BaseshowModel
}

func (mv *ActivityShow)SetThumbnailImg(img string){
	mv.ThumbnailImg = img;
}

func (mv *ActivityShow)SetBannerImages(imglist []string){
	mv.BannerImgList = imglist;
}