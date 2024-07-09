package utils

// making this extra file from the repo
import (
	"strconv"

	"github.com/rs/zerolog/log"
)

func ParseInt(str string) int64 {
	integer, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		log.Error().Err(err).Msg("Utils.Int64")
	}

	return integer
}
