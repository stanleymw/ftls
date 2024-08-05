package ftls

// ftls Client Library

import (
	"crypto/tls"
	"encoding/gob"
	"io"
	"log"
	// "os"
	// "time"

	"github.com/stanleymw/ftls/internal/protocol"
)

type Client struct {
	tls.Certificate
	*tls.Conn

	*gob.Encoder
	*gob.Decoder
}

func NewClient(cert string, key string, addr string) Client {
	cer, err := tls.LoadX509KeyPair(cert, key)

	if err != nil {
		log.Fatal(err)
	}

	con, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cer},
	})

	if err != nil {
		log.Fatal(err)
	}

	writer := gob.NewEncoder(con)
	reader := gob.NewDecoder(con)

	return Client{Certificate: cer, Conn: con, Encoder: writer, Decoder: reader}
}

func (client Client) RetrieveFile(dest io.ReaderFrom) {
	recv := protocol.File{}

	client.Encoder.Encode(protocol.RETRIEVE_FILE)
	client.Decoder.Decode(&recv)

	dest.ReadFrom(io.LimitReader(client.Conn, recv.Size))
}

func (client Client) GetDir() string {
	var recv string

	client.Encoder.Encode(protocol.GET_CURRENT_DIRECTORY)
	client.Decoder.Decode(&recv)

	return recv
}

func (client Client) ListDir() protocol.DirList {
	var recv protocol.DirList

	client.Encoder.Encode(protocol.GET_DIRECTORY_LIST)
	client.Decoder.Decode(&recv)

	return recv
}

func (client Client) CloseConnection() error {
	return client.Encoder.Encode(protocol.CLOSE_CONNECTION)
}
