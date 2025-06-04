package repository

import (
	"math"

	// Corrected import path for the dto package
	"github.com/Revprm/Nutrigrow-Backend/dto"
	"gorm.io/gorm"
)

// Paginate provides a GORM scope for pagination.
// It takes a dto.PaginationRequest and returns a function that GORM can use to apply limit and offset.
func Paginate(req dto.PaginationRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Calculate offset based on the current page and items per page.
		// Page numbers are typically 1-based.
		offset := (req.Page - 1) * req.PerPage
		return db.Offset(offset).Limit(req.PerPage)
	}
}

// TotalPage calculates the total number of pages based on the total count of items and items per page.
func TotalPage(count, perPage int64) int64 {
	if perPage <= 0 { // Avoid division by zero or negative perPage
		return 0
	}
	// Use math.Ceil to round up to the nearest whole number, ensuring all items are covered.
	totalPage := int64(math.Ceil(float64(count) / float64(perPage)))
	return totalPage
}
