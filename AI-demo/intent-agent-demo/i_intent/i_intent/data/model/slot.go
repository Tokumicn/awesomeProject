package models

import (
	"fmt"
	"strings"
)

type Slot struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func IsSlotFullyFilled(slots []Slot) bool {
	for _, slot := range slots {
		if slot.Value == "" {
			return false
		}
	}
	return true
}

func GetRawSlot(parameters []Slot) []Slot {
	output := make([]Slot, len(parameters))
	for i, item := range parameters {
		output[i] = Slot{
			Name:  item.Name,
			Desc:  item.Desc,
			Type:  item.Type,
			Value: "",
		}
	}
	return output
}

func UpdateSlot(newValues map[string]interface{}, targetSlots []Slot) {
	for name, val := range newValues {
		for i := range targetSlots {
			if targetSlots[i].Name == name {
				targetSlots[i].Value = val.(string)
				break
			}
		}
	}
}

func FormatNameValueForLogging(slots []Slot) string {
	logStrings := make([]string, len(slots))
	for i, slot := range slots {
		logStrings[i] = fmt.Sprintf("name: %s, Value: %s", slot.Name, slot.Value)
	}
	return strings.Join(logStrings, "\n")
}
