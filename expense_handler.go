package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func AddExpense(c echo.Context) (err error) {
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
		log.Println("AddExpense is error", err)
		return c.JSON(http.StatusInternalServerError, "AddExpense is error")
	}

	return c.JSON(http.StatusCreated, res)
}

func GetAllExpense(c echo.Context) (err error) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		
	var res []ResExpense
		
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		log.Println("Prepare statement is error", err)
		return c.JSON(http.StatusInternalServerError, "Prepare statement is error")
	}

	row, err := stmt.Query()
	if err != nil {
		log.Println("Query is error", err)
		return c.JSON(http.StatusInternalServerError, "Query is error")
	}

	for row.Next() {
		var r ResExpense
		err = row.Scan(&r.ID, &r.TITLE, &r.AMOUNT, &r.NOTE, (*pq.StringArray)(&r.TAGS))
		if err != nil {
			log.Println("GetAllExpense is error", err)
			return c.JSON(http.StatusInternalServerError, "GetAllExpense is error")
		}

		res = append(res, r)
	}

	return c.JSON(http.StatusOK, res)
}

func GetExpense(c echo.Context) (err error) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	res := new(ResExpense)

	uri := strings.Split(c.Request().RequestURI, "/")
	id := uri[len(uri)-1]

	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id=$1")
	if err != nil {
		log.Println("Prepare statement is error", err)
		return c.JSON(http.StatusInternalServerError, "Prepare statement is error")
	}

	row := stmt.QueryRow(id)
	err = row.Scan(&res.ID, &res.TITLE, &res.AMOUNT, &res.NOTE, (*pq.StringArray)(&res.TAGS))
	if err != nil {
		log.Println("GetExpense is error", err)
		return c.JSON(http.StatusInternalServerError, "GetExpense is error")
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateExpense(c echo.Context) (err error) {
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
		log.Println("UpdateExpense is error", err)
		return c.JSON(http.StatusInternalServerError, "UpdateExpense is error")
	}

	return c.JSON(http.StatusOK, res)
}