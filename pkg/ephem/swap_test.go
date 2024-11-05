package ephem

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	empty   = "Filename                                Type            Size    Used    Priority\n"
	corrupt = `Filename                                Type            Size    Used    Priority
foo bar baz
`
	badPartition = `Filename				Type		Size	Used	Priority
/var/lib/swap-1							foo        1048576	123		-3
`

	sample = `Filename				Type		Size	Used	Priority
/dev/sda6                               partition   4194300	0		-1
/var/lib/swap-1							file        1048576	123		-3
/var/lib/swap-2 						file        2097152	4567	-2`
)

func TestReadSwapFile(t *testing.T) {
	as := require.New(t)

	_, err := ReadSwapFile(strings.NewReader(""))
	as.Error(err, "swaps file is empty")

	_, err = ReadSwapFile(strings.NewReader("foo bar baz hello world\n"))
	as.Errorf(err, "header in swaps file is malformed: '%s'", "foo bar baz hello world\n")

	swaps, err := ReadSwapFile(strings.NewReader(empty))
	as.NoError(err)
	as.Empty(swaps)

	_, err = ReadSwapFile(strings.NewReader(corrupt))
	as.Errorf(err, "malformed entry in swaps file: '%s'", "foo bar baz")

	_, err = ReadSwapFile(strings.NewReader(badPartition))
	as.Errorf(err, "malformed entry in swaps file: '%s'", `/var/lib/swap-1							foo        1048576	123		-3`)

	swaps, err = ReadSwapFile(strings.NewReader(sample))
	as.NoError(err)
	as.Len(swaps, 3)

	as.Equal(swaps[0].Filename, "/dev/sda6")
	as.Equal(swaps[0].Type, SwapTypePartition)
	as.Equal(swaps[0].Size, uint64(4194300))
	as.Equal(swaps[0].Used, uint64(0))
	as.Equal(swaps[0].Priority, int32(-1))

	as.Equal(swaps[1].Filename, "/var/lib/swap-1")
	as.Equal(swaps[1].Type, SwapTypeFile)
	as.Equal(swaps[1].Size, uint64(1048576))
	as.Equal(swaps[1].Used, uint64(123))
	as.Equal(swaps[1].Priority, int32(-3))

	as.Equal(swaps[2].Filename, "/var/lib/swap-2")
	as.Equal(swaps[2].Type, SwapTypeFile)
	as.Equal(swaps[2].Size, uint64(2097152))
	as.Equal(swaps[2].Used, uint64(4567))
	as.Equal(swaps[2].Priority, int32(-2))
}
