package block

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Device struct {
	Alignment    int           `json:"alignment,omitempty"`
	IdLink       string        `json:"id-link,omitempty"`
	Id           string        `json:"id,omitempty"`
	DiscAln      int           `json:"disc-aln,omitempty"`
	Dax          bool          `json:"dax,omitempty"`
	DiscGran     string        `json:"disc-gran,omitempty"`
	DiskSeq      int           `json:"disk-seq,omitempty"`
	DiscMax      string        `json:"disc-max,omitempty"`
	DiscZero     bool          `json:"disc-zero,omitempty"`
	Fsavail      interface{}   `json:"fsavail,omitempty"`
	Fsroots      []interface{} `json:"fsroots,omitempty"`
	Fssize       interface{}   `json:"fssize,omitempty"`
	Fstype       interface{}   `json:"fstype,omitempty"`
	Fsused       interface{}   `json:"fsused,omitempty"`
	Fsuse        interface{}   `json:"fsuse%,omitempty"`
	Fsver        interface{}   `json:"fsver,omitempty"`
	Group        string        `json:"group,omitempty"`
	Hctl         interface{}   `json:"hctl,omitempty"`
	Hotplug      bool          `json:"hotplug,omitempty"`
	Kname        string        `json:"kname,omitempty"`
	Label        interface{}   `json:"label,omitempty"`
	LogSec       int           `json:"log-sec,omitempty"`
	MajMin       string        `json:"maj:min,omitempty"`
	MinIo        int           `json:"min-io,omitempty"`
	Mode         string        `json:"mode,omitempty"`
	Model        string        `json:"model,omitempty"`
	Mq           string        `json:"mq,omitempty"`
	Name         string        `json:"name,omitempty"`
	OptIo        int           `json:"opt-io,omitempty"`
	Owner        string        `json:"owner,omitempty"`
	Partflags    interface{}   `json:"partflags,omitempty"`
	Partlabel    interface{}   `json:"partlabel,omitempty"`
	Partn        interface{}   `json:"partn,omitempty"`
	Parttype     interface{}   `json:"parttype,omitempty"`
	Parttypename interface{}   `json:"parttypename,omitempty"`
	Partuuid     interface{}   `json:"partuuid,omitempty"`
	Path         string        `json:"path,omitempty"`
	PhySec       int           `json:"phy-sec,omitempty"`
	Pkname       interface{}   `json:"pkname,omitempty"`
	Pttype       string        `json:"pttype,omitempty"`
	Ptuuid       string        `json:"ptuuid,omitempty"`
	Ra           int           `json:"ra,omitempty"`
	Rand         bool          `json:"rand,omitempty"`
	Rev          string        `json:"rev,omitempty"`
	Rm           bool          `json:"rm,omitempty"`
	Ro           bool          `json:"ro,omitempty"`
	Rota         bool          `json:"rota,omitempty"`
	RqSize       int           `json:"rq-size,omitempty"`
	Sched        string        `json:"sched,omitempty"`
	Serial       string        `json:"serial,omitempty"`
	Size         string        `json:"size,omitempty"`
	Start        interface{}   `json:"start,omitempty"`
	State        string        `json:"state,omitempty"`
	Subsystems   string        `json:"subsystems,omitempty"`
	Mountpoint   interface{}   `json:"mountpoint,omitempty"`
	Mountpoints  []interface{} `json:"mountpoints,omitempty"`
	Tran         string        `json:"tran,omitempty"`
	Type         string        `json:"type,omitempty"`
	Uuid         interface{}   `json:"uuid,omitempty"`
	Vendor       interface{}   `json:"vendor,omitempty"`
	Wsame        string        `json:"wsame,omitempty"`
	Wwn          string        `json:"wwn,omitempty"`
	Zoned        string        `json:"zoned,omitempty"`
	ZoneSz       string        `json:"zone-sz,omitempty"`
	ZoneWgran    string        `json:"zone-wgran,omitempty"`
	ZoneApp      string        `json:"zone-app,omitempty"`
	ZoneNr       int           `json:"zone-nr,omitempty"`
	ZoneOmax     int           `json:"zone-omax,omitempty"`
	ZoneAmax     int           `json:"zone-amax,omitempty"`
	Children     []Device      `json:"children,omitempty"`

	// KernelModule is the result of probing /device/driver/module
	KernelModule string `json:"module,omitempty"`
}

type lsblkOutput struct {
	Devices []*Device `json:"blockdevices"`
}

func Scan() ([]*Device, error) {

	cmd := exec.Command("lsblk", "-O", "-a", "-J")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run lsblk: %w", err)
	}

	// todo unsure if this captures mmc

	var output lsblkOutput
	if err = json.Unmarshal(b, &output); err != nil {
		return nil, fmt.Errorf("failed to unmarshal lsblk output: %w", err)
	}

	for _, dev := range output.Devices {
		path, err := os.Readlink(dev.Path + "/device/driver/module")
		if err != nil {
			// todo add some logging and check error
			continue
		}

		println("setting module")
		dev.KernelModule = filepath.Base(path)
	}

	return output.Devices, nil
}
