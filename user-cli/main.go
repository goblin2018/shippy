package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/goblin2018/shippy/user-service/proto/user"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

func createUser(ctx context.Context, service micro.Service, user *pb.User) error {
	client := pb.NewUserService("shippy.user.service", service.Client())
	rsp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}
	fmt.Println("Response: ", rsp.User)
	return nil

}

func main() {
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your name",
			},
			&cli.StringFlag{
				Name:  "email",
				Usage: "E-mail",
			},
			&cli.StringFlag{
				Name:  "company",
				Usage: "Company name",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Password",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) error {
			log.Println(c)
			name := c.String("name")
			email := c.String("email")
			company := c.String("company")
			password := c.String("password")

			log.Println("test: ", name, email, company, password)

			ctx := context.Background()
			user := &pb.User{
				Name:    name,
				Email:   email,
				Company: company,
				Pasword: password,
			}

			if err := createUser(ctx, service, user); err != nil {
				log.Println("error creating user: ", err.Error())
				return err
			}
			return nil
		}),
	)

}
