package dto

import "gorm.io/gorm"

type Pagination struct {
	Next          int
	Previous      int
	RecordPerPage int
	CurrentPage   int
	TotalPage     int
}

func GetPaginated(db *gorm.DB, page int, pageSize int, result interface{}) (*Pagination, error) {
	var totalRows int64
	offset := (page - 1) * pageSize

	db.Model(result).Count(&totalRows)
	err := db.Limit(pageSize).Offset(offset).Find(result).Error
	if err != nil {
		return nil, err
	}

	totalPages := int(totalRows) / pageSize
	if int(totalRows)%pageSize != 0 {
		totalPages++
	}

	pagination := Pagination{
		Next:          min(page+1, totalPages),
		Previous:      max(page-1, 1),
		RecordPerPage: pageSize,
		CurrentPage:   page,
		TotalPage:     totalPages,
	}

	return &pagination, nil
}
