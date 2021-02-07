package base

type IdResponse struct {
	ID uint64 `json:"id"`
}

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type PageResponse struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Data     interface{} `json:"data"`
}
