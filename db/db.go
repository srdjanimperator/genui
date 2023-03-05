package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/srdjanimperator/genui/model"
)

type DbConnConfig struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

type Connection struct {
	Config *DbConnConfig
	db     *sql.DB
}

func (db *Connection) GetConnString() (string, error) {
	if db.Config == nil {
		return "", errors.New("Db config is empty!")
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.Config.Host, db.Config.Port, db.Config.User, db.Config.Pwd, db.Config.DbName)
	return connStr, nil
}

func NewConnection(cfg DbConnConfig) *Connection {
	dbConn := &Connection{
		Config: &cfg,
	}
	connStr, err := dbConn.GetConnString()
	if err != nil {
		panic("cannot connect to db, borting api startup!")
	}
	fmt.Println("Api connection to:", connStr)
	sqldb, err2 := sql.Open("postgres", connStr)
	if err2 != nil {
		fmt.Println(err2)
		panic("cannot connect to db, aborting api startup!")
	}
	dbConn.db = sqldb
	return dbConn
}

func (conn *Connection) FormsList() []*model.GuiForm {
	rows, err := conn.db.Query(`SELECT "code", "title" FROM "guiforms"`)
	if err != nil {
		panic("Db connection ping failed, aborting api startup!")

	}
	defer rows.Close()
	forms := make([]*model.GuiForm, 0)
	for rows.Next() {
		f := &model.GuiForm{}
		rows.Scan(&f.Code, &f.Title)
		forms = append(forms, f)
	}
	return forms
}
