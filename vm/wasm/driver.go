// +build wasm

package wasm

import (
	"github.com/skyvein-baas/go-contractsdk/code"
	"github.com/skyvein-baas/go-contractsdk/exec"
)

type vm struct {
}

// New returns a wasm vm
func New() code.vm {
	return new(vm)
}

func (d *vm) Serve(contract code.Contract) {
	initDebugLog()
	exec.RunContract(0, contract, syscall)
}
