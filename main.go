package main

import (
	"log"

	"github.com/call-forwarding/config"
	"github.com/call-forwarding/database"
	"github.com/call-forwarding/server"
)

func main() {
	conf := config.GetConfig()
	log.Printf("config Server :: %+v", *conf.Server)
	log.Printf("config Db :: %+v", *conf.Db)
	db := database.NewHarperDatabase(conf)

	server.NewEchoServer(conf, db).Start()

	// =====================================
	// cl := database.NewHarperDatabase(conf).GetDb()
	// // dt := []models.CallForwardStatusResp{}
	// dt := []map[string]interface{}{}

	// searchAttribute := "phoneNumber"

	// attributeList := []string{"*"}

	// searchValue := "+12345678"

	// err := cl.SearchByValue("data", "callforwards", &dt, harperdb.Attribute(searchAttribute), searchValue, harperdb.AttributeList(attributeList))
	// if err != nil {
	// 	log.Println("Something went wrong while reading list", err.Error())
	// 	return
	// }

	// log.Println("=====> DT ", dt)
	// err := cl.CreateDatabase("dev-1")
	// if err != nil {
	// 	log.Println("Error while creating database :: ", err.Error())
	// 	return
	// }
	// log.Println("dev-1 db successfully created!")

	// err = cl.DropDatabase("dev-1")
	// if err != nil {
	// 	log.Println("Error while dropping database :: ", err.Error())
	// 	return
	// }
	// log.Println("dev-1 db successfully dropped!")

}
