package container

var c *Container

func init() {
	c = New()
}

func New() *Container {
	c := new(Container)
	c.Reset()
	return c
}

func Get(key string) interface{} {
	return c.Get(key)
}

func Set(key string, value interface{}) {
	c.Set(key, value)
}

func Has(key string) bool {
	return c.Has(key)
}

func IsResolved(key string) bool {
	return c.IsResolved(key)
}

func IsUnresolved(key string) bool {
	return c.IsUnresolved(key)
}

func Remove(key string) {
	c.Remove(key)
}

func Reset(key string) {
	c.Reset()
}

type Container struct {
	unresolved map[string]func() (interface{}, error)
	resolved   map[string]interface{}
}

func (c *Container) Set(key string, value interface{}) {
	c.Remove(key)
	if c.isResolver(value) {
		c.unresolved[key] = value.(func() (interface{}, error))
		return
	}
	c.resolved[key] = value
}

func (c *Container) Has(key string) bool {
	if _, ok := c.resolved[key]; ok {
		return true
	}

	if _, ok := c.unresolved[key]; ok {
		return true
	}

	return false
}

func (c *Container) IsResolved(key string) bool {
	_, ok := c.resolved[key]
	return ok
}

func (c *Container) IsUnresolved(key string) bool {
	_, ok := c.unresolved[key]
	return ok
}

func (c *Container) Get(key string) interface{} {
	if val, ok := c.resolved[key]; ok {
		return val
	}

	if val, ok := c.unresolved[key]; ok {
		l, ok := interface{}(val).(func() (interface{}, error))
		if ok {
			i, err := l()
			if err != nil {
				panic(err)
				// return nil // error happening
			}
			c.resolved[key] = i
			delete(c.unresolved, key)
			return i
		}
		return val
	}

	return nil
}

func (c *Container) isResolver(value interface{}) bool {
	_, ok := interface{}(value).(func() (interface{}, error))
	return ok
}

func (c *Container) Remove(key string) {
	delete(c.resolved, key)
	delete(c.unresolved, key)
}

func (c *Container) Reset() {
	c.resolved = make(map[string]interface{})
	c.unresolved = make(map[string]func()(interface{}, error))
}