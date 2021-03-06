package units

import (
  "github.com/go-martini/martini"

  "./../unit"
)

func init() {
  g := unit.Group(func(router martini.Router) {

    router.Get("/1", func() string {
      return "v1 - 1!"
    })

    router.Get("/2", func() string {
      return "v1 - 2!"
    })

  })

  g.Register("/v1")

}
