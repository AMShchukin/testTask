package app

import (
	"github.com/stretchr/testify/assert"
	"testTask/internal/pkg/service"
	"testing"
	"time"
)

func TestSaveSortPointAndGetSortPointForDstOffice(t *testing.T) {
	dstOfficeId := int64(126)
	expectedSortPointId := int64(456)
	service.SaveSortPointForDstOfficeId(dstOfficeId, expectedSortPointId)
	// подождем 100 миллисекунд чтобы другой гоурутиной потом получить (первая гоурутина тратит 40 наносекунд на то чтобы отпустить блок плюс накладные расходы)
	time.Sleep(100 * time.Millisecond)
	actualSortPointId, err := service.GetSortPointIdByDstOfficeId(dstOfficeId)
	if err != nil {
		t.Errorf("при получении id СЦ возникла ошибка")
	}
	assert.Equal(t, expectedSortPointId, actualSortPointId)
}
