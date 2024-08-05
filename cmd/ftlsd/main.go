package main

import (
	"bufio"
	"crypto/tls"
	"encoding/gob"
	"log"
	"net"
	"os"

	"github.com/stanleymw/ftls/internal/protocol"
)

func main() {
	cer, err := tls.LoadX509KeyPair("./sc.pem", "./sk.unencrypted.pem")

	if err != nil {
		log.Fatal(err)
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{cer},
		// ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	ln, err := tls.Listen("tcp", ":1337", conf)

	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	log.Println("Listening...")
	for {

		conn, err := ln.Accept()
		log.Println("New connection...")
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	defer log.Println("Closing connection!")

	// r := bufio.NewReader(conn)
	writer := gob.NewEncoder(conn)
	reader := gob.NewDecoder(conn)
	for {
		var recv byte
		log.Println("waiting for opcode...")
		err := reader.Decode(&recv)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("decoded | got code %d", recv)

		switch recv {
		case protocol.GET_SERVER_INFO:
			writer.Encode(protocol.Response{Body: "Test data transmission"})

		case protocol.GET_DIRECTORY_LIST:
			read, _ := os.ReadDir(".")
			var ret []string
			for _, i := range read {
				ret = append(ret, i.Name())
			}

			log.Print(writer.Encode(ret))

		case protocol.GET_CURRENT_DIRECTORY:
			dir, _ := os.Getwd()

			log.Print(writer.Encode(dir))

		case protocol.CLOSE_CONNECTION:
			return

		case protocol.RETRIEVE_FILE:
			file, _ := os.Open("data.txt")
			defer file.Close()
			fio := bufio.NewReader(file)

			stat, _ := file.Stat()

			// z := protocol.File{Size: stat.Size(), Name: stat.Name(), IsDir: stat.IsDir()}

			log.Println(writer.Encode(stat))

			fio.WriteTo(conn)
			log.Print("File sent!")
		}
	}
}
