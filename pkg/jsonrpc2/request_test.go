package jsonrpc2

import (
	"encoding/json"
	"testing"

	. "github.com/black-desk/lib/go/errwrap"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidRequest(t *testing.T) {
	cases := []string{
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name",
                        "params" : {},
                        "id" : "a"
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name",
                        "params" : [],
                        "id" : null
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name"
                }`,
	}

	for _, testCase := range cases {
		Convey("Create a valid JSON-RPC request from json", t, func() {
			var req Request
			err := Annotate(json.Unmarshal([]byte(testCase), &req),
				"testCase %s", testCase)

			So(err, ShouldBeNil)

			Convey("Should pass validate", func() {
				err := Annotate(validate.Struct(&req),
					"testCase %s", testCase)

				So(err, ShouldBeNil)
			})
		})
	}
}

func TestInvalidRequest(t *testing.T) {
	cases := []string{
		`{
                        "jsonrpc" : "1.0",
                        "method"  : "some function name",
                        "params"  : {},
                        "id"      : "a"
                }`,
		`{
                        "method"  : "some function name",
                        "params"  : {},
                        "id"      : "a"
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name",
                        "params"  : "a",
                        "id"      : "a"
                }`,
	}

	for _, testCase := range cases {
		Convey("Create an invalid JSON-RPC request from json", t,
			func() {
				var req Request
				err := Annotate(
					json.Unmarshal([]byte(testCase), &req),
					"testCase %s", testCase)

				So(err, ShouldBeNil)

				Convey("Should not pass validate", func() {
					err := Annotate(validate.Struct(&req),
						"testCase %s", testCase)

					So(err, ShouldNotBeNil)
				})
			},
		)
	}
}

func TestRequestIsNotification(t *testing.T) {
	cases := []string{
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name",
                        "params"  : {}
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name",
                        "params"  : []
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name"
                }`,
	}

	for _, testCase := range cases {
		Convey("Create a JSON-RPC request from json without id", t,
			func() {
				var req Request
				err := Annotate(
					json.Unmarshal([]byte(testCase), &req),
					"testCase %s", testCase)

				So(err, ShouldBeNil)

				Convey("IsNotification should return true",
					func() {
						So(
							req.IsNotification(),
							ShouldBeTrue,
						)
					})
			})
	}

	cases = []string{
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name",
                        "params"  : {},
                        "id"      : 1
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name",
                        "params"  : [],
                        "id"      : "1"
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name",
                        "id"      : 1.0001
                }`,
	}

	for _, testCase := range cases {
		Convey("Create a JSON-RPC request from json with id", t,
			func() {
				var req Request
				err := Annotate(
					json.Unmarshal([]byte(testCase), &req),
					"testCase %s", testCase)

				So(err, ShouldBeNil)

				Convey("IsNotification should return false",
					func() {
						So(
							req.IsNotification(),
							ShouldBeFalse,
						)
					})
			})
	}
}
