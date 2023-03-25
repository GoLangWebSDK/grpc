package grpc

import (
	"github.com/PeteProgrammer/go-automapper"
)

// Map auto map source model to destination
func Map(source, dest interface{}) {

	automapper.MapLoose(source, dest)
}

// MapList auto map source model list to destination
func MapList(source, dest []interface{}) {

	automapper.MapLoose(source, dest)
}
