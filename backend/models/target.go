package models

type Target struct {
	targetConfig struct {
		Type   string `json:"type"`
		Source string `json:"source"`
	} `json:"target_config"`
	interval struct {
		Frequency int    `json:"frequency"`
		Unit      string `json:"unit"`
	} `json:"interval"`
}
