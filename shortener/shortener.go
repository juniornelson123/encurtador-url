package shortener

// package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiHost = "https://www.googleapis.com/urlshortener/v1/url?key=AIzaSyDVNkFiyghjNpliLYWUZQ8ZSpTGpey-QQU"

//Short is function than return url short
func Short(url string) (<-chan string, <-chan error) {

	// var errorFunc error
	c := make(chan string)

	err := make(chan error)

	go func() {

		result, errorF := getUrl(url)
		if errorF != nil {
			err <- errorF

		} else {
			err <- fmt.Errorf("false")

		}

		c <- result

	}()

	return c, err
}

func getUrl(url string) (string, error) {

	var body map[string]interface{}
	jsonBody := `{"longUrl":"` + url + `"}`
	resp, _ := http.Post(apiHost, "application/json", bytes.NewBufferString(jsonBody))

	body_byte, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(body_byte), &body)

	if body["id"] != nil {
		return body["id"].(string), nil
	} else {
		return "Valor invalido", fmt.Errorf("Dados invalidos\n")
	}

}

// Multiples Urls
// func main() {
// 	c := Short("google.com", "youtube.com")
// 	// c := result(Short("http://google.com"), Short("http://youtube.com"))
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// }

// func result(c1, c2 <-chan string) <-chan string {
// 	c := make(chan string)

// 	go func() {

// 		for {
// 			select {
// 			case s := <-c1:
// 				c <- s
// 			case s := <-c2:
// 				c <- s
// 			}
// 		}
// 	}()

// 	return c
// }
