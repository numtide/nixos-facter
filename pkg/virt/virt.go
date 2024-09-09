// Package virt contains utilities for detecting virtualized environments.
// It has been adapted from [systemd].
//
// [systemd]: https://github.com/systemd/systemd/blob/main/src/basic/virt.c
package virt

// Type represents various virtualization and container types.
//
//go:generate enumer -type=Type -json -transform=snake -trimprefix Type -output=./virt_enum_type.go
type Type int

const (
	TypeNone Type = iota
	TypeKvm
	TypeAmazon
	TypeQemu
	TypeBochs
	TypeXen
	TypeUml
	TypeVmware
	TypeOracle
	TypeMicrosoft
	TypeZvm
	TypeParallels
	TypeBhyve
	TypeQnx
	TypeAcrn
	TypePowerVM
	TypeApple
	TypeGoogle
	TypeSre
	TypeVmOther
	TypeSystemdNspawn
	TypeLxcLibvirt
	TypeLxc
	TypeOpenvz
	TypeDocker
	TypePodman
	TypeRkt
	TypeWsl
	TypeProot
	TypePouch
	TypeContainerOther
)

// Detect identifies the virtualization type of the current system.
// Returns the detected Type and an error if detection fails.
func Detect() (Type, error) {
	// todo do we care about detecting if we are in a container?
	return detectVM()
}
