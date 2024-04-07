package osinfo

type OsReport struct {
	System       string `json:"system"`
	Distribution string `json:"distribution"`
	Version      string `json:"version"`
	Hostname     string `json:"hostname"`
}
