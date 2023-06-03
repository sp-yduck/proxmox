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

func (vm *VirtualMachine) CurrentStatus() (*VirtualMachineStatus, error) {
	path := qemuPath(vm.Node.Name())
	var status *VirtualMachineStatus
	if err := vm.Client.Get(fmt.Sprintf("%s/%d/status/current", path, vm.VMID), &status); err != nil {
		return nil, err
	}
	return status, nil
}

func (vm *VirtualMachine) Reboot() {

}
