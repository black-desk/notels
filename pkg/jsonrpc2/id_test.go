package jsonrpc2_test

import (
	"encoding/json"
	"testing"

	. "github.com/black-desk/notels/pkg/jsonrpc2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestJsonRPCID(t *testing.T) {
	var id ID

	jsonNumbers := []string{`100`, `100.10001`}

	for _, jsonNumber := range jsonNumbers {
		Convey("We have a json number", t, func() {
			jsonNumber := json.RawMessage(jsonNumber)

			Convey("Unmarshal it to JsonRPCID", func() {
				err := json.Unmarshal(jsonNumber, &id)
				So(err, ShouldBeNil)

				Convey("JsonRPCID should be number", func() {
					So(id.IsNum(), ShouldBeTrue)
				})

				Convey("JsonRPCID should not be null", func() {
					So(id.IsNull(), ShouldBeFalse)
				})

				Convey(
					"JsonRPCID should not be string",
					func() {
						So(id.IsStr(), ShouldBeFalse)
					},
				)
			})
		})
	}

	jsonStrings := []string{`"null"`, `"aaa"`}

	for _, jsonString := range jsonStrings {
		Convey("We have a json string", t, func() {
			jsonString := json.RawMessage(jsonString)

			Convey("Unmarshal it to JsonRPCID", func() {
				err := json.Unmarshal(jsonString, &id)
				So(err, ShouldBeNil)

				Convey(
					"JsonRPCID should not be number",
					func() {
						So(id.IsNum(), ShouldBeFalse)
					},
				)

				Convey("JsonRPCID should not be null", func() {
					So(id.IsNull(), ShouldBeFalse)
				})

				Convey(
					"JsonRPCID should be string",
					func() {
						So(id.IsStr(), ShouldBeTrue)
					},
				)
			})
		})
	}
}
