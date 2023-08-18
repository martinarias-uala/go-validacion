package utils

import (
	"log"

	"github.com/google/uuid"
)

func GetUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return uuid.String()
}
