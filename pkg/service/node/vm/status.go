package vm

import "fmt"

type VirtualMachineStatus struct {
	// Ha             HAStatus
	Status         string
	VMID           int
	Agent          bool
	Cpus           int
	Lock           string
	MaxDisk        int
	MaxMem         int
	Name           string
	PID            int
	QMPStatus      string
	RunningMachine string
	RunningQEMU    string
	Spice          bool
	Tags           string
	UpTime         int
}

type RebootOption struct {
	TimeOut int `json:"timeout,omitempty"`
}

type ResumeOption struct {
	NoCheck  bool `json:"nocheck,omitempty"`
	SkipLock bool `json:"skiplock,omitempty"`
}

type StartOption struct {
	// override qemu's -cpu argument with the given string
	ForceCPU string `json:"force-cpu,omitempty"`
	// specifies the qemu machine type
	Machine string `json:"machine,omitempty"`
	// cluster node name
	MigratedFroom string `json:"migratedfrom,omitempty"`
	// cidr of (sub) network that is used for migration
	MigrationNetwork string `json:"migration_network,omitempty"`
	// migration traffic is ecrypted using an SSH tunnel by default.
	// On secure, completely private networks this can be disabled to increase performance.
	MigrationType string `json:"migration_type,omitempty"`
	SkipLock      bool   `json:"skiplock,omitempty"`
	// some command save/restore state from this location
	StateURI string `json:"stateuri,omitempty"`
	// Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage.
	// Providing the special value '1' will map each source storage to itself.
	TargetStoraage string `json:"targetstorage,omitempty"`
	TimeOut        int    `json:"timeout,omitempty"`
}

func (vm *VirtualMachine) CurrentStatus() (*VirtualMachineStatus, error) {
	path := qemuPath(vm.Node.Name())
	var status *VirtualMachineStatus
	if err := vm.Client.Get(fmt.Sprintf("%s/%d/status/current", path, vm.VMID), &status); err != nil {
		return nil, err
	}
	return status, nil
}

func (vm *VirtualMachine) Reboot(option RebootOption) error {
	path := qemuPath(vm.Node.Name())
	var upid string
	if err := vm.Client.Post(fmt.Sprintf("%s/%d/status/reboot", path, vm.VMID), option, &upid); err != nil {
		return err
	}
	return vm.Node.EnsureTaskDone(upid)
}

// func (vm *VirtualMachine) Reset() error {

// }

func (vm *VirtualMachine) Resume(option ResumeOption) error {
	path := qemuPath(vm.Node.Name())
	var upid string
	if err := vm.Client.Post(fmt.Sprintf("%s/%d/status/resume", path, vm.VMID), option, &upid); err != nil {
		return err
	}
	return vm.Node.EnsureTaskDone(upid)
}

// func (vm *VirtualMachine) Shutdown() error {

// }

func (vm *VirtualMachine) Start(option StartOption) error {
	path := qemuPath(vm.Node.Name())
	var upid string
	if err := vm.Client.Post(fmt.Sprintf("%s/%d/status/start", path, vm.VMID), option, &upid); err != nil {
		return err
	}
	return vm.Node.EnsureTaskDone(upid)
}
