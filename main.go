package main

import (
	"gingonic/route"
)

func main() {
	r := route.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
