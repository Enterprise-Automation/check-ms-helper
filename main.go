package checks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	//  loop checks
	for _, check := range checks {

		// compare check.name with env var "CHECK_ACTION"
		if check.name == os.Getenv("CHECK_ACTION") {

			// create post body var equal equal to
			postBody, _ := json.Marshal(check.function())

			// string res to buf
			requestBody := bytes.NewBuffer(postBody)

			// post to callback url (env var) with headxer format of json and that req.body is req.body (buf)
			resp, err := http.Post(os.Getenv("CHECK_CALLBACK_URL"), "application/json", requestBody)

			//
			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}

			//
			defer resp.Body.Close()

			//
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			//
			sb := string(body)
			fmt.Println(sb)

		}
	}
	//
}
