#ifndef BRIDGE_H
#define BRIDGE_H

#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

// ELRA Registry Bridge
int bridge_elra_init(const char* db_path);
int bridge_elra_create_instance(const char* id, const char* name, const char* arch,
                                 const char* mode, const char* path, int pid,
                                 int auto_inject, const char* status);
int bridge_elra_get_instance(const char* id, char* out, size_t out_size);
int bridge_elra_list_instances(char* out, size_t out_size);
int bridge_elra_update_status(const char* id, const char* status);
int bridge_elra_delete_instance(const char* id);
int bridge_elra_create_record(const char* id, const char* target_id,
                               const char* mode, const char* path, int auto_inject);
int bridge_elra_merge_record(const char* instance_id, const char* record_id);
void bridge_elra_close(void);

// LS Driver Bridge
int bridge_ls_init(const char* channel_id);
int bridge_ls_capture(char* out, size_t out_size);
int bridge_ls_detect_type(const unsigned char* data, size_t len);
int bridge_ls_send(const char* channel_id, const unsigned char* data, size_t len);

// Renderer Bridge
int bridge_renderer_init(const char* title, int w, int h);
int bridge_renderer_clear(void);
int bridge_renderer_draw_text(const char* text, int x, int y,
                               unsigned char r, unsigned char g, unsigned char b);
int bridge_renderer_present(void);
int bridge_renderer_should_close(void);
void bridge_renderer_poll_events(void);
void bridge_renderer_destroy(void);

#ifdef __cplusplus
}
#endif

#endif
