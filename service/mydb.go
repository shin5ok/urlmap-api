package service

type MyDB interface {
	Put(params *[]string) error
	Get(query string) (*[]string, error)
	List() (*[]string, error)
}
