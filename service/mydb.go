package service

type MyDB interface {
	Put(params []interface{}) error
	Get(query string) ([]string, error)
	List() ([]string, error)
}

var RedirectTableName = "redirects"

var RedirectTableColumn = []string{"user", "redirect_path", "org", "host", "comment", "active", "begin_at", "end_at", "created_at", "update_at", "deleted_at"}
