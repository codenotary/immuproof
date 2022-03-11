package meta

import "github.com/adrg/xdg"

const AppName = "ImmuProof"
const LedgerHeaderName = "lc-ledger"

var DefaultConfigFolder = xdg.ConfigHome + "/" + AppName
var DefaultStateFolder = xdg.StateHome + "/" + AppName

var DefaultCNCPort = 443
var DefaultCNCHost = "localhost"
