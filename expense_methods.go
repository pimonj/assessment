package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func getAllExpense(c echo.Context) (err error) {
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
			log.Println("Get data is error", err)
			return c.JSON(http.StatusInternalServerError, "Get data is error")
		}

		res = append(res, r)
	}

	return c.JSON(http.StatusOK, res)
}
