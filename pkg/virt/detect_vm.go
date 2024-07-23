package virt

import "errors"

var (
	detectVmDmi        = detectVmUnsupported
	detectVmCpuId      = detectVmUnsupported
	detectVmDeviceTree = detectVmUnsupported
)

func detectVmUnsupported() (Type, error) {
	return TypeNone, nil
}

func detectVM() (v Type, err error) {
	var dmi Type
	var dmiErr error

	var other, hyperv bool
	var xenDom0 bool

	finish := func() (Type, error) {
		// None of the checks gave us a clear answer, so we use this fallback logic: if hyperv
		// enlightenments are available but the VMM wasn't recognized as anything yet, it's probably
		// Microsoft.
		if v == TypeNone {
			if hyperv {
				v = TypeMicrosoft
			} else if other {
				v = TypeVmOther
			}
		}
		return v, nil
	}

	// We have to use the correct order here:
	//
	// → First, try to detect Oracle Virtualbox, Amazon EC2 Nitro, Parallels, and Google Compute Engine,
	//   even if they use KVM, as well as Xen, even if it cloaks as Microsoft Hyper-V. Attempt to detect
	//   UML at this stage too, since it runs as a user-process nested inside other VMs. Also check for
	//   Xen now, because Xen PV mode does not override CPUID when nested inside another hypervisor.
	//
	// → Second, try to detect from CPUID. This will report KVM for whatever software is used even if
	//   info in DMI is overwritten.
	//
	// → Third, try to detect from DMI.
	dmi, dmiErr = detectVmDmi()

	switch dmi {
	case TypeOracle, TypeXen, TypeAmazon, TypeParallels, TypeGoogle:
		v = dmi
		return finish()
	default:
		// do nothing, let's continue
	}

	// detect UML
	if v, err = detectUml(); err != nil {
		return 0, err
	} else if v != TypeNone {
		return finish()
	}

	// detect Xen
	if v, err = detectXen(); err != nil {
		return 0, err
	} else if v == TypeXen {
		// If we are Dom0, then we expect to not report as a VM. However, as we might be nested
		// inside another hypervisor which can be detected via the CPUID check, wait to report this
		// until after the CPUID check.
		if xenDom0, err = detectXenDom0(); err != nil {
			return 0, err
		} else if !xenDom0 {
			return finish()
		}
	} else if v != TypeNone {
		return 0, errors.New("should not reach this branch")
	}

	// detect from cpuid
	if v, err = detectVmCpuId(); err != nil {
		return 0, err
	} else if v == TypeMicrosoft {
		// QEMU sets the CPUID string to hyperv's, in case it provides hyperv enlightenments. Let's
		// hence not return Microsoft here but just use the other mechanisms first to make a better
		// decision. */
		hyperv = true
	} else if v == TypeVmOther {
		other = true
	} else if v != TypeNone {
		return finish()
	}

	// if we are in Dom0 and have not yet finished, finish with the result of detect_vm_cpuid */
	if xenDom0 {
		return finish()
	}

	// let's get back to DMI
	if dmiErr != nil {
		return 0, dmiErr
	} else if dmi == TypeVmOther {
		other = true
	} else if dmi != TypeNone {
		v = dmi
		return finish()
	}

	// check high-level hypervisor sysfs file
	if v, err = detectHypervisor(); err != nil {
		return 0, err
	} else if v == TypeVmOther {
		other = true
	} else if v != TypeNone {
		return finish()
	}

	if v, err = detectVmDeviceTree(); err != nil {
		return 0, err
	} else if v == TypeVmOther {
		other = true
	} else if v != TypeNone {
		return finish()
	}

	// todo detect_vm_zvm

	return finish()
}
