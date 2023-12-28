package checks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Function func() map[string]interface{}

type check struct {
	name     string
	function Function
}

var checks []check

func NewCheck(name string, function Function) {
	fmt.Println("New check: " + name)
	checks = append(checks, check{name: name, function: function})
}

func RegisterChecks() {
	fmt.Println("registering checks")
	for _, check := range checks {
		if check.name == os.Getenv("CHECK_ACTION") {
			fmt.Println("Performing check: " + check.name)
			postBody, _ := json.Marshal(check.function())
			requestBody := bytes.NewBuffer(postBody)
			fmt.Println("making request to: " + os.Getenv("CHECK_CALLBACK_URL"))
			fmt.Println("with body: \n" + postBody.String())
			resp, err := http.Post(os.Getenv("CHECK_CALLBACK_URL"), "application/json", requestBody)
			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			if resp.StatusCode != 200 {
				log.Fatalf("An Error Occured making the request. it returned a non 200 statuscode")
			}
			defer resp.Body.Close()
		}
	}
}
