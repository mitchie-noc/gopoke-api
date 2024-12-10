package utils

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Code string
	Message string
}

func GetOffsetAndLimit(context *gin.Context) (int, int, error){
	offset := 0
	limit := 20

	providedOffset := context.Query("offset")
	providedLimit := context.Query("limit")

	if providedOffset != "" {
		fmt.Println("Updating offset to request value")
		newOffset, err := strconv.Atoi(providedOffset)

		if err != nil {
			return -1, -1, err
		}
		offset = newOffset
	}

	if providedLimit != "" {
		fmt.Println("Updating limit to request value")
		newLimit, err := strconv.Atoi(providedLimit)

		if err != nil {
			return -1, -1, err
		}
		limit = newLimit
	}
	return offset, limit, nil
}