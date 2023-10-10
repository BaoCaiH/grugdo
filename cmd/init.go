/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Grug init config",
	Long:  "Grug check and create config dir, file, and table. Trust grug!",
	Run:   initGrugDb,
}

func dbExists(db *sql.DB) bool {
	var exists bool
	results, err := db.Query("SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='grug_command');")
	check(err)
	results.Next()
	err = results.Scan(&exists)
	results.Close()
	check(err)

	return exists
}

func initGrugDb(cmd *cobra.Command, args []string) {
	home := os.Getenv("HOME")
	grugDir := home + "/.grugdo"
	grugDb := grugDir + "/grugdo.db"

	if _, err := os.Stat(grugDir); os.IsNotExist(err) {
		fmt.Println("Grug tried to find config dir.")
		fmt.Printf("Grug can't find `%s` at $HOME\n", highlight(".grugdo"))
		var permission string
		fmt.Print("Can grug create dir? (Y/n) ")
		fmt.Scanln(&permission)
		if permission == "" || permission == "y" || permission == "Y" {
			fmt.Printf("Grug create `%s` dir now!\n", highlight(".grugdo"))
			err = os.MkdirAll(grugDir, 0755) // Users have read and execute, admin have all
			check(err)
			fmt.Printf("Grug created `%s`, let go!\n", highlight(".grugdo"))
		} else {
			fmt.Println("Grug can't work without it. Grug bye!")
			os.Exit(1)
		}
		fmt.Println()
	}

	if _, err := os.Stat(grugDb); os.IsNotExist(err) {
		fmt.Println("Grug tried to find config file.")
		fmt.Printf("Grug can't find `%s` at $HOME/.grugdo\n", highlight("grugdo.db"))
		var permission string
		fmt.Print("Can grug create file? (Y/n) ")
		fmt.Scanln(&permission)
		if permission == "" || permission == "y" || permission == "Y" {
			fmt.Printf("Grug create `%s` file now!\n", highlight("grugdo.db"))
			file, err := os.Create(grugDb)
			check(err)
			file.Close()
			fmt.Printf("Grug created `%s`, let go!\n", highlight("grugdo.db"))
		} else {
			fmt.Println("Grug can't work without it. Grug bye!")
			os.Exit(1)
		}
		fmt.Println()
	}

	db, err := sql.Open("sqlite3", grugDb)
	check(err)
	defer db.Close()

	// var exists bool
	// results, err := db.Query("SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='grug_command');")
	// check(err)
	// results.Next()
	// err = results.Scan(&exists)
	// results.Close()
	// check(err)

	if !dbExists(db) {
		_, err := db.Exec("CREATE TABLE grug_command(name TEXT PRIMARY KEY, type TEXT, command TEXT);")
		check(err)
		fmt.Printf("Grug created `%s` table.\n", highlight("grug_command"))
	}

	fmt.Println("Grug done init!")
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
