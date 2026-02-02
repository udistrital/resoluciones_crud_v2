package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type HorasTrabajadas_20230621_142612 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &HorasTrabajadas_20230621_142612{}
	m.Created = "20230621_142612"

	migration.Register("HorasTrabajadas_20230621_142612", m)
}

// Run the migrations
func (m *HorasTrabajadas_20230621_142612) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230621_142612_horas_trabajadas_up.sql")

	if err != nil {
		// handle error
		log.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *HorasTrabajadas_20230621_142612) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230621_142612_horas_trabajadas_down.sql")

	if err != nil {
		// handle error
		log.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		m.SQL(request)
		// do whatever you need with result and error
	}

}
