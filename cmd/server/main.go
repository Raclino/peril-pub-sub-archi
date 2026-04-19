package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

const connString = "amqp://guest:guest@localhost:5672/"

func main() {
	fmt.Println("Starting Peril server...")

	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, w io.Writer, args []string) error {
	_ = args

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	conn, err := amqp.Dial(connString)
	if err != nil {
		return fmt.Errorf("couldn't create RabbitMQ connection: %w", err)
	}
	defer conn.Close()

	fmt.Fprintln(w, "Connection to RabbitMQ was successful!")
	fmt.Fprintln(w, "Press Ctrl+C to stop the server.")

	<-ctx.Done()

	fmt.Fprintln(w, "Shutting down server...")

	return nil
}
