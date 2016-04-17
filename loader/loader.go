package loader

type Loader interface {
	Load() error
	setup() error
}
