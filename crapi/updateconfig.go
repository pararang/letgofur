// created by: Muhammad Ikhsan
// created on: 2025-04-30

package crapi

// UpdateConfig updates the details of an application on the Caprover instance.
// This is a public wrapper around the internal updateAppDetails method from original crapi code.
func (c *Caprover) UpdateConfig(data UpdateAppRequest) error {
	return c.updateAppDetails(data)
}
