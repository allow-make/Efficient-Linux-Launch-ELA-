package version

import (
"fmt"
"log"
)

func Check(rootDir string) {
if err := Load(rootDir); err != nil {
log.Printf("Version check failed, using default: %s", DefaultFull)
}
fmt.Printf("Version: %s\n", String())
}
