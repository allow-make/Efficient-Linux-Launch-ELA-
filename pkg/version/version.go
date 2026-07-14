package version

import (
"fmt"
"os"
"path/filepath"
"strings"
"sync"
)

const DefaultFull = "Efficient Linux Launcher"

var (
once sync.Once
full string
)

func Load(rootDir string) error {
var err error
once.Do(func() {
path := filepath.Join(rootDir, "version")
data, readErr := os.ReadFile(path)
if readErr != nil {
full = DefaultFull
_ = os.WriteFile(path, []byte(full+"\n"), 0644)
return
}
content := strings.TrimSpace(string(data))
if content == "" || !strings.HasPrefix(content, "Efficient Linux Launcher") {
full = DefaultFull
_ = os.WriteFile(path, []byte(full+"\n"), 0644)
return
}
full = content
})
return err
}

func String() string {
if full == "" {
return DefaultFull
}
return full
}

func Short() string {
s := String()
if s == DefaultFull {
return "ELA"
}
return "ELA " + strings.TrimPrefix(s, "Efficient Linux Launcher ")
}

func InstanceName(distro string) string {
return fmt.Sprintf("%s - %s", String(), distro)
}
