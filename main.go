package main

import (
	pb "cybershen-go/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// gRPC implementation
type NumberServiceServer struct {
	pb.UnimplementedNumberServiceServer
}

var (
	globalMax   int32
	globalMaxID int32
)

func (s *NumberServiceServer) StreamNumbers(stream pb.NumberService_StreamNumbersServer) error {
	// log details
	if p, ok := peer.FromContext(stream.Context()); ok {
		log.Printf("New stream from %v", p.Addr)
	}

	for {
		// reception of the client message
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Error during reception : %v", err)
			return err
		}

		log.Printf("Message received : ID=%d, num=%d", msg.Id, msg.Num)

		// updating global max value
		if msg.Num > globalMax {
			globalMax = msg.Num
			globalMaxID = msg.Id
			log.Printf("New global max : ID=%d, max=%d", globalMaxID, globalMax)
		}
		currentMax := globalMax
		currentMaxID := globalMaxID

		// sending message to client
		response := &pb.ServerMessage{
			MessageID: currentMaxID,
			Max:       currentMax,
		}
		if err := stream.Send(response); err != nil {
			log.Printf("Error during sending : %v", err)
			return err
		}

		log.Printf("Message sent to client : messageID=%d, max=%d", response.MessageID, response.Max)
	}
}

func main() {
	// Init server
	server := grpc.NewServer()
	pb.RegisterNumberServiceServer(server, &NumberServiceServer{})

	// start listener
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051 : %v", err)
	}

	log.Println("Listening on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to boot %v", err)
	}
}
