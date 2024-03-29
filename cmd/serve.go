/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"go-graphql-cli/domain/repositories"
	"go-graphql-cli/infra/db"
	"go-graphql-cli/usecase"
	"go-graphql-cli/usecase/converter"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gormDB, err := db.InitDB()
		if err != nil {
			log.Printf("Failed to initialize DB: %v", err)
			return
		}

		u := usecase.NewEntriesUseCase(
			converter.NewEntriesConverter(),
			repositories.NewEntryRepository(gormDB),
		)
		entry, err := u.Fetch(args[0])
		if err != nil {
			log.Printf("Failed to fetch entry: %v", err)
			return
		}
		log.Println(entry)
		u.CreateEntry(entry)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
