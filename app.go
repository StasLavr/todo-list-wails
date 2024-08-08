package main

import (
	"context"
	"fmt"
	"wails/backend"

)

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

// Greet returns a greeting for the given name

func (b *App) Add(name,data string){
	fmt.Println("add")
	fmt.Println(name,data)
	task.Add(name,data)
}

func (b *App) Get() []task.Item {
	fmt.Println("get")
	json := task.Get()
	return json
}


func (b *App) Del(id int) {
	fmt.Println("del")
	task.Del(id)
}


func (b *App) Done(id int) {
	fmt.Println("Done")
	task.Done(id)
}

func (b *App) Done_off(id int) {
	fmt.Println("Done")
	task.Done_off(id)
}

func (b *App) Get_done() []task.Item {
	fmt.Println("get_done")
	json := task.Get_done()
	return json
}

