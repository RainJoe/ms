package main

import (
	"github.com/RainJoe/ms/api/config"
	"github.com/RainJoe/ms/api/gateway"
	pb "github.com/RainJoe/ms/service/user/proto/user"
	"google.golang.org/grpc"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}
	log.Printf("loaded config %#v", config.Cfg)
	conn, err := grpc.Dial(config.Cfg.UserSeviceAddr.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userServiceClient := pb.NewUserServiceClient(conn)
	router := gateway.Handler(userServiceClient)

	if err := router.Run(config.Cfg.Port); err != nil {
		log.Fatalln(err)
	}
}
