package data

type DataStruct struct {
	Data map[string]interface{} `json:"data" swaggertype:"object"`
}

type KeysStruct struct {
	Keys []string `json:"keys"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
