/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	familytree "github.com/amitramachandran/family_tree"
	"github.com/spf13/cobra"
)

// wivesCmd represents the wives command
var wivesCmd = &cobra.Command{
	Use:   "wives of",
	Short: "Sub command for count",
	Long:  `Command for counting the number of wives.`,
	Run: func(cmd *cobra.Command, args []string) {
		l := log.Default()
		np := familytree.NewEntity(l)
		name := familytree.PreProcessName(args)
		count := np.CountWives(name)
		fmt.Println(count)
	},
}

func init() {
	countCmd.AddCommand(wivesCmd)
}
