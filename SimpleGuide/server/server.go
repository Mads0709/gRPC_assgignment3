package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"net"
	"os"
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
	cl := Clients{int64(rq.Id), int(rq.Port), 0, rc}
	list = append(list, cl)
	for i := 0; i < len(list); i++ {
		list[i].stream.Send(&proto.ResponsMessage{Respond: "", Id: rq.Id})
	}

	wg.Add(1)
	//---- this is to delete a user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if input == "exit" {
		wg.Done()
		wg.Done()
	}
	//----
	wg.Wait()
	//We wait and thereby keep the stream open
	return nil
}

func (c *Server) PopulateChatMessage(con context.Context, msg *proto.ChatMessage) (*proto.ErrorMessage, error) {
	for i := 0; i < len(list); i++ {
		if list[i].clientId != msg.Id {
			list[i].stream.Send(&proto.ResponsMessage{Respond: msg.Message, Id: msg.Id})
		}
	}

	log.Printf("message: %s from id: %d vectorclock: <%d>", msg.Message, msg.Id, msg.Vectorclock)
	// sends message onto the stream
	return &proto.ErrorMessage{Message: "error!"}, nil
}
