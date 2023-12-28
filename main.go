package checks

import (
	"bytes"
	"encoding/json"
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
	checks = append(checks, check{name: name, function: function})
}

func RegisterChecks() {
	for _, check := range checks {
		if check.name == os.Getenv("CHECK_ACTION") {
			postBody, _ := json.Marshal(check.function())
			requestBody := bytes.NewBuffer(postBody)
			resp, err := http.Post(os.Getenv("CHECK_CALLBACK_URL"), "application/json", requestBody)
			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			defer resp.Body.Close()
		}
	}
}
