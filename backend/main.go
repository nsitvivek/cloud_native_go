package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	pb "../api"
	"golang.org/x/net/context"

	//pb "github.com/campoy/just-for-func/say-grpc/api"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port o listen to")
	flag.Parse()
	log.Printf("listening to port %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("Could not listen t oport %d: %v", *port, err)
	}
	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("Could not serve: %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("fcould not create temp file %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("flite failed close  %s:%v", f.Name, err)
	}
	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed %s", data)
	}
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read temp file %v", err)
	}
	return &pb.Speech{Audio: data}, nil
}
