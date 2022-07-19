package app

import (
	"context"
	"testTask/internal/pkg/service"
	proto "testTask/proto/gen/proto"
)

type ProtoHandler struct {
	proto.UnimplementedTransferBoxApiServer
}

func (p *ProtoHandler) GetSortPointId(ctx context.Context, request *proto.GetSortPointIdRequest) (*proto.GetSortPointIdResponse, error) {
	val, err := service.GetSortPointIdByDstOfficeId(request.DstOfficeId)
	if err != nil {
		return nil, err
	}
	response := proto.GetSortPointIdResponse{SortPointId: val}
	return &response, nil
}
