package main

import (
	// "net/http"

	// "github.com/codegangsta/negroni"
	// "github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"

	"gopkg.in/mgo.v2"
)

var(
	renderer	*render.Render
	mongoSession	*mgo.Session
)

func init(){
	//render생성
	renderer = render.New()
	
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	mongoSession = s
}