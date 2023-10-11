/*
Copyright © 2023 BaoCaiH
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var outfile string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "Grug list saved commands",
	Long:  "Grug list command names, types, and command",
	Run:   listCommand,
}

func listCommand(cmd *cobra.Command, args []string) {

	fmt.Println()

	db := dbLoad()
	defer db.Close()

	commands := selectAll(db)

	if len(outfile) > 0 {
		writeFile(outfile, strings.ReplaceAll(commands, "\t", "|"))
		os.Exit(0)
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, commands)
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&outfile, "file", "f", "", "Grug write list to file.")
}
