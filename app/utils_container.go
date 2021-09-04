/*
Name: plume
File: utils_container.go
Author: Landers
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// 容器接口代理请求
func getContainers(api string) []Containers {
	var containers []Containers
	var cr []ContainersRaw
 	url := fmt.Sprintf("%s/containers/json?all=true", api)
 	res, e := Get(url)
 	if e != nil {
 		fmt.Printf("get containers failed: %v\n", e)
 		return containers
	}

	_ = json.Unmarshal(res, &cr)
	for _, c := range cr {
		containers = append(containers, Containers{
			ID:     c.ID,
			Name:   c.Names,
			Image:  c.Image,
			Date:   time.Unix(c.Created, 0).Format("2006-01-02 15:04"),
			Status: c.State,
		})
	}
	return containers
}

func getContainer(id, api string) Container {
	var container Container
	url := fmt.Sprintf("%s/containers/%s/json", api, id)

	res, e := Get(url)
	if e != nil {
		fmt.Printf("get container info failed: %v\n", e)
		return container
	}
	var cr ContainerRaw
	_ = json.Unmarshal(res, &cr)
	container = Container{
		Name:   cr.Name,
		Cmd:    cr.Config["Cmd"],
		Image:  cr.Config["Image"],
		User:   cr.Config["User"],
		Volume: cr.HostConfig["Binds"],
		Port:   cr.HostConfig["PortBindings"],
		WrkDir: cr.Config["WorkingDir"],
		Date:   cr.Created,
		Status: cr.State["Status"],
		Cpu:    cr.HostConfig["CpuPercent"],
		Mem:    cr.HostConfig["Memory"],
	}
	return container
}

func startContainer(id, api string) string {
	url := fmt.Sprintf("%s/containers/%s/start", api, id)
	return PostStatus(url, nil)
}

func stopContainer(id, api string) string {
	url := fmt.Sprintf("%s/containers/%s/stop", api, id)
	return PostStatus(url, nil)
}

func restartContainer(id, api string) string {
	url := fmt.Sprintf("%s/containers/%s/restart", api, id)
	return PostStatus(url, nil)
}

func createContainer(name, api string, data io.ReadCloser) string {
	url := fmt.Sprintf("%s/containers/create?name=%s", api, name)
	res := PostStatus(url, data)

	return res
}

func delContainer(id, api string) string {
	url := fmt.Sprintf("%s/containers/%s", api, id)
	return DeleteStatus(url)
}