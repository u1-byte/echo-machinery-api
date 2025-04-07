package main

import (
	"log"

	"github.com/RichardKnop/machinery/v2"
	"github.com/alecthomas/kong"
	"github.com/labstack/echo/v4"
)

var server *machinery.Server

type CLI struct {
	API      APICommand      `cmd:"" help:"Run API server"`
	Worker   WorkerCommand   `cmd:"" help:"Run background worker"`
	Producer ProducerCommand `cmd:"" help:"Run producer for testing"`
}

type APICommand struct{}

func (a *APICommand) Run() error {
	e := echo.New()
	e.GET("/add", SendTaskHandler("add"))
	e.GET("/multiply", SendTaskHandler("multiply"))

	log.Println("Listening on :8080")
	e.Logger.Fatal(e.Start(":8080"))
	return nil
}

type WorkerCommand struct{}

func (w *WorkerCommand) Run() error {
	worker := server.NewCustomQueueWorker("worker_machinery", 5, "test_queue")
	log.Println("Worker started...")
	if err := worker.Launch(); err != nil {
		log.Fatal(err)
	}
	return nil
}

type ProducerCommand struct{}

func (w *ProducerCommand) Run() error {
	log.Println("Producer started...")
	Produce()
	return nil
}

func main() {
	var cli CLI
	kongCtx := kong.Parse(&cli)
	var err error
	server, err = StartServer()
	if err != nil {
		log.Fatal(err)
	}
	kongCtx.Run()
	kongCtx.FatalIfErrorf(err)
}
