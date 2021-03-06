package routers

import (
	"TaobaoServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//setting up logger
	beego.SetLogger("file", `{"filename":"logs/default.log","daily":false, "maxsize":512000}`)
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/homepage/goodsdata", &controllers.HPGoodsController{})
	beego.Router("/homepage/goodstypemsg", &controllers.GoodsTypeController{})
	beego.Router("/goodsdeta", &controllers.GoodsDetailController{})
	beego.Router("/personal/data", &controllers.PersonalDataController{})
	beego.Router("/update", &controllers.UpdataMsgController{})
	beego.Router("/upload/newgoods", &controllers.UploadGoodsController{})
	beego.Router("/upload/images", &controllers.UploadImagesController{})
	beego.Router("/entrance", &controllers.EntranceController{})
	beego.Router("/smallupdate", &controllers.UpdateController{})
	beego.Router("/deleteapi", &controllers.DeleteController{})
	beego.Router("/public", &controllers.PublicController{})
	beego.Router("/postform", &controllers.PostFormController{})
}
