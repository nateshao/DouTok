// Code generated by hertz generator. DO NOT EDIT.

package main

import (
	router "github.com/TremblingV5/DouTok/applications/api/biz/router"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
