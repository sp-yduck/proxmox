package vm

import (
	// "fmt"
	// "net/url"
	// "os"
	// "path/filepath"
	// "strings"
	"fmt"
	"time"
	// "github.com/diskfs/go-diskfs/filesystem/iso9660"
)

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
}

type Net struct {
	Net0 string `json:"net0,omitempty"`
}

// wip n = 0~30
type Scsi struct {
	Scsi0 string `json:"scsi0,omitempty"`
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
	// cloud-init setup public ssh keys (one key per line, OpenSSH format)
	SSHKeys string `json:",omitempty"`
	// start VM after it was created successfully
	Start bool `json:"start,omitempty"`
	// tags of the VM. only for meta information
	Tags string `json:",omitempty"`
	// enable/disable template
	Template bool `json:",omitempty"`
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

// DefaultAgentWaitInterval is the polling interval when waiting for agent exec commands
var DefaultAgentWaitInterval = 100 * time.Millisecond

func qemuPath(node string) string {
	return fmt.Sprintf("/nodes/%s/qemu", node)
}

// to do : options
func (vm *VirtualMachine) Delete() (string, error) {
	path := fmt.Sprintf("%s/%d", qemuPath(vm.Node), vm.VMID)
	var res string
	if err := vm.Client.Delete(path, res); err != nil {
		return "", err
	}
	return res, nil
}

// func (v *VirtualMachine) Ping() error {
// 	return v.client.Get(fmt.Sprintf("/nodes/%s/qemu/%d/status/current", v.Node, v.VMID), &v)
// }

// func (v *VirtualMachine) Config(options ...VirtualMachineOption) (*Task, error) {
// 	var upid UPID
// 	data := make(map[string]interface{})
// 	for _, opt := range options {
// 		data[opt.Name] = opt.Value
// 	}
// 	err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/config", v.Node, v.VMID), data, &upid)
// 	return NewTask(upid, v.client), err
// }

// func (v *VirtualMachine) TermProxy() (vnc *VNC, err error) {
// 	return vnc, v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/termproxy", v.Node, v.VMID), nil, &vnc)
// }

// func (v *VirtualMachine) HasTag(value string) bool {
// 	if v.VirtualMachineConfig.Tags == "" {
// 		return false
// 	}

// 	if v.VirtualMachineConfig.TagsSlice == nil {
// 		v.SplitTags()
// 	}

// 	for _, tag := range v.VirtualMachineConfig.TagsSlice {
// 		if tag == value {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (v *VirtualMachine) AddTag(value string) (*Task, error) {
// 	if v.HasTag(value) {
// 		return nil, nil //noop
// 	}

// 	if v.VirtualMachineConfig.TagsSlice == nil {
// 		v.SplitTags()
// 	}

// 	v.VirtualMachineConfig.TagsSlice = append(v.VirtualMachineConfig.TagsSlice, value)
// 	v.VirtualMachineConfig.Tags = strings.Join(v.VirtualMachineConfig.TagsSlice, ";")

// 	return v.Config(VirtualMachineOption{
// 		Name:  "tags",
// 		Value: v.VirtualMachineConfig.Tags,
// 	})
// }

// func (v *VirtualMachine) RemoveTag(value string) (*Task, error) {
// 	if !v.HasTag(value) {
// 		return nil, nil //noop
// 	}

// 	if v.VirtualMachineConfig.TagsSlice == nil {
// 		v.SplitTags()
// 	}

// 	for i, tag := range v.VirtualMachineConfig.TagsSlice {
// 		if tag == value {
// 			v.VirtualMachineConfig.TagsSlice = append(
// 				v.VirtualMachineConfig.TagsSlice[:i],
// 				v.VirtualMachineConfig.TagsSlice[i+1:]...,
// 			)
// 		}
// 	}

// 	v.VirtualMachineConfig.Tags = strings.Join(v.VirtualMachineConfig.TagsSlice, ";")
// 	return v.Config(VirtualMachineOption{
// 		Name:  "tags",
// 		Value: v.VirtualMachineConfig.Tags,
// 	})
// }

// func (v *VirtualMachine) SplitTags() {
// 	v.VirtualMachineConfig.TagsSlice = strings.Split(v.VirtualMachineConfig.Tags, ";")
// }

// // CloudInit takes four yaml docs as a string and make an ISO, upload it to the data store as <vmid>-user-data.iso and will
// // mount it as a CDROM to be used with nocloud cloud-init. This is NOT how proxmox expects a user to do cloud-init
// // which can be found here: https://pve.proxmox.com/wiki/Cloud-Init_Support#:~:text=and%20meta.-,Cloud%2DInit%20specific%20Options,-cicustom%3A%20%5Bmeta
// // If you want to use the proxmox implementation you'll need to use the cloudinit APIs https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu/{vmid}/cloudinit
// func (v *VirtualMachine) CloudInit(device, userdata, metadata, vendordata, networkconfig string) error {
// 	isoName := fmt.Sprintf(UserDataISOFormat, v.VMID)
// 	// create userdata iso file on the local fs
// 	iso, err := makeCloudInitISO(isoName, userdata, metadata, vendordata, networkconfig)
// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		// _ = os.Remove(iso.Name())
// 	}()

// 	node, err := v.client.Node(v.Node)
// 	if err != nil {
// 		return err
// 	}

// 	storage, err := node.StorageISO()
// 	if err != nil {
// 		return err
// 	}

// 	task, err := storage.Upload("iso", iso.Name())
// 	if err != nil {
// 		return err
// 	}

// 	// iso should only be < 5mb so wait for it and then mount it
// 	if err := task.WaitFor(5); err != nil {
// 		return err
// 	}

// 	task, err = v.AddTag(MakeTag("cloud-init"))
// 	if err != nil {
// 		return err
// 	}
// 	task.WaitFor(2)

// 	task, err = v.Config(VirtualMachineOption{
// 		Name:  device,
// 		Value: fmt.Sprintf("%s:iso/%s,media=cdrom", storage.Name, isoName),
// 	}, VirtualMachineOption{
// 		Name:  "boot",
// 		Value: fmt.Sprintf("%s;%s", v.VirtualMachineConfig.Boot, device),
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	return task.WaitFor(2)
// }

// func makeCloudInitISO(filename, userdata, metadata, vendordata, networkconfig string) (iso *os.File, err error) {
// 	iso, err = os.Create(filepath.Join(os.TempDir(), filename))
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer func() {
// 		err = iso.Close()
// 	}()

// 	fs, err := iso9660.Create(iso, 0, 0, blockSize, "")
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := fs.Mkdir("/"); err != nil {
// 		return nil, err
// 	}

// 	cifiles := map[string]string{
// 		"user-data": userdata,
// 		"meta-data": metadata,
// 	}
// 	if vendordata != "" {
// 		cifiles["vendor-data"] = vendordata
// 	}
// 	if networkconfig != "" {
// 		cifiles["network-config"] = networkconfig
// 	}

// 	for filename, content := range cifiles {
// 		rw, err := fs.OpenFile("/"+filename, os.O_CREATE|os.O_RDWR)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if _, err := rw.Write([]byte(content)); err != nil {
// 			return nil, err
// 		}
// 	}

// 	if err = fs.Finalize(iso9660.FinalizeOptions{
// 		RockRidge:        true,
// 		VolumeIdentifier: volumeIdentifier,
// 	}); err != nil {
// 		return nil, err
// 	}

// 	return
// }

// // VNCWebSocket copy/paste when calling to get the channel names right
// // send, recv, errors, closer, errors := vm.VNCWebSocket(vnc)
// // for this to work you need to first setup a serial terminal on your vm https://pve.proxmox.com/wiki/Serial_Terminal
// func (v *VirtualMachine) VNCWebSocket(vnc *VNC) (chan string, chan string, chan error, func() error, error) {
// 	p := fmt.Sprintf("/nodes/%s/qemu/%d/vncwebsocket?port=%d&vncticket=%s",
// 		v.Node, v.VMID, vnc.Port, url.QueryEscape(vnc.Ticket))

// 	return v.client.VNCWebSocket(p, vnc)
// }

// func (v *VirtualMachine) IsRunning() bool {
// 	return v.Status == StatusVirtualMachineRunning && v.QMPStatus == StatusVirtualMachineRunning
// }

// func (v *VirtualMachine) Start() (*Task, error) {
// 	var upid UPID
// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/status/start", v.Node, v.VMID), nil, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) IsStopped() bool {
// 	return v.Status == StatusVirtualMachineStopped && v.QMPStatus == StatusVirtualMachineStopped
// }

// func (v *VirtualMachine) Reset() (task *Task, err error) {
// 	var upid UPID
// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/status/reset", v.Node, v.VMID), nil, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) Shutdown() (task *Task, err error) {
// 	var upid UPID
// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/status/shutdown", v.Node, v.VMID), nil, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) Stop() (task *Task, err error) {
// 	var upid UPID
// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/status/stop", v.Node, v.VMID), nil, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) IsPaused() bool {
// 	return v.Status == StatusVirtualMachineRunning && v.QMPStatus == StatusVirtualMachinePaused
// }

// func (v *VirtualMachine) Pause() (task *Task, err error) {
// 	var upid UPID
// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/status/suspend", v.Node, v.VMID), nil, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) IsHibernated() bool {
// 	return v.Status == StatusVirtualMachineStopped && v.QMPStatus == StatusVirtualMachineStopped && v.Lock == "suspended"
// }

// func (v *VirtualMachine) Hibernate() (task *Task, err error) {
// 	var upid UPID
// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/status/suspend", v.Node, v.VMID), map[string]string{"todisk": "1"}, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) Resume() (task *Task, err error) {
// 	var upid UPID
// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/status/resume", v.Node, v.VMID), nil, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) Reboot() (task *Task, err error) {
// 	var upid UPID
// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/status/reboot", v.Node, v.VMID), nil, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) Delete() (task *Task, err error) {
// 	if v.HasTag(MakeTag("cloud-init")) {
// 		node, err := v.client.Node(v.Node)
// 		if err != nil {
// 			return nil, err
// 		}
// 		isoStorage, err := node.StorageISO()
// 		if err != nil {
// 			return nil, err
// 		}
// 		iso, err := isoStorage.ISO(fmt.Sprintf(UserDataISOFormat, v.VMID))
// 		if err != nil {
// 			return nil, err
// 		}
// 		task, err = iso.Delete()
// 		if err != nil {
// 			return nil, err
// 		}
// 		if err := task.WaitFor(5); err != nil {
// 			return nil, err
// 		}
// 	}
// 	var upid UPID
// 	if err := v.client.Delete(fmt.Sprintf("/nodes/%s/qemu/%d", v.Node, v.VMID), &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) Migrate(target, targetstorage string) (task *Task, err error) {
// 	var upid UPID
// 	params := map[string]string{
// 		"target": target,
// 	}
// 	if targetstorage != "" {
// 		params["targetstorage"] = targetstorage
// 	}
// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/migrate", v.Node, v.VMID), params, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) Clone(params *VirtualMachineCloneOptions) (newid int, task *Task, err error) {
// 	var upid UPID

// 	if params == nil {
// 		params = &VirtualMachineCloneOptions{}
// 	}

// 	if params.NewID == 0 {
// 		cluster, err := v.client.Cluster()
// 		if err != nil {
// 			return newid, nil, err
// 		}

// 		newid, err = cluster.NextID()
// 		if err != nil {
// 			return newid, nil, err
// 		}
// 		params.NewID = newid
// 	}

// 	if err := v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/clone", v.Node, v.VMID), params, &upid); err != nil {
// 		return newid, nil, err
// 	}

// 	return newid, NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) ResizeDisk(disk, size string) (err error) {
// 	err = v.client.Put(fmt.Sprintf("/nodes/%s/qemu/%d/resize", v.Node, v.VMID), map[string]string{
// 		"disk": disk,
// 		"size": size,
// 	}, nil)
// 	if err != nil {
// 		return
// 	}

// 	return
// }

// func (v *VirtualMachine) UnlinkDisk(diskID string, force bool) (task *Task, err error) {
// 	var upid UPID

// 	params := map[string]string{"idlist": diskID}
// 	if force {
// 		params["force"] = "1"
// 	}
// 	err = v.client.Put(fmt.Sprintf("/nodes/%s/qemu/%d/unlink", v.Node, v.VMID), params, &upid)
// 	if err != nil {
// 		return
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) MoveDisk(disk string, params *VirtualMachineMoveDiskOptions) (task *Task, err error) {
// 	var upid UPID

// 	if params == nil {
// 		params = &VirtualMachineMoveDiskOptions{}
// 	}

// 	if disk != "" {
// 		params.Disk = disk
// 	}

// 	err = v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/move_disk", v.Node, v.VMID), params, &upid)
// 	if err != nil {
// 		return
// 	}

// 	return NewTask(upid, v.client), nil
// }

// func (v *VirtualMachine) AgentGetNetworkIFaces() (iFaces []*AgentNetworkIface, err error) {
// 	node, err := v.client.Node(v.Node)
// 	if err != nil {
// 		return
// 	}

// 	networks := map[string][]*AgentNetworkIface{}
// 	err = v.client.Get(fmt.Sprintf("/nodes/%s/qemu/%d/agent/network-get-interfaces", node.Name, v.VMID), &networks)
// 	if err != nil {
// 		return
// 	}
// 	if result, ok := networks["result"]; ok {
// 		for _, iface := range result {
// 			if "lo" == iface.Name {
// 				continue
// 			}
// 			iFaces = append(iFaces, iface)
// 		}
// 	}

// 	return

// }

// func (v *VirtualMachine) WaitForAgent(seconds int) error {
// 	timeout := time.After(time.Duration(seconds) * time.Second)
// 	ticker := time.NewTicker(DefaultWaitInterval)
// 	defer ticker.Stop()

// 	for {
// 		_, err := v.AgentOsInfo()
// 		if err == nil {
// 			return nil
// 		}
// 		if !strings.Contains(err.Error(), "500 QEMU guest agent is not running") {
// 			return err
// 		}

// 		select {
// 		case <-timeout:
// 			return ErrTimeout
// 		case <-ticker.C:
// 		}
// 	}
// }

// func (v *VirtualMachine) AgentExec(command, inputData string) (pid int, err error) {
// 	err = v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/agent/exec", v.Node, v.VMID),
// 		map[string]string{
// 			"command":    command,
// 			"input-data": inputData,
// 		},
// 		&pid)

// 	return
// }

// func (v *VirtualMachine) AgentExecStatus(pid int) (status *AgentExecStatus, err error) {
// 	err = v.client.Get(fmt.Sprintf("/nodes/%s/qemu/%d/agent/exec-status?pid=%d", v.Node, v.VMID, pid), &status)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return
// }

// func (v *VirtualMachine) WaitForAgentExecExit(pid, seconds int) (*AgentExecStatus, error) {
// 	timeout := time.After(time.Duration(seconds) * time.Second)
// 	ticker := time.NewTicker(DefaultAgentWaitInterval)
// 	defer ticker.Stop()

// 	for {
// 		status, err := v.AgentExecStatus(pid)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if status.Exited {
// 			return status, nil
// 		}

// 		select {
// 		case <-timeout:
// 			return nil, ErrTimeout
// 		case <-ticker.C:
// 		}
// 	}
// }

// func (v *VirtualMachine) AgentOsInfo() (info *AgentOsInfo, err error) {
// 	results := map[string]*AgentOsInfo{}
// 	err = v.client.Get(fmt.Sprintf("/nodes/%s/qemu/%d/agent/get-osinfo", v.Node, v.VMID), &results)
// 	if err != nil {
// 		return
// 	}

// 	info, ok := results["result"]
// 	if !ok {
// 		err = fmt.Errorf("result is empty")
// 	}

// 	return
// }

// func (v *VirtualMachine) AgentSetUserPassword(password string, username string) error {
// 	return v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/agent/set-user-password", v.Node, v.VMID), map[string]string{"password": password, "username": username}, nil)
// }

// func (v *VirtualMachine) FirewallOptionGet() (firewallOption *FirewallVirtualMachineOption, err error) {
// 	err = v.client.Get(fmt.Sprintf("/nodes/%s/qemu/%d/firewall/options", v.Node, v.VMID), firewallOption)
// 	return
// }

// func (v *VirtualMachine) FirewallOptionSet(firewallOption *FirewallVirtualMachineOption) error {
// 	return v.client.Put(fmt.Sprintf("/nodes/%s/qemu/%d/firewall/options", v.Node, v.VMID), firewallOption, nil)
// }

// func (v *VirtualMachine) FirewallGetRules() (rules []*FirewallRule, err error) {
// 	err = v.client.Get(fmt.Sprintf("/nodes/%s/qemu/%d/firewall/rules", v.Node, v.VMID), &rules)
// 	return
// }

// func (v *VirtualMachine) FirewallRulesCreate(rule *FirewallRule) error {
// 	return v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/firewall/rules", v.Node, v.VMID), rule, nil)
// }

// func (v *VirtualMachine) FirewallRulesUpdate(rule *FirewallRule) error {
// 	return v.client.Put(fmt.Sprintf("/nodes/%s/qemu/%d/firewall/rules/%d", v.Node, v.VMID, rule.Pos), rule, nil)
// }

// func (v *VirtualMachine) FirewallRulesDelete(rulePos int) error {
// 	return v.client.Delete(fmt.Sprintf("/nodes/%s/qemu/%d/firewall/rules/%d", v.Node, v.VMID, rulePos), nil)
// }

// func (v *VirtualMachine) NewSnapshot(name string) (task *Task, err error) {
// 	var upid UPID
// 	if err = v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/snapshot", v.Node, v.VMID), map[string]string{"snapname": name}, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }
// func (v *VirtualMachine) Snapshots() (snapshots []*Snapshot, err error) {
// 	err = v.client.Get(fmt.Sprintf("/nodes/%s/qemu/%d/snapshot", v.Node, v.VMID), &snapshots)
// 	return
// }

// func (v *VirtualMachine) SnapshotRollback(name string) (task *Task, err error) {
// 	var upid UPID
// 	if err = v.client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/snapshot/%s/rollback", v.Node, v.VMID, name), nil, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, v.client), nil
// }
