/*
Copyright © 2023 BaoCaiH
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var executable, overwrite bool

// addCmd represents the save command
var addCmd = &cobra.Command{
	Use:   "add [name] [command]",
	Short: "Grug add command",
	Long:  "Grug save command to DB. With name and type. Grug will default to remind",
	Run:   addCommand,
}

func addCommand(cmd *cobra.Command, args []string) {
	fmt.Println()

	if len(args) < 2 {
		fmt.Println("Grug need at least [name] and [command]. Bye!")
		os.Exit(1)
	}

	db := dbLoad()
	defer db.Close()

	name := args[0]
	command := strings.Join(args[1:], " ")

	saveMode := "R"
	if executable {
		saveMode = "E"
	}
	fmt.Printf("Received command: %s\n", command)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Mode: %s\n", saveMode)
	fmt.Printf("Overwrite: %t\n", overwrite)
	fmt.Println()

	if cmdExists(db, name) && !overwrite {
		fmt.Printf("Grug found saved `%s`\n", highlight(name))
		fmt.Println("but Grug not allow to overwrite. Grug bye!")
		os.Exit(0)
	}

	_, err := db.Exec(
		fmt.Sprintf("INSERT OR REPLACE INTO grug_command(name, type, command) VALUES('%s','%s','%s');", name, saveMode, command),
	)
	check(err)

	fmt.Printf("Grug saved `%s`\n", highlight(name))
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().BoolVarP(&overwrite, "overwrite", "o", false, "Grug overwrite exist save.")
	addCmd.PersistentFlags().BoolVarP(&executable, "executable", "e", false, "Grug can execute. Grug can always remind.")
}
