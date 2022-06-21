package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Pods struct {
}

func main() {

	var token string = "af71f0a4-45d8-8be1-4d4c-98536fbfe81f"
	sessionToken := getSessionToken(token)

	emptyBody, _ := json.Marshal(map[string]string{})
	podGet, err := http.NewRequest("GET", "https://flasharray1.testdrive.local/api/1.19/pod", bytes.NewBuffer(emptyBody))
	if err != nil {
		log.Fatal("cannot GET pod")
	}
	podGet.Header.Add("Cookie", sessionToken)

	client := &http.Client{}
	podData, err := client.Do(podGet)
	if err != nil {
		log.Fatal("unable to create http client")
	}
	defer podData.Body.Close()
	pods, err := io.ReadAll(podData.Body)
	fmt.Println(string(pods))
}

func getSessionToken(token string) string {
	requestBody, err := json.Marshal(map[string]string{
		"api_token": token,
	})

	resp, err := http.Post("https://flasharray1.testdrive.local/api/1.19/auth/session", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatal("cannot load")
	}
	cookies := resp.Header["Set-Cookie"][0]
	sessionToken := strings.Split(cookies, ";")[0]

	return sessionToken
}
