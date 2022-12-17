package jsonrpc2

import (
	"encoding/json"
	"testing"

	. "github.com/black-desk/lib/go/errwrap"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidResponse(t *testing.T) {
	cases := []string{
		`{
                        "jsonrpc" : "2.0",
                        "result"  : true,
                        "id": "a"
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "result"  : [],
                        "id" : null
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "error"  : {
                                "code": 0,
                                "message": ""
                        },
                        "id" : null
                }`,
	}

	for _, testCase := range cases {
		Convey("Create a valid JSON-RPC response from json", t, func() {
			var resp Response
			err := Annotate(
				json.Unmarshal([]byte(testCase), &resp),
				"testCase %s", testCase)

			So(err, ShouldBeNil)

			Convey("Should pass validate", func() {
				err := Annotate(validate.Struct(&resp),
					"testCase %s", testCase)

				So(err, ShouldBeNil)
			})
		})
	}
}

func TestInvalidResponse(t *testing.T) {
	cases := []string{
		`{
                        "jsonrpc" : "2.0",
                        "id": "a"
                }`,
		`{
                        "jsonrpc" : "2.0",
                        "result"  : [],
                        "error"  : {
                                "code": 0,
                                "message": ""
                        },
                        "id" : null
                }`,
	}

	for _, testCase := range cases {
		Convey(
			"Create an invalid JSON-RPC response from json",
			t,
			func() {
				var resp Response
				err := Annotate(
					json.Unmarshal([]byte(testCase), &resp),
					"testCase %s", testCase)

				So(err, ShouldBeNil)

				Convey("Should not pass validate", func() {
					err := validate.Struct(&resp)

					So(err, ShouldNotBeNil)
				})
			},
		)
	}
}
