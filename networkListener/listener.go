package networkListener

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

var Listener = &cobra.Command{
	Use: "listen",
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.Printf("Running a networkSender listener on port '%d' via the '%s' protocol.\n", Port, Protocol)
		}

		address := fmt.Sprintf("%s:%d", Host, Port)
		log.Printf("Address: '%s'", address)

		ln, err := net.Listen(Protocol, address)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("DEBUG: from main: ")
		defer func() {
			if Verbose {
				log.Println("Closing connection.", ln)
			}

			if ln != nil {
				err := ln.Close()
				if err != nil {
					panic(err)
				}
			}
		}()

		if Verbose {
			log.Println("Listening...")
		}

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}

			go func(c net.Conn) {
				defer func() {
					err := c.Close()
					if err != nil {
						panic(err)
					}
				}()

				buffer := make([]byte, 1024)
				n := 0
				n, err := c.Read(buffer)
				if err != nil {
					if err == io.EOF {
						log.Println("Connection closed by client.")
					} else {
						log.Fatal(err)
					}
				}

				if Verbose {
					log.Printf("Copied %d bytes into the buffer.\n", n)
				}

				if Verbose {
					log.Printf("%s", buffer)
				}

				// This version works.
				w, err := conn.Write(buffer[:n])
				// The below seems to not always work as I'd expect.
				// If you replace the above line with the below, then the `go run echo.go network` will send the
				// string, but it seems somewhat undeterministic what will happen.
				//w, err := conn.Write(buffer)
				if err != nil {
					log.Fatal("An error occurred writing the buffer: ", err)
				}

				if Verbose {
					log.Printf("Wrote %d bytes to the connection.\n", w)
				}
			}(conn)
		}
	},
}

func init() {
	Listener.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Set this flag to have more verbose output.")
	Listener.PersistentFlags().IntVarP(&Port, "port", "p", 3031, "The port to echo to.")
	Listener.PersistentFlags().StringVarP(&Protocol, "protocol", "t", "tcp", "The networkSender protocol.")
	Listener.PersistentFlags().StringVarP(&Host, "hostname", "n", "127.0.0.1", "The hostname or IP address to connect to.")
}
