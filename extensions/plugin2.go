package extensions

import (
    "github.com/go-martini/martini"

    app "./../app"
)


func init() {
    app.Register("/v2", app.Group(func (router martini.Router) {

        router.Get("/1", func() string {
            return "v2 - 1!"
        })

        router.Get("/2", func() string {
            return "v2 - 2!"
        })

    }))

}
