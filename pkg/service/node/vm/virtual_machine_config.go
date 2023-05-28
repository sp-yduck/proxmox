package vm

import (
	"reflect"
	"strconv"
	"strings"
)

type VirtualMachineConfig struct {
	// PVE Metadata
	Digest      string `json:"digest"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Meta        string `json:"meta,omitempty"`
	VMGenID     string `json:"vmgenid,omitempty"`
	Hookscript  string `json:"hookscript,omitempty"`
	Hotplug     string `json:"hotplug,omitempty"`
	Template    int    `json:"template,omitempty"`

	Tags      string   `json:"tags,omitempty"`
	TagsSlice []string `json:"-"` // internal helper to manage tags easier

	Protection int    `json:"protection,omitempty"`
	Lock       string `json:"lock,omitempty"`

	// Boot configuration
	Boot   string `json:"boot,omitempty"`
	OnBoot int    `json:"onboot,omitempty"`

	// Qemu general specs
	OSType  string `json:"ostype,omitempty"`
	Machine string `json:"machine,omitempty"`
	Args    string `json:"args,omitempty"`

	// Qemu firmware specs
	Bios     string `json:"bios,omitempty"`
	EFIDisk0 string `json:"efidisk0,omitempty"`
	SMBios1  string `json:"smbios1,omitempty"`
	Acpi     int    `json:"acpi,omitempty"`

	// Qemu CPU specs
	Sockets  int    `json:"sockets,omitempty"`
	Cores    int    `json:"cores,omitempty"`
	CPU      string `json:"cpu,omitempty"`
	CPULimit int    `json:"cpulimit,omitempty"`
	CPUUnits int    `json:"cpuunits,omitempty"`
	Vcpus    int    `json:"vcpus,omitempty"`
	Affinity string `json:"affinity,omitempty"`

	// Qemu memory specs
	Numa      int    `json:"numa,omitempty"`
	Memory    int    `json:"memory,omitempty"`
	Hugepages string `json:"hugepages,omitempty"`
	Balloon   int    `json:"balloon,omitempty"`

	// Other Qemu devices
	VGA       string `json:"vga,omitempty"`
	SCSIHW    string `json:"scsihw,omitempty"`
	TPMState0 string `json:"tpmstate0,omitempty"`
	Rng0      string `json:"rng0,omitempty"`
	Audio0    string `json:"audio0,omitempty"`

	// Disk devices
	IDEs map[string]string `json:"-"`
	IDE0 string            `json:"ide0,omitempty"`
	IDE1 string            `json:"ide1,omitempty"`
	IDE2 string            `json:"ide2,omitempty"`
	IDE3 string            `json:"ide3,omitempty"`

	SCSIs  map[string]string `json:"-"`
	SCSI0  string            `json:"scsi0,omitempty"`
	SCSI1  string            `json:"scsi1,omitempty"`
	SCSI2  string            `json:"scsi2,omitempty"`
	SCSI3  string            `json:"scsi3,omitempty"`
	SCSI4  string            `json:"scsi4,omitempty"`
	SCSI5  string            `json:"scsi5,omitempty"`
	SCSI6  string            `json:"scsi6,omitempty"`
	SCSI7  string            `json:"scsi7,omitempty"`
	SCSI8  string            `json:"scsi8,omitempty"`
	SCSI9  string            `json:"scsi9,omitempty"`
	SCSI10 string            `json:"scsi10,omitempty"`
	SCSI11 string            `json:"scsi11,omitempty"`
	SCSI12 string            `json:"scsi12,omitempty"`
	SCSI13 string            `json:"scsi13,omitempty"`
	SCSI14 string            `json:"scsi14,omitempty"`
	SCSI15 string            `json:"scsi15,omitempty"`
	SCSI16 string            `json:"scsi16,omitempty"`
	SCSI17 string            `json:"scsi17,omitempty"`
	SCSI18 string            `json:"scsi18,omitempty"`
	SCSI19 string            `json:"scsi19,omitempty"`
	SCSI20 string            `json:"scsi20,omitempty"`
	SCSI21 string            `json:"scsi21,omitempty"`
	SCSI22 string            `json:"scsi22,omitempty"`
	SCSI23 string            `json:"scsi23,omitempty"`
	SCSI24 string            `json:"scsi24,omitempty"`
	SCSI25 string            `json:"scsi25,omitempty"`
	SCSI26 string            `json:"scsi26,omitempty"`
	SCSI27 string            `json:"scsi27,omitempty"`
	SCSI28 string            `json:"scsi28,omitempty"`
	SCSI29 string            `json:"scsi29,omitempty"`
	SCSI30 string            `json:"scsi30,omitempty"`

	SATAs map[string]string `json:"-"`
	SATA0 string            `json:"sata0,omitempty"`
	SATA1 string            `json:"sata1,omitempty"`
	SATA2 string            `json:"sata2,omitempty"`
	SATA3 string            `json:"sata3,omitempty"`
	SATA4 string            `json:"sata4,omitempty"`
	SATA5 string            `json:"sata5,omitempty"`

	VirtIOs  map[string]string `json:"-"`
	VirtIO0  string            `json:"virtio0,omitempty"`
	VirtIO1  string            `json:"virtio1,omitempty"`
	VirtIO2  string            `json:"virtio2,omitempty"`
	VirtIO3  string            `json:"virtio3,omitempty"`
	VirtIO4  string            `json:"virtio4,omitempty"`
	VirtIO5  string            `json:"virtio5,omitempty"`
	VirtIO6  string            `json:"virtio6,omitempty"`
	VirtIO7  string            `json:"virtio7,omitempty"`
	VirtIO8  string            `json:"virtio8,omitempty"`
	VirtIO9  string            `json:"virtio9,omitempty"`
	VirtIO10 string            `json:"virtio10,omitempty"`
	VirtIO11 string            `json:"virtio11,omitempty"`
	VirtIO12 string            `json:"virtio12,omitempty"`
	VirtIO13 string            `json:"virtio13,omitempty"`
	VirtIO14 string            `json:"virtio14,omitempty"`
	VirtIO15 string            `json:"virtio15,omitempty"`

	Unuseds map[string]string `json:"-"`
	Unused0 string            `json:"unused0,omitempty"`
	Unused1 string            `json:"unused1,omitempty"`
	Unused2 string            `json:"unused2,omitempty"`
	Unused3 string            `json:"unused3,omitempty"`
	Unused4 string            `json:"unused4,omitempty"`
	Unused5 string            `json:"unused5,omitempty"`
	Unused6 string            `json:"unused6,omitempty"`
	Unused7 string            `json:"unused7,omitempty"`
	Unused8 string            `json:"unused8,omitempty"`
	Unused9 string            `json:"unused9,omitempty"`

	// Network devices
	Nets map[string]string `json:"-"`
	Net0 string            `json:"net0,omitempty"`
	Net1 string            `json:"net1,omitempty"`
	Net2 string            `json:"net2,omitempty"`
	Net3 string            `json:"net3,omitempty"`
	Net4 string            `json:"net4,omitempty"`
	Net5 string            `json:"net5,omitempty"`
	Net6 string            `json:"net6,omitempty"`
	Net7 string            `json:"net7,omitempty"`
	Net8 string            `json:"net8,omitempty"`
	Net9 string            `json:"net9,omitempty"`

	// NUMA topology
	Numas map[string]string `json:"-"`
	Numa0 string            `json:"numa0,omitempty"`
	Numa1 string            `json:"numa1,omitempty"`
	Numa2 string            `json:"numa2,omitempty"`
	Numa3 string            `json:"numa3,omitempty"`
	Numa4 string            `json:"numa4,omitempty"`
	Numa5 string            `json:"numa5,omitempty"`
	Numa6 string            `json:"numa6,omitempty"`
	Numa7 string            `json:"numa7,omitempty"`
	Numa8 string            `json:"numa8,omitempty"`
	Numa9 string            `json:"numa9,omitempty"`

	// Host PCI devices
	HostPCIs map[string]string `json:"-"`
	HostPCI0 string            `json:"hostpci0,omitempty"`
	HostPCI1 string            `json:"hostpci1,omitempty"`
	HostPCI2 string            `json:"hostpci2,omitempty"`
	HostPCI3 string            `json:"hostpci3,omitempty"`
	HostPCI4 string            `json:"hostpci4,omitempty"`
	HostPCI5 string            `json:"hostpci5,omitempty"`
	HostPCI6 string            `json:"hostpci6,omitempty"`
	HostPCI7 string            `json:"hostpci7,omitempty"`
	HostPCI8 string            `json:"hostpci8,omitempty"`
	HostPCI9 string            `json:"hostpci9,omitempty"`

	// Serial devices
	Serials map[string]string `json:"-"`
	Serial0 string            `json:"serial0,omitempty"`
	Serial1 string            `json:"serial1,omitempty"`
	Serial2 string            `json:"serial2,omitempty"`
	Serial3 string            `json:"serial3,omitempty"`

	// USB devices
	USBs  map[string]string `json:"-"`
	USB0  string            `json:"usb0,omitempty"`
	USB1  string            `json:"usb1,omitempty"`
	USB2  string            `json:"usb2,omitempty"`
	USB3  string            `json:"usb3,omitempty"`
	USB4  string            `json:"usb4,omitempty"`
	USB5  string            `json:"usb5,omitempty"`
	USB6  string            `json:"usb6,omitempty"`
	USB7  string            `json:"usb7,omitempty"`
	USB8  string            `json:"usb8,omitempty"`
	USB9  string            `json:"usb9,omitempty"`
	USB10 string            `json:"usb10,omitempty"`
	USB11 string            `json:"usb11,omitempty"`
	USB12 string            `json:"usb12,omitempty"`
	USB13 string            `json:"usb13,omitempty"`
	USB14 string            `json:"usb14,omitempty"`

	// Parallel devices
	Parallels map[string]string `json:"-"`
	Parallel0 string            `json:"parallel0,omitempty"`
	Parallel1 string            `json:"parallel1,omitempty"`
	Parallel2 string            `json:"parallel2,omitempty"`

	// Cloud-init
	CIType       string `json:"citype,omitempty"`
	CIUser       string `json:"ciuser,omitempty"`
	CIPassword   string `json:"cipassword,omitempty"`
	Nameserver   string `json:"nameserver,omitempty"`
	Searchdomain string `json:"searchdomain,omitempty"`
	SSHKeys      string `json:"sshkeys,omitempty"`
	CICustom     string `json:"cicustom,omitempty"`

	// Cloud-init interfaces
	IPConfigs map[string]string `json:"-"`
	IPConfig0 string            `json:"ipconfig0,omitempty"`
	IPConfig1 string            `json:"ipconfig1,omitempty"`
	IPConfig2 string            `json:"ipconfig2,omitempty"`
	IPConfig3 string            `json:"ipconfig3,omitempty"`
	IPConfig4 string            `json:"ipconfig4,omitempty"`
	IPConfig5 string            `json:"ipconfig5,omitempty"`
	IPConfig6 string            `json:"ipconfig6,omitempty"`
	IPConfig7 string            `json:"ipconfig7,omitempty"`
	IPConfig8 string            `json:"ipconfig8,omitempty"`
	IPConfig9 string            `json:"ipconfig9,omitempty"`
}

type VirtualMachineCloneOptions struct {
	NewID       int    `json:"newid"`
	BWLimit     uint64 `json:"bwlimit,omitempty"`
	Description string `json:"description,omitempty"`
	Format      string `json:"format,omitempty"`
	Full        uint8  `json:"full,omitempty"`
	Name        string `json:"name,omitempty"`
	Pool        string `json:"pool,omitempty"`
	SnapName    string `json:"snapname,omitempty"`
	Storage     string `json:"storage,omitempty"`
	Target      string `json:"target,omitempty"`
}

type VirtualMachineMoveDiskOptions struct {
	Disk         string `json:"disk"`
	BWLimit      uint64 `json:"bwlimit,omitempty"`
	Delete       uint8  `json:"delete,omitempty"`
	Digest       string `json:"digest,omitempty"`
	Format       string `json:"format,omitempty"`
	Storage      string `json:"storage,omitempty"`
	TargetDigest string `json:"target-digest,omitempty"`
	TargetDisk   string `json:"target-disk,omitempty"`
	TargetVMID   int    `json:"target-vmid,omitempty"`
}

func (vmc *VirtualMachineConfig) mergeIndexedDevices(prefix string) map[string]string {
	deviceMap := make(map[string]string, 0)
	t := reflect.TypeOf(*vmc)
	v := reflect.ValueOf(*vmc)
	count := v.NumField()

	for i := 0; i < count; i++ {
		fn := t.Field(i).Name
		fv := v.Field(i).String()
		if "" == fv {
			continue
		}
		if strings.HasPrefix(fn, prefix) {
			// Ignore non-numeric suffixes like SCSIHW
			suffix := strings.TrimPrefix(fn, prefix)
			if _, err := strconv.Atoi(suffix); err != nil {
				continue
			}
			deviceMap[strings.ToLower(fn)] = fv
		}
	}

	return deviceMap
}

func (vmc *VirtualMachineConfig) MergeIDEs() map[string]string {
	if nil == vmc.IDEs {
		vmc.IDEs = vmc.mergeIndexedDevices("IDE")
	}
	return vmc.IDEs
}

func (vmc *VirtualMachineConfig) MergeSCSIs() map[string]string {
	if nil == vmc.SCSIs {
		vmc.SCSIs = vmc.mergeIndexedDevices("SCSI")
	}
	return vmc.SCSIs
}

func (vmc *VirtualMachineConfig) MergeSATAs() map[string]string {
	if nil == vmc.SATAs {
		vmc.SATAs = vmc.mergeIndexedDevices("SATA")
	}
	return vmc.SATAs
}

func (vmc *VirtualMachineConfig) MergeNets() map[string]string {
	if nil == vmc.Nets {
		vmc.Nets = vmc.mergeIndexedDevices("Net")
	}
	return vmc.Nets
}

func (vmc *VirtualMachineConfig) MergeVirtIOs() map[string]string {
	if nil == vmc.VirtIOs {
		vmc.VirtIOs = vmc.mergeIndexedDevices("VirtIO")
	}
	return vmc.VirtIOs
}

func (vmc *VirtualMachineConfig) MergeUnuseds() map[string]string {
	if nil == vmc.Unuseds {
		vmc.Unuseds = vmc.mergeIndexedDevices("Unused")
	}
	return vmc.Unuseds
}

func (vmc *VirtualMachineConfig) MergeSerials() map[string]string {
	if nil == vmc.Serials {
		vmc.Serials = vmc.mergeIndexedDevices("Serial")
	}
	return vmc.Serials
}

func (vmc *VirtualMachineConfig) MergeUSBs() map[string]string {
	if nil == vmc.USBs {
		vmc.HostPCIs = vmc.mergeIndexedDevices("USB")
	}
	return vmc.USBs
}

func (vmc *VirtualMachineConfig) MergeHostPCIs() map[string]string {
	if nil == vmc.HostPCIs {
		vmc.HostPCIs = vmc.mergeIndexedDevices("HostPCI")
	}
	return vmc.HostPCIs
}

func (vmc *VirtualMachineConfig) MergeNumas() map[string]string {
	if nil == vmc.Numas {
		vmc.HostPCIs = vmc.mergeIndexedDevices("Numa")
	}
	return vmc.Numas
}

func (vmc *VirtualMachineConfig) MergeParallels() map[string]string {
	if nil == vmc.Parallels {
		vmc.Parallels = vmc.mergeIndexedDevices("Parallel")
	}
	return vmc.Parallels
}

func (vmc *VirtualMachineConfig) MergeIPConfigs() map[string]string {
	if nil == vmc.IPConfigs {
		vmc.IPConfigs = vmc.mergeIndexedDevices("IPConfig")
	}
	return vmc.IPConfigs
}
