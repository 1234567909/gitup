package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RegisterGRPC(port int, register func(s *grpc.Server)) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%d", port))
	if err != nil {
		log.Fatalf("faile to listen: %v", err)
		return nil, nil
	}

	s := grpc.NewServer()
	//反射接口支持查询
	reflection.Register(s)
	register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return nil, nil
	}
}
