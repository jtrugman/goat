/*
Copyright © 2022 Justin Trugman | @jtrugnan

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jtrugman/goat/model"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Read Network Configurations from a YAML file",
	Long: `File reads network configurations from a YAML file and applies them to the respective network interface
Example usage: goat file FILE_PATH

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")

		kid := readYaml(args)
		fmt.Println(kid)
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func readYaml(args []string) model.Kid {

	if len(args) != 1 {
		log.Fatal("Incorrect Number of Arguments")
	}

	yfile, err := ioutil.ReadFile(args[0])

	if err != nil {
		log.Fatal(err)
	}

	var kid model.Kid

	err2 := yaml.Unmarshal(yfile, &kid)

	if err != nil {
		log.Fatal(err2)
	}

	return (kid)
}
