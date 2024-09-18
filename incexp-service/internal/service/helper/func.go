package helper

import (
	"incexp-service/logger"

	"github.com/google/uuid"
)

func ParseUUID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		logger.Error("Error parsing UUID: ", err, id)
		return uuid.Nil, err
	}
	return parsedID, nil
}
