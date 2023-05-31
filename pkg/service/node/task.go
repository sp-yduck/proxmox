package node

func (c *Task) IsStatusOK() bool {
	return c.Status == "OK"
}
