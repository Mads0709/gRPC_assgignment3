package main

import (
	proto "chittyChat/gRPC"
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChatServer // Necessary
	name                          string
	port                          int
	vectorClock                   string
}

var port = flag.Int("port", 8101, "server port number")

func main() {
	// Get the port from the command line when the server is run
	flag.Parse()

	// Create a server struct
	server := &Server{
		name:        "ChittyChat",
		port:        *port,
		vectorClock: "<0,0,0>",
	}

	// Start the server
	go startServer(server)

	// Keep the server running until it is manually quit
	for {

	}
}

func startServer(server *Server) {

	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	proto.RegisterChatServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

// the medhod inside the service Chat in proto that connects the client and the server
func (c *Server) ChatService(ctx context.Context, in *proto.MessageFromClient) (*proto.MessageFromServer, error) {
	log.Printf("User: %s , Message: %s , Timespamp: %s", in.Name, in.Msg, in.Timestamp)
	return &proto.MessageFromServer{
		Name:      c.name,
		Timestamp: c.vectorClock,
	}, nil
}
