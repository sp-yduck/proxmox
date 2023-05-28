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
	Delete(p string, v interface{}) error
	Upload(p string, d map[string]string, f *os.File, v interface{}) error
}
