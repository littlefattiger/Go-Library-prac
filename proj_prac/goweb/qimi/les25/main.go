package main

import (
	"les25/dao"
	"les25/models"
	"les25/routers"
)

func main() {
	err := dao.InitMYSQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":9090")
}
