package networkSender

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
)

var Verbose bool
var Port int
var Host string
var Protocol string

var NetworkEcho = &cobra.Command{
	Use:  "network",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.Printf("Running a networkSender echo to port '%d' via the '%s' protocol.\n", Port, Protocol)
		}

		address := fmt.Sprintf("%s:%d", Host, Port)
		log.Printf("Address: '%s'", address)

		conn, err := net.Dial(Protocol, address)
		if err != nil {
			panic(err)
		}

		defer func() {
			if Verbose {
				log.Println("Closing connection.", conn)
			}

			if conn != nil {
				err := conn.Close()
				if err != nil {
					panic(err)
				}
			}
		}()

		if Verbose {
			log.Println("Got a successful connection: ", conn)
		}

		for i, v := range cmd.Flags().Args() {
			b := []byte(v)
			written, err := conn.Write(b)
			if err != nil {
				panic(err)
			}

			if Verbose {
				log.Printf("Wrote arg %d (%s) with %d bytes.\n", i, v, written)
			}
		}

		if Verbose {
			log.Println("Listening for response now...")
		}

		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			if err == io.EOF {
				log.Printf("Received an EOF.")
			} else {
				log.Fatal("An error occurred: ", err)
			}
		} else {
			log.Printf("Got response %d bytes: %s.", n, response[:n])
		}
	},
}

func init() {
	NetworkEcho.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Set this flag to have more verbose output.")
	NetworkEcho.PersistentFlags().IntVarP(&Port, "port", "p", 3031, "The port to echo to.")
	NetworkEcho.PersistentFlags().StringVarP(&Protocol, "protocol", "t", "tcp", "The networkSender protocol.")
	NetworkEcho.PersistentFlags().StringVarP(&Host, "hostname", "n", "127.0.0.1", "The hostname or IP address to connect to.")
}
