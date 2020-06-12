// +build wasm

package vm

import (
	"github.com/skyvein-baas/go-contractsdk/code"
	"github.com/skyvein-baas/go-contractsdk/vm/wasm"
)

// Serve run contract in wasm environment
func Serve(contract code.Contract) {
	vm := wasm.New()
	vm.Serve(contract)
}
