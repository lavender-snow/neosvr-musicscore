package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ScoreCreateController(scores []string) string {
	url := ""

	maxLength := 0
	for _, score := range scores {
		values := strings.Split(score, "_")
		i, _ := strconv.Atoi(values[1])
		if i > maxLength {
			maxLength = i
		}
	}

	units := make([]string, maxLength, maxLength)

	for _, score := range scores {
		values := strings.Split(score, "_")
		scale := values[0]
		unitNo, _ := strconv.Atoi(values[1])
		unitNo--
		if units[unitNo] == "" {
			units[unitNo] = scale
		} else {
			units[unitNo] = fmt.Sprintf("%s,%s", units[unitNo], scale)
		}
	}

	scoreData := fmt.Sprintf("%s;", strings.Join(units, ";"))
	log.Printf(scoreData)

	return url
}
