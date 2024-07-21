package model

type SuccessResponse struct {
	Data       interface{}   `json:"data"`
	Paging     *PageMetadata `json:"paging,omitempty"`
	Message    string        `json:"message"`
	StatusCode int           `json:"status_code"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	PerPage   int   `json:"per_page"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
