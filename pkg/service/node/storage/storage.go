package storage

import (
	"fmt"

	"github.com/sp-yduck/proxmox/pkg/api"
)

func (c *Storage) Delete() (string, error) {
	var res string
	if err := c.Client.Delete(fmt.Sprintf("/storage/%s", c.Storage), &res); err != nil {
		return "", err
	}
	return res, nil
}

func (c *Storage) Contents() ([]*Content, error) {
	path := contentPath(c.Node, c.Storage)
	var contents []*Content
	if err := c.Client.Get(path, &contents); err != nil {
		return nil, err
	}
	for _, content := range contents {
		content.Client = c.Client
		content.Storage = c.Storage
	}
	return contents, nil
}

func (c *Storage) GetContent(volID string) (*Content, error) {
	path := contentPath(c.Node, c.Storage)
	var contents []*Content
	if err := c.Client.Get(path, &contents); err != nil {
		return nil, err
	}
	for _, content := range contents {
		if content.VolID == volID {
			content.Client = c.Client
			content.Storage = c.Storage
			content.Node = c.Node
			return content, nil
		}
	}
	return nil, api.ErrNotFound
}

// to do : options
func (c *Storage) CreateContent(filename string, size, vmid int, format string) (*Content, error) {
	path := contentPath(c.Node, c.Storage)
	data := make(map[string]interface{})
	data["filename"] = filename
	data["size"] = size
	data["vmid"] = vmid
	switch format {
	case "raw", "qcow2", "subvlo":
		data["format"] = format
	}
	var res string
	if err := c.Client.Post(path, data, &res); err != nil {
		return nil, err
	}
	content, err := c.GetContent(fmt.Sprintf("%s:%s", c.Storage, filename))
	if err != nil {
		return nil, err
	}
	return content, nil
}

// func (s *Storage) Upload(content, file string) (*Task, error) {
// 	if !IsValidContent(content) {
// 		return nil, fmt.Errorf("only iso and vztmpl allowed")
// 	}

// 	stat, err := os.Stat(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if stat.IsDir() {
// 		return nil, fmt.Errorf("file is a directory %s", file)
// 	}

// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()

// 	var upid UPID
// 	if err := s.Client.Upload(fmt.Sprintf("/nodes/%s/storage/%s/upload", s.Node, s.Storage),
// 		map[string]string{"content": content}, f, &upid); err != nil {
// 		return nil, err
// 	}

// 	return NewTask(upid, s.Client), nil
// }

// to do : checksums
func (s *Storage) DownloadURL(content, filename, url string) (string, error) {
	if !IsValidContent(content) {
		return "", fmt.Errorf("only iso and vztmpl allowed")
	}
	var upid string
	if err := s.Client.Post(fmt.Sprintf("/nodes/%s/storage/%s/download-url", s.Node, s.Storage), map[string]string{
		"content":  content,
		"filename": filename,
		"url":      url,
	}, &upid); err != nil {
		return "", err
	}
	return upid, nil
}

// func (s *Storage) ISO(name string) (iso *ISO, err error) {
// 	err = s.Client.Get(fmt.Sprintf("/nodes/%s/storage/%s/content/%s:%s/%s", s.Node, s.Storage, s.Storage, "iso", name), &iso)
// 	if err != nil {
// 		return nil, err
// 	}

// 	iso.client = s.Client
// 	iso.Node = s.Node
// 	iso.Storage = s.Storage
// 	if iso.VolID == "" {
// 		iso.VolID = fmt.Sprintf("%s:iso/%s", iso.Storage, name)
// 	}
// 	return
// }

// func (s *Storage) VzTmpl(name string) (vztmpl *VzTmpl, err error) {
// 	err = s.Client.Get(fmt.Sprintf("/nodes/%s/storage/%s/content/%s:%s/%s", s.Node, s.Storage, s.Storage, "vztmpl", name), &vztmpl)
// 	if err != nil {
// 		return nil, err
// 	}

// 	vztmpl.client = s.Client
// 	vztmpl.Node = s.Node
// 	vztmpl.Storage = s.Storage
// 	if vztmpl.VolID == "" {
// 		vztmpl.VolID = fmt.Sprintf("%s:vztmpl/%s", vztmpl.Storage, name)
// 	}
// 	return
// }

// func (s *Storage) Backup(name string) (backup *Backup, err error) {
// 	err = s.Client.Get(fmt.Sprintf("/nodes/%s/storage/%s/content/%s:%s/%s", s.Node, s.Storage, s.Storage, "backup", name), &backup)
// 	if err != nil {
// 		return nil, err
// 	}

// 	backup.client = s.Client
// 	backup.Node = s.Storage
// 	backup.Storage = s.Storage
// 	return
// }

// func (v *VzTmpl) Delete() (*Task, error) {
// 	return deleteVolume(v.client, v.Node, v.Storage, v.VolID, v.Path, "vztmpl")
// }

// func (b *Backup) Delete() (*Task, error) {
// 	return deleteVolume(b.client, b.Node, b.Storage, b.VolID, b.Path, "backup")
// }

// func (i *ISO) Delete() (*Task, error) {
// 	return deleteVolume(i.client, i.Node, i.Storage, i.VolID, i.Path, "iso")
// }

// func deleteVolume(c *Client, n, s, v, p, t string) (*Task, error) {
// 	var upid UPID
// 	if v == "" && p == "" {
// 		return nil, fmt.Errorf("volid or path required for a delete")
// 	}

// 	if v == "" {
// 		// volid not returned in the volume endpoints, need to generate
// 		v = fmt.Sprintf("%s:%s/%s", s, t, filepath.Base(p))
// 	}

// 	err := c.Delete(fmt.Sprintf("/nodes/%s/storage/%s/content/%s", n, s, v), &upid)
// 	return NewTask(upid, c), err
// }
