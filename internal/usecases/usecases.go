package usecases

// Core is the core struct of the usecases layer.
// It contains all the usecases methods.
type Core struct {
}

// Setup is the setup struct of the usecases layer.
// It contains all the information necessary to initialize the usecases dependencies.
type Setup struct {
}

// New creates a brand new usecases Core with all dependencies initialized.
// It receives a pointer to a Setup struct with all the information necessary to initialize the usecases dependencies.
//
// It returns a pointer to the new instantiated Core.
func New(c *Setup) (*Core, error) {
	core := &Core{}
	return core, nil
}

// Cleanup cleans up all the usecases dependencies.
func (c *Core) Cleanup() {
}
