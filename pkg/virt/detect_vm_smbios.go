package virt

import "os"

type SmbiosBit int

const (
	SmbiosBitUnknown SmbiosBit = iota
	SmbiosBitSet
	SmbiosBitUnset
)

func detectSmbios() SmbiosBit {
	/* The SMBIOS BIOS Characteristics Extension Byte 2 (Section 2.1.2.2 of
	 * https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.4.0.pdf), specifies that
	 * the 4th bit being set indicates a VM. The BIOS Characteristics table is exposed via the kernel in
	 * /sys/firmware/dmi/entries/0-0. Note that in the general case, this bit being unset should not
	 * imply that the system is running on bare-metal.  For example, QEMU 3.1.0 (with or without KVM)
	 * with SeaBIOS does not set this bit. */

	b, err := os.ReadFile("/sys/firmware/dmi/entries/0-0/raw")
	if err != nil {
		// todo log error
		return SmbiosBitUnknown
	}

	if len(b) < 20 || b[1] < 20 {
		// todo log error
		/* The spec indicates that byte 1 contains the size of the table, 0x12 + the number of
		 * extension bytes. The data we're interested in is in extension byte 2, which would be at
		 * 0x13. If we didn't read that much data, or if the BIOS indicates that we don't have that
		 * much data, we don't infer anything from the SMBIOS. */
		return SmbiosBitUnknown
	}

	if b[19]&(1<<4) == 1 {
		// todo add logging
		return SmbiosBitSet
	}

	// todo add logging
	return SmbiosBitUnset
}
