package nni_test

import (
	"fmt"
	"nni"
	"testing"
)

/*
type Route struct { // the path is "/"

}
*/

type RouteApple struct { // the path is "/apple/:kind"
	nni.Route `url:"apple"`
	kind string
}

type RouteAppleGet struct { // the path is "/apple/:kind/get/:number"
	RouteApple `url:"get"`
	number string
	nni.GET
}

func TestExample1(t *testing.T){
	fmt.Println(nni.Parse(&RouteAppleGet{}))
}
