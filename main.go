package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/filters"

	"strings"
	"sync"
)
func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup

	srvs, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}
	for _, srv := range srvs {
		serviceID := srv.ID
		serviceImage := srv.Spec.TaskTemplate.ContainerSpec.Image
		serviceTAG := strings.Split(strings.Split(serviceImage, "@")[0], ":")[1]
		fmt.Printf("%s %s \n", srv.Spec.Name, serviceTAG)
		wg.Add(1)
		go func() {
			filters := filters.NewArgs()
			filters.Add("label", "com.docker.swarm.service.id="+serviceID)

			containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Filters: filters})
			if err != nil {
				panic(err)
			}
			for _, container := range containers {
				if serviceImage == container.Image {
					fmt.Printf("%s %s %s\n", container.ID[:10], container.Image, container.State)
				}
			}
			wg.Done()
		}()

	}
	wg.Wait()
}