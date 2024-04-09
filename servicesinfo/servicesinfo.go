package servicesinfo

// ServiceStruct reprezentuje informacje o pojedynczej usłudze
type ServiceStruct struct {
	Name      string `json:"name"`
	State     string `json:"state"`
	StartType string `json:"start_type"`
}

type ServiceReport struct {
	Services []ServiceStruct `json:"services"`
}