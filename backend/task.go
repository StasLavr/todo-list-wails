package task

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
    ID   int    
    Field1 string 
    Field2 string 
    Field3 string 
}

func Add(name, date string) {
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04")
	database, _ := sql.Open("sqlite3", "database.db")
	defer database.Close()
	stat, err := database.Prepare("INSERT INTO tasks (name,date,done_task,date_create) VALUES (?,?,'0',?)")
	if err != nil {
		panic(err)
	}
	stat.Exec(name, date, formattedTime)
}

func Del(id int) {
	database, _ := sql.Open("sqlite3", "database.db")
	stat, err := database.Prepare("DELETE FROM tasks WHERE id = ? ")
	if err != nil {
		panic(err)
	}
	stat.Exec(id)
}

func Done(id int) {
	database, _ := sql.Open("sqlite3", "database.db")
	stat, err := database.Prepare("UPDATE tasks set done_task = '1' WHERE id = ?")
	if err != nil {
		panic(err)
	}
	stat.Exec(id)
}

func Done_off(id int) {
	database, _ := sql.Open("sqlite3", "database.db")
	stat, err := database.Prepare("UPDATE tasks set done_task = '0' WHERE id = ?")
	if err != nil {
		panic(err)
	}
	stat.Exec(id)
}

func Get() []Item {
	database, _ := sql.Open("sqlite3", "database.db")
	query := "SELECT id,name,date,date_create FROM tasks WHERE done_task = '0'"
	row, err := database.Query(query)
	if err != nil {
		if nil == sql.ErrNoRows {
			fmt.Println("Нету")
		}
	}
	tasks := []Item{}
	for row.Next() {
		p := Item{}
		err := row.Scan(&p.ID, &p.Field1, &p.Field2,&p.Field3)
		if err != nil{
            fmt.Println(err)
			continue
        }
		tasks = append(tasks, p)
	}
	return tasks
}

func Get_done() []Item {
	database, _ := sql.Open("sqlite3", "database.db")
	query := "SELECT id,name,date,date_create FROM tasks WHERE done_task = '1'"
	row, err := database.Query(query)
	if err != nil {
		if nil == sql.ErrNoRows {
			fmt.Println("Нету")
		}
	}
	tasks := []Item{}
	for row.Next() {
		p := Item{}
		err := row.Scan(&p.ID, &p.Field1, &p.Field2, &p.Field3)
		if err != nil{
            fmt.Println(err)
			continue
        }
		tasks = append(tasks, p)
	}
	return tasks
}
