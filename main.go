package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	token    *string
	ip       *string
	endpoint *string
)

func main() {

	token = flag.String("t", "", "User token")
	ip = flag.String("i", "", "Array IP")
	endpoint = flag.String("e", "", "Endpoint")

	flag.Parse()

	if *token == "" || *ip == "" || *endpoint == "" {
		usage()
	}

	// postRequest("/api/v1/auth/session")
	getRequest()
}

func usage() {
	fmt.Println("you messed up")
	os.Exit(1)
}

func postRequest(endpoint string) {

	data := url.Values{
		"api_token": {"John Doe"},
	}

	postRequest := "https://" + *ip + endpoint

	fmt.Println(postRequest)

	resp, err := http.PostForm(postRequest, data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["form"])
}

func getRequest() {

	getRequest := *ip + *endpoint

	resp, err := http.Get(getRequest)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
