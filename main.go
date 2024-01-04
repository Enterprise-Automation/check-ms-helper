package checks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"context"
)

type Function func(ctx context.Context) map[string]interface{}

type check struct {
	name     string
	function Function
	item	context.Context
}

var checks []check

func NewCheck(name string, function Function, item context.Context) {
	fmt.Println("New check: " + name)
	checks = append(checks, check{name: name, function: function, item: item})
}

func RegisterChecks() {
	fmt.Println("registering checks")
	for _, check := range checks {
		if check.name == os.Getenv("CHECK_ACTION") {
			fmt.Println("Performing check: " + check.name )
			postBody, _ := json.Marshal(check.function(check.item))
			requestBody := bytes.NewBuffer(postBody)
			fmt.Println("making request to: " + os.Getenv("CHECK_CALLBACK_URL"))
			fmt.Println("with body: \n" + string(postBody))
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
