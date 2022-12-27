package main

import (
	"log"
	"net/http"

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