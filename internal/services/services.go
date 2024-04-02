package services

import (
	"sort"
	pb "github.com/silverflin/go-rpc/goguide"
	"context"
	"github.com/silverflin/go-rpc/internal/model"
    "log"
)


type ProductListServer struct {
	pb.UnimplementedProductsServer
}


func (s ProductListServer) GetProductsByPrice(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductList, error) {
    log.Printf("New Request %v", req.ProductName)
	filteredProductList := &pb.ProductList{Products: make([]*pb.Product, 0)}
	for _, prod := range model.GetAllProducts() {
		if prod.Name == req.ProductName {
			filteredProductList.Products = append(filteredProductList.Products, prod)
		}
	}
    OrderProductListByPrice(filteredProductList)

	return filteredProductList, nil
}

func  OrderProductListByPrice(l *pb.ProductList){
    sort.Slice(l.Products, func(i, j int) bool {
        return l.Products[i].CurrentPrice > l.Products[j].CurrentPrice
    })
}