package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "net/http"
)

var db *sql.DB

func init() {
    db, _ = sql.Open("mysql", "godev:godev_pwd@tcp(192.168.1.103:3306)/YD_IdGenerator?charset=utf8")
    db.SetMaxOpenConns(2000)
    db.SetMaxIdleConns(1000)
    db.Ping()
}

func main() {
    startHttpServer()
}

func startHttpServer() {
    http.HandleFunc("/pool", pool)
    err := http.ListenAndServe(":9092", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func pool(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("select * from YT_SeqADXInvoiceId")
    defer rows.Close()
    checkErr(err)
    
    columns, _ := rows.Columns()
    scanArgs := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
    for j := range values {
        scanArgs[j] = &values[j]
    }
    
    record := make(map[string]string)
    for rows.Next() {
        //将行数据保存到record字典
        err = rows.Scan(scanArgs...)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
    }
    
    fmt.Println(record)
    fmt.Fprintln(w, "finish")
}

func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}
