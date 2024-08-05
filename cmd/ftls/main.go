package main

import (
	"log"
	"os"
	"time"

	"github.com/stanleymw/ftls/libclient"
)

func main() {
	client := ftls.NewClient("cc.pem", "ck.unencrypted.pem", ":1337")
	// Lets simulate a session
	start := time.Now()

	file, _ := os.Create("data.txt")
	delta := time.Since(start)

	client.RetrieveFile(file)
	// log.Printf("Transferred %d bytes in %s (%f MB/s)", recv.Size, delta, float64(recv.Size)/delta.Seconds()/1000000)

	log.Printf("Took %s", delta)

	// log.Println(client.GetDir())
	// log.Println(client.ListDir())

	// writer.Encode(protocol.CLOSE_CONNECTION)
	log.Println("done!")
}
