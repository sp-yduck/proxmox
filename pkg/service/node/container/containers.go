package container

import (
	"fmt"
)

type Container struct {
	Name    string
	Node    string
	Client  Client
	CPUs    int
	Status  string
	VMID    int
	Uptime  uint64
	MaxMem  uint64
	MaxDisk uint64
	MaxSwap uint64
}

type Client interface {
	Get(p string, v interface{}) error
	Post(p string, d interface{}, v interface{}) error
	Delete(p string, v interface{}) error
}

type VNC struct {
	Cert   string
	Port   int
	Ticket string
	UPID   string
	User   string
}

type Containers []*Container

type ContainerStatuses []*ContainerStatus
type ContainerStatus struct {
	Data string `json:",omitempty"`
}

func (c *Container) Start() (status string, err error) {
	return status, c.Client.Post(fmt.Sprintf("/nodes/%s/lxc/%d/status/start", c.Node, c.VMID), nil, &status)
}

func (c *Container) Stop() (status *ContainerStatus, err error) {
	return status, c.Client.Post(fmt.Sprintf("/nodes/%s/lxc/%d/status/stop", c.Node, c.VMID), nil, &status)
}

func (c *Container) Suspend() (status *ContainerStatus, err error) {
	return status, c.Client.Post(fmt.Sprintf("/nodes/%s/lxc/%d/status/suspend", c.Node, c.VMID), nil, &status)
}

func (c *Container) Reboot() (status *ContainerStatus, err error) {
	return status, c.Client.Post(fmt.Sprintf("/nodes/%s/lxc/%d/status/reboot", c.Node, c.VMID), nil, &status)
}

func (c *Container) Resume() (status *ContainerStatus, err error) {
	return status, c.Client.Post(fmt.Sprintf("/nodes/%s/lxc/%d/status/resume", c.Node, c.VMID), nil, &status)
}

func (c *Container) TermProxy() (vnc *VNC, err error) {
	return vnc, c.Client.Post(fmt.Sprintf("/nodes/%s/lxk/%d/termproxy", c.Node, c.VMID), nil, &vnc)
}

// func (c *Container) VNCWebSocket(vnc *VNC) (chan string, chan string, chan error, func() error, error) {
// 	p := fmt.Sprintf("/nodes/%s/lxc/%d/vncwebsocket?port=%d&vncticket=%s",
// 		c.Node, c.VMID, vnc.Port, url.QueryEscape(vnc.Ticket))

// 	return c.Client.VNCWebSocket(p, vnc)
// }
