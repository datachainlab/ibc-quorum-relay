package main

import (
	"log"

	"github.com/datachainlab/ibc-quorum-relay/module"
	"github.com/hyperledger-labs/yui-relayer/cmd"
)

func main() {
	if err := cmd.Execute(
		module.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
