package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/samwho/livesplit"
	"github.com/samwho/streamdeck"
)

type InitFunc func(sd *streamdeck.Client, ls *livesplit.Client)

var inits []InitFunc

func RegisterInit(init InitFunc) {
	inits = append(inits, init)
}

func ActionName(s string) string {
	return fmt.Sprintf("dev.samwho.streamdeck-livesplit.%s", s)
}

func main() {
	f, err := ioutil.TempFile("", "streamdeck-livesplit.log")
	if err != nil {
		log.Fatalf("error creating temp file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	params, err := streamdeck.ParseRegistrationParams(os.Args)
	if err != nil {
		log.Fatalf("error parsing registration params: %v", err)
	}

	sd := streamdeck.NewClient(context.Background(), params)
	defer sd.Close()

	ls, err := livesplit.NewClient()
	if err != nil {
		log.Fatalf("error creating livesplit client: %v", err)
	}
	defer ls.Close()

	for _, init := range inits {
		init(sd, ls)
	}

	if err := sd.Run(); err != nil {
		log.Fatalf("error running streamdeck client: %v\n", err)
	}
}
