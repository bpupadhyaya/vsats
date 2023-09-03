package block_pkg

type Block struct {
	Timestamp         int64  // time when the block was created
	PreviousBlockHash []byte // hash of the previous block
	MyBlockHash       []byte // has of the current block
	AllData           []byte // body (data or transactions)
}

type Blockchain struct {
	Blocks []*Block // A blockchain contains blocks
}
