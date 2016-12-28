package store

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)

func TestDump(t *testing.T) {
	Convey("dump%load data", t, func() {
		m := map[string]interface{}{
			"aaaaaa": "111",
			"bbbbbb": "222",
		}
		Dump(m)
		d_str, _ := json.Marshal(Load())
		e_str, _ := json.Marshal(m)
		fmt.Println(e_str)
		So(string(d_str), ShouldEqual, string(e_str))
	})

}
