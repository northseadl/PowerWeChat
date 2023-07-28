package PowerWeChat

type Config interface {
	Default()
	Validate() error
}
