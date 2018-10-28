package godock

import (
	"fmt"
	"context"
	"log"
	"io"
	// "encoding/json"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/events"
)


func PrintList(){
	// cli, err := client.NewEnvClient()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		// fmt.Printf("%#v", container)
		fmt.Println(container.Ports)
		fmt.Printf("%s %s\n", container.ID, container.Image)
	}
}

func PrintEvents(){
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
	for{
		// e e' type Message
		log.Println("Waiting for event...")
		e := <- body
        if err != nil && err == io.EOF {
            break
		}
		
		log.Println(e.Status, e.From, e.Action)
	}
}