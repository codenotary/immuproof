package meta

import (
	"github.com/adrg/xdg"
	"time"
)

const AppName = "immuproof"
const LedgerHeaderName = "lc-ledger"

var DefaultConfigFolder = xdg.ConfigHome + "/" + AppName
var DefaultStateFolder = xdg.StateHome + "/" + AppName
var StateMapFileName = xdg.StateHome + "/" + AppName + "/" + "state-map.json"

var DefaultCNCPort = 443
var DefaultCNCHost = "localhost"

var DefaultAuditInterval, _ = time.ParseDuration("1h")
