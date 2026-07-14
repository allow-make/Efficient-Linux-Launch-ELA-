package ui

import (
"fmt"

"ela/pkg/i18n"
)

type Wizard struct {
step int
}

func NewWizard() *Wizard {
return &Wizard{step: 0}
}

func (w *Wizard) Start() {
w.showLicense()
}

func (w *Wizard) showLicense() {
fmt.Println("========================================")
fmt.Println(i18n.T("license_title"))
fmt.Println("========================================")
fmt.Println()
fmt.Println("1. " + i18n.T("license_accept"))
fmt.Println("2. " + i18n.T("license_reject"))
fmt.Println()
}

func (w *Wizard) ShowMainMenu() {
fmt.Println("========================================")
fmt.Println(i18n.T("wizard_title"))
fmt.Println("========================================")
fmt.Println()
fmt.Printf("1. %s\n", i18n.T("wizard_download"))
fmt.Printf("2. %s\n", i18n.T("wizard_existing"))
fmt.Printf("3. %s\n", i18n.T("wizard_usb"))
}
