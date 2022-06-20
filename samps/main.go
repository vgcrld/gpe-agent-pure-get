package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	var token string = "c6325da8-65ce-7cd9-8a08-a4e9d0204328"

	requestBody, err := json.Marshal(map[string]string{
		"api_token": token,
	})

	emptyBody, err := json.Marshal(map[string]string{})

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://flasharray1.testdrive.local/api/1.19/auth/session", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatal("cannot load")
	}
	cookies := resp.Header["Set-Cookie"]

	podGet, err := http.NewRequest("GET", "https://flasharray1.testdrive.local/api/1.19/pod", bytes.NewBuffer(emptyBody))
	podGet.Header.Add("Cookie", cookies[0])

	spew.Dump(podGet)

}
