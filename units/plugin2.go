package units

import (
  "github.com/go-martini/martini"

  "./../unit"

  "github.com/martini-contrib/render" //- unit-dep
)

func init() {
  g := unit.Group(func(router martini.Router) {

    // Suppress "imported and not used" warning
    render.Renderer()

    router.Get("/1", func() string {
      return "v2 - 1!"
    })

    router.Get("/2", func() string {
      return "v2 - 2!"
    })

  })

  g.Register("/v2")
}
