package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RichardKnop/machinery/v2"
	amqpbackend "github.com/RichardKnop/machinery/v2/backends/amqp"
	amqpbroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/labstack/echo/v4"
	tasksInit "github.com/u1/echo-machinery-api/tasks"
)

func SendTaskHandler(opName string) echo.HandlerFunc {
	return func(c echo.Context) error {
		aStr := c.QueryParam("a")
		bStr := c.QueryParam("b")

		a, err1 := strconv.Atoi(aStr)
		b, err2 := strconv.Atoi(bStr)

		if err1 != nil || err2 != nil {
			return c.String(http.StatusBadRequest, "Parameters 'a' and 'b' must be integers")
		}
		signature := &tasks.Signature{
			Name:       opName,
			RoutingKey: "test_queue",
			Args: []tasks.Arg{
				{Type: "int", Value: a},
				{Type: "int", Value: b},
			},
		}

		_, err := server.SendTask(signature)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to send task")
		}
		return c.String(http.StatusOK, fmt.Sprintf("Task %s sent!", opName))
	}
}

func StartServer() (*machinery.Server, error) {
	cnf := &config.Config{
		Broker:          "amqp://guest:guest@localhost:5672/",
		DefaultQueue:    "test_queue",
		ResultBackend:   "amqp://guest:guest@localhost:5672/",
		ResultsExpireIn: 1,
		AMQP: &config.AMQPConfig{
			Exchange:      "machinery_exchange",
			ExchangeType:  "direct",
			BindingKey:    "test_queue",
			PrefetchCount: 10,
		},
	}

	// Create server instance
	broker := amqpbroker.New(cnf)
	backend := amqpbackend.New(cnf)
	lock := eagerlock.New()
	server := machinery.NewServer(cnf, broker, backend, lock)

	// Register worker
	tasksMap := map[string]interface{}{
		"add":      tasksInit.Add,
		"multiply": tasksInit.Multiply,
	}
	err := server.RegisterTasks(tasksMap)
	if err != nil {
		log.Fatal(err)
	}
	return server, err
}
