package controller

import (
	"database/sql"
)

type Basecontroller struct {
	db *sql.DB
}

func NewBasecontroller(db *sql.DB) *Basecontroller {
	return &Basecontroller{db}
}

func (c *Basecontroller) Status(name string) float64 {
	rows, err := c.db.Query("show global status where variable_name=?", name)
	if err != nil {
		panic(err)
	}
	var (
		vname  string
		values int
	)
	for rows.Next() {
		rows.Scan(&vname, &values)
	}
	return float64(values)
}

func (c *Basecontroller) Variables(name string) float64 {
	rows, err := c.db.Query("show global variables where variable_name=?", name)
	if err != nil {
		panic(err)
	}
	var (
		vname  string
		values int
	)
	for rows.Next() {
		rows.Scan(&vname, &values)
	}
	return float64(values)
}
