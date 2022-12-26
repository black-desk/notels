package jsonrpc2

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMessageAs(t *testing.T) {
	cases := []string{
		`{
                        "jsonrpc" : "2.0",
                        "method"  : "some function name",
                        "params" : {},
                        "id" : "a"
                }`,
	}
	for _, testCase := range cases {
		Convey("We have some JSON-RPC request messages", t, func() {
			msg := json.RawMessage(testCase)
			So(json.Valid(msg), ShouldBeTrue)

			Convey(
				"messageAs should return true for Request and false otherwise",
				func() {
					tmpRequest := Request{}
					tmpResponse := Response{}

					So(
						messageAs(msg, &tmpResponse),
						ShouldBeFalse,
					)
					So(
						messageAs(msg, &tmpRequest),
						ShouldBeTrue,
					)
					So(tmpRequest.Id, ShouldNotBeNil)
				},
			)
		})
	}
	cases = []string{
		`{
                        "jsonrpc" : "2.0",
                        "result"  : true,
                        "id" : "a"
                }`,
	}
	for _, testCase := range cases {
		Convey("We have some JSON-RPC response messages", t, func() {
			msg := json.RawMessage(testCase)
			So(json.Valid(msg), ShouldBeTrue)

			Convey(
				"messageAs should return true for Response and false otherwise",
				func() {
					tmpRequest := Request{}
					tmpResponse := Response{}

					So(
						messageAs(msg, &tmpResponse),
						ShouldBeTrue,
					)
					So(tmpResponse.Id, ShouldNotBeNil)
					So(
						messageAs(msg, &tmpRequest),
						ShouldBeFalse,
					)
				},
			)
		})
	}
}
