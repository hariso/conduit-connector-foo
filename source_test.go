package foo_test

import (
	"context"
	"testing"

	foo "github.com/hariso/conduit-connector-foo"
	"github.com/matryer/is"
)

func TestTeardownSource_NoOpen(t *testing.T) {
	is := is.New(t)
	con := foo.NewSource()
	err := con.Teardown(context.Background())
	is.NoErr(err)
}
