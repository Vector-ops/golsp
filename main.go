package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/Vector-ops/golsp/lsp"
	"github.com/Vector-ops/golsp/rpc"
)

func main() {
	logger := getLogger("G:/dev/golang/golsp/log.txt")
	logger.Println("Hey, I started!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Recieved message with method: %s", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, we could not parse this: %s", err)
		}

		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		msg := lsp.NewIntializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout

		writer.Write([]byte(reply))
		logger.Print("Sent the reply")
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you did not give me a good file.")
	}

	return log.New(logfile, "[golsp(edu)]", log.Ldate|log.Ltime|log.Lshortfile)
}
