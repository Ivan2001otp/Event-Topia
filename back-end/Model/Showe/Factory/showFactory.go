package factory

import (
	show "Backend/Model/Showe"
	"Backend/Util"
	"fmt"
)

func CreateShowVariantByFactory(showType string,baseShowStructParams show.BaseshowModel,) interface{}{
	switch(showType){

	case string(Util.Movie):
		return &show.Movie{
			BaseshowModel: baseShowStructParams,
			Movie_rating: 3,
			Movie_votes: 100,
			Movie_experience:"2D,3D" ,
		}

	case string(Util.Activity):
		return &show.ActivityShow{
			BaseshowModel: baseShowStructParams,
		}

	case string(Util.Event):
		return &show.Eventshow{
			BaseshowModel: baseShowStructParams,
		}

	case string(Util.LiveShow):
		return &show.Liveshow{
			BaseshowModel: baseShowStructParams,
		}
	}

	return fmt.Errorf("The instance does not exist!");
}