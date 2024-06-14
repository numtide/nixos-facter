package nix

import (
	"fmt"
	"strings"
)

func ToNixList(args []string) string {
	return strings.Join(args, " ")
}

func ToNixStringList(args []string) string {
	wrapped := make([]string, len(args))
	for idx := range args {
		wrapped[idx] = fmt.Sprintf(`"%s"`, args[idx])
	}
	return ToNixList(wrapped)
}

func MultiLineList(indent string, args []string) string {
	if len(args) == 0 {
		return " [ ]"
	}
	res := "\n" + indent
	for idx, arg := range args {
		if idx > 0 {
			res = res + indent + "  "
		}
		res += arg + "\n"
	}
	res += indent + "]"
	return res
}
