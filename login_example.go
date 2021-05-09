package main

import (
	"./mangadex"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func checkLoginStatus(client *mangadex.APIClient, ctx context.Context) mangadex.CheckResponse {
	authResp, resp, err := client.AuthApi.GetAuthCheck(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("%v", resp)
	}
	return authResp
}

func main() {

	// Directory configuration
	fileSession := "session.json"
	userUsername := flag.String("username", "", "mangadex username")
	userPassword := flag.String("password", "", "mangadex password")
	flag.Parse()
	if *userUsername == "" || *userPassword == "" {
		log.Fatalf("username and password are required!!")
		os.Exit(1)
	}

	// Create client
	config := mangadex.NewConfiguration()
	client := mangadex.NewAPIClient(config)
	config.UserAgent = "similar-manga v2.0"
	config.HTTPClient = &http.Client{
		Timeout: 60 * time.Second,
	}
	ctx := context.Background()

	// Left first try to login
	token := mangadex.LoginResponseToken{}
	if _, err := os.Stat(fileSession); err == nil {
		fileManga, _ := ioutil.ReadFile(fileSession)
		_ = json.Unmarshal([]byte(fileManga), &token)
		config.AddDefaultHeader("Authorization", "Bearer "+token.Session)
	}
	authResp := checkLoginStatus(client, ctx)
	fmt.Println(authResp)

	// On first ever start we will get our session!
	if !authResp.IsAuthenticated && len(token.Session) == 0 {
		fmt.Println("Performing first time login!")
		bodyData := map[string]string{
			"username": *userUsername,
			"password": *userPassword,
		}
		opts := mangadex.AuthApiPostAuthLoginOpts{}
		opts.Body = optional.NewInterface(bodyData)
		authResp, resp, err := client.AuthApi.PostAuthLogin(ctx, &opts)
		if err != nil {
			log.Fatalf("%v\n%v", resp, err)
		}
		if resp.StatusCode != 200 {
			log.Fatalf("%v\n%v", resp, err)
		}
		file, _ := json.MarshalIndent(authResp.Token, "", " ")
		_ = ioutil.WriteFile(fileSession, file, 0644)
		token = *authResp.Token
		config.AddDefaultHeader("Authorization", "Bearer "+token.Session)
	} else if !authResp.IsAuthenticated {
		fmt.Println("Performing session refresh!")
		bodyData := map[string]string{
			"token": token.Refresh,
		}
		opts := mangadex.AuthApiPostAuthRefreshOpts{}
		opts.Body = optional.NewInterface(bodyData)
		authResp, resp, err := client.AuthApi.PostAuthRefresh(ctx, &opts)
		if resp.StatusCode == 401 {
			os.Remove(fileSession)
			log.Fatalf("need to perform re-login, please re-run!")
			os.Exit(1)
		}
		if err != nil {
			log.Fatalf("%v\n%v", resp, err)
		}
		if resp.StatusCode != 200 {
			log.Fatalf("%v\n%v", resp, err)
		}
		file, _ := json.MarshalIndent(authResp.Token, "", " ")
		_ = ioutil.WriteFile(fileSession, file, 0644)
		token = *authResp.Token
		config.AddDefaultHeader("Authorization", "Bearer "+token.Session)
	}
	authResp = checkLoginStatus(client, ctx)
	fmt.Println(authResp)

}
