// Binary gotdchats implements chat list request example using testing server.
package main

import (
	"context"
	"crypto/rand"
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/proxy"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/tg"
)

func run(ctx context.Context) error {
	logger, _ := zap.NewDevelopment(zap.IncreaseLevel(zapcore.DebugLevel))
	defer func() { _ = logger.Sync() }()

	var (
		persisted = flag.Bool("persisted", false, "persist session")
		dcID      = flag.Int("dc", 2, "ID of DC")
	)
	flag.Parse()

	var storage telegram.SessionStorage
	if *persisted {
		// Setting up session storage.
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		sessionDir := filepath.Join(home, ".td")
		if err := os.MkdirAll(sessionDir, 0700); err != nil {
			return err
		}
		storage = &telegram.FileSessionStorage{
			Path: filepath.Join(sessionDir, "session-user.json"),
		}
	}

	client := telegram.NewClient(telegram.TestAppID, telegram.TestAppHash, telegram.Options{
		Logger:         logger,
		SessionStorage: storage,
		Resolver:       dcs.Plain(dcs.PlainOptions{Dial: proxy.Dial}),
		DCList:         dcs.Test(),
	})

	return client.Run(ctx, func(ctx context.Context) error {
		if self, err := client.Self(ctx); err != nil || self.Bot {
			if err := auth.NewFlow(
				auth.Test(rand.Reader, *dcID), auth.SendCodeOptions{},
			).Run(ctx, client.Auth()); err != nil {
				return errors.Wrap(err, "auth")
			}
		}

		c := tg.NewClient(client)
		for range time.NewTicker(time.Second * 5).C {
			chats, err := c.MessagesGetAllChats(ctx, nil)

			if d, ok := telegram.AsFloodWait(err); ok {
				// Server told us to wait N seconds before sending next message.
				logger.Info("Sleeping", zap.Duration("duration", d))
				time.Sleep(d)
				continue
			}

			if err != nil {
				return errors.Wrap(err, "get chats")
			}

			switch chats.(type) {
			case *tg.MessagesChats: // messages.chats#64ff9fd5
				logger.Info("Chats")
			case *tg.MessagesChatsSlice: // messages.chatsSlice#9cd81144
				logger.Info("Slice")
			}
		}

		return nil
	})
}

func main() {
	if err := run(context.Background()); err != nil {
		panic(err)
	}
}
