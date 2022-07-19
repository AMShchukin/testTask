package service

import (
	"fmt"
	"testTask/internal/pkg/repository"
)

type getSortPointResult struct {
	Val   int64
	Error error
}

var repo = repository.Storage{
	Data: make(map[int64]int64),
}

func GetSortPointIdByDstOfficeId(dstOfficeId int64) (int64, error) {
	ch := make(chan getSortPointResult)
	go func(dstOfficeId int64) {
		val, ok := repo.GetSortPointIdByDstOfficeId(dstOfficeId)
		if ok == true {
			ch <- getSortPointResult{Val: val, Error: nil}
		} else {
			ch <- getSortPointResult{Val: 0, Error: fmt.Errorf("sort point for dst office id=%v wasn't found", dstOfficeId)}
		}
	}(dstOfficeId)
	getSortPointResult := <-ch
	if getSortPointResult.Error != nil {
		return 0, getSortPointResult.Error
	}
	return getSortPointResult.Val, nil
}

func SaveSortPointForDstOfficeId(dstOfficeId int64, sortPointId int64) {
	go func(dstOfficeId int64, sortPointId int64) {
		repo.SaveSortPointIdForDstOffice(dstOfficeId, sortPointId)
	}(dstOfficeId, sortPointId)
}
