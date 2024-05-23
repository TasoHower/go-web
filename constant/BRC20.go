package constant

const (
	FIRST_BRC20_Block = 779832
	FirstJubilee      = 824544
)

// brc20 protocal
const (
	BRC20_P = "brc-20"
)

// brc20 op
const (
	BRC20_OP_DEPLOY   = "deploy"
	BRC20_OP_MINT     = "mint"
	BRC20_OP_TRANSFER = "transfer"

	BRC20_OP_SEND = "send"
)

const (
	BRC20_OP_N_DEPLOY   = 0
	BRC20_OP_N_MINT     = 1
	BRC20_OP_N_TRANSFER = 2
)

const (
	BRC20_VALID_DEFAULT = 0
	BRC20_VALID_VALID   = 1
	BRC20_VALID_INVALID = 2
	BRC20_VALID_CURSED  = 3
	BRC20_VALID_WRONG   = 4
)
