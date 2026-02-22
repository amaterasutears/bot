package machine

import "context"

// Storage defines the interface for persisting state and data.
type Storage interface {
	// SetState stores the state identified by uid.
	SetState(ctx context.Context, uid int64, state string) error

	// State retrieves the state identified by uid.
	State(ctx context.Context, uid int64) (string, error)

	// DeleteState removes the state identified by uid.
	DeleteState(ctx context.Context, uid int64) error

	// SetData stores a data item identified by uid.
	SetData(ctx context.Context, uid int64, data Data) error

	// Data retrieves all data items identified by uid.
	Data(ctx context.Context, uid int64) (DataMap, error)

	// DeleteData removes all data items identified by uid.
	DeleteData(ctx context.Context, uid int64) error
}
