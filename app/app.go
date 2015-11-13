package app

import (
    "github.com/go-martini/martini"
)


var App *martini.ClassicMartini


type Group func(router martini.Router)


func Register(path string, grp Group) {
    println("register")
    App.Router.Group(path, grp)
}


func init() {
    println("setup")
    App = martini.Classic()
}
