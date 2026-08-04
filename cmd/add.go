/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-task-manager/model"
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tl := model.TaskList{}

		err := tl.LoadFromFile("tasks.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		title, err := cmd.Flags().GetString("title")
		if err != nil {
			fmt.Println(err)
			return
		}
		desc, err := cmd.Flags().GetString("description")
		if err != nil {
			fmt.Println(err)
			return
		}
		t := model.Task{
			Title:       title,
			Description: desc,
			Done:        false,
		}
		tl.AddTask(t)
		err = tl.SaveToFile("tasks.json")
		if err != nil {

			return
		}
		fmt.Printf("")
	},
}

func init() {
	taskCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("title", "t", "default title", "Title of the task")
	addCmd.Flags().StringP("description", "d", "default description", "Description of the task")
	addCmd.Flags().Bool("done", false, "is the task completed")
}
