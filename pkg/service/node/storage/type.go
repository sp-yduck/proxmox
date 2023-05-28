package storage

import "os"

type Storage struct {
	Client
	Node         string
	Active       int
	Avail        int
	Content      string
	Enabled      int
	Shared       int
	Storage      string
	Total        int
	Type         string
	Used         int
	UsedFraction float64 `json:"used_fraction"`
}

type Content struct {
	Client
	Node      string
	Storage   string `json:",omitempty"`
	Content   string `json:",omitempty"`
	CTime     string `json:",omitempty"`
	Encrypted string
	Format    string
	Notes     string
	Parent    string
	Protected bool
	Size      int
	Used      int
	// to do : Verificateion
	VMID  int
	VolID string `josn:"volid,omitempty"`
}

type UPID string

type Client interface {
	Get(p string, v interface{}) error
	Post(p string, d interface{}, v interface{}) error
	Delete(p string, v interface{}) error
	Upload(p string, d map[string]string, f *os.File, v interface{}) error
}
