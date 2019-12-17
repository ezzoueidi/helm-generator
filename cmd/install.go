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
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	//"os/exec"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install your helm charts",
	Long: `

Helm-generator generates for you your helm charts
with an interactive shell to fill all your environment
variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("install called")
		// get the name of the repository
		scanner_repo := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter the name of the folder: ")
		scanner_repo.Scan()
		repo := scanner_repo.Text()
		obj, err := os.Open(".")
		if err != nil {
			log.Fatalf("failed opening directory: %s", err)
		}
		defer obj.Close()
		list, _ := obj.Readdirnames(0)
		for _, name := range list {
			if name == repo {
				scanner_answer := bufio.NewScanner(os.Stdin)
				fmt.Print("It looks like there is a folder already with that name. \nDo you want me to use it? [Y/n]: ")
				scanner_answer.Scan()
				answer := scanner_answer.Text()
				if answer == "Y" || answer == "y" {
					if _, err := os.Stat(".git/"); err != nil {
						if os.IsNotExist(err) {
							scanner_git := bufio.NewScanner(os.Stdin)
							fmt.Println("This folder is not a git repository. \n Do you want me ")
						} else {
							// other error
						}
					}

				} else if answer == "N" || answer == "n" {

				}
			}

		}
		// scans the number of charts
		scanner_chart := bufio.NewScanner(os.Stdin)
		fmt.Print("How much helm charts you'd like to create? : ")
		scanner_chart.Scan()
		number_charts := scanner_chart.Text()
		nb_chart, err := strconv.Atoi(number_charts)
		arr := make([]string, 0)
		if len(number_charts) != 0 && err == nil {
			for i := 1; i <= nb_chart; i++ {
				scanner_nb := bufio.NewScanner(os.Stdin)
				fmt.Print("Enter the name of the ", i, " chart: ")
				scanner_nb.Scan()
				name_chart := scanner_nb.Text()
				if len(name_chart) != 0 && err == nil {
					arr = append(arr, name_chart)
				}
			}

		} else if len(number_charts) != 0 && err != nil {
			fmt.Println("You have entered a non integer value for your charts!")
		}
		//create the folder and files needed for the helm charts

	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
