package networkSender

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
)

var verbose bool
var port int
var host string
var protocol string

var NetworkEcho = &cobra.Command{
	Use:  "network",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.Printf("Running a networkSender echo to port '%d' via the '%s' protocol.\n", port, protocol)
		}

		address := fmt.Sprintf("%s:%d", host, port)
		log.Printf("Address: '%s'", address)

		conn, err := net.Dial(protocol, address)
		if err != nil {
			panic(err)
		}

		defer func() {
			if verbose {
				log.Println("Closing connection.", conn)
			}

			if conn != nil {
				err := conn.Close()
				if err != nil {
					panic(err)
				}
			}
		}()

		if verbose {
			log.Println("Got a successful connection: ", conn)
		}

		for i, v := range cmd.Flags().Args() {
			b := []byte(v)
			w, err := conn.Write(b)
			if err != nil {
				panic(err)
			}

			if verbose {
				log.Printf("Wrote arg %d (%s) with %d bytes.\n", i, v, w)
			}

			if verbose {
				log.Println("Listening for response now...")
			}

			response := make([]byte, w)
			n, err := conn.Read(response)
			if err != nil {
				if err == io.EOF {
					log.Printf("Received an EOF.")
					continue
				} else {
					log.Fatal("An unexpected error occurred: ", err)
				}
			}

			log.Printf("Wrote %d bytes, got %d bytes: %s.", w, n, string(response[:n]))
		}
	},
}

func init() {
	NetworkEcho.Flags().BoolVarP(&verbose, "verbose", "v", false, "Set this flag to have more verbose output.")
	NetworkEcho.PersistentFlags().IntVarP(&port, "port", "p", 3031, "The port to echo to.")
	NetworkEcho.PersistentFlags().StringVarP(&protocol, "protocol", "t", "tcp", "The networkSender protocol.")
	NetworkEcho.PersistentFlags().StringVarP(&host, "hostname", "n", "127.0.0.1", "The hostname or IP address to connect to.")
}
