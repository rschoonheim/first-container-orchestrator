package main

import (
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
)

func init() {

	var storagePaths = []string{
		"storage/containers",
		"storage/volumes",
		"storage/networks",
		"storage/images",
	}

	for _, path := range storagePaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(path, 0755)
			slog.Info("Created storage directory", "path", path)
		}
	}
}

func main() {

	// Create listener
	//
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
	defer listener.Close()

	// Create gRPC server
	//
	server := grpc.NewServer()

	server.Serve(listener)
}

//cmd := exec.Cmd{
//	Path:   "/bin/sh",
//	Dir:    "/",
//	Stdin:  os.Stdin,
//	Stdout: os.Stdout,
//	Stderr: os.Stderr,
//	SysProcAttr: &syscall.SysProcAttr{
//		Chroot:     "/home/romano/Downloads/alpine-minirootfs-3.20.3-x86",
//		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
//	},
//}
//
//if err := cmd.Run(); err != nil {
//	fmt.Println(err)
//}
