#include "bridge.h"
#include "../elra/elra.h"
#include "../ls/ls.h"
#include "../renderer/renderer.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// ============================================================
// ELRA Registry Bridge
// ============================================================

int bridge_elra_init(const char* db_path) {
    return elra_init(db_path);
}

int bridge_elra_create_instance(const char* id, const char* name, const char* arch,
                                 const char* mode, const char* path, int pid,
                                 int auto_inject, const char* status) {
    ELRAInstance inst;
    strncpy(inst.id, id, sizeof(inst.id) - 1);
    strncpy(inst.name, name, sizeof(inst.name) - 1);
    strncpy(inst.arch, arch, sizeof(inst.arch) - 1);
    strncpy(inst.mode, mode, sizeof(inst.mode) - 1);
    strncpy(inst.path, path, sizeof(inst.path) - 1);
    inst.pid = pid;
    inst.auto_inject = auto_inject;
    strncpy(inst.status, status, sizeof(inst.status) - 1);
    return elra_create_instance(&inst);
}

int bridge_elra_get_instance(const char* id, char* out, size_t out_size) {
    ELRAInstance inst;
    if (elra_get_instance(id, &inst) != 0) {
        return -1;
    }
    snprintf(out, out_size,
        "{\"id\":\"%s\",\"name\":\"%s\",\"arch\":\"%s\",\"mode\":\"%s\","
        "\"path\":\"%s\",\"pid\":%d,\"auto_inject\":%d,\"status\":\"%s\"}",
        inst.id, inst.name, inst.arch, inst.mode,
        inst.path, inst.pid, inst.auto_inject, inst.status);
    return 0;
}

int bridge_elra_list_instances(char* out, size_t out_size) {
    ELRAInstance instances[64];
    size_t count = 64;
    if (elra_list_instances(instances, &count) != 0) {
        return -1;
    }
    out[0] = '[';
    size_t pos = 1;
    for (size_t i = 0; i < count && pos < out_size - 10; i++) {
        if (i > 0) {
            out[pos++] = ',';
        }
        pos += snprintf(out + pos, out_size - pos,
            "{\"id\":\"%s\",\"name\":\"%s\",\"arch\":\"%s\",\"mode\":\"%s\",\"status\":\"%s\"}",
            instances[i].id, instances[i].name, instances[i].arch,
            instances[i].mode, instances[i].status);
    }
    out[pos++] = ']';
    out[pos] = '\0';
    return 0;
}

int bridge_elra_update_status(const char* id, const char* status) {
    return elra_update_status(id, status);
}

int bridge_elra_delete_instance(const char* id) {
    return elra_delete_instance(id);
}

int bridge_elra_create_record(const char* id, const char* target_id,
                               const char* mode, const char* path, int auto_inject) {
    ELRARecord rec;
    strncpy(rec.id, id, sizeof(rec.id) - 1);
    strncpy(rec.target_id, target_id, sizeof(rec.target_id) - 1);
    strncpy(rec.mode, mode, sizeof(rec.mode) - 1);
    strncpy(rec.path, path, sizeof(rec.path) - 1);
    rec.auto_inject = auto_inject;
    return elra_create_record(&rec);
}

int bridge_elra_merge_record(const char* instance_id, const char* record_id) {
    return elra_merge_record(instance_id, record_id);
}

void bridge_elra_close(void) {
    elra_close();
}

// ============================================================
// LS Driver Bridge
// ============================================================

int bridge_ls_init(const char* channel_id) {
    return ls_init(channel_id);
}

int bridge_ls_capture(char* out, size_t out_size) {
    LSOutput output = {0};
    if (ls_capture(&output) < 0) {
        return -1;
    }
    size_t len = output.size;
    if (len > out_size - 1) {
        len = out_size - 1;
    }
    memcpy(out, output.buffer, len);
    out[len] = '\0';
    ls_free(&output);
    return (int)len;
}

int bridge_ls_detect_type(const unsigned char* data, size_t len) {
    LSDataType type = ls_detect_type(data, len);
    return (int)type;
}

int bridge_ls_send(const char* channel_id, const unsigned char* data, size_t len) {
    return ls_send(channel_id, data, len);
}

// ============================================================
// Renderer Bridge
// ============================================================

static Renderer g_renderer = {0};

int bridge_renderer_init(const char* title, int w, int h) {
    return renderer_init(&g_renderer, title, w, h);
}

int bridge_renderer_clear(void) {
    return renderer_clear(&g_renderer);
}

int bridge_renderer_draw_text(const char* text, int x, int y,
                               unsigned char r, unsigned char g, unsigned char b) {
    Color fg = {r, g, b, 255};
    Color bg = {0, 0, 0, 255};
    return renderer_draw_text(&g_renderer, text, x, y, fg, bg);
}

int bridge_renderer_present(void) {
    return renderer_present(&g_renderer);
}

int bridge_renderer_should_close(void) {
    return renderer_should_close(&g_renderer);
}

void bridge_renderer_poll_events(void) {
    renderer_poll_events(&g_renderer);
}

void bridge_renderer_destroy(void) {
    renderer_destroy(&g_renderer);
}
