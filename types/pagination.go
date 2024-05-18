package types

const DefaultPageNumber int = 1
const DefaultPageSize int = 10

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func NewPagination(page int, size int) Pagination {
	var pagination Pagination
	if page == 0 || page < 0 {
		pagination.Page = DefaultPageNumber
	} else {
		pagination.Page = page
	}
	if size == 0 || size < 0 {
		pagination.Size = DefaultPageSize
	} else {
		pagination.Size = size
	}

	return pagination
}

type PaginatedResponse[DataType any] struct {
	Data       []DataType `json:"data"`
	Pagination Pagination `json:"pagination"`
	Prev       int        `json:"prev"`
	Next       int        `json:"next"`
	TotalPages int        `json:"totalPages"`
}

func NewPaginatedResponse[DataType any](data []DataType, pagination Pagination, totalPages int) PaginatedResponse[DataType] {
	prev := max(pagination.Page-1, 1)
	next := max(pagination.Page+1, 1)

	if pagination.Page == prev || totalPages == pagination.Page {
		prev = 0
	}

	if next > totalPages {
		next = 0
	}

	return PaginatedResponse[DataType]{
		Data:       data,
		Pagination: pagination,
		Prev:       prev,
		Next:       next,
		TotalPages: totalPages,
	}
}
