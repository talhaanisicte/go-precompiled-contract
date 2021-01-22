package constants

import (
	"os"
	"path/filepath"
	"time"
)

// Home -
var Home string = os.Getenv("HOME")

// PluginSecret -
var PluginSecret int

// CallbackInterval -
const CallbackInterval = 100 * time.Millisecond

// Paths
var (
	JavaNativePath = "/usr/java/packages/lib/"
	TempDir        = "/tmp/ipfsTemp/"
	GoRoot         = filepath.Join(Home, "go/src/")
)

// JavaClassName -
const JavaClassName = "io_icte_go_GoJni"

// Filenames
const (
	SolidityFileName = "solidityContract.sol"
	HTMLFileName     = "FormLayout.html"
)

// Wrapper MsgType
const (
	MsgTypeInitiator = byte(1)
	MsgTypeAcceptor  = byte(2)
)

// MsgTypes
const (
	MsgTypeStartPlugin = byte(0)
	MsgTypeGetOrderID  = byte(1)
	MsgTypeCrypto      = byte(250)
	MsgTypeCompile     = byte(251)
	MsgTypeIPFS        = byte(252)
	MsgTypeBlockchain  = byte(253)
	MsgTypeReceive     = byte(255)
)

// Crypto MsgSubtypes
const (
	MsgSubtypeGenerateBTC = byte(0)
)

// IPFS MsgSubtypes
const (
	MsgSubtypeGet         = byte(0)
	MsgSubtypeAdd         = byte(1)
	MsgSubtypeLoadLibrary = byte(2)
)

// MsgSubtypeCompile
const (
	MsgSubtypeCompileSol  = byte(0)
	MsgSubtypeCompileGo   = byte(1)
	MsgSubtypeCompileHTML = byte(2)
	MsgSubtypeDebugGo     = byte(3)
)
