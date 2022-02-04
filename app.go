package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/keybase/go-keychain"
	touchid "github.com/lox/go-touchid"
	"github.com/masahide/addsshkey/pkg/config"
	"github.com/masahide/addsshkey/pkg/opcli"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx    context.Context
	Config *config.Config
	op     *opcli.OpCli
	msgCh  chan string
}

// NewApp creates a new App application struct
func NewApp(appName string) *App {
	userConfDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(userConfDir, appName)
	c, err := config.New(dir)
	if err != nil {
		log.Fatal(err)
	}
	return &App{
		Config: c,
		op:     opcli.NewOpCli(c.OpConfig),
		msgCh:  make(chan string),
	}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (b *App) PutTwoFactorCode(code string) {
	select {
	case b.msgCh <- code:
	case <-b.ctx.Done():
	}
}

func (b *App) twoFactor() string {
	runtime.EventsEmit(b.ctx, "twofactor")
	select {
	case m := <-b.msgCh:
		return m
	case <-b.ctx.Done():
	}
	return ""
}

func (b *App) GetOpAccounts() ([]opcli.OpAccountInfo, error) {
	return b.op.GetAccountInfo()
}

func (b *App) AddAccount(info opcli.OpAccountInfo, pass string) error {
	return b.op.AddAccount(info, pass, b.twoFactor)
}

// Greet returns a greeting for the given name
func (b *App) Greet(name string) string {
	service := "MyService"
	account := "test-yamasaki"
	accessGroup := "A123456789.group.com.mycorp"
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(service)
	query.SetAccount(account)
	query.SetAccessGroup(accessGroup)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return err.Error()
	} else if len(results) != 1 {
		return err.Error()
	}

	ok, err := touchid.Authenticate("access llamas")
	if err != nil {
		return err.Error()
	}

	if !ok {
		return "Failed to authenticate"
	}

	pass := string(results[0].Data)
	return fmt.Sprintf("Hello %s, It's show time!", name+":"+pass)
}
