/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"quanty/ast"
	quantyerror "quanty/error"
	"quanty/parser"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// astCmd represents the ast command
var astCmd = &cobra.Command{
	Use:   "ast <path>",
	Short: "Print file's AST",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Args:    cobra.ExactArgs(1),
	Version: "0.0.1",

	PreRunE: func(cmd *cobra.Command, args []string) error {

		path := args[0]
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			return fmt.Errorf("Path [%s] doesn't exist", path)
		}

		if info.IsDir() {
			var files []string
			err := filepath.Walk(path, findFiles(&files))
			if err != nil {
				return fmt.Errorf("Not Quanty files found in [%s]", path)
			}
		}

		// for _, arg := range args {
		// 	info, err := os.Stat(arg)

		// 	if info.IsDir() {
		// 		return fmt.Errorf("[%s] isn't a Quanty file", arg)
		// 	}
		// }

		if cmd.HasAvailableFlags() {
			var flag *pflag.Flag
			flag = cmd.Flag("format")
			if flag != nil && flag.Value.String() != "json" && flag.Value.String() != "dump" {
				return fmt.Errorf("Format type [%s] is not valid", flag.Value.String())
			}
		}

		return nil
	},
	Run: run,
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// astCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// astCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(astCmd)

	astCmd.Flags().StringP("format", "f", "dump", "Format to print <json|dump>")
}

func run(cmd *cobra.Command, args []string) {

	path := args[0]

	var files []string
	filepath.Walk(path, findFiles(&files))

	flagFormat := cmd.Flag("format").Value.String()

	for _, file := range files {
		fileRead, errRead := ioutil.ReadFile(file)
		check(errRead)

		fileParsed, errParsed := parser.ParseFile(
			&ast.Source{
				Name:  file,
				Input: string(fileRead),
			},
		)
		print(fileParsed, errParsed, flagFormat)
	}
}

func print(file *ast.File, err *quantyerror.Error, format string) {

	var data string = "No data"
	var toPrint interface{} = file

	if err != nil {
		toPrint = err
	}

	if format == "json" {
		jsonD, _ := json.MarshalIndent(toPrint, "", "  ")
		data = string(jsonD)
	} else {
		data = ast.Dump(toPrint)
	}

	fmt.Println(data)
}
