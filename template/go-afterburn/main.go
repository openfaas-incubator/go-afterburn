package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	//"github.com/openfaas/go-afterburn/template/go-afterburn/function"
	"handler/function"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {

		req, err := http.ReadRequest(reader)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
		requestBody, _ := ioutil.ReadAll(req.Body)
		resultBody := function.Handle(requestBody)
		r := makeResponse(&resultBody)

		r.Write(os.Stdout)
	}
}

func makeResponse(resultBody *[]byte) http.Response {
	httpVer := "HTTP/1.1"
	status := http.StatusOK
	r := http.Response{
		Proto:      httpVer,
		Status:     strconv.Itoa(status),
		StatusCode: status,
	}

	buf := bytes.NewReader(*resultBody)
	r.Body = ioutil.NopCloser(buf)
	r.ContentLength = int64(buf.Len())

	return r
}
