package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"
	"sync"

	proto "simpleguide/grpc"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedRegisterClientServer
	name string
	port int
}

type Clients struct {
	clientId    int64
	clientPort  int
	vectorClock int
	stream      proto.RegisterClient_RegisterToServerServer
}

var port = flag.Int("port", 0, "server port number")
var list []Clients

var wg sync.WaitGroup

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
	log.SetFlags(0)
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	proto.RegisterRegisterClientServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}

}

func (c *Server) RegisterToServer(rq *proto.Request, rc proto.RegisterClient_RegisterToServerServer) error {
	log.Printf("ID %d Connected to server", rq.Id)
	cl := Clients{int64(rq.Id), int(rq.Port), 0, rc}
	list = append(list, cl)
	for i := 0; i < len(list); i++ {
		list[i].stream.Send(&proto.ResponsMessage{Respond: "", Id: rq.Id})
	}

	wg.Add(1)
	wg.Wait()
	//We wait and thereby keep the stream open
	return nil
}

func (c *Server) PopulateChatMessage(con context.Context, msg *proto.ChatMessage) (*proto.ErrorMessage, error) {
	for i := 0; i < len(list); i++ {
		if list[i].clientId != msg.Id {
			list[i].stream.Send(&proto.ResponsMessage{Respond: msg.Message, Id: msg.Id, Vectorclock: msg.Vectorclock})
		}
	}

	log.Printf("%s from id: %d <%d>", msg.Message, msg.Id, msg.Vectorclock)
	// sends message onto the stream
	return &proto.ErrorMessage{Message: "error!"}, nil
}
