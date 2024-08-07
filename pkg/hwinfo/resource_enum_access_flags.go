// Code generated by "enumer -type=AccessFlags -json -transform=snake -trimprefix AccessFlags -output=./resource_enum_access_flags.go"; DO NOT EDIT.

package hwinfo

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _AccessFlagsName = "unknownread_onlywrite_onlyread_write"

var _AccessFlagsIndex = [...]uint8{0, 7, 16, 26, 36}

const _AccessFlagsLowerName = "unknownread_onlywrite_onlyread_write"

func (i AccessFlags) String() string {
	if i >= AccessFlags(len(_AccessFlagsIndex)-1) {
		return fmt.Sprintf("AccessFlags(%d)", i)
	}
	return _AccessFlagsName[_AccessFlagsIndex[i]:_AccessFlagsIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _AccessFlagsNoOp() {
	var x [1]struct{}
	_ = x[AccessFlagsUnknown-(0)]
	_ = x[AccessFlagsReadOnly-(1)]
	_ = x[AccessFlagsWriteOnly-(2)]
	_ = x[AccessFlagsReadWrite-(3)]
}

var _AccessFlagsValues = []AccessFlags{AccessFlagsUnknown, AccessFlagsReadOnly, AccessFlagsWriteOnly, AccessFlagsReadWrite}

var _AccessFlagsNameToValueMap = map[string]AccessFlags{
	_AccessFlagsName[0:7]:        AccessFlagsUnknown,
	_AccessFlagsLowerName[0:7]:   AccessFlagsUnknown,
	_AccessFlagsName[7:16]:       AccessFlagsReadOnly,
	_AccessFlagsLowerName[7:16]:  AccessFlagsReadOnly,
	_AccessFlagsName[16:26]:      AccessFlagsWriteOnly,
	_AccessFlagsLowerName[16:26]: AccessFlagsWriteOnly,
	_AccessFlagsName[26:36]:      AccessFlagsReadWrite,
	_AccessFlagsLowerName[26:36]: AccessFlagsReadWrite,
}

var _AccessFlagsNames = []string{
	_AccessFlagsName[0:7],
	_AccessFlagsName[7:16],
	_AccessFlagsName[16:26],
	_AccessFlagsName[26:36],
}

// AccessFlagsString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AccessFlagsString(s string) (AccessFlags, error) {
	if val, ok := _AccessFlagsNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _AccessFlagsNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AccessFlags values", s)
}

// AccessFlagsValues returns all values of the enum
func AccessFlagsValues() []AccessFlags {
	return _AccessFlagsValues
}

// AccessFlagsStrings returns a slice of all String values of the enum
func AccessFlagsStrings() []string {
	strs := make([]string, len(_AccessFlagsNames))
	copy(strs, _AccessFlagsNames)
	return strs
}

// IsAAccessFlags returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AccessFlags) IsAAccessFlags() bool {
	for _, v := range _AccessFlagsValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for AccessFlags
func (i AccessFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccessFlags
func (i *AccessFlags) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AccessFlags should be a string, got %s", data)
	}

	var err error
	*i, err = AccessFlagsString(s)
	return err
}
