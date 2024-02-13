package context

import (
	"context"
)

var (
	// Ctx represents the context for the load generator.
	Ctx context.Context
	// Cancel represents the cancel function for the load generator context.
	Cancel context.CancelFunc
)
