/*
Name: plume
File: utils_images.go
Author: Landers
*/

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func getImages(api string) []Images {
	var imgs []Images
	url := fmt.Sprintf("%s/images/json", api)
	res, e := Get(url)
	if e != nil {
		return []Images{}
	}
	var raw []ImagesRaw
	_ = json.Unmarshal(res, &raw)
	for _, r := range raw {
		imgs = append(imgs, Images{
			ID:         r.ID,
			Tags:       r.RepoTags,
			Date:       time.Unix(r.Created, 0).Format("2006-01-02 15:04"),
			Size:       calcSize(r.Size),
			Containers: r.Containers,
		})
	}

	return imgs
}

func calcSize(s int) string {
	gb := s / (1024*1024*1024)
	mb := s / (1024*1024)
	kb := s / 1024

	if gb >= 1 {
		return fmt.Sprintf("%dgb", gb)
	} else if mb >= 1 {
		return fmt.Sprintf("%dmb", mb)
	} else if kb >= 1 {
		return fmt.Sprintf("%dkb", kb)
	} else {
		return fmt.Sprintf("%db", s)
	}
}
