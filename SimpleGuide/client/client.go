package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"math"
	"os"
	proto "simpleguide/grpc"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var lamportTimeStamp = int64(0)

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

//go run server/server.go -port 5400
//go run client/client.go -cPort 8080 -sPort 5400 -cId x
//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/proto.proto

func main() {
	// Parse the flags to get the port for the client
	flag.Parse()

	// Create a client
	client := &Client{
		id:          *clientId,
		portNumber:  *clientPort,
		vectorClock: 1, //note that we call it a vector clock but it is a lambport time stamp
	}
	registerToServer(client)
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
		log.SetFlags(0)
		log.Printf("%d Connected to Server", m.Id)
	}

	go listenOnServer(serverStream, client)

	for scanner.Scan() {
		input := scanner.Text()
		client.vectorClock += 1
		lamportTimeStamp += 1
		if input == "logoff" {
			log.Printf("--logging off--")
			serverConnection.LogOffServer(context.Background(), &proto.LogOffMessage{Id: int64(*clientId)})
			return
		} else {
			log.Printf("my message: %s <%d>", input, client.vectorClock) //
			serverConnection.PopulateChatMessage(context.Background(), &proto.ChatMessage{
				Message:     input,
				Id:          int64(client.id),
				Vectorclock: int64(client.vectorClock),
			})
		}
	}
}

func listenOnServer(serverStream proto.RegisterClient_RegisterToServerClient, client *Client) {
	for {
		resp, err := serverStream.Recv()

		if err != nil {
			log.Printf("Error %s", err)
		}
		if resp.Respond == "" { //then we know that it is someone that connected to the server
			log.Printf("%d Connected to Server", resp.Id)

		} else {
			lamportTimeStamp = int64(math.Max(float64(resp.Vectorclock), float64(lamportTimeStamp)) + 1) //lamport timestamp
			client.vectorClock = int(lamportTimeStamp)
			log.Printf("Message from %d: %s <%d>", resp.Id, resp.Respond, lamportTimeStamp)
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
