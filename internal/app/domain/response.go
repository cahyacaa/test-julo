package domain

type Response[responseDataStruct any] struct {
	StatusCode int                `json:"-"`
	Data       responseDataStruct `json:"data"`
	Status     interface{}        `json:"status"`
}
