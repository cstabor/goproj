package main

import (
	"database/sql" //这包一定要引用，是底层的sql驱动
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv" //把int转换为string
)

func main() {
	db, err := sql.Open("mysql", "dsp_dev:yd!@#dsp_dev@tcp(192.168.1.103:3306)/YD_IdGenerator?charset=utf8")
	//数据库连接字符串，别告诉我看不懂。端口一定要写/
	if err != nil { //连接成功 err一定是nil否则就是报错
		panic(err.Error())       //抛出异常
		fmt.Println(err.Error()) //仅仅是显示异常
	}
	defer db.Close() //只有在前面用了 panic 这时defer才能起作用，如果链接数据的时候出问题，他会往err写数据

	rows, err := db.Query("select YF_seq_adx_invoice_id, YF_stub from YT_SeqADXInvoiceId")
	//判断err是否有错误的数据，有err数据就显示panic的数据
	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var id int     //定义一个id 变量
	var lvs string //定义lvs 变量
	for rows.Next() { //开始循环
		rerr := rows.Scan(&id, &lvs) //数据指针，会把得到的数据，往刚才id 和 lvs引入
		if rerr == nil {
			fmt.Println("id:", strconv.Itoa(id)+" stub: "+lvs)
		}
	}
	insert_sql := "INSERT INTO xiaorui(lvs) VALUES(?)"
	_, e4 := db.Exec(insert_sql, "nima")
	fmt.Println(e4)
	db.Close() //关闭数据库
}
