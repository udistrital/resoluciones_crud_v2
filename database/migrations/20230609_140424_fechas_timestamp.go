package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type FechasTimestamp_20230609_140424 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &FechasTimestamp_20230609_140424{}
	m.Created = "20230609_140424"

	migration.Register("FechasTimestamp_20230609_140424", m)
}

// Run the migrations
func (m *FechasTimestamp_20230609_140424) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230609_140424_fechas_timestamp_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *FechasTimestamp_20230609_140424) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230609_140424_fechas_timestamp_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
