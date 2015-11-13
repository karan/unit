package unit

import (
  "github.com/go-martini/martini"
)

// A group type that encapsulates martini router type.
type Group func(router martini.Router)

// Register a router group
func (g Group) Register(path string) {
  println("register")
  m.Router.Group(path, g)
}

// The main app
// Private
var m *martini.ClassicMartini

func init() {
  println("setup")
  m = martini.Classic()
}

func Run() {
  m.Run()
}
