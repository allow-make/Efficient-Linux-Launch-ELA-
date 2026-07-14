#ifndef RENDERER_H
#define RENDERER_H

#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    void* window;
    int width;
    int height;
    int font_size;
    int initialized;
    char theme_name[32];
} Renderer;

typedef struct {
    unsigned char r;
    unsigned char g;
    unsigned char b;
    unsigned char a;
} Color;

int renderer_init(Renderer* r, const char* title, int w, int h);
int renderer_set_theme(Renderer* r, const char* theme_name);
int renderer_clear(Renderer* r);
int renderer_draw_text(Renderer* r, const char* text, int x, int y, Color fg, Color bg);
int renderer_draw_cursor(Renderer* r, int x, int y, int blink);
int renderer_present(Renderer* r);
int renderer_should_close(Renderer* r);
void renderer_poll_events(Renderer* r);
void renderer_destroy(Renderer* r);

#ifdef __cplusplus
}
#endif

#endif
