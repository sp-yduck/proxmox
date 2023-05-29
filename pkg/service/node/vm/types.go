package vm

import ()

type VirtualMachine struct {
	Client    Client
	Node      string
	Cpu       float32
	Cpus      int
	Disk      int
	DiskRead  int
	DiskWrite int
	MaxDisk   int
	MaxMem    int
	Mem       int
	Name      string
	NetIn     int
	NetOut    int
	Status    string
	Template  int
	UpTime    int
	VMID      int
}

type Client interface {
	Get(p string, v interface{}) error
	Post(p string, d interface{}, v interface{}) error
	Delete(p string, v interface{}) error
}

type Arch string
type OSType string
type ScsiHw string

const (
	X86_64  Arch = "x86_64"
	Aarch64 Arch = "aarch64"
)

const (
	Other OSType = "other"
	Wxp
	W2k
	W2k3
	W2k8
	Wvista
	Win7
	Win8
	Win10
	Win11
	// linux 2.4 kernel
	L24 OSType = "l24"
	// linux 2.6-6 kernel
	L26     OSType = "l26"
	Solaris OSType = "solaris"
)

const (
	Lsi              = "lsi"
	Lsi53c810        = "lsi53c810"
	VirtioScsiPci    = "virtio-scsi-pci"
	VirtioScsiSingle = "virtio-scsi-single"
	Megasas          = "megasas"
	Pvscsi           = "pvscsi"
)

type Ide struct {
	Ide0 string `json:"ide0,omitempty"`
	Ide1 string `json:"ide1,omitempty"`
	Ide2 string `json:"ide2,omitempty"`
}

type Net struct {
	Net0 string `json:"net0,omitempty"`
}

// wip n = 0~30
type Scsi struct {
	Scsi0 string `json:"scsi0,omitempty"`
}

type Serial struct {
	Serial0 string `json:"serial0,omitempty"`
}

// wip
// reference : https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu
type VirtualMachineCreateOptions struct {
	Acpi     bool   `json:"acpi,omitempty"`
	Affinity string `json:"affinity,omitempty"`
	Agent    string `json:"agent,omitempty"`
	Arch     Arch   `json:"arch,omitempty"`
	// boot order. ";" separated. : 'order=device1;device2;device3'
	Boot string `json:"boot,omitempty"`
	// cloud-init custom files
	CiCustom   string `json:"cicustom,omitempty"`
	CiPassword string `json:"cipassword,omitempty"`
	CiType     string `json:"citype,omitempty"`
	CiUser     string `json:"ciuser,omitempty"`
	// number of cores : 1 ~
	Cores int `json:"cores,omitempty"`
	// emulated cpu type
	Cpu string `json:"cpu,omitempty"`
	// limit of cpu usage : 0 ~
	// 0 indicated no limit
	CpuLimit    int    `json:",omitempty"`
	Description string `json:",omitempty"`

	// allow to overwrite existing VM
	Force bool `json:",omitempty"`
	// Use volume as IDE hard disk or CD-ROM (n is 0 to 3).
	// Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	// Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Ide
	// enable/disable KVM hardware virtualization
	Kvm bool `json:",omitempty"`
	// specifies the QEMU machine type
	Machine string `json:",omitempty"`
	// amount of RAM for the VM in MiB : 16 ~
	Memory int `json:"memory,omitempty"`
	// name for VM. Only used on the configuration web interface
	Name string `json:"name,omitempty"`
	// network device
	Net
	// specifies whether a VM will be started during system bootup
	OnBoot bool `json:",omitempty"`
	// quest OS
	OSType OSType `json:"ostype,omitempty"`
	// use volume as scsi hard disk or cd-rom
	// use special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume
	// use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Scsi
	// SCSI controller model
	ScsiHw ScsiHw `json:"scsihw,omitempty"`
	Serial
	// number of sockets
	Sockets int `json:"sockets,omitempty"`
	// cloud-init setup public ssh keys (one key per line, OpenSSH format)
	SSHKeys string `json:",omitempty"`
	// start VM after it was created successfully
	Start bool `json:"start,omitempty"`
	// tags of the VM. only for meta information
	Tags string `json:",omitempty"`
	// enable/disable template
	Template bool   `json:",omitempty"`
	VGA      string `json:"vga,omitempty"`
	// vm id
	VMID int `json:"vmid,omitempty"`
}

const (
	StatusVirtualMachineRunning = "running"
	StatusVirtualMachineStopped = "stopped"
	StatusVirtualMachinePaused  = "paused"

	UserDataISOFormat = "user-data-%d.iso"

	volumeIdentifier = "cidata"
	fileMode         = 0700
	blockSize        = 2048
)
