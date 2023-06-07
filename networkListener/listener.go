package networkListener

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

var Listener = &cobra.Command{
	Use: "listen",
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.Printf("Running a networkSender listener on port '%d' via the '%s' protocol.\n", port, protocol)
		}

		address := fmt.Sprintf("%s:%d", host, port)
		log.Printf("Address: '%s'", address)

		ln, err := net.Listen(protocol, address)
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if verbose {
				log.Println("Closing connection.", ln)
			}

			if ln != nil {
				err := ln.Close()
				if err != nil {
					panic(err)
				}
			}
		}()

		if verbose {
			log.Println("Listening...")
		}

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}

			go handleConn(conn)
		}
	},
}

func handleConn(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n := 0
		n, err2 := conn.Read(buffer)
		if err2 != nil {
			if err2 == io.EOF {
				break
			}

			log.Printf("Unexpected error: %s.", err2)
			break
		}

		if verbose {
			log.Printf("Copied %d bytes into the buffer.\n", n)
			log.Printf("%s\n", buffer)
		}

		w, err := conn.Write(buffer[:n])
		if err != nil {
			log.Fatal("Error writing: ", err)
		}

		if verbose {
			log.Printf("Wrote %d bytes to the connection.\n", w)
		}
	}
}

func init() {
	Listener.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Set this flag to have more verbose output.")
	Listener.PersistentFlags().IntVarP(&port, "port", "p", 3031, "The port to echo to.")
	Listener.PersistentFlags().StringVarP(&protocol, "protocol", "t", "tcp", "The networkSender protocol.")
	Listener.PersistentFlags().StringVarP(&host, "hostname", "n", "127.0.0.1", "The hostname or IP address to connect to.")
}
