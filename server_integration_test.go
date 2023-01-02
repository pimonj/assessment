package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	//"github.com/stretchr/testify/assert"
)

type Response struct {
	*http.Response
	err error
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return  &Response{res, err}
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")

}

func TestAddUser(t *testing.T) {

	var res ResExpense
	body := bytes.NewBufferString(`{
		"amount": 407,
		"note": "delivery",
		"title": "pizza",
		"tags": ["food"]
	}`)

	err := request(http.MethodPost, uri("expenses"), body).Decode(&res)
	if err != nil {
		t.Fatal("can't create expense:", err.Error())
	}

	// var tags = []string{
	// 	"food",
	// }
	// assert.Nil(t, err)
	// assert.Equal(t, http.StatusCreated, r)
	// assert.NotEqual(t, 0, res.ID)
	// assert.Equal(t, "delivery", res.NNOTE)
	// assert.Equal(t, 407, res.AMOUNT)
	// assert.Equal(t, "pizza", res.TITLE)
	// assert.Equal(t, tags, res.TTAGS)

}


func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}