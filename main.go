package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/jedib0t/go-pretty/v6/table"
)

func contains(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Image", "Name", "State", "Status", "Ports", "Docker compose path"})

	for _, container := range containers {
		if path, ok := container.Labels["com.docker.compose.project.working_dir"]; ok {
			ports := make([]int, len(container.Ports))
			portsList := ""

			for index, containerPort := range container.Ports {
				port := int(containerPort.PublicPort)
				if !contains(ports, port) && port != 0 {
					ports[index] = port
					portsList = portsList + fmt.Sprintf(" %d", port)
				}
			}

			t.AppendRows([]table.Row{
				{container.ID[:6], container.Image, container.Names, container.State, container.Status, portsList, path},
			})
		}
	}

	t.Render()

}
