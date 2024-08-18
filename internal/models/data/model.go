package data

type Write struct {
	Data map[string]interface{} `json:"data"`
}

type Read struct {
	Keys []string `json:"keys"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
