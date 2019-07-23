package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/samwho/livesplit"
	"github.com/samwho/streamdeck"
)

type InitFunc func(sd *streamdeck.Client, ls *livesplit.Client)

var inits []InitFunc
var running []func() error

func RegisterInit(init InitFunc) {
	inits = append(inits, init)
}

func ActionName(s string) string {
	return fmt.Sprintf("dev.samwho.streamdeck-livesplit.%s", s)
}

func WhileRunning(f func() error) {
	running = append(running, f)
}

func main() {
	f, err := ioutil.TempFile("", "streamdeck-livesplit.log")
	if err != nil {
		log.Fatalf("error creating temp file: %v", err)
	}
	defer f.Close()

	//streamdeck.Log().SetOutput(f)
	livesplit.Log().SetOutput(f)

	params, err := streamdeck.ParseRegistrationParams(os.Args)
	if err != nil {
		log.Fatalf("error parsing registration params: %v", err)
	}

	sd := streamdeck.NewClient(context.Background(), params)
	defer sd.Close()

	ls := livesplit.NewClient()
	defer ls.Close()

	for _, init := range inits {
		init(sd, ls)
	}

	go func() {
		for {
			<-time.After(time.Duration(16) * time.Millisecond)
			phase, err := ls.GetCurrentTimerPhase()
			if err != nil || phase != livesplit.Running {
				continue
			}

			for _, f := range running {
				if err := f(); err != nil {
					log.Printf("error running realtime function: %v", err)
				}
			}
		}
	}()

	if err := sd.Run(); err != nil {
		log.Fatalf("error running streamdeck client: %v\n", err)
	}
}
