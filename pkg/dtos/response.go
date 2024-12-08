package dtos

type MetaData struct {
	Total int64 `json:"total"`
	Count int64 `json:"count"`
}

type Links struct {
	Self string `json:"self"`
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type DataResponseDTO struct {
	Data interface{} `json:"data,omitempty"`
}

type ResponseDTO struct {
	DataResponseDTO
	Meta  MetaData `json:"meta"`
	Links Links    `json:"links"`
}

type ErrorResponseDTO struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
