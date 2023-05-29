package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var get_user_data_CMD = &cobra.Command{
	Use:   "uname",
	Short: "Get user data",
	Long:  "Get user data",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := get_github_user_data(username)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		fmt.Println("Name: ", user.Name)
		fmt.Println("Bio: ", user.Bio)
	},
}

var username string

func init() {
	rootCmd.AddCommand(get_user_data_CMD)
	get_user_data_CMD.Flags().StringVarP(&username, "username", "u", "", "Github username")
	get_user_data_CMD.MarkFlagRequired("username")
}

type userData struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func get_github_user_data(username string) (*userData, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API request failed with status: %s", resp.Status)
	}

	var user userData
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
