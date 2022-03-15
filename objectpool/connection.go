package objectpool

type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}
