package cli

import (
    "fmt"
    "time"
    "github.com/spf13/cobra"

    "timesync/config"
    "timesync/internal/core"
)


var rootCmd = &cobra.Command{
	Use: "timesync",
	Short: "A time management and scheduling CLI tool",
	Long: 	`TimeSync helps you schedule meetings and team timezones efficiently`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init(){
	rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use: "export",
	Short: "Export schedules",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Exporting schedules...")
	},
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule a meeting",
	Run: func(cmd *cobra.Command, args []string) {
		details, _ := cmd.Flags().GetString("details")
		timeStr, _ := cmd.Flags().GetString("time")
		tz, _ := cmd.Flags().GetString("timezone")

		parsedTime, err := time.Parse("2006-01-02T15:04", timeStr)
		if err != nil {
			fmt.Printf("Error parsing time: %v\n", err)
			return
		}

		config := config.LoadConfig()
		schedule, err := core.ScheduleMeeting(config.DefaultTimezone, tz, details, parsedTime)
		if err != nil {
			fmt.Printf("Error scheduling meeting: %v\n", err)
			return
		}

		fmt.Printf("Meeting scheduled: %v\n", schedule)
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
	scheduleCmd.Flags().String("details", "", "Details of the meeting")
	scheduleCmd.Flags().String("time", "", "Time of the meeting (format: 2006-01-02T15:04)")
	scheduleCmd.Flags().String("timezone", "", "Timezone of the meeting")
}
