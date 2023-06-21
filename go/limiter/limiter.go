package limiter

type Limiter interface {
	Take() bool
}
