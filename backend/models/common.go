package models

import "backend/errcode"

// Pagination Pagination.
type Pagination struct {
	PageSize int `form:"page_size" json:"page_size"`
	PageNum  int `form:"page_num" json:"page_num"`
}

// PaginationRep Pagination Response.
type PaginationRep struct {
	PageSize int   `json:"page_size"`
	PageNum  int   `json:"page_num"`
	Total    int64 `json:"total"`
}

// Verify verify the value of pageNum and pageSize.
func (p *Pagination) Verify() int {
	if p.PageNum < 0 {
		return errcode.IllegalPageNumErr
	} else if p.PageNum == 0 {
		p.PageNum = DefaultPageNum
	}
	if p.PageSize < 0 {
		return errcode.IllegalPageSizeErr
	} else if p.PageSize == 0 {
		p.PageSize = DefaultPageSize
	}
	return 0
}
