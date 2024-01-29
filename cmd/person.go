/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	familytree "github.com/amitramachandran/family_tree"
	"github.com/spf13/cobra"
)

// personCmd represents the person command
var personCmd = &cobra.Command{
	Use:     "person",
	Short:   "create persons",
	Long:    `Sub command for add which is used to create persons`,
	Example: "family-tree add person Amit",
	Run: func(cmd *cobra.Command, args []string) {
		l := log.Default()
		np := familytree.NewEntity(l)
		name := familytree.PreProcessName(args)
		np.AddPerson(name)

	},
}

func init() {
	addCmd.AddCommand(personCmd)

}
