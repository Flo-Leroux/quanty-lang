package main

import (
	"encoding/json"
	"fmt"
	"quanty/language"
)

// Func main should be as small as possible and do as little as possible by convention
// func main() {
// 	// Generate our config based on the config supplied
// 	// by the user in the flags
// 	cfgPath, err := server.ParseFlags()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	cfg, err := server.NewConfig(cfgPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Run the server
// 	cfg.Run()
// }

func main() {
	mainPath := "./app/main.qy"

	file := language.NewFile(mainPath)

	jsonRes, err := json.MarshalIndent(file, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonRes))
	fmt.Println()
}
