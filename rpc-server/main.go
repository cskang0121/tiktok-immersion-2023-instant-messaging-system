package main

import (
	"log"

	rpc "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc/imservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	/* --------------- Modularization --------------- */ 	

	// Initialise connection to database "tiktok"
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/tiktok")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	// Drop table "messages" if previously created 
	drop, err := db.Query("DROP TABLE IF EXISTS messages")
	if err != nil {
        panic(err)
    }
	defer drop.Close()

	// Create new table "messages"
	create, err := db.Query("CREATE TABLE messages (id int PRIMARY KEY AUTO_INCREMENT, chat VARCHAR(255), sender VARCHAR(255), send_time INT, message TEXT);")
    if err != nil {
        panic(err)
    }
	defer create.Close()

	r, err := etcd.NewEtcdRegistry([]string{"etcd:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := rpc.NewServer(new(IMServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "demo.rpc.server",
	}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
