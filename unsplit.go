package main

import (
	"context"

	"github.com/samwho/livesplit"
	"github.com/samwho/streamdeck"
)

func init() {
	RegisterInit(func(sd *streamdeck.Client, ls *livesplit.Client) {
		action := sd.Action(ActionName("unsplit"))
		action.RegisterHandler(streamdeck.KeyDown, func(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
			return ls.Unsplit()
		})
	})
}
