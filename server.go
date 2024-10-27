package main

import (
	"context"
	"log"
	"net"

	pb "github.com/kaungmyathan22/golang-coffee-grpc/coffeeshop_proto"
	"google.golang.org/grpc"
)

// Create a struct and embed our UnimplementCofeeShopServer
// We provide a full implementation to the methods that this embedded struct specifies down below
type server struct {
	pb.UnimplementedCoffeeShopServer
}

// Get a menu, stream the response back to the client
func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "Black Coffee",
		},
		&pb.Item{
			Id:   "2",
			Name: "Americano",
		},
		&pb.Item{
			Id:   "3",
			Name: "Vanilla Soy Chai Latte",
		},
	}

	// weird little gimmicky way to "simulate" streaming data back to the client
	// ideally this is representing sending the pieces of data we have back as we get them
	for i, _ := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}

	return nil
}

// Place an order
func (s *server) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABC123",
	}, nil
}

// Get order status
func (s *server) GetOrderStatus(context context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "IN PROGRESS",
	}, nil
}

func main() {

	// setup a listener on port 9001
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a new grpc server
	grpcServer := grpc.NewServer()

	// register our server struct as a handle for the CoffeeShopService rpc calls that come in through grpcServer
	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	// Serve traffic
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
