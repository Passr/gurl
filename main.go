package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/anth1y/gurl/requestbuilder"
)

func main() {
	request, err := requestbuilder.New(os.Args[1:]).Build()
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
