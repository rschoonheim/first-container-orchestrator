package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/gofiber/fiber/v3/client"
	"gopkg.in/yaml.v2"
	"os"
)

func init() {
	os.MkdirAll("storage/cli", 0755)

	cc := client.New()

	cert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	caCert, err := os.ReadFile("certs/server-ca.crt")
	if err != nil {
		panic(err)
	}
	certPool.AppendCertsFromPEM(caCert)

	cc.SetTLSConfig(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	})

	resp, err := cc.Get("https://127.0.0.1:7331/")
	if err != nil {
		panic(err)
	}

	println(resp.Status())

}

func main() {
	println("Container Orchestrator CLI\n")

	argc := len(os.Args)
	if argc < 2 {
		println("Usage: container-orchestrator-cli <command> [options]")

		println("Commands:")
		println("  apply	-	Apply configuration")
		os.Exit(1)
	}

	command := os.Args[1]

	// todo: Implement CLI functionality
	// todo: Refactor the switch statement to maintainable code

	switch command {
	case "apply":
		if argc < 3 {
			println("Usage: container-orchestrator-cli apply <project-path>")
			os.Exit(1)
		}
		projectPath := os.Args[2]

		println("Applying configuration from project path:", projectPath)

		// Read main.yaml
		//
		file, err := os.Open(projectPath + "/main.yaml")
		if err != nil {
			println("Failed to apply configuration. main.yaml not found in project path.")
			os.Exit(1)
		}
		defer file.Close()

		var mainConfig mainConfiguration
		err = yaml.NewDecoder(file).Decode(&mainConfig)
		if err != nil {
			println("Failed to apply configuration. main.yaml is invalid.")
			os.Exit(1)
		}

		// Validate the received configuration

		// Start application of the configuration

	}

}
