// +build wasm

package wasm

import (
	"git.skyvein.net/service/contractsdk/go/code"
	"git.skyvein.net/service/contractsdk/go/exec"
)

type driver struct {
}

// New returns a wasm driver
func New() code.Driver {
	return new(driver)
}

func (d *driver) Serve(contract code.Contract) {
	initDebugLog()
	exec.RunContract(0, contract, syscall)
}
