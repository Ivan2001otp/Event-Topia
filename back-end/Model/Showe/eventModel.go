package showe

type Eventshow struct{
	BaseshowModel
}

func (mv *Eventshow)SetThumbnailImg(img string){
	mv.ThumbnailImg = img;
}

func (mv *Eventshow)SetBannerImages(imglist []string){
	mv.BannerImgList = imglist;
}