package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionEsquemaResolucionesNew_20211028_131735 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionEsquemaResolucionesNew_20211028_131735{}
	m.Created = "20211028_131735"

	migration.Register("CreacionEsquemaResolucionesNew_20211028_131735", m)
}

// Run the migrations
func (m *CreacionEsquemaResolucionesNew_20211028_131735) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/CreacionEsquemaResolucionesNew_20211028_131735_up.sql")

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
func (m *CreacionEsquemaResolucionesNew_20211028_131735) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/CreacionEsquemaResolucionesNew_20211028_131735_down.sql")

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
