package docker

import (
	"strings"
)

func NewContainer(id string, image string)(*Container, error){
	tag := strings.Split(strings.Split(image, "@")[0], ":")[1]
	container := Container{
		Id:id,
		Tag:tag,
		Image:image,
	}

	return &container, nil
}

