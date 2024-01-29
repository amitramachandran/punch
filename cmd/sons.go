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

// sonsCmd represents the sons command
var sonsCmd = &cobra.Command{
	Use:   "sons of",
	Short: "Sub command for count",
	Long:  `Command for counting the number of sons.`,
	Run: func(cmd *cobra.Command, args []string) {
		l := log.Default()
		np := familytree.NewEntity(l)
		name := familytree.PreProcessName(args)
		count := np.CountSons(name)
		fmt.Println(count)

	},
}

func init() {
	countCmd.AddCommand(sonsCmd)
}
