/*
Copyright © 2022 Justin Trugman | @jtrugnan

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

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
		kid, err := readYaml(args)

		if err != nil {
			log.Fatal(err)
		}

		cmdProgram, cmdArray := executeTC(kid)

		executeCommand(cmdProgram, cmdArray)
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

func readYaml(args []string) (model.Kid, error) {

	if len(args) != 1 {
		err := fmt.Errorf("incorrect number of arguments")
		return model.Kid{}, err
	}

	yfile, err := ioutil.ReadFile(args[0])

	if err != nil {
		return model.Kid{}, err
	}

	var kid model.Kid

	err2 := yaml.Unmarshal(yfile, &kid)

	if err2 != nil {
		return model.Kid{}, err2
	}

	return (kid), err
}

func executeTC(kid model.Kid) (string, []string) {
	// TODO: Change log.fatal to return error string
	cmdArray := []string{"qdisc"}
	cmdProgram := "tc"

	switch kid.Job.Command.Operation {
	case "delete":
		cmdArray = append(cmdArray, kid.Job.Command.Operation, "dev", kid.Job.Command.Port, "root")
		return cmdProgram, cmdArray
	case "add", "change":
		cmdArray = append(cmdArray, kid.Job.Command.Operation)
	default:
		log.Fatal("Operation not supported")
	}

	cmdArray = append(cmdArray, "dev", kid.Job.Command.Port, "root", "netem")

	if len(kid.Job.Command.Bitrate.BitrateUnit) > 0 {
		// TODO: Clean up nested if statement
		switch kid.Job.Command.Bitrate.BitrateUnit {
		case "kbit", "mbit", "gbit":
			cmdArray = append(cmdArray, "rate", fmt.Sprintf("%f", kid.Job.Command.Bitrate.BitrateValue)+kid.Job.Command.Bitrate.BitrateUnit)

		default:
			log.Fatal("Bitrate Unit not supported")
		}
	}

	if kid.Job.Command.PktLoss != 0 {
		// TODO: Clean up nested if statement
		switch kid.Job.Command.PktLoss {
		case kid.Job.Command.PktLoss:
			cmdArray = append(cmdArray, "loss", fmt.Sprintf("%f", kid.Job.Command.PktLoss)+"%")
		default:
			log.Fatal("Incorrect PKT Loss type")
		}
	}

	if kid.Job.Command.Latency != 0 {
		cmdArray = append(cmdArray, "delay", fmt.Sprintf("%f", kid.Job.Command.Latency))
		if kid.Job.Command.Jitter != 0 {
			cmdArray = append(cmdArray, fmt.Sprintf("%f", kid.Job.Command.Jitter))
		}
	}

	fmt.Print(cmdProgram, cmdArray)

	return cmdProgram, cmdArray

}

func executeCommand(cmdProgram string, cmdArray []string) string {
	cmd := exec.Command(cmdProgram, cmdArray...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	return string(output)
}
