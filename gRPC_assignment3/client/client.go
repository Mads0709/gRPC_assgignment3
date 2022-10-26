package main

import (
	"bufio"
	proto "chittyChat/gRPC"
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	name        string
	portNumber  int
	vectorClock string
}

var (
	clientPort = flag.Int("cPort", 0, "client port number")
	serverPort = flag.Int("sPort", 0, "server port number (should match the port used for the server)")
)

func main() {
	// Parse the flags to get the port for the client
	flag.Parse()
	// Create a client
	client1 := &Client{
		name:        "Mads",
		portNumber:  *clientPort,
		vectorClock: "<0,0,0>",
	}

	client2 := &Client{
		name:        "Mathias",
		portNumber:  *clientPort,
		vectorClock: "<0,0,0>",
	}

	/*client2 := &Client{
		name:        "user2",
		portNumber:  *clientPort,
		vectorClock: "<0,0,0>",
	}*/

	// Wait for the client (user) to ask for the time
	//log.Printf(client.name, client.portNumber, client.vectorClock)
	go waitForTimeRequest1(client1)
	go waitForTimeRequest2(client2)

	for {

	}
}

func waitForTimeRequest1(client *Client) {
	// Connect to the server
	serverConnection, _ := connectToServer()

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		log.Printf("Client asked for the message with input: %s\n", input)

		// Ask the server for the time
		messageReturnMessage, err := serverConnection.ChatService(context.Background(), &proto.MessageFromClient{
			Name:      client.name,
			Msg:       "hej",
			Timestamp: client.vectorClock,
		})

		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Printf("Server %s says the time is %s\n", messageReturnMessage.Name, messageReturnMessage.Timestamp)
		}
	}
}

func waitForTimeRequest2(client *Client) {
	// Connect to the server
	serverConnection, _ := connectToServer()

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		log.Printf("Client asked for the message with input: %s\n", input)

		// Ask the server for the time
		messageReturnMessage, err := serverConnection.ChatService(context.Background(), &proto.MessageFromClient{
			Name:      client.name,
			Msg:       "hej",
			Timestamp: client.vectorClock,
		})

		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Printf("Server %s says the time is %s\n", messageReturnMessage.Name, messageReturnMessage.Timestamp)
		}
	}
}

func connectToServer() (proto.ChatClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *serverPort)
	} else {
		log.Printf("Connected to the server at port %d\n", *serverPort)
	}
	return proto.NewChatClient(conn), nil
}
