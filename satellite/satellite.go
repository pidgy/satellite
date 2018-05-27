package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"github.com/trashbo4t/satellite/common"
	"net"
	"os"
)

var (
	tcp, udp bool
)

func init() {
	debug   := flag.Bool("d", false, "set log level to debug")
	info    := flag.Bool("i", false, "set log level to info")
	warning := flag.Bool("w", false, "set log level to warning")
	error   := flag.Bool("e", false, "set log level to error")
	stderr  := flag.Bool("s", false, "output to stderr")
	tcp     = *flag.Bool("t", true,  "listen on TCP")
	udp     = *flag.Bool("u", false, "listen on UDP")
	flag.Parse()

	switch {
	case *debug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case *info:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case *warning:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case *error:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
	if *stderr {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		zerolog.TimeFieldFormat = ""
	} else {
		// Redirect your logging to a file descriptor here
	}
}

// handle the data stream, send the response and close the connection
func handleTcp(c chan net.Conn) {
	log.Debug().Msg("TCP handler is up")
	restart:
	for {
		conn := <- c
		data := make([]byte, 1024)
		reqLen, err := conn.Read(data)
		if err != nil {
			log.Error().Err(err).Msg("Could not read message")
			conn.Close()
			continue restart
		}
		log.Debug().Msgf("Received Message: %s", string(data))
		log.Debug().Msgf("Total bytes: %d", reqLen)
		data, err = common.HandleJSON(data)
		if err != nil {
			log.Error().Err(err).Msg("Handling JSON")
			conn.Close()
			continue restart
		} else {
			log.Debug().Msgf("Sending reply: %s", string(data))
			conn.Write(data)
		}
		conn.Close()
	}
}

// Launch 5 handlers to accept multiple connections synchronously
func launchTcpHandlers() chan net.Conn {
	log.Debug().Msg("Launching TCP Handlers")
	connChan := make(chan net.Conn)
	i := 0
	for i < 5 {
		go handleTcp(connChan)
		i++
	}
	return connChan
}

// Accept connections and send them down the connection channel
func spinTcp(listener net.Listener) {
	defer listener.Close()
	log.Info().Msgf("Listening on %s:%s", common.CONN_HOST, common.CONN_PORT)
	connChan := launchTcpHandlers()
	for {
		conn, err := listener.Accept()
		log.Debug().Msg("Accepted incoming connection")
		if err != nil {
			log.Panic().Err(err).Msg("Could not accept")
		}
		connChan <- conn
	}
}

// Spin
// Start up the server
func main() {
	log.Info().Msg("Started Satellite \"Beep Boop Bop Beep\" ")
	if tcp {
		l, err := common.TcpConn()
		if err != nil {
			log.Panic().Err(err).Msg("Cannot listen")
		}
		spinTcp(l)
	} else if udp {
		c, err := common.UdpConn()
		if err != nil {
			log.Panic().Err(err).Msg("Cannot listen")
		}
		// TODO handle UDP
		_ = c
	} else {
		log.Panic().Msg("No TCP or UDP listener....something has gone horribly wrong!")
	}
}
