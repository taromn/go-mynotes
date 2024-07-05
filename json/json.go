package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	bt := []byte(`{"country":"japan", "localtime":["7:00am", "8:00am"],"decimal":1.1, "cities": {"shibuya": "hachiko", "ikebukuro":"ikehukuro"}}`)

	var decoded map[string]interface{}

	if err := json.Unmarshal(bt, &decoded); err != nil {
		panic(err)
	}

	fmt.Println(decoded)

	cn := decoded["country"]
	fmt.Println(cn)

	lt := decoded["localtime"]
	fmt.Println(lt)

	ct := decoded["cities"]
	fmt.Println(ct) // returns map[ikebukuro:ikehukuro shibuya:hachiko]

	dm := decoded["decimal"]
	fmt.Printf("%T", dm)
}
