package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "gett2/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)


func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/driver/asd", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestDriverGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 404", func() {
			So(w.Code, ShouldEqual, 404)
		})
	})
}
