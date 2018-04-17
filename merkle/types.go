package merkle

type HashOp uint8

const (
	RIPEMD160 = HashOp(iota)
	SHA224
	SHA256
	SHA284
	SHA512
	SHA3_224
	SHA3_256
	SHA3_384
	SHA3_512
	SHA256_X2
)

type ProofType uint8

type Node struct {
	Prefix []byte
	Suffix []byte
	Op     HashOp
}

type ExistsData struct {
	Prefix []byte
	Suffix []byte
	Op     HashOp
}

type ExistsProof struct {
	PrefixLeaf  []byte
	PrefixInner []byte
	Data        ExistsData
	Nodes       []Node
	RootHash    []byte
}

type AbsentProof struct {
}

type RangeProof struct {
}

type KeyProof interface {
	Verify(leaf []byte) error
	Root() []byte
}

type SubTreeProof struct {
}

type MultiProof struct {
	KeyProof  KeyProof
	SubProofs []ExistsProof
}