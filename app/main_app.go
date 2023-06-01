package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Service
func MainApp(username string) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}
	// Decode response to struct
	var user userData
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return
	}
	LangsAndRepos(url)
	// +===================== View ================+
	// Profile Picture
	img, err := profile_image(user.Avatar)
	if err != nil {
		return
	}
	prof_pict := widgets.NewImage(nil)
	prof_pict.Image = img
	// Name
	name := widgets.NewParagraph()
	name.Title = "Name"
	name.Text = user.Name
	name.BorderStyle.Fg = ui.ColorBlue
	// Bio
	bio := widgets.NewParagraph()
	bio.Title = "Bio"
	bio.Text = user.Bio
	// Most Used Languages
	lng := []string{}
	for key, value := range LangToFetch {
		text := fmt.Sprintf("%s: %.2f", key, value)
		lng = append(lng, text+"%")
	}
	most_lang := widgets.NewList()
	most_lang.Title = "Most used languages"
	most_lang.Rows = lng
	most_lang.TextStyle = ui.NewStyle(ui.ColorYellow)
	// Grid layout
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(2.0/5,
			ui.NewCol(1.4/5, prof_pict),
			ui.NewCol(1.0/5, most_lang),
			ui.NewCol(2.6/5,
				ui.NewRow(1.0/4, name),
				ui.NewRow(1.0/4, bio),
			),
		),
	)
	ui.Render(grid)
	return
}
