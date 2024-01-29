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

// daughtersCmd represents the daughters command
var daughtersCmd = &cobra.Command{
	Use:   "daughters of",
	Short: "Sub command for count",
	Long:  `Command for counting the number of daughters.`,
	Run: func(cmd *cobra.Command, args []string) {
		l := log.Default()
		np := familytree.NewEntity(l)
		name := familytree.PreProcessName(args)
		count := np.CountDaughters(name)
		fmt.Println(count)
	},
}

func init() {
	countCmd.AddCommand(daughtersCmd)

}
