/*
Copyright © 2023 BaoCaiH
*/
package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	home      = os.Getenv("HOME")
	grugDir   = home + "/.grugdo"
	grugDb    = grugDir + "/grugdo.db"
	highlight = color.New(color.FgCyan).SprintFunc()
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

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

func cmdSelect(db *sql.DB, name string) (string, string) {
	var commandType, command string
	results, err := db.Query(
		fmt.Sprintf("SELECT type, command FROM grug_command WHERE name='%s';", name),
	)
	check(err)
	results.Next()
	err = results.Scan(&commandType, &command)
	results.Close()
	check(err)

	return commandType, command
}

func cmdRemove(db *sql.DB, name string) bool {
	_, err := db.Exec(
		fmt.Sprintf("DELETE FROM grug_command WHERE name='%s';", name),
	)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Printf("Grug removed %s if it exists\n", highlight(name))
	return true
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

func dbLoad() *sql.DB {
	db, err := sql.Open("sqlite3", grugDb)
	check(err)
	if !dbExists(db) {
		fmt.Printf("Grug can't find DB. Try run:\n\t%s\n", highlight("grugdo init"))
		os.Exit(1)
	}

	return db
}

func selectAll(db *sql.DB) string {
	var name, commandType, command string

	commands := "name\ttype\tcommand\n"

	results, err := db.Query("SELECT * FROM grug_command ORDER BY 1;")
	check(err)
	defer results.Close()
	for results.Next() {
		err = results.Scan(&name, &commandType, &command)
		check(err)
		commands += fmt.Sprintf("%s\t%s\t%s\n", name, commandType, command)
	}

	return commands
}

func writeFile(filename, contents string) {
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	b, err := w.WriteString(contents)
	check(err)

	fmt.Printf("Wrote %s bytes\n", highlight(b))

	w.Flush()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grugdo",
	Short: "Grug do stuffs",
	Long: "Grug forget stuffs. Grug can't remember commands. Machine remembers commands " +
		"for Grug. Grug say what Grug want. Machine do it for Grug. Or remind Grug the commands.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Grug help")
}
