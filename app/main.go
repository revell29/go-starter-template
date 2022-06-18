package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	_authHttp "github.com/questizen/core-system/auth/delivery/http"
	_authRepo "github.com/questizen/core-system/auth/repository/postgres"
	_authUsecase "github.com/questizen/core-system/auth/usecase"
)

func main() {
	fmt.Println("Running services, and loading env configuration")
	godotenv.Load()
	config := GetConfig()

	configList := fmt.Sprintf("Loading config %+v", config)
	fmt.Println(configList)

	db, err := SetupDb()

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to PlanetScale!")

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Minute*2, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	fmt.Println(os.Getenv("NODE_PORT"))

	e := echo.New()

	timeoutContext := 5 * time.Second

	ar := _authRepo.NewPostgreAuthRepository(db)
	au := _authUsecase.NewAuthUseCase(ar, timeoutContext)
	_authHttp.NewAuthHandler(e, au)

	fmt.Println(timeoutContext)

	// Graceful Shutdown
	go func() {

		if err := e.Start(":8080"); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	e.Shutdown(ctx)

	log.Println("Shutting down server")
	os.Exit(0)
}
