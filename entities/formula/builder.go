package formula

// Builder is used to create formula objects.
type Builder struct {}

// NewBuilder returns a new object used to Build Formula objects.
func NewBuilder() *Builder {
	return &Builder{}
}

// Build creates a new Formula object.
func (b *Builder) Build() Arbitrary {
	return &Identity{}
}