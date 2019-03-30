package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)  // Decode stdin.
	enc := json.NewEncoder(os.Stdout) // Encode to stdout.
	type em struct{}
	badWords := map[string]em{
		"bollocks": em{},
		"grody":    em{},
	}
	for {
		var in map[string]interface{}
		if err := dec.Decode(&in); err != nil {
			log.Println(err)
			return
		}
		for k := range in { // Remove bad keys from input.
			if _, ok := badWords[k]; ok {
				delete(in, k)
			}
		}
		if err := enc.Encode(&in); err != nil {
			log.Println(err)
		}
	}
}
