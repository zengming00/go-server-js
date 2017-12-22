package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-sql-driver/mysql"
)

func handErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// db, err := sql.Open("mysql", "root:root@/test?parseTime=true&loc="+url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", "root:root@/test2")
	handErr(err)
	rows, err := db.Query("select * from users")
	handErr(err)
	defer rows.Close()

	cols, err := rows.Columns()
	handErr(err)
	ct, err := rows.ColumnTypes()
	handErr(err)

	arr := make([]interface{}, len(ct))
	for i, v := range ct {
		t := v.ScanType()
		v := reflect.New(t).Interface()
		arr[i] = v
		fmt.Println(cols[i], t)
	}

	for rows.Next() {
		err = rows.Scan(arr...)
		handErr(err)

		m := make(map[string]interface{})
		for i, col := range cols {
			if col == "template_info" || col == "state" {
				m[col] = ""
				continue
			}
			v := arr[i]
			switch vv := v.(type) {
			case *int32:
				m[col] = *vv
			case *sql.NullString:
				m[col] = *vv
			case *sql.NullBool:
				m[col] = *vv
			case *sql.NullFloat64:
				m[col] = *vv
			case *sql.NullInt64:
				m[col] = *vv
			case *sql.RawBytes:
				m[col] = string(*vv)
			case *mysql.NullTime:
				m[col] = *vv
			default:
				m[col] = vv
				panic("unknow type")
			}
		}

		if bts, err := json.MarshalIndent(m, "", "  "); err != nil {
			panic(err)
		} else {
			fmt.Println(string(bts))
		}
	}
	err = rows.Err()
	handErr(err)
}
