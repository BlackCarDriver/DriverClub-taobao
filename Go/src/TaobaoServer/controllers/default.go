package controllers

import (
	"TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type HPGoodsController struct {
	beego.Controller
}
type GoodsTypeController struct {
	beego.Controller
}
type GoodsDetailController struct {
	beego.Controller
}

//自带例子
func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

//主页商品封面数据
func (this *HPGoodsController) Post() {
	PostBody := models.PostBody1{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &PostBody); err != nil {
		return
	}
	//需要根据不同的类型，标签和页数从数据库中获取真实数据
	fmt.Println(PostBody.GoodsTag, "------------------", PostBody.GoodsIndex)
	this.Data["json"] = &models.MockGoodsData
	this.ServeJSON()
}

//返回商品分类和标签列表数据
func (this *GoodsTypeController) Get() {
	//需要从数据库获取真实数据返回
	this.Data["json"] = &models.MockTypeData
	this.ServeJSON()
}

func (this *GoodsDetailController) Get() {
	//需要从请求主体获取商品参数再从数据库取对应信息
	//需要从数据库获取真实数据返回
	this.Data["json"] = &models.MockGoodsDetail
	this.ServeJSON()
}
