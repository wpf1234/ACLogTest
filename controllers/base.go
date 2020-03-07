package controllers

import (
	"ACTest/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type baseController struct {
	beego.Controller
}

func (c *baseController) Prepare() {
	//c.TokenAuth()
	//c.TokenExist()
	//c.TokenUnmarshal()
	beego.Info("method:", c.Ctx.Request.Method,
		" url:", c.Ctx.Request.RequestURI,
		" data:", string(c.Ctx.Input.RequestBody))
}

//验证token是否合法,调用接口
func (c *baseController) TokenAuth() {
	tkn := c.Ctx.Input.Header("authorization")
	url := beego.AppConfig.String("tokenauth")
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		beego.Error("验证token合法————声明请求出错：", err.Error())
	}
	request.Header.Set("Authorization", tkn)
	beego.Info("token: ", tkn)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		beego.Error("验证token合法————获取响应出错：", err.Error())
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		beego.Error("验证token合法————解析body出错：", err.Error())
		return
	}
	var resp models.TokenResp
	json.Unmarshal(body, &resp)
	if resp.Success != true {
		beego.Error("tokenauth:", resp.Msg)
		c.Ctx.Output.Status = 403
		c.Data["json"] = models.Result{Code: 403, Message: resp.Msg, Data: []struct{}{}}
		c.ServeJSON()
		return
	}
}

//验证token是否存在

func (c *baseController) TokenExist() {

	tkn := c.Ctx.Input.Header("authorization")
	url := beego.AppConfig.String("tokenexist")

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {

		beego.Error("验证token存在————声明请求出错：", err.Error())
		return
	}

	request.Header.Set("Authorization", tkn)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {

		beego.Error("验证token存在————获取响应出错：", err.Error())
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {

		beego.Error("验证token存在————解析body出错：", err.Error())
		return
	}
	var resp models.TokenResp
	json.Unmarshal(body, &resp)
	if resp.Success != true {
		beego.Error(`tokenexist:`, resp.Msg)
		c.Ctx.Output.Status = 403
		c.Data["json"] = models.Result{Code: 403, Message: resp.Msg, Data: []struct{}{}}
		c.ServeJSON()
		return

	}
}

//解析token
func (c *baseController) TokenUnmarshal() {
	tkn := c.Ctx.Input.Header("authorization")
	url := beego.AppConfig.String("token")

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {

		beego.Error("解析token————声明请求出错：", err.Error())
		return
	}

	request.Header.Set("Authorization", tkn)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {

		beego.Error("解析token————获取响应出错：", err.Error())
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {

		beego.Error("解析token————解析body出错：", err.Error())
		return
	}
	var resp models.TokenInfo
	json.Unmarshal(body, &resp)
	if resp.Code == 200 {
		beego.Info(resp.Data.Name)
	}
}
