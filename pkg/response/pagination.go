package response

import "math"

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

func (p *Pagination) CalculateTotalPages(totalItems int) {
	p.TotalPages = int(math.Ceil(float64(totalItems) / float64(p.Limit)))
}

func (p *Pagination) CalculateOffset() int {
	return (p.Page - 1) * p.Limit
}
