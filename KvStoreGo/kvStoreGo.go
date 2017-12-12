package main
import (
"fmt"
"io/ioutil"
"log"
"net/http"
"os"

)
func main() {
	response, err := http.Get("http://localhost:9001/api/values/")
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

