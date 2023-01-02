package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGetExpenseByID(t *testing.T) {

	var e ResExpense
	id := "1"

	res := request(http.MethodGet, uri("expenses", id), nil)
	err := res.Decode(&e)

	var tags = []string{
		"food",
		"beverage",
	}
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, id, e.ID)
	assert.Equal(t, "night market promotion discount 10 bath", e.NOTE)
	assert.Equal(t, float64(79), e.AMOUNT)
	assert.Equal(t, "strawberry smoothie", e.TITLE)
	assert.Equal(t, tags, e.TAGS)
}


func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}