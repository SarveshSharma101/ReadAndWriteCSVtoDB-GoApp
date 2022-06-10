package controller

import (
	service "GoServerSaveToDB-App/Service"

	"github.com/gorilla/mux"
)

//Initialize the mux router and set the mapping for the end point to its respective handler function
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/BakeryShop", service.SaveBakeryShopData).Methods("POST")
	router.HandleFunc("/MobileShop", service.SaveMobileShopData).Methods("POST")
	router.HandleFunc("/ElectronicShop", service.SaveElectronicShopData).Methods("POST")
	router.HandleFunc("/SuperMarketShop", service.SaveSuperMarketShopData).Methods("POST")
	return router
}
