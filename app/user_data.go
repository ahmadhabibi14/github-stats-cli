package app

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// User entities
type userData struct {
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar_url"`
}

// Service
func GetGithubUserData(username string) {
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
	fetchImage, err := http.Get(user.Avatar)
	if err != nil {
		return
	}
	var images []image.Image
	img, _, err := image.Decode(fetchImage.Body)
	if err != nil {
		return
	}
	images = append(images, img)

	// +===================== View ================+
	// Profile Picture
	prof_pict := widgets.NewImage(nil)
	prof_pict.Image = images[0]
	// Name
	name := widgets.NewParagraph()
	name.Title = "Name"
	name.Text = user.Name
	name.BorderStyle.Fg = ui.ColorBlue
	// Bio
	bio := widgets.NewParagraph()
	bio.Title = "Bio"
	bio.Text = user.Bio
	// Grid layout
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(2.0/5,
			ui.NewCol(1.0/4, prof_pict),
			ui.NewCol(3.0/4,
				ui.NewRow(1.0/2, name),
				ui.NewRow(1.0/2, bio),
			),
		),
	)
	ui.Render(grid)
	return
}
