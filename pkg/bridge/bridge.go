package bridge

/*
#cgo CFLAGS: -I${SRCDIR}/../../src
#cgo LDFLAGS: -L${SRCDIR}/../../build -lbridge -lsqlite3 -lGL -lGLU

#include <stdlib.h>
#include "bridge/bridge.h"
*/
import "C"
import (
"unsafe"
)

func ELRAInit(dbPath string) int {
cs := C.CString(dbPath)
defer C.free(unsafe.Pointer(cs))
return int(C.bridge_elra_init(cs))
}

func ELRACreateInstance(id, name, arch, mode, path string, pid int, autoInject int, status string) int {
cid := C.CString(id)
cname := C.CString(name)
carch := C.CString(arch)
cmode := C.CString(mode)
cpath := C.CString(path)
cstatus := C.CString(status)
defer func() {
C.free(unsafe.Pointer(cid))
C.free(unsafe.Pointer(cname))
C.free(unsafe.Pointer(carch))
C.free(unsafe.Pointer(cmode))
C.free(unsafe.Pointer(cpath))
C.free(unsafe.Pointer(cstatus))
}()
return int(C.bridge_elra_create_instance(cid, cname, carch, cmode, cpath, C.int(pid), C.int(autoInject), cstatus))
}

func ELRAGetInstance(id string) string {
cid := C.CString(id)
defer C.free(unsafe.Pointer(cid))
buf := make([]byte, 4096)
C.bridge_elra_get_instance(cid, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
return string(buf)
}

func ELRAListInstances() string {
buf := make([]byte, 8192)
C.bridge_elra_list_instances((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
return string(buf)
}

func ELRAUpdateStatus(id, status string) int {
cid := C.CString(id)
cstatus := C.CString(status)
defer func() {
C.free(unsafe.Pointer(cid))
C.free(unsafe.Pointer(cstatus))
}()
return int(C.bridge_elra_update_status(cid, cstatus))
}

func ELRADeleteInstance(id string) int {
cid := C.CString(id)
defer C.free(unsafe.Pointer(cid))
return int(C.bridge_elra_delete_instance(cid))
}

func ELRACreateRecord(id, targetID, mode, path string, autoInject int) int {
cid := C.CString(id)
ctarget := C.CString(targetID)
cmode := C.CString(mode)
cpath := C.CString(path)
defer func() {
C.free(unsafe.Pointer(cid))
C.free(unsafe.Pointer(ctarget))
C.free(unsafe.Pointer(cmode))
C.free(unsafe.Pointer(cpath))
}()
return int(C.bridge_elra_create_record(cid, ctarget, cmode, cpath, C.int(autoInject)))
}

func ELRAMergeRecord(instanceID, recordID string) int {
cinst := C.CString(instanceID)
crec := C.CString(recordID)
defer func() {
C.free(unsafe.Pointer(cinst))
C.free(unsafe.Pointer(crec))
}()
return int(C.bridge_elra_merge_record(cinst, crec))
}

func ELRAClose() {
C.bridge_elra_close()
}

func LSInit(channelID string) int {
cs := C.CString(channelID)
defer C.free(unsafe.Pointer(cs))
return int(C.bridge_ls_init(cs))
}

func LSCapture() string {
buf := make([]byte, 4096)
C.bridge_ls_capture((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
return string(buf)
}

func LSDetectType(data []byte) int {
if len(data) == 0 {
return 2
}
return int(C.bridge_ls_detect_type((*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data))))
}

func LSSend(channelID string, data []byte) int {
cid := C.CString(channelID)
defer C.free(unsafe.Pointer(cid))
if len(data) == 0 {
return -1
}
return int(C.bridge_ls_send(cid, (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data))))
}

func RendererInit(title string, width, height int) int {
ctitle := C.CString(title)
defer C.free(unsafe.Pointer(ctitle))
return int(C.bridge_renderer_init(ctitle, C.int(width), C.int(height)))
}

func RendererClear() int {
return int(C.bridge_renderer_clear())
}

func RendererDrawText(text string, x, y int, r, g, b byte) int {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(ctext))
return int(C.bridge_renderer_draw_text(ctext, C.int(x), C.int(y), C.uchar(r), C.uchar(g), C.uchar(b)))
}

func RendererPresent() int {
return int(C.bridge_renderer_present())
}

func RendererShouldClose() int {
return int(C.bridge_renderer_should_close())
}

func RendererPollEvents() {
C.bridge_renderer_poll_events()
}

func RendererDestroy() {
C.bridge_renderer_destroy()
}
