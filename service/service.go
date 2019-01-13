package service

import (
	"errors"
	"log"
	. "moderation-service/repository"
	"strings"
)

type Service struct {
	Repository
}

func (s Service) ValidateText(text string) error {
	objectionableData := s.Repository.GetObjectionableData()
	if isObjectionable(text, objectionableData) {
		return errors.New("Objectionable data found")
	} else {
		return nil
	}
}

func isObjectionable(text string, objectionableData map[string][]string) bool {
	textArray := strings.Split(text, " ")
	for _, element := range textArray {
		if _, ok := objectionableData[element]; ok {
			log.Println("Objectionable content found: " + element)
			return true
		}
	}
	return false
}
