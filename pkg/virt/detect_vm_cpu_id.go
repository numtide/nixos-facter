//go:build 386 || amd64

package virt

import (
	"github.com/klauspost/cpuid/v2"
)

// See https://github.com/systemd/systemd/blob/main/src/basic/virt.c

var vendorMapping = map[cpuid.Vendor]Type{
	cpuid.XenHVM: TypeXen,
	cpuid.KVM:    TypeKvm,
	cpuid.QEMU:   TypeQemu,
	cpuid.VMware: TypeVmware,
	cpuid.MSVM:   TypeMicrosoft,
	cpuid.Bhyve:  TypeBhyve,
	cpuid.QNX:    TypeQnx,
	cpuid.ACRN:   TypeAcrn,
	cpuid.SRE:    TypeSre,
	cpuid.Apple:  TypeApple,
}

func init() {
	// https://lwn.net/Articles/301888/
	detectVmCpuId = func() (Type, error) {
		if !cpuid.CPU.VM() {
			return TypeNone, nil
		}
		vm, ok := vendorMapping[cpuid.CPU.HypervisorVendorID]
		if !ok {
			return TypeVmOther, nil
		}
		return vm, nil
	}
}
