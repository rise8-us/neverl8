package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/drewfugate/neverl8/application"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Start app
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}
