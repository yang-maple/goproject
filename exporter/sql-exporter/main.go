package main

import (
	"database/sql" // sql 包
	"exporter/sql-exporter/conf"
	"exporter/sql-exporter/controller"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mysql 驱动包
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=PRC&parseTime=true", conf.Username, conf.Passwd, conf.Ip, conf.Port, conf.Dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	prometheus.MustRegister(controller.NewUpcontroller(db))
	prometheus.MustRegister(controller.NewConcontroller(db))
	prometheus.MustRegister(controller.NewCommand(db))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}
