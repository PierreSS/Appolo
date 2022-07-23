package config

import (
	"fmt"
	"strconv"
)

// Pagination helps chosing which page u want
func Pagination(page string) (string, error) {
	offset, err := strconv.Atoi(page)
	if err != nil {
		return "", err
	}
	if offset > 1 {
		offset *= 25 - 25
	} else if offset == 1 {
		offset = 0
	} else {
		return "", fmt.Errorf("Wrong input for number of pages : %d", offset)
	}

	return strconv.Itoa(offset), nil
}
