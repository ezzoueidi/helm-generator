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
	"os/exec"
	"runtime"
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
		if runtime.GOOS == "linux" {
			helmFolder := "/home/nzoueidi/.helm" // TODO will be filled as environment variable
			if stat, err := os.Stat(helmFolder); err == nil && stat.IsDir() /*&& isInstalled("helm")*/ {
				//helm is installed and initialized - nothing todo
				fmt.Println("Helm is installed and inited")
			} else {
				fmt.Println("test not valid")
				fmt.Println(helmFolder)

			}

		}

	},
}

func isInstalled(toolname string) bool {
	cmd := exec.Command("command", "-v", toolname)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func init() {
	rootCmd.AddCommand(initCmd)

	//initCmd.Flags().IntVar(&debug, "debug", 0, "debug level")
}
