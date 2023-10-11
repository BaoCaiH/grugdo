/*
Copyright © 2023 BaoCaiH
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	execute    bool
	extraFlags []string
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [name] [args]",
	Short: "Grug run saved command.",
	Long: `Grug will remind. Grug will run if asked nicely.
Grug can save bash file too (Grug can't handle things like *).`,
	Run: runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
	fmt.Println()

	if len(args) < 1 {
		fmt.Println("Grug need command name to run. Bye!")
		os.Exit(1)
	}

	db := dbLoad()
	defer db.Close()

	name := args[0]

	if !cmdExists(db, name) {
		fmt.Printf("Grug can't find %s", name)
		os.Exit(1)
	}
	commandType, command := cmdSelect(db, name)

	commandParts := append(strings.Split(command, " "), args[1:]...)
	commandParts = append(commandParts, extraFlags...)
	fmt.Println(commandParts)

	if len(outfile) > 0 {
		writeFile(outfile, strings.Join(commandParts, " "))
		os.Exit(0)
	}

	if !execute || commandType != "E" {
		os.Exit(0)
	}

	executor := exec.Command(commandParts[0], commandParts[1:]...)
	executor.Stdout = os.Stdout

	if err := executor.Run(); err != nil {
		fmt.Println("Grug can't run command. Nooooo! ", err)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().BoolVarP(&execute, "execute", "e", false, "Grug execute the command.")
	runCmd.PersistentFlags().StringVarP(&outfile, "file", "f", "", "Grug save the command to file.")
	runCmd.PersistentFlags().StringSliceVarP(&extraFlags, "extra-flags", "x", []string{}, "Grug can add flag to command.")
}
