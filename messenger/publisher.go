package messenger

import (
	"log"
	"os"
	"github.com/nats-io/go-nats"
	"github.com/KoganezawaRyouta/hey-test/http/client"
    "time"
)

func NewPublisher(urls, subj string) error {

	nc, err := nats.Connect(urls)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for {
		data := client.GetIssues()
		for _, m := range data {
            nc.Publish(subj, []byte(m.Title))
            nc.Publish(subj, []byte(string(m.Number)))
            time.Sleep(1 * time.Second)
		}
		nc.Flush()
		if err := nc.LastError(); err != nil {
			return err
		}
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.SetPrefix("[Nats pub] ")
}
