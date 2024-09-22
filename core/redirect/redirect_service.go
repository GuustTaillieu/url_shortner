package core

type RedirectService interface {
	Find(code string) (*RedirectModel, error)
	Store(redirect *RedirectModel) error
}


