package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	pb "grupo5/proto"
)

func generateId() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Intn(100))
}

type Server struct {
	pb.UnimplementedWishListServiceServer
}

func (s *Server) create(ctx context.Context, in *pb.CreateWishlistReq) (*pb.CreateWishListRes, error) {
	fmt.Println("Creating wishlist" + in.Wishlist.Name)
	return &pb.CreateWishListRes{
		WishlistId: generateId(),
	}, nil
}

func (s *Server) add(ctx context.Context, in *pb.AddItemReq) (*pb.AddItemRes, error) {
	return &pb.AddItemRes{
		ItemId: generateId(),
	}, nil
}
