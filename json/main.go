package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `
  {
    "key" : "value",
    "key2" : {
      "key4" : 123
    }
  }

  `

	// empty interface can take any type
	// everything that was Unmarshalled using empy interface is of type interface
	// therefore, if I want to add integers together i need to specify that i'm adding an
	// integer like so  obj.(float64)
	var obj map[string]interface{}

	// takes a string and turns it into a slice of bytes.
	// then it will take that slice of bytes and put it in obj map
	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(obj)

	// json.Marshal returns a slice of bytes, which need to be converted to string
	// in order to look like json.
	bs, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	println(string(bs))

}
