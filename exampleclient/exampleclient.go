package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/trashbo4t/satellite/common"
	"os"
)

func main() {
	// The common json object shared between client/server
	obj := common.MyObject{
		Magic: 1,
		Msg:   "This is a test message",
	}

	log.Logger = log.Output(os.Stderr)
	zerolog.TimeFieldFormat = ""
	log.Info().Msg("This is an example test of the satellite Json server")

	conn, err := common.SendJson(obj)
	if err != nil {
		log.Error().Err(err).Msg("Error sending Json")
		return
	}
	log.Info().Msg("Object sent to the server without error")
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		log.Error().Err(err).Msg("Error receiving Json")
		return
	}
	log.Info().Msgf("Received %d bytes from satellite", n)
	obj, err = common.AsJson(data)
	if err != nil {
		log.Error().Err(err).Msg("Error converting Json")
		return
	}
	log.Info().Msgf("Json response from satellite: %s", obj.Response)
}