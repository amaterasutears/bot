// Package machine provides logic for managing state and data.
package machine

import "context"

// Machine manages the state and data using a storage backend.
type Machine struct {
	storage Storage
}

// New creates a new Machine instance with the provided storage.
func New(storage Storage) *Machine {
	return &Machine{
		storage: storage,
	}
}

// SetState updates the state by uid.
func (m *Machine) SetState(ctx context.Context, uid int64, state string) error {
	return m.storage.SetState(ctx, uid, state)
}

// State retrieves the current state by uid.
func (m *Machine) State(ctx context.Context, uid int64) (string, error) {
	return m.storage.State(ctx, uid)
}

// SetData updates the data by uid.
func (m *Machine) SetData(ctx context.Context, uid int64, data Data) error {
	return m.storage.SetData(ctx, uid, data)
}

// Data retrieves the data by uid.
func (m *Machine) Data(ctx context.Context, uid int64) (DataMap, error) {
	return m.storage.Data(ctx, uid)
}

// Reset clears both the state and data by uid.
func (m *Machine) Reset(ctx context.Context, uid int64) error {
	if err := m.storage.DeleteState(ctx, uid); err != nil {
		return err
	}

	if err := m.storage.DeleteData(ctx, uid); err != nil {
		return err
	}

	return nil
}
