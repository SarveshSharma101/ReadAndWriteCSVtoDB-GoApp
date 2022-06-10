package service

import (
	shops "GoServerSaveToDB-App/DataModel/Shop"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbUrl = "?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

//define the structure of the expected request body for each of the shop
type (
	BakeryShopReqBody struct {
		Type string        `json:"type"`
		Data *shops.Bakery `json:"data"`
	}
	Mobile struct {
		Type string        `json:"type"`
		Data *shops.Mobile `json:"data"`
	}
	Electronic struct {
		Type string            `json:"type"`
		Data *shops.Electronic `json:"data"`
	}
	SuperMarket struct {
		Type string             `json:"type"`
		Data *shops.SuperMarket `json:"data"`
	}
)

//Initialize the db connect and auto-migrate the table schema
func InitDBConnection(uname, password, url, dbname string) {
	dbUrl = uname + ":" + password + "@" + url + dbname + dbUrl

	var err error
	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	DB.AutoMigrate(&shops.Bakery{})
	DB.AutoMigrate(&shops.Electronic{})
	DB.AutoMigrate(&shops.Mobile{})
	DB.AutoMigrate(&shops.SuperMarket{})
}

//handle the post req for saving bakery shop data and save the data to DB.
//If there is an error while saving send back error response with status code as 400(Bad-request)
func SaveBakeryShopData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var reqBody BakeryShopReqBody
	json.NewDecoder(r.Body).Decode(&reqBody)
	if err := DB.Create(reqBody.Data).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("There was some error while saving Bakery Shop data")
		// log.Fatal(err)
	}

	json.NewEncoder(w).Encode(reqBody)
}

//handle the post req for saving mobile shop data and save the data to DB.
//If there is an error while saving send back error response with status code as 400(Bad-request)
func SaveMobileShopData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var reqBody Mobile
	json.NewDecoder(r.Body).Decode(&reqBody)
	if err := DB.Create(reqBody.Data).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("There was some error while saving Mobile Shop data")
		// log.Fatal(err)
	}

	json.NewEncoder(w).Encode(reqBody)
}

//handle the post req for saving electronics shop data and save the data to DB.
//If there is an error while saving send back error response with status code as 400(Bad-request)
func SaveElectronicShopData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var reqBody Electronic
	json.NewDecoder(r.Body).Decode(&reqBody)
	if err := DB.Create(reqBody.Data).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("There was some error while saving Electronics Shop data")
		// log.Fatal(err)
	}

	json.NewEncoder(w).Encode(reqBody)
}

//handle the post req for saving super-market shop data and save the data to DB.
//If there is an error while saving send back error response with status code as 400(Bad-request)
func SaveSuperMarketShopData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var reqBody SuperMarket
	json.NewDecoder(r.Body).Decode(&reqBody)
	if err := DB.Create(reqBody.Data).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("There was some error while saving Suer Market Shop data")
		// log.Fatal(err)
	}

	json.NewEncoder(w).Encode(reqBody)
}
