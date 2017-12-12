package main
import (
"fmt"
"io/ioutil"
"log"
"net/http"
"os"


	"encoding/json"
	"bytes"
)
func main() {
	// Create two requests (PUT and POST)
	putRequest := newValueRequest(http.MethodPut, "http://localhost:9002/api/values/2", "foo")
	postRequest := newValueRequest(http.MethodPost, "http://localhost:9001/api/values", "bar")

	// Setup a HTTP client
	c := &http.Client{}

	// Send the two requests
	c.Do(putRequest)
	c.Do(postRequest)

	response, err := http.Get("http://localhost:9001/api/values ")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}


func newValueRequest(method, rawurl, value string) *http.Request {
	b, err := json.Marshal(value)
	if err != nil {
		return nil
	}

	req, err := http.NewRequest(method, rawurl, bytes.NewReader(b))
	if err != nil {
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	return req
}

