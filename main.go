package main

import (
	"encoding/json"
	"fmt"
	_ "unsafe"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func main() {
	// pod := v1.Pod{}

	// json.Marshal(pod)

	test := appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "test",
						},
					},
				},
			},
		},
	}

	data, err := json.Marshal(test)
	if err != nil {
		panic(fmt.Sprintf("json.Marshal failed: %v", err))
	}

	fmt.Println(string(data))
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
