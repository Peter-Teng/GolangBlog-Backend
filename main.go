package main

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/model"
	"fmt"
)

func main() {
	visitor := model.Visitor{Id: "3"}
	if err := common.Db.Delete(&visitor).Error; err != nil{
		fmt.Println(err)
	}
	//gin.SetMode(common.AppMode)
	//router := common.SetupRouter()
	//router.Run(common.ServerPort)
}
