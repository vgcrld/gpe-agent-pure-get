package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	token    *string
	ip       *string
	endpoint *string

	query             string
	post_auth_session string
	api_version       string
)

type versions struct {
	Version []string `json:"version"`
}

func main() {

	setFlags()

	api_version = getLatestVersion1()

	post_auth_session = "https://" + *ip + "/api/" + api_version + "/auth/session"
	sessionToken := getSessionToken(post_auth_session, *token)

	ep := "https://" + *ip + "/api/" + api_version + *endpoint
	// log.Println("attempting endpoing: ", ep)
	resp, err := getPureData(ep, sessionToken)
	if err != nil {
		log.Fatal("unable to fetch ", *endpoint)
	}
	defer resp.Body.Close()

	pods, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("can read response body", err)
	}
	fmt.Println(string(pods))

}

func setFlags() {
	token = flag.String("token", "", "User token")
	ip = flag.String("ip", "", "Array IP")
	endpoint = flag.String("endpoint", "", "Endpoint")

	flag.Parse()

	if *token == "" || *ip == "" || *endpoint == "" {
		usage()
	}
}

func getLatestVersion1() string {
	var data []byte
	data = getRequest("/api/api_version")

	var vers versions
	err := json.Unmarshal(data, &vers)
	if err != nil {
		log.Fatal("can unmarshal vers")
	}

	var v1s []string

	for _, v := range vers.Version {
		s := strings.Split(v, ".")
		if s[0] == "1" {
			v1s = append(v1s, v)
		}
	}

	lv := v1s[len(v1s)-1]

	return lv
}

func usage() {
	command := os.Args[0]
	msg := fmt.Sprintf("%s -t <token> -i <ip> -e <endpoint>", command)
	fmt.Println(msg)
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

func getRequest(ep string) []byte {

	var request string

	if ep != "" {
		request = *ip + ep
	} else {
		request = *ip + *endpoint
	}

	resp, err := http.Get("https://" + request)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func getSessionToken(sessionEndpoint, token string) string {
	requestBody, err := json.Marshal(map[string]string{
		"api_token": token,
	})

	resp, err := http.Post(sessionEndpoint, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatal("cannot load")
	}
	cookies := resp.Header["Set-Cookie"][0]
	sessionToken := strings.Split(cookies, ";")[0]

	return sessionToken
}

func getPureData(endpoint, sessionToken string) (*http.Response, error) {

	emptyBody, _ := json.Marshal(map[string]string{})
	podGet, err := http.NewRequest("GET", endpoint, bytes.NewBuffer(emptyBody))
	if err != nil {
		log.Fatal("cannot GET pod")
	}
	podGet.Header.Add("Cookie", sessionToken)

	client := &http.Client{}
	podData, err := client.Do(podGet)
	if err != nil {
		log.Fatal("unable to create http client")
	}

	return podData, err
}
