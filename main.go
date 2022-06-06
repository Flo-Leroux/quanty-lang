package main

import (
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
	rootPath := "./app"

	// tree :=
	language.NewTree(rootPath)

	// jsonRes, err := json.MarshalIndent(tree, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(jsonRes))
	// fmt.Println()
}
