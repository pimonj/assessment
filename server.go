package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pimonj/assessment/auth"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error

	url := os.Getenv("DATABASE_URL")
	if len(url) == 0 {
		log.Fatal("Not found DATABASE_URL", err)
	}

	db, err = sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS expenses ( id SERIAL PRIMARY KEY, title TEXT, amount FLOAT, note TEXT, tags TEXT[] );`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(auth.Authentication)

	g := e.Group("/expenses")

	g.POST("", AddExpense)
	g.GET("/:id", GetExpense)
	g.PUT("/:id", UpdateExpense)
	g.GET("", GetAllExpense)

	port := os.Getenv("PORT")
	fmt.Println("PORT =", port)
	
	if len(port) == 0 {
		e.Logger.Info("Server will be start at port:2565", err)		
		port = ":2565"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	fmt.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}