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
	"quanty/build"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve <path>",
	Short: "A brief description of your command",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	// Args:    cobra.ExactArgs(1),
	Version: "0.0.1",

	PreRunE: func(cmd *cobra.Command, args []string) error {

		// path := args[0]
		// info, err := os.Stat(path)
		// if os.IsNotExist(err) {
		// 	return fmt.Errorf("Path [%s] doesn't exist", path)
		// }

		// if info.IsDir() {
		// 	var files []string
		// 	err := filepath.Walk(path, findFiles(&files))
		// 	if err != nil {
		// 		return fmt.Errorf("Not Quanty files found in [%s]", path)
		// 	}
		// }

		return nil
	},
	Run: serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serve(cmd *cobra.Command, args []string) {

	// path := args[0]

	build.Serve()
}
