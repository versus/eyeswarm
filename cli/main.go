package main

import (
	"flag"
	"log"
	"github.com/BurntSushi/toml"
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/filters"
	"github.com/versus/eyeswarm/docker"
)

const (
	Version = "v0.0.1"
	Author  = " by Valentyn Nastenko [nastenko@zeoalliance.com]"
)

var conf docker.Config

func main()  {
	log.Println("eyeswarm client ", Version, Author)

	flagConfigFile := flag.String("c", "./config.toml", "config: path to config file")
	flag.Parse()

	if _, err := toml.DecodeFile(*flagConfigFile, &conf); err != nil {
		log.Fatalln("Error parse config.toml", err.Error())
	}

	ctx := context.Background()
	sig := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	tickChan := time.NewTicker(time.Second * 5).C

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatalln("err: ", err)
		os.Exit(1)
	}


	go func() {
		for {
			select {
			case <- tickChan:
				fmt.Println("Ticker ticked")
				filters := filters.NewArgs()
				containers, err := cli.ContainerList(ctx, types.ContainerListOptions{Filters: filters})
				if err != nil {
					log.Fatalln("err: ", err)
				}
				for _, container := range containers {
					    instance ,_ := docker.NewContainer(container.ID[:10], container.Image)
						log.Printf("%s %s %s\n", instance.Id, instance.Image, instance.Tag)

				}
			case <-sig:
				done <- true
				return
			}
		}
	}()

	<-done
	log.Println("exiting")
	os.Exit(0)
}