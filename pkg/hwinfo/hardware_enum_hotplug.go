// Code generated by "enumer -type=Hotplug -json -transform=snake -trimprefix Hotplug -output=./hardware_enum_hotplug.go"; DO NOT EDIT.

package hwinfo

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _HotplugName = "nonepcmciacardbuspciusbfirewire"

var _HotplugIndex = [...]uint8{0, 4, 10, 17, 20, 23, 31}

const _HotplugLowerName = "nonepcmciacardbuspciusbfirewire"

func (i Hotplug) String() string {
	if i < 0 || i >= Hotplug(len(_HotplugIndex)-1) {
		return fmt.Sprintf("Hotplug(%d)", i)
	}
	return _HotplugName[_HotplugIndex[i]:_HotplugIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _HotplugNoOp() {
	var x [1]struct{}
	_ = x[HotplugNone-(0)]
	_ = x[HotplugPcmcia-(1)]
	_ = x[HotplugCardbus-(2)]
	_ = x[HotplugPci-(3)]
	_ = x[HotplugUsb-(4)]
	_ = x[HotplugFirewire-(5)]
}

var _HotplugValues = []Hotplug{HotplugNone, HotplugPcmcia, HotplugCardbus, HotplugPci, HotplugUsb, HotplugFirewire}

var _HotplugNameToValueMap = map[string]Hotplug{
	_HotplugName[0:4]:        HotplugNone,
	_HotplugLowerName[0:4]:   HotplugNone,
	_HotplugName[4:10]:       HotplugPcmcia,
	_HotplugLowerName[4:10]:  HotplugPcmcia,
	_HotplugName[10:17]:      HotplugCardbus,
	_HotplugLowerName[10:17]: HotplugCardbus,
	_HotplugName[17:20]:      HotplugPci,
	_HotplugLowerName[17:20]: HotplugPci,
	_HotplugName[20:23]:      HotplugUsb,
	_HotplugLowerName[20:23]: HotplugUsb,
	_HotplugName[23:31]:      HotplugFirewire,
	_HotplugLowerName[23:31]: HotplugFirewire,
}

var _HotplugNames = []string{
	_HotplugName[0:4],
	_HotplugName[4:10],
	_HotplugName[10:17],
	_HotplugName[17:20],
	_HotplugName[20:23],
	_HotplugName[23:31],
}

// HotplugString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func HotplugString(s string) (Hotplug, error) {
	if val, ok := _HotplugNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _HotplugNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Hotplug values", s)
}

// HotplugValues returns all values of the enum
func HotplugValues() []Hotplug {
	return _HotplugValues
}

// HotplugStrings returns a slice of all String values of the enum
func HotplugStrings() []string {
	strs := make([]string, len(_HotplugNames))
	copy(strs, _HotplugNames)
	return strs
}

// IsAHotplug returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Hotplug) IsAHotplug() bool {
	for _, v := range _HotplugValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Hotplug
func (i Hotplug) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Hotplug
func (i *Hotplug) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Hotplug should be a string, got %s", data)
	}

	var err error
	*i, err = HotplugString(s)
	return err
}
