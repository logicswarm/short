package util

import (
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Int32Ptr(i int32) *int32 {
	return &i
}

func IntOrStringPtr(i intstr.IntOrString) *intstr.IntOrString {
	return &i
}