package main

import (
	"consume-rest/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	GetRestClient()
}

func GetRestClient() {
	responses, err := http.Get("http://dummy.restapiexample.com/api/v1/employees")

	if err != nil {
		log.Printf("Error when get api %s", err.Error())
	}

	// log.Printf("Response : %s", responses.StatusCode)
	data, _ := ioutil.ReadAll(responses.Body)

	var modelResponse model.Response

	errorUnmarshallingJSON := json.Unmarshal(data, &modelResponse)
	if errorUnmarshallingJSON != nil {
		log.Printf("Error when decode json : %s", errorUnmarshallingJSON.Error)
	}

	// log.Printf("Responses : %s", modelResponse.Data)
	for _, result := range modelResponse.Data {
		log.Printf("Employee details : %s", result.EmployeeName)
	}
}
