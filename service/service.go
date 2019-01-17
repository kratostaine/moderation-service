package service

import (
	"errors"
	"log"
	"strings"

	. "github.com/kratostaine/moderation-service/repository"
)

type Service struct {
	Repository
}

func (s Service) ValidateText(text string) error {
	objectionableData := s.Repository.GetObjectionableData()
	maxLength := s.Repository.GetMaxLengthOfObjectionableText()
	if isObjectionable(text, objectionableData, maxLength) {
		return errors.New("Objectionable data found")
	} else {
		return nil
	}
}

func isObjectionable(text string, objectionableData map[string][]string, maxLength int) bool {
	text = strings.ToLower(text)
	textArray := strings.Split(text, " ")
	for _, element := range textArray {
		if _, ok := objectionableData[element]; ok {
			log.Println("Objectionable content found: " + element)
			return true
		}
	}
	return isObjectionableOnDeepSearch(text, objectionableData, maxLength)
}

func isObjectionableOnDeepSearch(text string, objectionableData map[string][]string, maxLength int) bool {
	text = strings.Replace(text, " ", "", -1)
	textLen := len(text)
	for pos, char := range text {
		var tempText strings.Builder
		tempText.WriteString(string(char))
		for i := pos + 1; i < textLen && i < maxLength; i++ {
			tempText.WriteString(string(text[i]))
			if _, ok := objectionableData[tempText.String()]; ok {
				ignorableText := objectionableData[tempText.String()]
				falsePositive := false
				for _, element := range ignorableText {
					if strings.Contains(text, element) {
						log.Println("False positive found: " + element)
						falsePositive = true
						break
					}
				}
				if !falsePositive {
					log.Println("Objectionable content found: " + tempText.String())
					return true
				}
			}
		}
	}
	return false
}
