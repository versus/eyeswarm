package types


type Service struct {
	Id string `json:"service_id"`
	Name string `json:"service_name"`
	Image string `json:"service_image"`
	Tag string `json:"service_tag"`
}

type Container struct {
	Id string `json:"container_id"`
	Image string `json:"container_image"`
	Tag string `json:"container_tag"`
}

type Config struct {
	Server server
}

type server struct {
	IP       string
	Port     string
	Endpoint string
}