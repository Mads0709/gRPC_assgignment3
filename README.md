# gRPC_assgignment3
To run the program open a terminal and write the following command: 
`go run server/server.go -port 5400`
This comman will start a server on port 5400

Then open any number of new terminals and write the following command: 
`go run client/client.go -cPort 8080 -sPort 5454 -cId x`
Here x has to be substituted with a unique integer
