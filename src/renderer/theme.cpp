#include "theme.h"
#include <string.h>

const Theme THEME_DEFAULT = {
    .bg_r = 0, .bg_g = 0, .bg_b = 0,
    .fg_r = 0, .fg_g = 255, .fg_b = 0,
    .accent_r = 0, .accent_g = 255, .accent_b = 0,
    .name = "default"
};

const Theme THEME_DARK = {
    .bg_r = 0, .bg_g = 0, .bg_b = 0,
    .fg_r = 255, .fg_g = 255, .fg_b = 255,
    .accent_r = 255, .accent_g = 255, .accent_b = 255,
    .name = "dark"
};

const Theme THEME_LIGHT = {
    .bg_r = 255, .bg_g = 255, .bg_b = 255,
    .fg_r = 0, .fg_g = 0, .fg_b = 0,
    .accent_r = 0, .accent_g = 0, .accent_b = 0,
    .name = "light"
};

// Termux official colors
// Background: #0c0c0c (very dark, slight black)
// Foreground: #ffffff (white)
// Prompt green: #00ff00
// Accent/orange: #ff6f00
const Theme THEME_TERMUX = {
    .bg_r = 12, .bg_g = 12, .bg_b = 12,    // #0c0c0c
    .fg_r = 255, .fg_g = 255, .fg_b = 255, // #ffffff
    .accent_r = 255, .accent_g = 111, .accent_b = 0, // #ff6f00
    .name = "termux"
};

const Theme* theme_get_by_name(const char* name) {
    if (name == NULL) return &THEME_DEFAULT;
    if (strcmp(name, "default") == 0) return &THEME_DEFAULT;
    if (strcmp(name, "dark") == 0) return &THEME_DARK;
    if (strcmp(name, "light") == 0) return &THEME_LIGHT;
    if (strcmp(name, "termux") == 0) return &THEME_TERMUX;
    return &THEME_DEFAULT;
}

void theme_apply(const Theme* theme) {
    if (theme == NULL) return;
    glClearColor(
        theme->bg_r / 255.0f,
        theme->bg_g / 255.0f,
        theme->bg_b / 255.0f,
        1.0f
    );
}
