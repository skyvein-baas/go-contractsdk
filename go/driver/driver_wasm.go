// +build wasm

package driver

import (
	"git.skyvein.net/service/contractsdk/go/code"
	"git.skyvein.net/service/contractsdk/go/driver/wasm"
)

// Serve run contract in wasm environment
func Serve(contract code.Contract) {
	driver := wasm.New()
	driver.Serve(contract)
}
