package controllers

import (
	"lian/util"
)

type MainController struct {
	BaseController
}

var queryp = util.P{}

func (c *MainController) Get() {

	c.TplName = "form.html"

}
func (c *MainController)Getdata(){
	//page_size
	page_size := 10
	explorerSearch := c.GetString("explorerSearch")
	page := c.GetString("page")
	curlpage := util.ToInt(page)
	if curlpage < 1 {
		curlpage = 1
	}
	if len(explorerSearch) > 0 {
		queryp["data"] = explorerSearch
	} else {
		delete(queryp, "data")
	}
	datalist := util.D("test").Find(queryp).Page((curlpage-1)*page_size, page_size-1).AllData()
	totalcount := util.D("test").Find(queryp).Count()

	c.Data["json"] = map[string]interface{}{"totalcount":totalcount,"data":datalist,"culpage":page}
	c.ServeJSON()
	return

}