package godock

import (
	"context"
	"fmt"
	"io"
	"log"

	// "encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	// "github.com/docker/docker/api/types/events"
)

type PortMap struct {
	Service       string
	HostPort      int
	Host          string
	ContainerPort int
}

func InspectContainer(c types.Container) []PortMap {
	// j, _ := json.Marshal(c)
	// fmt.Println(string(j))
	services := []PortMap{}
	for _, p := range c.Ports {
		if p.IP != "" {
			mymap := PortMap{
				Service:       c.Names[0],
				HostPort:      int(p.PublicPort),
				Host:          p.IP,
				ContainerPort: int(p.PrivatePort),
			}
			services = append(services, mymap)
		}
	}
	// fmt.Printf("%#v \n", services)
	return services
}

func GetDockerHosts() []PortMap {
	fmt.Printf("%#v \n", PrintList())
	return PrintList()
}

func PrintList() []PortMap {
	// cli, err := client.NewEnvClient()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	services := []PortMap{}

	for _, container := range containers {
		services = append(services, InspectContainer(container)...)
	}

	return services
}

func PrintEvents() {
	// see https://stackoverflow.com/questions/38759427/how-to-watch-docker-event-with-engine-api

	// cli, err := client.NewEnvClient()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	body, oerr := cli.Events(context.Background(), types.EventsOptions{})
	if oerr != nil {
		// fmt.Println(body)
		// fmt.Println(oerr)
		// panic(oerr)
	}
	for {
		// e e' type Message
		log.Println("Waiting for event...")
		e := <-body
		if err != nil && err == io.EOF {
			break
		}

		log.Println(e.Status, e.From, e.Action)
	}
}
