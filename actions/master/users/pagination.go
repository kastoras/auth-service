package master_users

import (
	"math"
)

func (muc *MasterUsersController) CalculatePagination(limit, page int) (Pagination, error) {
	var pagination Pagination

	totalEntries, err := muc.TotalUsers()
	if err != nil {
		return pagination, err
	}
	pagination.Limit = limit
	pagination.TotalEntries = totalEntries
	pagination.TotalPages = calculateTotalPages(pagination.TotalEntries, limit)

	if page != 0 {
		pagination.CurrentPage = page
		pagination.Offset = calculateCurrentOffSet(page, pagination.TotalEntries, pagination.Limit)
	}

	return pagination, nil
}

func calculateCurrentOffSet(currentPage, totalEntries, limit int) int {
	if currentPage < 1 {
		currentPage = 1
	}
	offset := (currentPage - 1) * limit
	if offset >= totalEntries {
		return totalEntries
	}
	return offset
}

func calculateTotalPages(totalEntries, limitPerPage int) int {
	return int(math.Ceil(float64(totalEntries) / float64(limitPerPage)))
}
