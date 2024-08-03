package main

import (
	"crypto/tls"
	"encoding/gob"
	"io"
	"log"
	"os"
	"time"

	"github.com/stanleymw/ftls/protocol"
)

func main() {
	cer, err := tls.LoadX509KeyPair("./cc.pem", "./ck.unencrypted.pem")

	con, err := tls.Dial("tcp", ":1337", &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cer},
	})

	if err != nil {
		log.Fatal(err)
	}

	// Lets simulate a session
	writer := gob.NewEncoder(con)
	reader := gob.NewDecoder(con)

	start := time.Now()
	writer.Encode(protocol.RETRIEVE_FILE)

	recv := protocol.FtlsFile{}
	reader.Decode(&recv)

	file, err := os.Create("data.txt")

	file.ReadFrom(io.LimitReader(con, recv.Size))

	delta := time.Since(start)
	log.Printf("Transferred %d bytes in %s (%f MB/s)", recv.Size, delta, float64(recv.Size)/delta.Seconds()/1000000)

	writer.Encode(protocol.CLOSE_CONNECTION)
	log.Println("done!")
}
