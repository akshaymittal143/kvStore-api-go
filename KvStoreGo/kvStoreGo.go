package main
import (
"fmt"
"io/ioutil"
"log"
"net/http"
"os"

	"strings"

	"io"
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


func putRequest(url string, data io.Reader)  {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, data)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	_, err = client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

}

func httpPutExample()  {
	putRequest("http://localhost:9001/api/values/2", strings.NewReader("foo"))
}

