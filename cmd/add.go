// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"github.com/mitchellh/go-homedir"
	"github.com/s0rbus/tri/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item in the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := homedir.Dir()
		if err != nil {
			log.Fatalf("%v", err)
		}
		items, err := todo.ReadItems(homeDir + "/.tridos.json")
		if err != nil {
			log.Fatalf("%v", err)
		}
		for _, x := range args {
			items = append(items, todo.Item{Text: x})
		}
		err = todo.SaveItems(homeDir+"/.tridos.json", items)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
