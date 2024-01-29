/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	familytree "github.com/amitramachandran/family_tree"
	"github.com/spf13/cobra"
)

// relationshipCmd represents the relationship command
var relationshipCmd = &cobra.Command{
	Use:     "relationship",
	Short:   "create relations",
	Long:    `Sub command for add which is used to create relationship`,
	Example: "family-tree add relationship son",
	Run: func(cmd *cobra.Command, args []string) {
		l := log.Default()
		np := familytree.NewEntity(l)
		relation_name := familytree.PreProcessName(args)
		np.AddRelation(relation_name)
	},
}

func init() {
	addCmd.AddCommand(relationshipCmd)
}
