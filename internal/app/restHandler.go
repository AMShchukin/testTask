package app

import (
	json "encoding/json"
	"github.com/labstack/echo/v4"
	spec "testTask/internal/pkg/models"
	"testTask/internal/pkg/service"
)

type RestHandler struct {
}

func (r RestHandler) PostOffice(ctx echo.Context) error {
	request := spec.SaveSortPointIdForDstOfficeRequest{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	service.SaveSortPointForDstOfficeId(int64(*request.DstOfficeId), int64(*request.SortPointId))
	return nil
}
