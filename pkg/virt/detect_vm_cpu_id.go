//go:build 386 || amd64

package virt

import (
	"github.com/klauspost/cpuid/v2"
)

// See https://github.com/systemd/systemd/blob/main/src/basic/virt.c

var cpuMapping = map[string]Type{
	"XenVMMXenVMM": TypeXen,
	"KVMKVMKVM":    TypeKvm, // qemu with KVM
	"Linux KVM Hv": TypeKvm,
	"TCGTCGTCGTCG": TypeQemu, // qemu without KVM
	/* http://kb.vmware.com/selfservice/microsites/search.do?language=en_US&cmd=displayKC&externalId=1009458 */
	"VMwareVMware": TypeVmware,
	/* https://docs.microsoft.com/en-us/virtualization/hyper-v-on-windows/reference/tlfs */
	"Microsoft Hv": TypeMicrosoft,
	/* https://wiki.freebsd.org/bhyve */
	"bhyve bhyve ": TypeBhyve,
	"QNXQVMBSQG":   TypeQnx,
	/* https://wiki.freebsd.org/bhyve */
	"ACRNACRNACRN": TypeAcrn,
	/* https://www.lockheedmartin.com/en-us/products/Hardened-Security-for-Intel-Processors.html */
	"SRESRESRESRE": TypeSre,
	"Apple VZ":     TypeApple,
}

func init() {
	// https://lwn.net/Articles/301888/
	detectVmCpuId = _detectVmCpuId
}

func _detectVmCpuId() (Type, error) {
	if !cpuid.CPU.VM() {
		return TypeNone, nil
	}

	_, b, c, d := cpuid.Cpuid(0x40000000)
	v := string(cpuid.ValAsString(b, c, d))

	if vmType, ok := cpuMapping[cpuid.CPU.VendorString]; ok {
		return vmType, nil
	}
	vm, ok := cpuMapping[v]
	if !ok {
		return TypeVmOther, nil
	}
	return vm, nil
}
