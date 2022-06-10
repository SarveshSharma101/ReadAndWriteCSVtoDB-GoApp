package main

import (
	controller "GoServerSaveToDB-App/Controller"
	service "GoServerSaveToDB-App/Service"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	//Env variable to read data of the respective environment from AppConfig.yaml file
	env := "Local"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	uname, password, url, dbname := readConfigs("./Config/AppConfig.yaml", env)
	//Initialize db connection with the db details provide in config for the given env
	service.InitDBConnection(uname, password, url, dbname)
	//Setup the router and get a reference to it
	router := controller.GetRouter()

	//Initialize the server and provide the router to it
	fmt.Println("Server is listening at 6000...")
	log.Fatal(http.ListenAndServe(":6000", router))
}

/**
Read the Db configurations from the config file for the given environment
*/
func readConfigs(configFilePath string, env string) (uname, password, url, dbname string) {
	//Read the yaml file for configs
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	//convert the byte array to map
	configMap := map[string]map[string]map[string]string{}
	if err := yaml.Unmarshal(yamlFile, &configMap); err != nil {
		panic(err)
	}
	//fetch the value of env
	dbConfigs := configMap["Environment"][env]
	if len(dbConfigs) <= 0 {
		fmt.Println("No such env: ", env)
		panic("Please provide valid env")
	}

	//validate required values are present in the config files else panic
	if len(dbConfigs["dbName"]) <= 0 {
		fmt.Println("DB name not found on config file, for env: ", env)
		panic("Please provide DB name in configs")
	} else if len(dbConfigs["url"]) <= 0 {
		fmt.Println("DB url not found on config file, for env: ", env)
		panic("Please provide DB url in configs")
	} else if len(dbConfigs["password"]) <= 0 {
		fmt.Println("DB password not found on config file, for env: ", env)
		panic("Please provide DB password in configs")
	} else if len(dbConfigs["username"]) <= 0 {
		fmt.Println("DB username not found on config file, for env: ", env)
		panic("Please provide DB username in configs")
	} else {
		dbname = dbConfigs["dbName"]
		url = dbConfigs["url"]
		password = dbConfigs["password"]
		uname = dbConfigs["username"]
	}

	return uname, password, url, dbname
}
