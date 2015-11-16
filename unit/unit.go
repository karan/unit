package unit

import (
  "github.com/go-martini/martini"
)

// A group type that encapsulates martini router type.
type Group func(router martini.Router)

// Register a router group
func (g Group) Register(path string) {
  m.Router.Group(path, g)
}

// The main app
// Private
var m *martini.ClassicMartini

func init() {
  m = martini.Classic()
}

func Run() {
  m.RunOnAddr(":5000")
}
