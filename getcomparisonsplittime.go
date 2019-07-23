package main

import (
	"net"

	"github.com/samwho/livesplit"
	"github.com/samwho/streamdeck"
)

func init() {
	RegisterInit(func(sd *streamdeck.Client, ls *livesplit.Client) {
		action := sd.Action(ActionName("getcomparisonsplittime"))

		update := func(cmd []string) error {
			t, err := ls.GetComparisonSplitTime()
			if err, ok := err.(net.Error); ok && err.Timeout() {
				for _, ctx := range action.Contexts() {
					if err := sd.SetTitle(ctx, "N/A", streamdeck.HardwareAndSoftware); err != nil {
						return err
					}
				}
			}

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
		}

		ls.OnStartTimer(update)
		ls.OnSplit(update)
		ls.OnUnsplit(update)
		ls.OnSkipSplit(update)
		ls.OnReset(update)
	})
}
