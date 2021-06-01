package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	api "phudt/week4/internal/api"
	"phudt/week4/internal/repo"
	"phudt/week4/internal/service"

	cli "github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	port    = "0.0.0.0:50000"
	connStr = "localhost:50000"
)

const (
	KEY = "Authentication"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "server",
				Usage:  "Start server",
				Action: serverAction,
			},
			{
				Name:   "client",
				Usage:  "start client",
				Action: clientAction,
			},
		}}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func clientAction(c *cli.Context) error {
	conn, err := grpc.Dial(connStr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := api.NewPatientServiceClient(conn)
	req := &api.AddRequest{
		Fullname: "Tran Phong Phu",
		Address:  "HCM",
		Birthday: "2021-05-27 01:01:01",
		Gender:   api.Gender_MALE,
		Age:      100,
	}
	res, err := client.Add(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func serverAction(c *cli.Context) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	// Connect DB
	dsn := "root:root@tcp(127.0.0.1:3306)/covid?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db.AutoMigrate(&api.Patient{})
	if err != nil {
		return err
	}
	patientRepo := repo.NewPatientRepo(db)
	srv := service.NewService(patientRepo)
	//
	myInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			keys := md.Get(KEY)
			if len(keys) > 0 {
				authentication := keys[0]
				fmt.Println(authentication)
			}
		}
		return handler(ctx, req)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(myInterceptor))
	api.RegisterPatientServiceServer(s, &srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
