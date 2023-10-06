package controllers

import (
	"GinNGorm/pojo"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

//type Search struct {
//	Name string `json:"name"`
//	Cid  int    `json:"cid"`
//}

func (o OrderController) GetOrderInfo(c *gin.Context) {

	ReturnSuccess(c, 0, "success", "order interface", 1)

}

func (o OrderController) GetOderList(c *gin.Context) {

	//param := make(map[string]interface{})

	search := &pojo.Search{}

	err := c.BindJSON(&search)

	if err == nil {
		ReturnSuccess(c, 0, search.Name, search.Cid, 1)
		return
	}

	ReturnError(c, 404, gin.H{"Err": "err"})

}
