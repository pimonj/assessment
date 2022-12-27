package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func updateExpense(c echo.Context) (err error) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	req := new(ReqExpense)
	res := new(ResExpense)

	uri := strings.Split(c.Request().RequestURI, "/")
	id := uri[len(uri)-1]

	err = c.Bind(req)
	if err != nil {
		log.Println("Request is invalid", err)
		return c.JSON(http.StatusBadRequest, "Request is invalid")
	}

	stmt, err := db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1 RETURNING id, title, amount, note, tags")
	if err != nil {
		log.Println("Prepare statement is error", err)
		return c.JSON(http.StatusInternalServerError, "Prepare statement is error")
	}

	row := stmt.QueryRow(id, req.TITLE, req.AMOUNT, req.NOTE, pq.Array(req.TAGS))

	err = row.Scan(&res.ID, &res.TITLE, &res.AMOUNT, &res.NOTE, (*pq.StringArray)(&res.TAGS))
	if err != nil {
		log.Println("Get data is error", err)
		return c.JSON(http.StatusInternalServerError, "Get data is error")
	}

	return c.JSON(http.StatusOK, res)
}