package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	proto "simpleguide/grpc"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id         int
	portNumber int
}

var (
	clientPort = flag.Int("cPort", 0, "client port number")
	serverPort = flag.Int("sPort", 0, "server port number (should match the port used for the server)")
	clientId   = flag.Int("cId", 0, "client id number")
)

//go run server/server.go -port 5454
//go run client/client.go -cPort 8080 -sPort 5454 -cId x

func main() {
	// Parse the flags to get the port for the client
	flag.Parse()

	// Create a client
	client := &Client{
		id:         *clientId,
		portNumber: *clientPort,
	}

	// Wait for the client (user) to ask for the time
	go registerToServer(client)

	for {
	}
}

func listenOnServer(server *grpc.Server) {
	
}

func registerToServer(client *Client) {
	// Connect to the server
	serverConnection, _ := connectToServer()

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)

	registerReturnMessage, err := serverConnection.RegisterToServer(context.Background(), &proto.Request{
		Id:   int64(client.id),
		Port: int64(client.portNumber),
	})

	if err != nil {
		log.Printf(err.Error())
	} else {
		//log.Printf("Server %s says the time is %s\n", timeReturnMessage.ServerName, timeReturnMessage.Time)
		m, e := registerReturnMessage.Recv()
		if e != nil {
			log.Printf("Error %s \n", err.Error())
			return
		}
		log.Printf("Server %s", m.Respond)
	}
	for scanner.Scan() {
		input := scanner.Text()
		log.Printf("my message: %s", input)
		serverConnection.PopulateChatMessage(context.Background(), &proto.ChatMessage{
			Message: input,
			Id:      int64(client.id),
		})
	}

}

func connectToServer() (proto.RegisterClientClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *serverPort)
	} else {
		log.Printf("Connected at port %d\n", *serverPort)
	}
	return proto.NewRegisterClientClient(conn), nil
}
