// +build !cuda

package gorgonia

// CUDA indicates if this build is using CUDA
const CUDA = false

// ExternMetadata is used to hold metadata about external execution devices.
// In this build, it's an empty struct because the default build doesn't use external devices to execute the graph on
type ExternMetadata struct {
	b batchedBLAS
}

// HasFunc will always return false in this build
func (m ExternMetadata) HasFunc(name string) bool { return false }

// WorkAvailable returns a channel of empty struct, which is used to signal to the VM when there is work available. The VM will then call the DoWork method.
func (m *ExternMetadata) WorkAvailable() <-chan struct{} {
	if m.b != nil {
		return m.b.WorkAvailable()
	}

	return nil
}

// DoWork flushes any batched cgo calls. In this build it only flushes the batched BLAS calls.
func (m *ExternMetadata) DoWork() error {
	if m.b != nil {
		m.b.DoWork()
	}
	return nil
}

// Get gets a previously allocated memory slab of the provided size. If no memories of that size exist,
// it returns a NoOpError. The caller is then responsible for allocating the memory themselves.
func (m *ExternMetadata) Get(dev Device, size int64) (Memory, error) { return nil, noopError{} }

// Put puts a previously allocated memory slab of the provided size back into the pool. Currently this is a No-op in this build.
func (m *ExternMetadata) Put(dev Device, mem Memory, size int64) {}

// Cleanup cleans up the ancillary allocations made during the calling of batched external device function.
//
// The reason for this method is due to the fact that there is currently no way to free memory while the context is still running without causing
// some weirdness to the CUDA calls.
//
// This is a No-op in this build
func (m *ExternMetadata) Cleanup() {}

func (m *ExternMetadata) signal() {}
