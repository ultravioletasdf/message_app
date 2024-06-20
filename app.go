package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

var client = http.DefaultClient
var API_URL = "http://127.0.0.1:3991"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
func (a *App) IsSignedIn() bool {
	session, err := os.ReadFile("./.session")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	res, err := get(API_URL+"/check_session", string(session))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return string(res) == "1"
}
func (a *App) SignIn(username string, password string) string {
	res, err := get(API_URL+"/set_session", "")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(res))
	err = os.WriteFile("./.session", res, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ""
}
func (a *App) SignOut() error {
	_, err := get(API_URL, "/clear_session")
	if err != nil {
		return err
	}
	err = os.WriteFile("./.session", nil, 0644)
	if err != nil {
		return err
	}
	return nil
}
func get(url string, session string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if session != "" {
		req.Header.Add("Authorization", session)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
