package main

import (
	"github.com/RainJoe/ms/service/user/config"
	"github.com/RainJoe/ms/service/user/service"
	pb "github.com/RainJoe/ms/service/user/proto/user"
	"google.golang.org/grpc"
	_"github.com/lib/pq"
	"log"
	"net"
)


func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", ":8080")
	 if err != nil {
		 log.Fatal(err)
	}
	srv := grpc.NewServer()
	svc := service.NewUserService()
	pb.RegisterUserServiceServer(srv, svc)
	 if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
