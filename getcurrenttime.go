package main

import (
	"github.com/samwho/livesplit"
	"github.com/samwho/streamdeck"
)

func init() {
	RegisterInit(func(sd *streamdeck.Client, ls *livesplit.Client) {
		action := sd.Action(ActionName("getcurrenttime"))

		WhileRunning(func() error {
			t, err := ls.GetCurrentTime()
			if err != nil {
				return err
			}

			s := ""
			if t > 0 {
				s = livesplit.DurationToString(t)
			}

			for _, ctx := range action.Contexts() {
				if err := sd.SetTitle(ctx, s, streamdeck.HardwareAndSoftware); err != nil {
					return err
				}
			}

			return nil
		})
	})
}
