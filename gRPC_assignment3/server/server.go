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

var port = flag.Int("port", 0, "server port number")

func main() {
	// Get the port from the command line when the server is run
	flag.Parse()

	// Create a server struct
	server := &Server{
		name:        "serverName",
		port:        *port,
		vectorClock: "vectorClock",
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
	log.Printf("Client with Name: %d asked for the time\n", in.Name)
	return &proto.MessageFromServer{
		Name: c.name,
	}, nil
}
