package renderer

import "ela/pkg/bridge"

type Color struct {
    R, G, B, A byte
}

func Init(title string, width, height int) int {
    return bridge.RendererInit(title, width, height)
}

func SetTheme(name string) int {
    return bridge.RendererSetTheme(name)
}

func Clear() int {
    return bridge.RendererClear()
}

func DrawText(text string, x, y int, color Color) int {
    return bridge.RendererDrawText(text, x, y, color.R, color.G, color.B)
}

func DrawCursor(x, y int, blink int) int {
    return bridge.RendererDrawCursor(x, y, blink)
}

func Present() int {
    return bridge.RendererPresent()
}

func ShouldClose() int {
    return bridge.RendererShouldClose()
}

func PollEvents() {
    bridge.RendererPollEvents()
}

func Destroy() {
    bridge.RendererDestroy()
}
