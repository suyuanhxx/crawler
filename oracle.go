package main

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"log"
	"fmt"
)

type User struct {
	ID     uint64
	Mobile string
	MEMO1  string
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Oracle Driver example")

	//os.Setenv("NLS_LANG", "")

	// 用户名/密码@IP:端口/实例名
	db, err := sql.Open("oci8", "KSCF_WEB_DEV/KSCF_WEB_DEV@192.168.52.101:1521/orcl")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select ID,Mobile,MEMO1 from EC_U_BASIC")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Mobile, &user.MEMO1)
		fmt.Println(user)
	}
}
