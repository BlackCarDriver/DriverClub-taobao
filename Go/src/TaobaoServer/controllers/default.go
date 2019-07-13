package controllers

import (
	"TaobaoServer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	imgPath = `E:\tempfile\taobaosource\`
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
type PersonalDataController struct {
	beego.Controller
}

type UpdataMsgController struct {
	beego.Controller
}

type UploadGoodsController struct {
	beego.Controller
}

type UploadImagesController struct {
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

//商品详情获取数据接口
func (this *GoodsDetailController) Post() {
	postBody := models.GoodsPostBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	goodId := postBody.GoodId
	datatype := postBody.DataType
	if goodId == 0 || datatype == "" {
		return
	}
	fmt.Println("GoodsDetail postBody :", postBody)
	switch datatype {
	case "message":
		this.Data["json"] = &models.MockGoodsMessage
	case "detail":
		this.Data["json"] = &models.MockGoodsDetail
	}

	this.ServeJSON()
}

//个人详情页面信息获取接口
func (this *PersonalDataController) Post() {
	postBody := models.PersonalPostBody{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	userName := postBody.Name
	dataTag := postBody.Tag
	if userName == "" || dataTag == "" {
		return
	}
	fmt.Println(userName, " -------------- ", dataTag)
	switch dataTag {
	case "mymsg":
		this.Data["json"] = &models.MockUserMessage
	case "mygoods":
		this.Data["json"] = &models.MockGoodsShort
	case "mycollect":
		this.Data["json"] = &models.MockGoodsShort
	case "message":
		this.Data["json"] = &models.MockMyMessage
	case "rank":
		this.Data["json"] = &models.MockRank
	case "mycare":
		this.Data["json"] = &models.MockCare
	case "naving":
		this.Data["json"] = &models.MockMystatus
	}
	this.ServeJSON()
}

//更新个人信息
func (this *UpdataMsgController) Post() {
	postBody := models.UpdeteMsg{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	updateType := postBody.UpdataType
	switch updateType {
	case "basemsg":
		fmt.Println("Updata base message ...")
		this.Data["json"] = &models.MockUpdateResult
	case "connect":
		fmt.Println("Updata connect ways...")
		this.Data["json"] = &models.MockUpdateResult
	case "headimg":
		fmt.Println("Updata head img ...")
		this.Data["json"] = &models.MockUpdateResult
	}
	this.ServeJSON()
}

//上传商品
func (this *UploadGoodsController) Post() {
	postBody := models.UploadGoodsData{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &postBody); err != nil {
		return
	}
	fmt.Println(postBody)
	this.Data["json"] = &models.MockUpLoadResult
	this.ServeJSON()
}

//上传图片
func (this *UploadImagesController) Post() {
	f, h, err := this.GetFile("file")
	if err != nil {
		return
	}
	defer f.Close()
	err = this.SaveToFile("file", imgPath+h.Filename)
	if err != nil {
		fmt.Println("savetofile error : ", err)
		return
	}
	this.Data["json"] = &models.MockUpLoadResult
	this.ServeJSON()
}
