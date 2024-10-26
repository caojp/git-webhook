package git

type GitOperation interface {
	Execute() error
}
