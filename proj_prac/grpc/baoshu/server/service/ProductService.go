package service

import "context"

type ProdService struct {
	UnimplementedProdServiceServer
}

func (ps *ProdService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{ProdStock: request.ProdId}, nil
}
