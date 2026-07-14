package elra

import (
"encoding/json"
"ela/pkg/bridge"
)

type Instance struct {
ID         string `json:"id"`
Name       string `json:"name"`
Arch       string `json:"arch"`
Mode       string `json:"mode"`
Path       string `json:"path"`
PID        int    `json:"pid"`
AutoInject bool   `json:"auto_inject"`
Status     string `json:"status"`
}

type RecordPackage struct {
ID         string `json:"id"`
TargetID   string `json:"target_id"`
Mode       string `json:"mode"`
Path       string `json:"path"`
AutoInject bool   `json:"auto_inject"`
}

func Init(dbPath string) int {
return bridge.ELRAInit(dbPath)
}

func Close() {
bridge.ELRAClose()
}

func CreateInstance(inst *Instance) int {
auto := 0
if inst.AutoInject {
auto = 1
}
return bridge.ELRACreateInstance(
inst.ID, inst.Name, inst.Arch, inst.Mode, inst.Path,
inst.PID, auto, inst.Status,
)
}

func GetInstance(id string) (*Instance, error) {
data := bridge.ELRAGetInstance(id)
var inst Instance
if err := json.Unmarshal([]byte(data), &inst); err != nil {
return nil, err
}
return &inst, nil
}

func ListInstances() ([]Instance, error) {
data := bridge.ELRAListInstances()
var list []Instance
if err := json.Unmarshal([]byte(data), &list); err != nil {
return nil, err
}
return list, nil
}

func UpdateStatus(id, status string) int {
return bridge.ELRAUpdateStatus(id, status)
}

func DeleteInstance(id string) int {
return bridge.ELRADeleteInstance(id)
}

func CreateRecord(rec *RecordPackage) int {
auto := 0
if rec.AutoInject {
auto = 1
}
return bridge.ELRACreateRecord(rec.ID, rec.TargetID, rec.Mode, rec.Path, auto)
}

func MergeRecord(instanceID, recordID string) int {
return bridge.ELRAMergeRecord(instanceID, recordID)
}
