package helper

import (
	"budget-service/logger"

	"github.com/google/uuid"
)

func ParseUUID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		logger.Error("Error parsing UUID: ", err)
		return uuid.Nil, err
	}
	return parsedID, nil
}
