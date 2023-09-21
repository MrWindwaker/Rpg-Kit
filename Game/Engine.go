package game

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const GAME_SIZE int = 4

type Engine struct {
	should_close bool
	game_size    int

	sky rl.Texture2D

	// Assets
	sprites map[string]Sprite
}

var eng_lock = &sync.Mutex{}
var eng_instance *Engine

func Get_Engine() *Engine {
	if eng_instance == nil {
		eng_lock.Lock()
		defer eng_lock.Unlock()

		if eng_instance == nil {
			eng_instance = &Engine{
				should_close: false,
				game_size:    GAME_SIZE,
				sprites:      map[string]Sprite{},
			}
		}
	}

	return eng_instance
}

func (e *Engine) Run() {
	e.init()

	for !e.should_close {
		e.update()
		e.draw()
	}

	e.close()
}

func (e *Engine) init() {
	rl.SetConfigFlags(rl.FlagInterlacedHint)
	//rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(1280, 720, "Rpg Kit")

	e.sky = rl.LoadTexture("Assets/sky.png")
	e.sprites["ground"] = New_Sprite(rl.LoadTexture("Assets/ground.png"), 1, 1, 0, rl.NewVector2(0, 0), rl.NewVector2(320, 180))
	e.sprites["hero"] = New_Sprite(
		rl.LoadTexture("Assets/hero-sheet.png"),
		3,
		8,
		1,
		rl.NewVector2(float32((16*6)*GAME_SIZE), float32((16*5)*GAME_SIZE)),
		rl.NewVector2(32, 32),
	)
}

func (e *Engine) update() {
	e.should_close = rl.WindowShouldClose()
}

func (e *Engine) draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.GetColor(0x333333FF))

	rl.DrawTextureEx(e.sky, rl.NewVector2(0, 0), 0, float32(GAME_SIZE), rl.White)

	for _, sp := range e.sprites {
		sp.Draw()
	}

	rl.EndDrawing()
}

func (e *Engine) close() {
	rl.UnloadTexture(e.sky)

	for _, sp := range e.sprites {
		sp.Unload()
	}

	rl.CloseWindow()
}
