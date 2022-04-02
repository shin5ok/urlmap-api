package service

type MyDB interface {
	Put(string) error
	Get(string) (*[]string, error)
	List() (*[]string, error)
}
