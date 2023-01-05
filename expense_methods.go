package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func addExpense(c echo.Context) (err error) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	req := new(ReqExpense)
	res := new(ResExpense)

	err = c.Bind(req)
	if err != nil {
		log.Println("Request is invalid", err)
		return c.JSON(http.StatusBadRequest, "Request is invalid")
	}

	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4) RETURNING id, title, amount, note, tags", req.TITLE, req.AMOUNT, req.NOTE, pq.Array(req.TAGS))
	
	err = row.Scan(&res.ID, &res.TITLE, &res.AMOUNT, &res.NOTE, (*pq.StringArray)(&res.TAGS))
	if err != nil {
		log.Println("Insert is error", err)
		return c.JSON(http.StatusInternalServerError, "Insert is error")
	}

	return c.JSON(http.StatusCreated, res)
}

func getExpense(c echo.Context) (err error) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	res := new(ResExpense)

	uri := strings.Split(c.Request().RequestURI, "/")
	id := uri[len(uri) - 1]

	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id=$1")
	if err != nil {
		log.Println("Prepare statement is error", err)
		return c.JSON(http.StatusInternalServerError, "Prepare statement is error")
	}

	row := stmt.QueryRow(id)

	err = row.Scan(&res.ID, &res.TITLE, &res.AMOUNT, &res.NOTE, (*pq.StringArray)(&res.TAGS))
	if err != nil {
		log.Println("Get data is error", err)
		return c.JSON(http.StatusInternalServerError, "Get data is error")
	}

	return c.JSON(http.StatusOK, res)
}