/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var highlight = color.New(color.FgCyan).SprintFunc()

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grugdo",
	Short: "Grug do stuffs",
	Long: "Grug forget stuffs. Grug can't remember commands. Machine remembers commands " +
		"for Grug. Grug say what Grug want. Machine do it for Grug. Or remind Grud the commands.",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	check(err)
}

// func highlight(text string) string {
// 	return color.New(color.FgYellow).SprintFunc(text)
// }

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grugdo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	home := os.Getenv("HOME")

	if _, err := os.Stat(home + "/.grugdo"); os.IsNotExist(err) {
		fmt.Println("Grug tried to find config dir.")
		fmt.Printf("Grug can't find `%s` at $HOME\n", highlight(".grugdo"))
		var permission string
		fmt.Print("Can grug create dir? (Y/n) ")
		fmt.Scanln(&permission)
		if permission == "" || permission == "y" || permission == "Y" {
			fmt.Printf("Grug create `%s` dir now!\n", highlight(".grugdo"))
			err = os.MkdirAll(home+"/.grugdo", 0755) // Users have read and execute, admin have all
			check(err)
			fmt.Printf("Grug created `%s`, let go!\n", highlight(".grugdo"))
		} else {
			fmt.Println("Grug can't work without it. Grug bye!")
			os.Exit(1)
		}
		fmt.Println()
	}

	if _, err := os.Stat(home + "/.grugdo/grugdo.db"); os.IsNotExist(err) {
		fmt.Println("Grug tried to find config file.")
		fmt.Printf("Grug can't find `%s` at $HOME/.grugdo\n", highlight("grugdo.db"))
		var permission string
		fmt.Print("Can grug create file? (Y/n) ")
		fmt.Scanln(&permission)
		if permission == "" || permission == "y" || permission == "Y" {
			fmt.Printf("Grug create `%s` file now!\n", highlight("grugdo.db"))
			_, err = os.Create(home + "/.grugdo/grugdo.db")
			check(err)
			fmt.Printf("Grug created `%s`, let go!\n", highlight("grugdo.db"))
		} else {
			fmt.Println("Grug can't work without it. Grug bye!")
			os.Exit(1)
		}
		fmt.Println()
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
