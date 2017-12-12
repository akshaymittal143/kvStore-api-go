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
	countW:=0
	countG:=0

	//putRequest := newValueRequest(http.MethodPut, "http://localhost:9001/api/values/2", "foo")

	postRequest := newValueRequest(http.MethodPost, "http://localhost:9001/api/values", "bar")
	countW++

	addrs := []string{":9002", ":9003", ":9004", ":9005"}
	// Setup a HTTP client
	c := &http.Client{}

	//if one node1 write then do write for other nodes
	if countW<=5  {
		for _, addr := range addrs{
			postRequest := newValueRequest(http.MethodPost, "http://localhost"+addr+"/api/values", "bar")
			c.Do(postRequest)
			countW++
		}
	}
	// Send the two requests
	//c.Do(putRequest)
	c.Do(postRequest)

	response, err := http.Get("http://localhost:9001/api/values")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	countG=1
	fmt.Print("Server ",countG)
	fmt.Println(string(responseData))
	countG++

	if countG<=5  {
		for _, addr := range addrs{
			response, err := http.Get("http://localhost"+addr+"/api/values")
			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}
			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print("Server ",countG)
			fmt.Println(string(responseData))
			countG++
		}
	}

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

