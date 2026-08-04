/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-task-manager/model"
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
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
		var task model.Task
		var index int
		id := cmd.Flag("id").Value.String()
		for i, t := range tl.Tasks {
			if t.ID == id {
				index = i
				task = t
				break
			}
		}

		if cmd.Flags().Changed("title") {
			tittle, err := cmd.Flags().GetString("title")
			if err != nil {
				fmt.Println(err)
				return
			}
			task.Title = tittle
		}
		if cmd.Flags().Changed("description") {
			desc, err := cmd.Flags().GetString("description")
			if err != nil {
				fmt.Println(err)
				return
			}
			task.Description = desc
		}
		if cmd.Flags().Changed("is-done") {
			isDone, err := cmd.Flags().GetString("is-done")
			fmt.Println(isDone)
			if err != nil {
				fmt.Println(err)
				return
			}
			task.Done = isDone == "true"
		}
		// it's likely that with concurrent usage of task-list
		// the code below run into problem and cause incorrect placement, ignored for this proj
		tl.Tasks[index] = task
		tl.SaveToFile("tasks.json")

		return
	},
}

func init() {
	taskCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.Flags().StringP("id", "i", "", "task id")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringP("title", "t", "", "new Title of the task")
	updateCmd.Flags().StringP("description", "d", "", "new Description of the task")
	updateCmd.Flags().String("is-done", "false", "is the task completed")
}
