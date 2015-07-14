# callfire-go-api-helper

This is a simple set of code for interacting with the CallFire
API, made for use with the Go code samples in the new REST API
documentation.

## Usage

```go
package main

import (
	"fmt"
	"github.com/settermjd/callfire"
	"io/ioutil"
	"log"
)

func main() {
    // Initialise a request options object
    requestOptions := callfire.CallFireRequestOptions{
        Url:            "https://www.callfire.com/api/1.1/rest/call/?",
        Login:          "YOUR_LOGIN",
        Secret:         "YOUR_PASSWORD",
        ReqType:        "GET",
        RequestOptions: "MaxResults=10&FromNumber=2092084589&ToNumber=2092084589&LabelName=TestBroadcast&State=FINISHED",
    }

    // Initialise both the client and request objects
	client, req := callfire.InitHttpClient(requestOptions)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Dump the response
	fmt.Printf("%s", body)
}
```

## Tests

Not implemented yet

## Release History

- 0.0.1 Initial release
