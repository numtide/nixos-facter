package ephem

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var (
	swapHeaderRegex = regexp.MustCompile(`^Filename\s+Type\s+Size\s+Used\s+Priority$`)
	swapEntryRegex  = regexp.MustCompile(`^(.*?)\s+(partition|file)\s+(.*?)\s+(.*?)\s+(.*?)$`)
)

//go:generate enumer -type=SwapType -json -transform=snake -trimprefix SwapType -output=./swap_enum_type.go
type SwapType uint

const (
	SwapTypeFile SwapType = iota
	SwapTypePartition
)

type SwapEntry struct {
	Filename string   `json:"path"`
	Type     SwapType `json:"type"`
	Size     uint64   `json:"size"`
	Used     uint64   `json:"used"`
	Priority int32    `json:"priority"`
}

func SwapEntries() ([]*SwapEntry, error) {
	f, err := os.Open("/proc/swaps")
	if err != nil {
		return nil, err
	}

	devices, err := ReadSwapFile(f)
	if err != nil {
		return nil, err
	}

	for idx := range devices {
		// try to resolve stable device paths for each swap device
		stablePath, err := StableDevicePath(devices[idx].Filename)
		if err != nil {
			return nil, err
		}
		devices[idx].Filename = stablePath
	}

	return devices, nil
}

func ReadSwapFile(reader io.Reader) ([]*SwapEntry, error) {
	scanner := bufio.NewScanner(reader)
	if !scanner.Scan() {
		return nil, fmt.Errorf("swaps file is empty")
	} else if b := scanner.Bytes(); !swapHeaderRegex.Match(b) {
		return nil, fmt.Errorf("header in swaps file is malformed: '%s'", string(b))
	}

	var result []*SwapEntry
	for scanner.Scan() {
		line := scanner.Text()

		matches := swapEntryRegex.FindAllStringSubmatch(line, 1)
		if len(matches) != 1 {
			return nil, fmt.Errorf("malformed entry in swaps file: '%s'", line)
		}

		fields := matches[0]

		swapType, err := SwapTypeString(fields[2])
		if err != nil {
			return nil, fmt.Errorf("malformed swap type: '%s'", fields[2])
		}

		size, err := strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("malformed size value: '%s'", fields[3])
		}

		used, err := strconv.ParseUint(fields[4], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("malformed used value: '%s'", fields[4])
		}

		priority, err := strconv.ParseInt(fields[5], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("malformed priority value: '%s'", fields[5])
		}

		result = append(result, &SwapEntry{
			Filename: fields[1],
			Type:     swapType,
			Size:     size,
			Used:     used,
			Priority: int32(priority),
		})
	}

	return result, nil
}
