//go:build integration

package main

import (
	"bytes"
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


func TestAddExpense(t *testing.T) {

	var e ResExpense
	body := bytes.NewBufferString(`{
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"title": "strawberry smoothie",
		"tags": ["food", "beverage"]
	}`)

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&e)
	if err != nil {
		t.Fatal("can't create expense:", err.Error())
	}

	var tags = []string{
		"food",
		"beverage",
	}
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, e.ID)
	assert.Equal(t, "night market promotion discount 10 bath", e.NOTE)
	assert.Equal(t, float64(79), e.AMOUNT)
	assert.Equal(t, "strawberry smoothie", e.TITLE)
	assert.Equal(t, tags, e.TAGS)

}

func TestGetExpenseByID(t *testing.T) {

	var e ResExpense
	id := "1"

	res := request(http.MethodGet, uri("expenses", id), nil)
	err := res.Decode(&e)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, id, e.ID)
}

func TestUpdateExpense(t *testing.T) {

	var e ResExpense
	body := bytes.NewBufferString(`{
		"amount": 1500,
		"note": "black friday discount 45%",
		"title": "sweater",
		"tags": ["clothes"]
	}`)

	id := "1"

	res := request(http.MethodPut, uri("expenses", id), body)
	err := res.Decode(&e)
	if err != nil {
		t.Fatal("can't update expense:", err.Error())
	}

	var tags = []string{
		"clothes",
	}
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, id, e.ID)
	assert.Equal(t, "black friday discount 45%", e.NOTE)
	assert.Equal(t, float64(1500), e.AMOUNT)
	assert.Equal(t, "sweater", e.TITLE)
	assert.Equal(t, tags, e.TAGS)
}


func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}