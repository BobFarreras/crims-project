package ports

// ListOptions defineix paginacio i filtre.
type ListOptions struct {
	Page    int
	PerPage int
	Filter  string
}

func (o ListOptions) Normalize() ListOptions {
	page := o.Page
	perPage := o.PerPage
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 20
	}
	return ListOptions{Page: page, PerPage: perPage, Filter: o.Filter}
}
