package pagination

import (
	"math"

	"gorm.io/gorm"
)

type SQLPaging struct {
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

var DefaultSQLPaging SQLPaging = SQLPaging{
	Limit: 25,
	Page:  1,
}

// Set pagination limit and page if not provided or negative
func (paging *SQLPaging) SetPaging() {
	if paging.Limit <= 0 {
		paging.Limit = DefaultSQLPaging.Limit
	}

	if paging.Page <= 0 {
		paging.Page = DefaultSQLPaging.Page
	}
}

// Get pagination offset from provided limit and page
func (paging *SQLPaging) GetOffset() int {
	if paging.Page > 0 && paging.Limit > 0 {
		return (paging.Page - 1) * paging.Limit
	} else {
		return 0
	}
}

// Provide GORM pagination scopes
func GormPaginate(model interface{}, paging *SQLPaging, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	// set paging if not provided
	paging.SetPaging()
	offset := paging.GetOffset()
	limit := paging.Limit

	// count total of records
	var totalRows int64
	db.Model(model).Count(&totalRows)

	// calculate page size of data
	pageSize := math.Ceil(float64(totalRows) / float64(limit))
	paging.PageSize = int(pageSize)

	return func(db *gorm.DB) *gorm.DB {
		// paginate request data
		return db.Offset(offset).Limit(limit)
	}
}
