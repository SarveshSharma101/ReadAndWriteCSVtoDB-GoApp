package main

import (
	shops "ReadWriteXlApp/ReadWriteApp/DataModel"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const csvFolderPath = "../XlFiles"
const destFolderPath = "../ReadFiles/"

/**
Reading csv files from '../XlFiles/' folder one by one, for each line of data converting it to a valid json req Body
and making a post request to respective API
*/
func main() {
	//Read the folder content

	csvFolder, err := os.Open(csvFolderPath)
	panicIfError(err)

	csvFiles, _ := csvFolder.ReadDir(0)
	var csvFile *os.File

	//close the folder at the end of the function call
	defer csvFolder.Close()

	var reqBody map[string]interface{}
	var url string
	var success bool
	//Iterate over each file inside the folder
	for i, f := range csvFiles {
		//Open the files one by one and read the all the lines
		csvFileName := f.Name()
		fmt.Println(i, "-->", csvFileName)
		csvFile, err = os.Open(csvFolderPath + "/" + csvFileName)
		panicIfError(err)

		lines, err := csv.NewReader(csvFile).ReadAll()
		panicIfError(err)

		//Iterate over the lines read inside the file
		for _, line := range lines {
			fmt.Println("-------->", line)
			//Generate a json request body and get the appropriate url to make a post request to save the data
			reqBody, url = getRequestBody(strings.Split(csvFileName, ".")[0], &line)
			//make a post request
			success = postShopData(reqBody, url)
		}
		//Close the open file
		csvFile.Close()
		//Check if all the post request were successful and move the file from XlFiles to ReadFiles folder
		if success {
			err = os.Rename(csvFolderPath+"/"+csvFileName, destFolderPath+csvFileName)
			panicIfError(err)
		}
		fmt.Println("---------------------------------------------------------------------------")
	}
}

/*
Make a post request to the api url provide with the given request body
and return a boolean value to check if the request was success or failure
*/
func postShopData(reqBody map[string]interface{}, url string) bool {

	//convert requestBody to json
	jsonReqBody, err := json.Marshal(reqBody)
	panicIfError(err)

	//make a post request
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonReqBody))
	panicIfError(err)

	//read response body
	body, err := io.ReadAll(response.Body)
	panicIfError(err)
	//print response body and status code
	fmt.Println("Response: ", string(body))
	fmt.Println("Response status: ", response.StatusCode)

	//validate status code to be 200
	return response.StatusCode == 200
}

/**
Generate appropriate request body and url based on the file being read
*/
func getRequestBody(shop string, data *[]string) (reqBody map[string]interface{}, url string) {
	//object of shop interface
	var shopData shops.Shop
	url = "http://localhost:6000/"
	//store shop name
	reqBody = map[string]interface{}{
		"type": shop,
	}

	//Generate shop and api url data based on the shop name
	switch shop {
	case "supermarket":
		id, err := strconv.Atoi((*data)[0])
		panicIfError(err)
		//initialize shop to super-market
		shopData = shops.SuperMarket{
			Id:      id,
			Name:    (*data)[1],
			Gender:  (*data)[2],
			Email:   (*data)[3],
			PhoneNo: (*data)[4],
		}
		url = url + "SuperMarketShop"

	case "electronics":
		id, err := strconv.Atoi((*data)[0])
		panicIfError(err)
		//initialize shop to electronic
		shopData = shops.Electronic{
			Id:       id,
			Email:    (*data)[1],
			Name:     (*data)[2],
			PhoneNo:  (*data)[3],
			Address:  (*data)[4],
			Location: (*data)[5],
		}
		url = url + "ElectronicShop"

	case "mobile":
		id, err := strconv.Atoi((*data)[0])
		panicIfError(err)
		//initialize shop to mobile
		shopData = shops.Mobile{
			Id:       id,
			Name:     (*data)[1],
			Gender:   (*data)[2],
			Email:    (*data)[3],
			Address:  (*data)[4],
			Location: (*data)[5],
			PhoneNo:  (*data)[6],
			Hobby:    (*data)[7],
			Interest: (*data)[8],
		}
		url = url + "MobileShop"

	case "bakery":
		id, err := strconv.Atoi((*data)[0])
		panicIfError(err)
		//initialize shop to bakery
		shopData = shops.Bakery{
			Id:      id,
			Name:    (*data)[1],
			Email:   (*data)[2],
			Gender:  (*data)[3],
			PhoneNo: (*data)[4],
		}
		url = url + "BakeryShop"
	}
	//Add shop-data to request body
	reqBody["data"] = shopData
	return reqBody, url
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
