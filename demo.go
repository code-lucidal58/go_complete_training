package main

import (
	"fmt"
	"github.com/ugorji/go/codec"
	"io/ioutil"
	"reflect"
)

// create and configure Handle
var (
	js codec.JsonHandle
)

type RestApiV2 struct {
	Last  string `json:"last"`
	First string `json:"first"`
}

func main() {
	js.SliceType = reflect.TypeOf([]interface{}(nil))
	var restapi []interface{}
	// configure extensions
	// e.g. for msgpack, define functions and enable Time support for tag 1
	// mh.SetExt(reflect.TypeOf(time.Time{}), 1, myExt)

	// create and use decoder/encoder
	//r, _ := os.Open("restapi-v2.j¬son")
	b, _ := ioutil.ReadFile("restapi-v2.json")
	h := &js

	//dec := codec.NewDecoder(r, h)
	dec := codec.NewDecoderBytes(b, h)
	err := dec.Decode(&restapi)
	fmt.Println(err)
	fmt.Println(restapi)

}
