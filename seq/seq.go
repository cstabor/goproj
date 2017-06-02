package main

import (
	"fmt"
	"net/http"
	//"strings"
	"log"
	"database/sql" //这包一定要引用，是底层的sql驱动
	_ "github.com/go-sql-driver/mysql"
	"strconv" //把int转换为string
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()       //解析参数，默认是不会解析的
	//fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//	fmt.Println("key:", k)
	//	fmt.Println("val:", strings.Join(v, ""))
	//}
	fmt.Fprintf(w, "Hello Go\n") //这个写入到w的是输出到客户端的

	// ===============================
	db, err := sql.Open("mysql", "godev:godev_pwd@tcp(192.168.1.103:3306)/YD_IdGenerator?charset=utf8")
	//数据库连接字符串
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
			fmt.Fprintf(w, "id:"+strconv.Itoa(id)+" stub: "+lvs)
		}
	}
	//insert_sql := "INSERT INTO xiaorui(lvs) VALUES(?)"
	//_, e4 := db.Exec(insert_sql, "nima")
	//fmt.Println(e4)
	db.Close() //关闭数据库
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
