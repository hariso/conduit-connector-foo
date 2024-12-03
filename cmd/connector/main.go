package main

import (
	foo "github.com/hariso/conduit-connector-foo"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {
	sdk.Serve(foo.Connector)
}
