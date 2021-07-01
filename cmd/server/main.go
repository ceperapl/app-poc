package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/ceperapl/app-poc/dist"
	deliveryGrpc "github.com/ceperapl/app-poc/pkg/delivery/grpc"
	"github.com/ceperapl/app-poc/pkg/delivery/grpc/pb"
	"github.com/ceperapl/app-poc/pkg/repository/memory"
	"github.com/ceperapl/app-poc/pkg/usecase"
	grpc_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	// gRPC server
	defaultServerPort = "9090"

	// gRPC gateway (RESTful API)
	defaultGatewayPort = "8080"

	// Web-UI
	defaultWebUIPort = "8088"
)

func main() {
	serverPort := flag.String("server.port", defaultServerPort, "port of gRPC server")
	gatewayPort := flag.String("gateway.port", defaultGatewayPort, "port of gateway server")
	webUIPort := flag.String("webui.port", defaultWebUIPort, "port of web ui")

	flag.Parse()

	memoryTaskRepo, err := memory.NewTaskRepo()
	if err != nil {
		log.Fatal(err)
	}
	taskService := usecase.NewTaskService(memoryTaskRepo)
	list, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", *serverPort))
	if err != nil {
		log.Fatalf("SOMETHING HAPPEN: %v", err)
	}

	server := grpc.NewServer()
	deliveryGrpc.NewTaskServerGrpc(server, taskService)
	log.Println("GRPC Server Run at ", fmt.Sprintf("0.0.0.0:%s", *serverPort))

	doneC := make(chan error)

	go func() {
		doneC <- server.Serve(list)
	}()

	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%s", *serverPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("SOMETHING HAPPEN: %v", err)
	}
	gatewayMux := grpc_runtime.NewServeMux()
	err = pb.RegisterTasksHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		log.Fatalf("SOMETHING HAPPEN: %v", err)
	}

	log.Println("HTTP Gateway Run at ", fmt.Sprintf("0.0.0.0:%s", *gatewayPort))
	go func() {
		doneC <- http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", *gatewayPort), gatewayMux)
	}()

	webUImux := http.NewServeMux()
	staticFiles := dist.GetEmbedFS()

	var staticFS = http.FS(staticFiles)
	fs := http.FileServer(staticFS)
	// Serve static files
	webUImux.Handle("/assets/", fs)
	// Handle all other requests
	webUImux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var path = req.URL.Path
		log.Println("Serving request for path", path)
		w.Header().Add("Content-Type", "text/html")

		file, err := staticFS.Open("/index.html")
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(w, file)
	})
	log.Println("Web UI Run at ", fmt.Sprintf("0.0.0.0:%s", *webUIPort))
	go func() {
		doneC <- http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", *webUIPort), webUImux)
	}()

	if err := <-doneC; err != nil {
		log.Fatal(err)
	}
}
