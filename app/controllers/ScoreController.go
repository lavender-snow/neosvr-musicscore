package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lavender-snow/neosvr-musicscore/app/models"
)

// CreateScoreController 引き渡された入力情報を基に譜面を作成
func CreateScoreController(scores []string) (string, error) {

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
	id, err := models.InsertScoreData(scoreData)
	if err != nil {
		log.Printf(err.Error())
		return "", err
	}

	return id, nil
}

// GetScoreController 引き渡されたIDに対応する譜面情報を返却
func GetScoreController(id string) (string, error) {
	score, err := models.SelectScoreData(id)
	if err != nil {
		return "", err
	}
	return score, nil
}
