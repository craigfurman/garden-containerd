package main

import (
	"errors"
	"flag"
	"fmt"
	"garden-containerd/nerd"
	"io/ioutil"
	"log"
	"os"
	"time"

	"code.cloudfoundry.org/garden/server"
	"code.cloudfoundry.org/lager"
	"github.com/containerd/containerd"
)

func main() {
	containerdAddress := flag.String("containerd-address", containerd.DefaultAddress, "")
	pidfile := flag.String("pidfile", "", "")
	flag.Parse()

	if *pidfile == "" {
		bail("must set --pidfile")
	}

	must("waiting for containerd socket to exist", awaitFileExistence(*containerdAddress))
	must("writing pidfile", ioutil.WriteFile(*pidfile, []byte(fmt.Sprintf("%d", os.Getpid())), 0644))

	containerdClient, err := containerd.New(*containerdAddress)
	mustNot("connecting containerd client", err)
	defer containerdClient.Close()

	nerd := nerd.New(containerdClient, log.New(os.Stdout, "[nerd] ", log.LstdFlags))

	lagerLogger := lager.NewLogger("garden-containerd")
	lagerLogger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	gardenServer := server.New(
		"tcp", "0.0.0.0:7777",
		time.Minute,
		nerd,
		lagerLogger,
	)

	fmt.Println("starting...")
	must("starting garden server", gardenServer.Start())
	fmt.Println("started...")
	block := make(chan struct{})
	<-block
}

func must(action string, err error) {
	if err != nil {
		fmt.Printf("ERROR %s: %s\n", action, err)
		os.Exit(1)
	}
}

var mustNot = must

func bail(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func awaitFileExistence(path string) error {
	existsCh := make(chan error)
	go func(exists chan<- error) {
		for {
			_, err := os.Stat(path)
			if err == nil {
				close(existsCh)
				return
			}

			if !os.IsNotExist(err) {
				exists <- err
				return
			}

			time.Sleep(time.Millisecond * 100)
		}
	}(existsCh)

	select {
	case err := <-existsCh:
		return err
	case <-time.After(time.Second * 30):
		return errors.New("timed out after 30 seconds")
	}
}
