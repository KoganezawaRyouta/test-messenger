package cmd

import (
	"log"
	"os"

	"github.com/KoganezawaRyouta/hey-test/messenger"
	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
    "github.com/nats-io/go-nats"
)

var subCmd = &cobra.Command{
	Use:   "subscriber",
	Short: "hey",
	Long:  "hey",
	Run: func(cmd *cobra.Command, args []string) {
		messenger.NewSubscriber(nats.DefaultURL, "hey")
	},
}

var pubCmd = &cobra.Command{
	Use:   "publisher",
	Short: "hey",
	Long:  "hey",
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[Publisher Server] ")
		go func() {
			pid := os.Getpid()
			pidInfo, _ := ps.FindProcess(pid)
			log.Printf("start")
			log.Printf(" PID          : %d\n", pidInfo.Pid())
			log.Printf(" PPID         : %d\n", pidInfo.PPid())
			log.Printf(" Process name : %s\n", pidInfo.Executable())
			pp, _ := ps.FindProcess(pidInfo.PPid())
			log.Printf(" Parent process name : %s\n", pp.Executable())
			errsCh <- messenger.NewPublisher(nats.DefaultURL, "hey")
		}()
		log.Fatal("terminated", <-errsCh)
	},
}

func init () {
    RootCmd.AddCommand(subCmd)
    RootCmd.AddCommand(pubCmd)
}