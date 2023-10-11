# ibc-quorum-relay
A prover module of [yui-relayer](https://github.com/hyperledger-labs/yui-relayer). A corresponding Chain Module can be found in [ethereum-ibc-relay-chain](https://github.com/datachainlab/ethereum-ibc-relay-chain).

## Setup Relayer
Add this module to yui-relayer and activate it.

```golang
package main

import (
	"log"
	
	"github.com/hyperledger-labs/yui-relayer/cmd"
	quorum "github.com/datachainlab/ibc-quorum-relay/module"
)

func main() {
	if err := cmd.Execute(
		// counterparty.Module{}, //counter party
		quorum.Module{}, // Quorum Prover Module 
    ); err != nil {
		log.Fatal(err)
	}
}
```
