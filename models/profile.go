package models

type Profile struct {
	Name    string   `json:"name"`
	Hobbies []string `json:"hobbies"`
	Image   string   `json:"image"`
}

type HealthCheck struct {
	Status    string   `json:"status"`
	Version   string   `json:"version"`
	AppName   string   `json:"app_name"`
	Endpoints []string `json:"endpoints"`
}
