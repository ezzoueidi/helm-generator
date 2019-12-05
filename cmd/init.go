/*
Copyright © 2019 Naeil Ezzoueidi naeil@nzoueidi.com

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
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

//var debug int

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init for your current environment",
	Long: `
Check if you have all the existing tools (e.g helm, tiller),
otherwise it will install them for you.`,

	Run: func(cmd *cobra.Command, args []string) {
		//section for debug implementation
		if _, err := os.Stat(os.Getenv("HOME") + "/.helm/"); os.IsNotExist(err) {
			fmt.Println("your home dir is")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	//initCmd.Flags().IntVar(&debug, "debug", 0, "debug level")
}
