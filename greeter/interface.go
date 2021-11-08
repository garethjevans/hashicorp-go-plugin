package greeter

// Greeter is the interface that we're exposing as a plugin.
type Greeter interface {
	Greet() string
}
