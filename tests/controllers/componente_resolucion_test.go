package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	pgUser := os.Getenv("RESOLUCIONES_CRUD_V2_PGUSER")
	pgPass := os.Getenv("RESOLUCIONES_CRUD_V2_PGPASS")
	pgUrls := os.Getenv("RESOLUCIONES_CRUD_V2_PGHOST")
	pgDb := os.Getenv("RESOLUCIONES_CRUD_V2_PGDB")
	pgPort := os.Getenv("RESOLUCIONES_CRUD_V2_PGPORT")
	pgSchema := os.Getenv("RESOLUCIONES_CRUD_V2_PGSCHEMA")
	orm.RegisterDataBase("testbd", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+":"+pgPort+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")
}

func TestGetAllComponenteResolucion(t *testing.T) {
	Convey("Test: / - Componente resolucion GetAll", t, func() {
		r, err := http.NewRequest("GET", "/v1/componente_resolucion/", nil)
		if err != nil {
			t.Fatal("error", err)
		}
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})

}

func TestGetOneComponenteResolucionById(t *testing.T) {
	Convey("Test: / - ComponenteResolucion GetOne", t, func() {
		r, err := http.NewRequest("GET", "/v1/componente_resolucion/980", nil)
		if err != nil {
			t.Fatal("error", err)
		}
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

		// petición con datos incorrectos
		Convey("Status Code Should Be 404", func() {
			r, err = http.NewRequest("GET", "/v1/componente_resolucion/1", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 404)
		})
	})

}
