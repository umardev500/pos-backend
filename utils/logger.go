package utils

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func LogError(err error) string {
	refCode := uuid.New().String()
	log.Error().Str("refCode", refCode).Err(err).Msg("an error occurred")
	return refCode
}
