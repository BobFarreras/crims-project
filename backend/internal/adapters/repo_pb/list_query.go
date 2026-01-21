package repo_pb

import (
	"net/url"
	"strconv"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

func buildListQueryOptions(options ports.ListOptions) url.Values {
	normalized := options.Normalize()
	params := url.Values{}
	params.Set("page", strconv.Itoa(normalized.Page))
	params.Set("perPage", strconv.Itoa(normalized.PerPage))
	if normalized.Filter != "" {
		params.Set("filter", normalized.Filter)
	}
	return params
}
