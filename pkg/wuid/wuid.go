package wuid

import (
	"database/sql"
	"fmt"

	"github.com/edwingeng/wuid/mysql/wuid"
)

var w *wuid.WUID

func InitWUID(dataSource string) {
	newDB := func() (*sql.DB, bool, error) {
		db, err := sql.Open("mysql", dataSource)
		if err != nil {
			return nil, false, err
		}
		return db, true, nil
	}

	w = wuid.NewWUID("default", nil)
	err := w.LoadH28FromMysql(newDB, "wuid")
	if err != nil {
		panic(err)
	}
}

func GenerateUserID(dsn string) string {
	if w == nil {
		InitWUID(dsn)
	}
	return fmt.Sprintf("%016x", w.Next())
}
