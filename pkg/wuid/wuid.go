package wuid

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"

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

func CombineUserID(userId1, userId2 string) string {
	ids := []string{userId1, userId2}

	sort.Slice(ids, func(i, j int) bool {
		a, _ := strconv.ParseUint(ids[i], 0, 64)
		b, _ := strconv.ParseUint(ids[j], 0, 64)
		return a < b
	})

	return fmt.Sprintf("%s_%s", ids[0], ids[1])
}
