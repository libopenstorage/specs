package volume

import (
	"time"
)

// VolumeID driver specific system wide unique volume identifier.
type VolumeID string

// Labels a name-value map
type Labels map[string]string

// VolumeSpec has the properties needed to create a volume.
type VolumeSpec struct {
	// Ephemeral storage
	Ephemeral bool

	// Thin provisioned volume size in bytes
	Size uint64

	// Format disk with this FileSystem
	Format Filesystem

	// BlockSize for file system
	BlockSize int

	// HA Level specifies the number of nodes that are
	// allowed to fail, and yet data is availabel.
	// A value of 0 implies that data is not erasure coded,
	// a failure of a node will lead to data loss.
	HALevel int

	// This disk's CoS
	Cos VolumeCos

	// Perform dedupe on this disk
	Dedupe bool

	// SnapshotInterval in minutes, set to 0 to disable Snapshots
	SnapshotInterval int

	// Volume configuration labels
	ConfigLabels Labels
}

// VolumeLocator is a structure that is attached to a volume and is used to
// carry opaque metadata.
type VolumeLocator struct {
	// Name user friendly identifier
	Name string

	// VolumeLabels set of name-value pairs that acts as search filters.
	VolumeLabels Labels
}

// Volume represents a live, created volume.
type Volume struct {
	// Self referential VolumeID
	ID VolumeID

	// User specified locator
	Locator VolumeLocator

	// Volume creation time
	Ctime time.Time

	// User specified disk configuration
	Spec *VolumeSpec

	// Volume usage
	Usage uint64

	// Last time a scan was run
	LastScan time.Time

	// Filesystem type if any
	Format Filesystem

	// Volume Status
	Status VolumeStatus

	// VolumeState
	State VolumeState

	// Attached On - for clustered storage arrays
	AttachedOn interface{}

	// Device path
	DevicePath string

	// Attach path
	AttachPath string

	// Set of nodes no which this Volume is erasure coded - for clustered storage arrays
	ReplicaSet []interface{}

	// Last Recorded Error
	ErrorNum int

	// Error String
	ErrorString string
}

// CreateOptions are passed in with a CreateRequest
type CreateOptions struct {
	// FailIfExists fail create request if a volume with matching Locator already exists.
	FailIfExists bool

	// CreateFromSnap will create a volume with specified SnapID
	CreateFromSnap SnapID
}

type VolumeDriver interface {
	// ProtoDriver is the basic driver interface needed for managing a volume's state.
	ProtoDriver

	// BlockDriver is an interface you would implement only if the storage solution is block based.
	// Filesystem volume providers would not need to implement this interface.
	BlockDriver

	// MountDriver would need to be implemented if the storage driver wants to do something
	// specific during a volume's attachment to a container namespace.
	MountDriver

	// Enumerator is an interface to list the provisioned volumes.  A storage provider would need
	// to implement this interface if something specific needs to be done during volume
	// enumeration.
	Enumerator
}

type ProtoDriver interface {
	// String description of this driver.
	String() string

	// Create a new Vol for the specific volume spec.
	// It returns a system generated VolumeID that uniquely identifies the volume
	// If CreateOptions.FailIfExists is set and a volume matching the locator
	// exists then this will fail with ErrEexist. Otherwise if a matching available
	// volume is found then it is returned instead of creating a new volume.
	Create(locator VolumeLocator, options *CreateOptions, spec *VolumeSpec) (VolumeID, error)

	// Delete volume.
	// Errors ErrEnoEnt, ErrVolHasSnaps may be returned.
	Delete(volumeID VolumeID) error

	// Snap specified volume. IO to the underlying volume should be quiesced before
	// calling this function.
	// Errors ErrEnoEnt may be returned
	Snapshot(volumeID VolumeID, lables Labels) (SnapID, error)

	// SnapDelete snap specified by snapID.
	// Errors ErrEnoEnt may be returned
	SnapDelete(snapID SnapID) error

	// Stats for specified volume.
	// Errors ErrEnoEnt may be returned
	Stats(volumeID VolumeID) (VolumeStats, error)

	// Alerts on this volume.
	// Errors ErrEnoEnt may be returned
	Alerts(volumeID VolumeID) (VolumeAlerts, error)

	// Shutdown and cleanup.
	Shutdown()
}

type Enumerator interface {
	// Inspect specified volumes.
	// Errors ErrEnoEnt may be returned.
	Inspect(volumeIDs []VolumeID) ([]Volume, error)

	// Enumerate volumes that map to the volumeLocator. Locator fields may be regexp.
	// If locator fields are left blank, this will return all volumes.
	Enumerate(locator VolumeLocator, labels Labels) ([]Volume, error)

	// SnapInspect provides details on this snapshot.
	// Errors ErrEnoEnt may be returned
	SnapInspect(snapID []SnapID) ([]VolumeSnap, error)

	// Enumerate snaps for specified volume
	// Count indicates the number of snaps populated.
	SnapEnumerate(locator VolumeLocator, labels Labels) ([]VolumeSnap, error)
}

type BlockDriver interface {
	// Attach map device to the host.
	// On success the devicePath specifies location where the device is exported
	// Errors ErrEnoEnt, ErrVolAttached may be returned.
	Attach(volumeID VolumeID) (string, error)

	// Format volume according to spec provided in Create
	// Errors ErrEnoEnt, ErrVolDetached may be returned.
	Format(volumeID VolumeID) error

	// Detach device from the host.
	// Errors ErrEnoEnt, ErrVolDetached may be returned.
	Detach(volumeID VolumeID) error
}

type MountDriver interface {
	// Mount volume at specified path
	// Errors ErrEnoEnt, ErrVolDetached may be returned.
	Mount(volumeID VolumeID, mountpath string) error

	// Unmount volume at specified path
	// Errors ErrEnoEnt, ErrVolDetached may be returned.
	Unmount(volumeID VolumeID, mountpath string) error
}
