package handlers

import (
	"context"
	"garination.com/db/sdk/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetUserDealership(context.Context, *proto.GetUserDealershipRequest) (*proto.GetUserDealershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDealership not implemented")
}
func (h *Handler) GetUserMeta(context.Context, *proto.GetUserMetaRequest) (*proto.GetUserMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMeta not implemented")
}

func (h *Handler) InsertUserMeta(context.Context, *proto.InsertUserMetaRequest) (*proto.InsertUserMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUserMeta not implemented")
}
func (h *Handler) UpdateUserMeta(context.Context, *proto.UpdateUserMetaRequest) (*proto.UpdateUserMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMeta not implemented")
}
