package dto

func IsValidSortField(field string) bool {
	validSortFields := map[string]bool{
		"title":        true,
		"description":  true,
		"genre":        true,
		"release_date": true,
	}
	return validSortFields[field]
}
