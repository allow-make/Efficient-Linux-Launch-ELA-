package main

import (
"log"
"os"

"ela/pkg/settings"
"ela/pkg/ui"
"ela/pkg/version"
)

func main() {
// 加载版本信息
if err := version.Load("."); err != nil {
log.Printf("Version check: %v", err)
}

// 加载设置
sm, err := settings.NewSettingsManager("")
if err != nil {
log.Printf("Settings error: %v", err)
os.Exit(1)
}

// 启动 GUI
gui := ui.NewELAUI(sm)
gui.Start()
}
