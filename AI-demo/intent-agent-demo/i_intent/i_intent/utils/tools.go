package utils

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ExtractContinuousDigits(input string) []string {
	re := regexp.MustCompile(`\d+`)
	return re.FindAllString(input, -1)
}

func ExtractFloat(input string) float64 {
	re := regexp.MustCompile(`\d+\.?\d*`)
	match := re.FindString(input)
	if match != "" {
		val, err := strconv.ParseFloat(match, 64)
		if err == nil {
			return val
		}
	}
	return 0
}

func ExtractJSONFromString(input string) map[string]interface{} {
	//re := regexp.MustCompile(`\{.*?\}`)
	//matches := re.FindAllString(input, -1)

	var validJSONs map[string]interface{}
	err := json.Unmarshal([]byte(input), &validJSONs)
	if err != nil {
		return nil
	}

	return validJSONs
}

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func FormatSlotJSON(slot map[string]interface{}) ([]byte, error) {
	return json.Marshal(slot)
}

func FixJSON(badJSON string) map[string]interface{} {
	fixedJSON := strings.ReplaceAll(badJSON, "'", "\"")
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(fixedJSON), &result); err != nil {
		return nil
	}
	return result
}
