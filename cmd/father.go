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

// fatherCmd represents the father command
var fatherCmd = &cobra.Command{
	Use:   "father of",
	Short: "Get the father name",
	Long:  `Command to get the fathers name of given person`,
	Run: func(cmd *cobra.Command, args []string) {
		l := log.Default()
		np := familytree.NewEntity(l)
		name := familytree.PreProcessName(args)
		fatherName, err := np.GetFatherName(name)
		if err != nil {
			log.Fatalf("error happened %v", err)
		}
		fmt.Println(fatherName)

	},
}

func init() {
	rootCmd.AddCommand(fatherCmd)

}
