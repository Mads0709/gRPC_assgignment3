package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	proto "simpleguide/grpc"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedRegisterClientServer
	name string
	port int
}

type Clients struct {
	clientId   int
	clientPort int
}

type ClientList struct {
	clientsList []Clients
}

var port = flag.Int("port", 0, "server port number")
var list []Clients

func main() {
	flag.Parse()

	server := &Server{
		name: "serverName",
		port: *port,
	}

	go startServer(server)

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

	proto.RegisterRegisterClientServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}

}

func (c *Server) RegisterToServer(rq *proto.Request, rc proto.RegisterClient_RegisterToServerServer) error {
	log.Printf("Client ID %d Client port %d ", rq.Id, rq.Port)
	cl := Clients{int(rq.Id), int(rq.Port)}
	list = append(list, cl)
	return rc.Send(&proto.ResponsMessage{Respond: "register success!"})
}

func (c *Server) PopulateChatMessage(con context.Context, msg *proto.ChatMessage) (*proto.ErrorMessage, error) {
	var e error = nil
	log.Printf("message: %s from: %d", msg.Message, msg.Id)
	// sends message onto the stream
	return &proto.ErrorMessage{Message: "error!"}, e
}
