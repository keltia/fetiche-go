// This is the main monitoring daemon

package main

import (
	"os"
	"path/filepath"
)

var (
	// MyName is both package & CLI command name
	MyName = filepath.Base(os.Args[0])

	// Global flags
	fDebug   bool
	fVerbose bool
)

func main() {

}
