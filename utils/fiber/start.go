package fiber

import (
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/utils/errors"
)

func Start(app *fiber.App, port int) {
	addr := "127.0.0.1:" + strconv.Itoa(port)
	if err := app.Listen(addr); err != nil {
		panic(errors.Wrap("error starting server", err))
	}
}

func StartWithGracefulShutdown(app *fiber.App, port int) {
	// Start the HTTP server in a separate goroutine
	go func() {
		addr := "127.0.0.1:" + strconv.Itoa(port)
		if err := app.Listen(addr); err != nil {
			panic(errors.Wrap("error starting server", err))
		}
	}()

	// Use a channel to receive interrupt signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Block until an interrupt signal is received
	<-interrupt

	// Attempt to gracefully shut down the server
	if err := app.ShutdownWithTimeout(time.Minute * 1); err != nil {
		logger.Error("error occured during server shutdown", "error", err.Error())
	} else {
		logger.Debug("server gracefully shut down")
		os.Exit(0)
	}
}
