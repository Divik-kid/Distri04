package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	ping "github.com/Divik-kid/Distri04/ping"
	"google.golang.org/grpc"
)

var mu sync.Mutex

func main() {
	//Selects the port for each user starting at 5000 with the argument 0
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id: ownPort,
		//map of clients ID's and how many pings recieved from that client
		amountOfPings: make(map[int32]int32),

		clients: make(map[int32]ping.PingClient),
		ctx:     ctx,
	}

	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	ping.RegisterPingServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	//connects to all clients except self (maximum 3 clients)
	for i := 0; i < 3; i++ {
		port := int32(5000) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := ping.NewPingClient(conn)
		p.clients[port] = c
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p.sendPingToAll()
	}
}

type peer struct {
	ping.UnimplementedPingServer
	id            int32
	amountOfPings map[int32]int32
	clients       map[int32]ping.PingClient
	ctx           context.Context
	LampTime      int32
}

// when pinged
func (p *peer) Ping(ctx context.Context, req *ping.Request) (*ping.Reply, error) {
	id := req.Id
	p.amountOfPings[id] += 1

	rep := &ping.Reply{Amount: p.amountOfPings[id], Access: false}

	//Determine if this nodes' id is greater than the requests' author
	if req.LogTime < p.LampTime {
		//faster lamport time wins
		//fmt.Println("YES YOU CAN ACCESS")
		rep.Access = true
	} else if req.LogTime == p.LampTime {
		//bigger ID wins
		if p.id < id {
			rep.Access = true
		}
	}

	return rep, nil
}

func (p *peer) CriticalState() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("%v Has accessed the critical state", p.id)
	fmt.Println()
}

// when pinging
func (p *peer) sendPingToAll() {

	request := &ping.Request{Id: p.id, LogTime: p.LampTime}
	p.LampTime += 1
	var accessCount int
	for id, client := range p.clients {
		reply, err := client.Ping(p.ctx, request)
		if err != nil {
			fmt.Println("something went wrong")
		}
		if reply.Access {
			//count until all have said yes then access
			accessCount += 1
		} else {
			//wipe the count and dont access
			accessCount = 0
		}
		if accessCount == len(p.clients) {
			p.CriticalState()
		}

		fmt.Printf("id %v said %v\n", id, reply.Access)
	}

}
