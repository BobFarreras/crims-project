package repo_pb

import (
	"net/url"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

func TestBuildListQuery_Defaults(t *testing.T) {
	params := buildListQuery(ports.ListOptions{})
	if params.Get("page") != "1" {
		t.Fatalf("expected page 1, got %s", params.Get("page"))
	}
	if params.Get("perPage") != "20" {
		t.Fatalf("expected perPage 20, got %s", params.Get("perPage"))
	}
}

func TestBuildListQuery_Custom(t *testing.T) {
	params := buildListQuery(ports.ListOptions{Page: 2, PerPage: 50, Filter: "gameId='game-1'"})
	if params.Get("page") != "2" {
		t.Fatalf("expected page 2, got %s", params.Get("page"))
	}
	if params.Get("perPage") != "50" {
		t.Fatalf("expected perPage 50, got %s", params.Get("perPage"))
	}
	if params.Get("filter") != "gameId='game-1'" {
		t.Fatalf("expected filter, got %s", params.Get("filter"))
	}
}

func buildListQuery(options ports.ListOptions) url.Values {
	return buildListQueryOptions(options)
}
