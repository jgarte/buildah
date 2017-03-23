package buildah

import (
	"fmt"
)

// Delete removes the working container.  The Builder object should not be used
// after this method is called.
func (b *Builder) Delete() error {
	if err := b.store.DeleteContainer(b.ContainerID); err != nil {
		return fmt.Errorf("error deleting build container: %v", err)
	}
	b.MountPoint = ""
	b.Container = ""
	b.ContainerID = ""
	return nil
}
