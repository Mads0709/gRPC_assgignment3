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
	id          int
	portNumber  int
	vectorClock int
}

var (
	clientPort        = flag.Int("cPort", 0, "client port number")
	serverPort        = flag.Int("sPort", 0, "server port number (should match the port used for the server)")
	clientId          = flag.Int("cId", 0, "client id number")
	serverStream, err proto.RegisterClient_RegisterToServerClient
)

//go run server/server.go -port 5454
//go run client/client.go -cPort 8080 -sPort 5454 -cId x

func main() {
	// Parse the flags to get the port for the client
	flag.Parse()

	// Create a client
	client := &Client{
		id:          *clientId,
		portNumber:  *clientPort,
		vectorClock: 0,
	}

	go registerToServer(client)
	for {
		//This infinit forloop keeps the program running
	}
}

func registerToServer(client *Client) {
	// Connect to the server
	serverConnection, _ := connectToServer()

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)

	serverStream, err := serverConnection.RegisterToServer(context.Background(), &proto.Request{
		Id:   int64(client.id),
		Port: int64(client.portNumber),
	})

	if err != nil {
		log.Printf(err.Error())
	} else {
		m, e := serverStream.Recv()
		if e != nil {
			log.Printf("Error %s \n", err.Error())
			return
		}
		log.Printf("%d Connected to Server", m.Id)
	}

	go listenOnServer(serverStream)

	for scanner.Scan() {
		input := scanner.Text()
		client.vectorClock += 1
		log.Printf("my message: %s", input) //
		serverConnection.PopulateChatMessage(context.Background(), &proto.ChatMessage{
			Message:     input,
			Id:          int64(client.id),
			Vectorclock: int64(client.vectorClock),
		})
	}
}

func listenOnServer(serverStream proto.RegisterClient_RegisterToServerClient) {
	for {
		resp, err := serverStream.Recv()

		if err != nil {
			log.Printf("Error %s", err)
		}
		if resp.Respond == "" { //then we know that it is someone that connected to the server
			log.Printf("%d Connected to Server", resp.Id)
		} else {
			log.Printf("Message from %d: %s ", resp.Id, resp.Respond)
		}

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
