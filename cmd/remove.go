/*
Copyright © 2023 BaoCaiH
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var removeCmd = &cobra.Command{
	Use:   "rm [name]",
	Short: "Grug remove saved command",
	Long:  "Too many commands hurt Grug. Grug can remove some. If you want.",
	Run:   removeCommand,
}

func removeCommand(cmd *cobra.Command, args []string) {
	fmt.Println()

	db := dbLoad()
	defer db.Close()

	for _, command := range args {
		cmdRemove(db, command)
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
