package pagination

import (
	"math"
)

type Pager struct {
	// 页码
	CurrentPage int `json:"current_page"`
	// 每页数量
	PageSize int `json:"page_size"`
}

type FullPager struct {
	// 页码
	CurrentPage int `json:"current_page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
	// 总页数
	TotalPages int `json:"total_pages"`
}

func PageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}

func MaxPage(pageSize, totalRows int) int {
	if pageSize <= 0 {
		return 1
	}

	return int(math.Ceil(float64(totalRows / pageSize)))
}

func NewPager(currentPage int, pageSize int) *Pager {
	return &Pager{CurrentPage: currentPage, PageSize: pageSize}
}

func NewFullPager(page, pageSize, totalRows int) FullPager {
	return FullPager{
		CurrentPage: page,
		PageSize:    pageSize,
		TotalPages:  MaxPage(pageSize, totalRows),
		TotalRows:   totalRows,
	}
}
