package print_struct

import (
	"encoding/json"
	"fmt"
)

/*
Print a struct with string
*/
func PrintStructString(s interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("PrintStructString==>json.Marshal==>failed==>[%s] \n", err.Error())
	}
	// fmt.Printf("PrintStructString==>[%s] \n ", string(b))
	return string(b)
}

/*
Print a struct with json
*/
func PrintStructJson(s interface{}) string {
	b, err := json.MarshalIndent(s, "", "")
	if err != nil {
		fmt.Printf("PrintStructJson==>json.Marshal==>failed==>[%s] \n", err.Error())
	}
	// fmt.Printf("PrintStructJson==>[%s] \n", string(b))
	return string(b)
}
