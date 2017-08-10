package service

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHelloService(t *testing.T) {
	Convey("Given user Nick", t, func() {
		user := "Nick"

		Convey("When we see the user", func() {
			Convey("The greeting should be displayed to the user", func() {
				So(SayHello(user), ShouldEqual, "Hello Nick")
			})
		})
	})
}
