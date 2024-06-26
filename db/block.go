package db

type Status int

const (
	Processed Status = 0
	Verified  Status = 1 // each block's blobs will be verified by the post-verification process
	Skipped   Status = 2
)

type Block struct {
	Id            int64
	Root          string `gorm:"NOT NULL;index:idx_block_root;size:64"`
	ParentRoot    string
	StateRoot     string
	BodyRoot      string
	ProposerIndex uint64
	Signature     string
	Slot          uint64 `gorm:"NOT NULL;uniqueIndex:idx_block_slot"`
	ELBlockHeight uint64 // the eth1 block height
	BlobCount     int

	BundleName string `gorm:"NOT NULL"`
	Status     Status `gorm:"index:idx_block_status"`
}

func (*Block) TableName() string {
	return "block"
}
