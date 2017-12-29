package shortener

// package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const apiHost = "https://www.googleapis.com/urlshortener/v1/url?key=AIzaSyDVNkFiyghjNpliLYWUZQ8ZSpTGpey-QQU"

//Short is function than return url short
func Short(url string) <-chan string {

	c := make(chan string)

	go func() {

		c <- getUrl(url)

	}()

	return c
}

func getUrl(url string) string {

	var body map[string]interface{}
	jsonBody := `{"longUrl":"` + url + `"}`
	resp, _ := http.Post(apiHost, "application/json", bytes.NewBufferString(jsonBody))

	body_byte, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(body_byte), &body)

	return body["id"].(string)
}

func main() {

}
