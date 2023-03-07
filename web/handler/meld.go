package handler

import (
	"sheet-retrieve/pkg/grpc/proto/pb"
	"sheet-retrieve/utils"
	"context"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	// Data    interface{} `json:"data,omitempty"`
}

func (h *handler) RetrieveSheet(ctx *gin.Context) {
	fmt.Println("coming inside rsheet handler")
	context, cancel := context.WithTimeout(context.Background(), 1000000000000)
	defer cancel()
	req := &pb.RetrieveSheetRequest{}
	resp, err := h.grpcClient.RetrieveSheet(context, req)
	if err != nil {
		s, _ := status.FromError(err)
		utils.APIResponse(ctx, s.Message(), s.Code(), http.MethodGet, nil)
		return
	}
	var msg = defaultSuccessMsg
	var code = codes.OK
	var find_emails = make([]*pb.Email, 0)
	if resp.Email == nil {
		resp.Email = find_emails
	}
	utils.APIResponse(ctx, msg, code, http.MethodGet, resp)
	return
}

