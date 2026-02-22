package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/amaterasutears/bot"
	"github.com/amaterasutears/bot/machine"
	"github.com/amaterasutears/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// NOTE: Replace `nil` with your Storage implementation (e.g., memory.New(), redis.New(), postgres.New()).
	m := machine.New(nil)

	opts := []bot.Option{
		bot.WithMachine(m),
	}

	b, err := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)
	if nil != err {
		panic(err)
	}

	b.RegisterHandlerMatchState("old state", handler)

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	uid, ok := bot.ExtractUserIDFromUpdate(update)
	if !ok {
		// handle
	}

	err := b.Machine().SetState(ctx, uid, "new state")
	if err != nil {
		// handle
	}
}
