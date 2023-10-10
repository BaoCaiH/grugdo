/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	name                  string
	executable, overwrite bool
)

func cmdExists(db *sql.DB, name string) bool {
	var exists bool
	results, err := db.Query(
		fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM grug_command WHERE name='%s');", name),
	)
	check(err)
	results.Next()
	err = results.Scan(&exists)
	results.Close()
	check(err)

	return exists
}

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: saveCommand,
}

func saveCommand(cmd *cobra.Command, args []string) {
	home := os.Getenv("HOME")
	grugDir := home + "/.grugdo"
	grugDb := grugDir + "/grugdo.db"
	fmt.Println()

	db, err := sql.Open("sqlite3", grugDb)
	check(err)
	defer db.Close()
	if !dbExists(db) {
		fmt.Printf("Grug can't find DB. Try run:\n\t%s\n", highlight("grugdo init"))
		os.Exit(1)
	}

	command := strings.Join(args, " ")

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

	_, err = db.Exec(
		fmt.Sprintf("INSERT OR REPLACE INTO grug_command(name, type, command) VALUES('%s','%s','%s');", name, saveMode, command),
	)
	check(err)

	fmt.Printf("Grug saved `%s`\n", highlight(name))
}

func init() {
	rootCmd.AddCommand(saveCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	saveCmd.PersistentFlags().StringVarP(&name, "name", "n", "grug", "Name this command for Grug.")
	saveCmd.PersistentFlags().BoolVarP(&overwrite, "overwrite", "o", false, "Grug overwrite exist save.")
	saveCmd.PersistentFlags().BoolVarP(&executable, "executable", "e", false, "Grug can execute. Grug can always remind.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
