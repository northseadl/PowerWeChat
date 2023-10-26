package weconfig

type Config interface {
	Default()
	Validate() error
}
