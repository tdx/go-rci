package api

// Service ...
type Service interface {
	Run(hook string) ([]byte, error)
	Command(hook string) *Hook
}
