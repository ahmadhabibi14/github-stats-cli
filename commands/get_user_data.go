package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/spf13/cobra"
)

var username string

func init() {
	rootCmd.AddCommand(get_user_data_CMD)
	get_user_data_CMD.Flags().StringVarP(&username, "username", "u", "", "Github username")
	get_user_data_CMD.MarkFlagRequired("username")
}

var get_user_data_CMD = &cobra.Command{
	Use:   "get-data",
	Short: "Get user data",
	Long:  "Get user data",
	Run: func(cmd *cobra.Command, args []string) {
		get_github_user_data(username)
	},
}

// User entities
type userData struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

// Service
func get_github_user_data(username string) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	var user userData
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return
	}

	// +===================== View ================+

	// Name
	name := widgets.NewParagraph()
	name.Title = "Name"
	name.Text = user.Name

	// Bio
	bio := widgets.NewParagraph()
	bio.Title = "Bio"
	bio.Text = user.Bio

	// Grid layout
	grid := ui.NewGrid()
	termWidth, _ := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, 3)
	grid.Set(
		ui.NewRow(2.0/2,
			ui.NewCol(1.0/2, name),
			ui.NewCol(1.0/2, bio),
		),
	)
	ui.Render(grid)

	return
}
