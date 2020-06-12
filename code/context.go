package code

import (
	"math/big"

	pb "github.com/skyvein-baas/go-contractsdk/pb"
)

// 区块链上下问
type Context interface {
	// 调用参数
	Args() map[string][]byte
	//
	Caller() string
	// 获取合约创建者
	Initiator() string
	// 权限列表
	AuthRequire() []string

	// 操作数据
	PutObject(key []byte, value []byte) error
	GetObject(key []byte) ([]byte, error)
	DeleteObject(key []byte) error
	NewIterator(start, limit []byte) Iterator

	// 查询交易
	QueryTx(txid string) (*pb.Transaction, error)
	// 查询区块
	QueryBlock(blockid string) (*pb.Block, error)
	// 转账
	Transfer(to string, amount *big.Int) error
	// 调用合约
	Call(module, contract, method string, args map[string][]byte) (*Response, error)
}

// Iterator iterates over key/value pairs in key order
type Iterator interface {
	Key() []byte
	Value() []byte
	Next() bool
	Error() error
	// Iterator 必须在使用完毕后关闭
	Close()
}

// PrefixRange returns key range that satisfy the given prefix
func PrefixRange(prefix []byte) ([]byte, []byte) {
	var limit []byte
	for i := len(prefix) - 1; i >= 0; i-- {
		c := prefix[i]
		if c < 0xff {
			limit = make([]byte, i+1)
			copy(limit, prefix)
			limit[i] = c + 1
			break
		}
	}
	return prefix, limit
}
