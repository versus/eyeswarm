package docker

import (
	"github.com/versus/eyeswarm/types"
	"strings"
)

func NewContainer(id string, image string)(*types.Container, error){
	tag := strings.Split(strings.Split(image, "@")[0], ":")[1]
	container := types.Container{
		Id:id,
		Tag:tag,
		Image:image,
	}

	return &container, nil
}