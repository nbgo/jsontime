package jsontime

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
	"time"
	"fmt"
)

type TypeWithJSONTimeRFC3339Nano struct {
	Time RFC3339Nano `json:"time"`
}

type TypeWithJSONTimeRFC3339NanoRef struct {
	Time *RFC3339Nano `json:"time"`
}

func TestRFC3339Nano(t *testing.T) {
	tm := time.Now()
	ts := fmt.Sprintf("\"%v\"", tm.Format(time.RFC3339Nano))

	Convey("Direct converstion JsonTime -> JSON", t, func() {
		jt := RFC3339Nano(tm)
		res, err := json.Marshal(jt)
		So(err, ShouldBeNil)
		So(string(res), ShouldEqual, ts)
	})

	Convey("Direct conversion JSON -> JsonTime", t, func() {
		var jt RFC3339Nano
		err := json.Unmarshal([]byte(ts),  &jt)
		So(err, ShouldBeNil)
		So(time.Time(jt).UnixNano(), ShouldEqual, tm.UnixNano())
	})

	Convey("Convert JSON -> TypeWithJsonTime", t, func() {
		var jt TypeWithJSONTimeRFC3339Nano
		ts := fmt.Sprintf("{\"time\":%v}", ts)
		err := json.Unmarshal([]byte(ts), &jt)
		So(err, ShouldBeNil)
		So(time.Time(jt.Time).UnixNano(), ShouldEqual, tm.UnixNano())
	})

	Convey("Convert JSON (no value) -> TypeWithJsonTime", t, func() {
		var jt TypeWithJSONTimeRFC3339Nano
		var zeroTime time.Time
		ts := "{}"
		err := json.Unmarshal([]byte(ts), &jt)
		So(err, ShouldBeNil)
		So(time.Time(jt.Time).UnixNano(), ShouldEqual, zeroTime.UnixNano())
	})

	Convey("Convert JSON (null value) -> TypeWithJsonTime", t, func() {
		var jt TypeWithJSONTimeRFC3339Nano
		ts := fmt.Sprintf("{\"time\":%v}", "null")
		err := json.Unmarshal([]byte(ts), &jt)
		So(err, ShouldNotBeNil)
	})

	Convey("Convert JSON -> TypeWithJsonTimeRef", t, func() {
		var jt TypeWithJSONTimeRFC3339NanoRef
		ts := fmt.Sprintf("{\"time\":%v}", ts)
		err := json.Unmarshal([]byte(ts), &jt)
		So(err, ShouldBeNil)
		So(time.Time(*jt.Time).UnixNano(), ShouldEqual, tm.UnixNano())
	})

	Convey("Convert JSON (null value) -> TypeWithJsonTimeRef", t, func() {
		var jt TypeWithJSONTimeRFC3339NanoRef
		ts := fmt.Sprintf("{\"time\":%v}", "null")
		err := json.Unmarshal([]byte(ts), &jt)
		So(err, ShouldBeNil)
		So(jt.Time, ShouldBeNil)
	})

	Convey("Convert JSON (no value) -> TypeWithJsonTimeRef", t, func() {
		var jt TypeWithJSONTimeRFC3339NanoRef
		ts := "{}"
		err := json.Unmarshal([]byte(ts), &jt)
		So(err, ShouldBeNil)
		So(jt.Time, ShouldBeNil)
	})
}