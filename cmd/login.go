package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zmb3/spotify"

	"github.com/rpsl/spotify-automaton/config"
)

func Login(config *config.Config) {
	auth := spotify.NewAuthenticator(
		config.Spotify.RedirectURL,
		spotify.ScopeUserReadPrivate,
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopeUserLibraryRead)

	auth.SetAuthInfo(config.Spotify.ClientID, config.Spotify.ClientSecret)

	url := auth.AuthURL("automaton")

	fmt.Printf("Open then next url on your browser and authorize app: \n\n%v\n\n", url)

	authToken := question("Enter authorization code")

	token, err := auth.Exchange(authToken)

	if err != nil {
		fmt.Printf("\nCannot fetch token\n\n")
		os.Exit(1)

		return
	}

	config.WriteTokens(token.AccessToken, token.RefreshToken)

	fmt.Printf("\n\nAccess Token: %v\n", token.AccessToken)
	fmt.Printf("Refresh Token: %v\n\n", token.RefreshToken)
	fmt.Printf("\nValues saved into tokens.toml file\n\n")
}

// question make question to use and retrieve answer
func question(question string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s: ", question)

	answer, _ := reader.ReadString('\n')

	return strings.TrimSuffix(answer, "\n")
}
