package main
import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "encoding/json"
)


func weather_resp(url string, auth string)(string){



	client := http.Client{}
	req , err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header = http.Header{
		"content-type": {"application/json"},
		"authorization": {auth},
	}


	response , err := client.Do(req)
	if err != nil {
		   log.Fatalln(err)
	}

	fmt.Println(response)

	fmt.Println("----------------------------------------------------------------------------------")


	// read response body
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
	  fmt.Println(error)
	}
	// close response body
	response.Body.Close()

	var resp string = string(body)

	// print response body
	fmt.Println(resp)

	return resp
}



