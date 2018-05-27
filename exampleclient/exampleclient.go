package main

import (
	"github.com/rs/zerolog"
	"os"
	"github.com/rs/zerolog/log"
	"github.com/trashbo4t/satellite"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	zerolog.TimeFieldFormat = ""
	log.Info().Msg("This is an example test of the satellite Json server")
	obj := satellite.MyObject{
		Magic: 1,
		Msg:   "This is a test message",
	}
	conn, err := satellite.SendJson(obj)
	if err != nil {
		log.Error().Err(err).Msg("Error sending Json")
		return
		log.Info().Msg("Object sent to the server bytes without error")
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			log.Error().Err(err).Msg("Error receiving Json")
			return
		}
		log.Info().Msgf("Received %d bytes from satellite", n)
		msg, err := satellite.AsJson(data)
		if err != nil {
			log.Error().Err(err)
			return
		}
		log.Info().Msgf("Received message: %s", msg)
	}
}