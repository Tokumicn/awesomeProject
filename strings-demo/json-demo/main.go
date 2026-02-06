package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	input := `{
  "phone": "",
  "bill_period": ""
}`

	var validJSONs map[string]interface{}
	err := json.Unmarshal([]byte(input), &validJSONs)
	if err != nil {
		fmt.Println(err)
		return
	}

}
