package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type HttpRequest struct {
	Url    string
	Method string
	Param  string
	Data   map[string]interface{}
}

func RequestApi(options *HttpRequest) chan map[string]interface{} {
	ch := make(chan map[string]interface{})
	go func() {
		myBody := bytes.NewBuffer([]byte(""))
		if options.Data != nil {
			postBody, _ := json.Marshal(options.Data)
			myBody = bytes.NewBuffer(postBody)
		}
		url := options.Url + "?" + options.Param
		c := http.Client{Timeout: time.Duration(1) * 5 * time.Second}
		req, err := http.NewRequest(options.Method, url, myBody)

		if err != nil {
			panic(err)
		}

		req.Header.Add("Content-Type", "application/json")
		resp, err := c.Do(req)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		var data map[string]interface{}
		json.Unmarshal(body, &data)

		ch <- data
	}()

	return ch
}
