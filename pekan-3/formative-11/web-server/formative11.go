package webserver

import (
	"fmt"
	"math"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	radius = 7
	height = 10
)

func areaOfBase(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func circumferenceOfBase(radius float64) float64 {
	return 2 * math.Pi * radius
}

func volumeOfCylinder(radius, height float64) float64 {
	return math.Pi * math.Pow(radius, 2) * height
}

func WebServer() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, response *http.Request, _ httprouter.Params) {
		area := areaOfBase(radius)
		circumference := circumferenceOfBase(radius)
		volume := volumeOfCylinder(radius, height)

		fmt.Fprintf(writer, "jariJari: %d, tinggi: %d, volume: %.2f, luas alas: %.2f, keliling alas: %.2f",
			radius, height, volume, area, circumference)
	})

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
	server.ListenAndServe()

}
