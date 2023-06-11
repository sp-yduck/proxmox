package node

import (
	"os"
)

type Node struct {
	Client
	Cpu            float32
	Disk           int
	ID             string
	Level          string
	MaxCpu         int
	MaxMem         int
	Mem            int
	Node           string
	SSLFingerprint string
	Stauts         string
	Type           string
	UpTime         int
}

type Client interface {
	Get(p string, v interface{}) error
	Post(p string, d interface{}, v interface{}) error
	Put(p string, d interface{}, v interface{}) error
	Delete(p string, v interface{}) error
	Upload(p string, d map[string]string, f *os.File, v interface{}) error
	Req(m string, p string, d []byte, v interface{}) error
}

type Task struct {
	Client
	Endtime   int    `json:"endtime"`
	Id        string `json:"id,omitempty"`
	Node      string `json:"node"`
	PID       int    `json:"pid"`
	PStart    int    `json:"pstart"`
	StartTime int    `json:"starttime"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	UPID      string `json:"upid"`
	User      string `json:"user"`
}

type TaskStatus struct {
	Exitstatus string `json:"exitstatus"`
	Id         string `json:"id"`
	Node
}
