package base

import (
	parsglobal "cncy/cfg"
	"database/sql"
	"fmt"
)

type TableData struct {
	Name        string
	RecordCount int64
	Size        int64
}

func Printsql() {

	username := parsglobal.Globalconfig().DBUserName
	password := parsglobal.Globalconfig().DBPasswd
	address := parsglobal.Globalconfig().DatabaseHost
	port := parsglobal.Globalconfig().DatabasePort
	fmt.Println(username, password, address, port)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/mdm", username, password, address, port))
	if err != nil {
		fmt.Println("连接数据库出错")
		return
	}

	rows, err := db.Query("SELECT table_name, table_rows, data_length FROM information_schema.tables WHERE table_schema = 'mdm'")
	if err != nil {
		fmt.Println("读取数据库出错！")
		return
	}
	defer rows.Close()

	var tableData []TableData

	for rows.Next() {
		var tableName string
		var recordCount int64
		var size int64

		err = rows.Scan(&tableName, &recordCount, &size)
		if err != nil {
			fmt.Println("扫描记录出错")
			return
		}

		tableData = append(tableData, TableData{Name: tableName, RecordCount: recordCount, Size: size})
	}
	fmt.Println(tableData)
}
