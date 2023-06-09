package utils

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func NewPagination(page, pageSize int) *Pagination {
	return &Pagination{page, pageSize}
}
