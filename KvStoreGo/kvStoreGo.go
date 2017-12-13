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
	// Send the two requests
	//c.Do(putRequest)
	c.Do(postRequest)
	//if one node1 write then do write for other nodes
	if countW<=5  {
		for _, addr := range addrs{
			postRequest := newValueRequest(http.MethodPost, "http://localhost"+addr+"/api/values", "bar")
			//write operation for other nodes
			c.Do(postRequest)
			countW++
		}
	}

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

/*
Algorithm for sloppy quorum:
use vectorclock

func SloppyQuorum(id, nodes)  {
read the current node, pass nodes in array W/R
acknowledge
setTime to run the node

func SloppyQuorum.Write(key, value, write)
{
{
  if(!key || !value || !writeFactor) {
    throw new Error('write() requires a key, value and a write factor.')
  }
  if(!value.clock) {
    throw new Error('write() value must contain a vector clock.');
  }

  VClock.increment(value, this.id);

  this._execute({
      op: 'write',
      key: key,
      value: value
    }, writeFactor, function(err, responses) {
      callback(err);
    });
};
}

SloppyQuorumRead(key, readFactor, callback) {
  if(!key || !readFactor) {
    throw new Error('write() requires a key and a write factor.');
  }
  this._execute({
    op: 'read',
    key: key
  }, readFactor, function(err, responses) {
    // read repair using vector clocks
    // sort the responses by the vector clocks
    // then compare them to the topmost (in sequential order, the greatest) item
      // if they are concurrent with that item, then there is a conflict
      // that we cannot resolve, so we need to return the item.
      if(VClock.isConcurrent(item, repaired[0]) && !VClock.isIdentical(item, repaired[0])) {
        repaired.push(item);
      } else {
        // console.log('filtering', item, VClock.isConcurrent(item, repaired[0]), !VClock.isIdentical(item, repaired[0]));
      }
    });
    // combine with own read
    //repaired
  });
};
*/
