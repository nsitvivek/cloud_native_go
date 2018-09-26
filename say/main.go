package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "../api"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav file where the output will be written")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("bad usage")
		os.Exit(1)
	}
	// conn, err := net.Dial("tcp", *backend)
	// if err != nil {
	// 	log.Fatalf("Could not connect to the backend %s: %v", *backend, err)
	// }
	// defer conn.Close()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	client := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{Text: os.Args[1]}
	res, err := client.Say(context.Background(), text)

	if err != nil {
		log.Fatalf("Could not say text %s: %v", text.Text, err)
	}

	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("Could not write to file %s: %v", *output, err)
	}
}
