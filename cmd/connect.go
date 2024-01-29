/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	familytree "github.com/amitramachandran/family_tree"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:     "connect",
	Short:   "connect persons with relation",
	Long:    `Create connections between persons with the relationships to create family tree`,
	Example: "family-tree connect Amit as son of Rama",
	Run: func(cmd *cobra.Command, args []string) {
		l := log.Default()
		np := familytree.NewEntity(l)
		from, relation, to := familytree.PreProcessConnection(args)
		err := np.CheckEntries(from, to, relation)
		if err != nil {
			l.Fatalf("The entries are not in DB: %s", err)
		}
		np.AddRelationship(from, relation, to)
		// if (relation == "son") || (relation == "daughter") {
		// 	np.UpdateFatherName(from , to)
		// }
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

}
