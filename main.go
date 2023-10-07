package main

import (
	_ "unsafe"

	_ "k8s.io/api/core/v1"
)

func main() {
	// pod := v1.Pod{}

	// json.Marshal(pod)
}

//go:linkname unsafe_New reflect.unsafe_New
func unsafe_New() {}

//go:linkname typedmemmove reflect.typedmemmove
func typedmemmove() {}

//go:linkname typedslicecopy reflect.typedslicecopy
func typedslicecopy() {}

//go:linkname unsafe_NewArray reflect.unsafe_NewArray
func unsafe_NewArray() {}

//go:linkname unsafe_makemap reflect.makemap
func unsafe_makemap() {}

//go:linkname unsafe_mapassign reflect.mapassign
func unsafe_mapassign() {}
