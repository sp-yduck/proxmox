package node

import (
	"fmt"

	"github.com/sp-yduck/proxmox/pkg/api"
	storageapi "github.com/sp-yduck/proxmox/pkg/service/node/storage"
	"github.com/sp-yduck/proxmox/pkg/service/node/vm"
	versionapi "github.com/sp-yduck/proxmox/pkg/service/version"
)

func qemuPath(node string) string {
	return fmt.Sprintf("/nodes/%s/qemu", node)
}

func (c *Node) Name() string {
	return c.Node
}

func (c *Node) VirtualMachines() ([]*vm.VirtualMachine, error) {
	path := qemuPath(c.Node)
	var vms []*vm.VirtualMachine
	if err := c.Client.Get(path, &vms); err != nil {
		return nil, err
	}
	for _, vm := range vms {
		vm.Client = c.Client
		vm.Node = c
	}
	return vms, nil
}

func (c *Node) VirtualMachine(vmid int) (*vm.VirtualMachine, error) {
	path := qemuPath(c.Node)
	var vms []*vm.VirtualMachine
	if err := c.Client.Get(path, &vms); err != nil {
		return nil, err
	}
	for _, vm := range vms {
		if vm.VMID == vmid {
			vm.Client = c.Client
			vm.Node = c
			return vm, nil
		}
	}
	return nil, api.ErrNotFound
}

func (c *Node) CreateVirtualMachine(vmid int, options vm.VirtualMachineCreateOptions) (*vm.VirtualMachine, error) {
	path := qemuPath(c.Node)
	options.VMID = vmid
	var upid string
	if err := c.Client.Post(path, options, &upid); err != nil {
		return nil, err
	}
	if err := c.EnsureTaskDone(upid); err != nil {
		return nil, err
	}
	return c.VirtualMachine(vmid)
}

// to do : options
func (c *Node) DeleteVirtualMachine(vmid int) (string, error) {
	path := fmt.Sprintf("%s/%d", qemuPath(c.Node), vmid)
	var res string
	if err := c.Client.Delete(path, res); err != nil {
		return "", err
	}
	return res, nil
}

func storagePath(nodeName string) string {
	return fmt.Sprintf("/nodes/%s/storage", nodeName)
}

func (c *Node) Storages() ([]*storageapi.Storage, error) {
	var storages []*storageapi.Storage
	if err := c.Client.Get(storagePath(c.Node), &storages); err != nil {
		return nil, err
	}
	for _, s := range storages {
		s.Client = c.Client
		s.Node = c.Node
	}
	return storages, nil
}

func (c *Node) Storage(name string) (*storageapi.Storage, error) {
	var storages []*storageapi.Storage
	if err := c.Client.Get(storagePath(c.Node), &storages); err != nil {
		return nil, err
	}
	for _, s := range storages {
		if s.Storage == name {
			s.Client = c.Client
			s.Node = c.Node
			return s, nil
		}
	}
	return nil, api.ErrNotFound
}

func (n *Node) Version() (version *versionapi.Version, err error) {
	return version, n.Client.Get("/nodes/%s/version", &version)
}

// func (n *Node) TermProxy() (vnc *VNC, err error) {
// 	return vnc, n.client.Post(fmt.Sprintf("/nodes/%s/termproxy", n.Name), nil, &vnc)
// }

// // VNCWebSocket send, recv, errors, closer, error
// func (n *Node) VNCWebSocket(vnc *VNC) (chan string, chan string, chan error, func() error, error) {
// 	p := fmt.Sprintf("/nodes/%s/vncwebsocket?port=%d&vncticket=%s",
// 		n.Name, vnc.Port, url.QueryEscape(vnc.Ticket))

// 	return n.client.VNCWebSocket(p, vnc)
// }

// func (n *Node) Containers() (c Containers, err error) {
// 	if err = n.client.Get(fmt.Sprintf("/nodes/%s/lxc", n.Name), &c); err != nil {
// 		return
// 	}

// 	for _, container := range c {
// 		container.client = n.client
// 		container.Node = n.Name
// 	}

// 	return
// }

// func (n *Node) Container(vmid int) (*Container, error) {
// 	var c Container
// 	if err := n.client.Get(fmt.Sprintf("/nodes/%s/lxc/%d/status/current", n.Name, vmid), &c); err != nil {
// 		return nil, err
// 	}
// 	c.client = n.client
// 	c.Node = n.Name

// 	return &c, nil
// }

// func (n *Node) Appliances() (appliances Appliances, err error) {
// 	err = n.client.Get(fmt.Sprintf("/nodes/%s/aplinfo", n.Name), &appliances)
// 	if err != nil {
// 		return appliances, err
// 	}

// 	for _, t := range appliances {
// 		t.client = n.client
// 		t.Node = n.Name
// 	}

// 	return appliances, nil
// }

// func (n *Node) DownloadAppliance(template, storage string) (ret string, err error) {
// 	return ret, n.client.Post(fmt.Sprintf("/nodes/%s/aplinfo", n.Name), map[string]string{
// 		"template": template,
// 		"storage":  storage,
// 	}, &ret)
// }

// func (n *Node) VzTmpls(storage string) (templates VzTmpls, err error) {
// 	return templates, n.client.Get(fmt.Sprintf("/nodes/%s/storage/%s/content?content=vztmpl", n.Name, storage), &templates)
// }

// func (n *Node) VzTmpl(template, storage string) (*VzTmpl, error) {
// 	templates, err := n.VzTmpls(storage)
// 	if err != nil {
// 		return nil, err
// 	}

// 	volid := fmt.Sprintf("%s:vztmpl/%s", storage, template)
// 	for _, t := range templates {
// 		if t.VolID == volid {
// 			return t, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("could not find vztmpl: %s", template)
// }

// func (n *Node) StorageISO() (*Storage, error) {
// 	return n.findStorageByContent("iso")
// }

// func (n *Node) StorageVZTmpl() (*Storage, error) {
// 	return n.findStorageByContent("vztmpl")
// }

// func (n *Node) StorageBackup() (*Storage, error) {
// 	return n.findStorageByContent("backup")
// }

// func (n *Node) StorageRootDir() (*Storage, error) {
// 	return n.findStorageByContent("rootdir")
// }

// func (n *Node) StorageImages() (*Storage, error) {
// 	return n.findStorageByContent("images")
// }

// // findStorageByContent takes iso/backup/vztmpl/rootdir/images and returns the storage that type of content should be on
// func (n *Node) findStorageByContent(content string) (storage *Storage, err error) {
// 	storages, err := n.Storages()
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, storage := range storages {
// 		if strings.Contains(storage.Content, content) {
// 			storage.Node = n.Name
// 			storage.client = n.client
// 			return storage, nil
// 		}
// 	}

// 	return nil, api.ErrNotFound
// }

// func (c *Node) Networks() ([]*NodeNetwork, error) {
// 	var networks []*NodeNetwork
// 	if err := c.Client.Get(fmt.Sprintf("/nodes/%s/network", c.Node), &networks); err != nil {
// 		return nil, err
// 	}
// 	for _, n := range networks {
// 		n.client = c.Client
// 		n.Node = c.Node
// 	}
// 	return networks, nil
// }

// func (n *Node) Network(iface string) (network *NodeNetwork, err error) {
// 	err = n.client.Get(fmt.Sprintf("/nodes/%s/network/%s", n.Name, iface), &network)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if nil != network {
// 		network.client = n.client
// 		network.Node = n.Name
// 		network.NodeAPI = n
// 		network.Iface = iface
// 	}

// 	return network, nil
// }

// func (n *Node) NewNetwork(network *NodeNetwork) (task *Task, err error) {
// 	err = n.client.Post(fmt.Sprintf("/nodes/%s/network", n.Name), network, network)
// 	if nil != err {
// 		return
// 	}

// 	network.client = n.client
// 	network.Node = n.Name
// 	network.NodeAPI = n
// 	return n.NetworkReload()
// }

// func (n *Node) NetworkReload() (*Task, error) {
// 	var upid UPID
// 	err := n.client.Put(fmt.Sprintf("/nodes/%s/network", n.Name), nil, &upid)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, n.client), nil
// }

// func (n *Node) FirewallOptionGet() (firewallOption *FirewallNodeOption, err error) {
// 	err = n.client.Get(fmt.Sprintf("/nodes/%s/firewall/options", n.Name), firewallOption)
// 	return
// }

// func (n *Node) FirewallOptionSet(firewallOption *FirewallNodeOption) error {
// 	return n.client.Put(fmt.Sprintf("/nodes/%s/firewall/options", n.Name), firewallOption, nil)
// }

// func (n *Node) FirewallGetRules() (rules []*FirewallRule, err error) {
// 	err = n.client.Get(fmt.Sprintf("/nodes/%s/firewall/rules", n.Name), &rules)
// 	return
// }

// func (n *Node) FirewallRulesCreate(rule *FirewallRule) error {
// 	return n.client.Post(fmt.Sprintf("/nodes/%s/firewall/rules", n.Name), rule, nil)
// }

// func (n *Node) FirewallRulesUpdate(rule *FirewallRule) error {
// 	return n.client.Put(fmt.Sprintf("/nodes/%s/firewall/rules/%d", n.Name, rule.Pos), rule, nil)
// }

// func (n *Node) FirewallRulesDelete(rulePos int) error {
// 	return n.client.Delete(fmt.Sprintf("/nodes/%s/firewall/rules/%d", n.Name, rulePos), nil)
// }
