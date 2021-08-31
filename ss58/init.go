package ss58

import "sync"

// https://github.com/paritytech/substrate/blob/master/ss58-registry.json
var (
	SSPrefix          = []byte{0x53, 0x53, 0x35, 0x38, 0x50, 0x52, 0x45}
	PolkadotPrefix    = []byte{0x00}
	KsmPrefix         = []byte{0x02}
	KatalPrefix       = []byte{0x04}
	PlasmPrefix       = []byte{0x05}
	BifrostPrefix     = []byte{0x06}
	EdgewarePrefix    = []byte{0x07}
	KaruraPrefix      = []byte{0x08}
	ReynoldsPrefix    = []byte{0x09}
	AcalaPrefix       = []byte{0x0a}
	LaminarPrefix     = []byte{0x0b}
	PolymathPrefix    = []byte{0x0c}
	SubstraTEEPrefix  = []byte{0x0d}
	KulupuPrefix      = []byte{0x10}
	DarkPrefix        = []byte{0x11}
	DarwiniaPrefix    = []byte{0x12}
	StafiPrefix       = []byte{0x14}
	DockTestNetPrefix = []byte{0x15}
	DockMainNetPrefix = []byte{0x16}
	ShiftNrgPrefix    = []byte{0x17}
	SubsocialPrefix   = []byte{0x1c}
	PhalaPrefix       = []byte{0x1e}
	RobonomicsPrefix  = []byte{0x20}
	DataHighwayPrefix = []byte{0x21}
	CentrifugePrefix  = []byte{0x24}
	MathMainPrefix    = []byte{0x27}
	MathTestPrefix    = []byte{0x28}
	SubstratePrefix   = []byte{0x2a}
	ChainXPrefix      = []byte{0x2c}
	SubGamePrefix     = []byte{0x1b}

	//testnet for pcx
	ChainXTestNetPrefix = SubstratePrefix
)

//
var Prefixes = map[string][]byte{
	"SSPrefix":          {0x53, 0x53, 0x35, 0x38, 0x50, 0x52, 0x45},
	"PolkadotPrefix":    {0x00},
	"KsmPrefix":         {0x02},
	"KatalPrefix":       {0x04},
	"PlasmPrefix":       {0x05},
	"BifrostPrefix":     {0x06},
	"EdgewarePrefix":    {0x07},
	"KaruraPrefix":      {0x08},
	"ReynoldsPrefix":    {0x09},
	"AcalaPrefix":       {0x0a},
	"LaminarPrefix":     {0x0b},
	"PolymathPrefix":    {0x0c},
	"SubstraTEEPrefix":  {0x0d},
	"KulupuPrefix":      {0x10},
	"DarkPrefix":        {0x11},
	"DarwiniaPrefix":    {0x12},
	"StafiPrefix":       {0x14},
	"DockTestNetPrefix": {0x15},
	"DockMainNetPrefix": {0x16},
	"ShiftNrgPrefix":    {0x17},
	"SubsocialPrefix":   {0x1c},
	"PhalaPrefix":       {0x1e},
	"RobonomicsPrefix":  {0x20},
	"DataHighwayPrefix": {0x21},
	"CentrifugePrefix":  {0x24},
	"MathMainPrefix":    {0x27},
	"MathTestPrefix":    {0x28},
	"SubstratePrefix":   {0x2a},
	"ChainXPrefix":      {0x2c},
	"SubGamePrefix":     {0x1b},

	//testnet for pcx
	"ChainXTestNetPrefix": {0x2a},
}
var once sync.Once

//note, prefix can only be added instead of updated.
func AddPrefix(name string, prefix []byte) bool {
	if name == "" || prefix == nil {
		return false
	}
	if Prefixes[name] != nil {
		return false
	}
	once.Do(func() {
		Prefixes[name] = prefix
	})
	return true
}
