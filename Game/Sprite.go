package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Sprite struct {
	texture rl.Texture2D

	hFrames int
	vFrames int
	frame   int

	position  rl.Vector2
	f_size    rl.Vector2
	frame_map map[int]rl.Vector2
}

func New_Sprite(txt rl.Texture2D, hf, vf, tf int, pos, fsize rl.Vector2) Sprite {
	s := Sprite{
		texture:   txt,
		hFrames:   hf,
		vFrames:   vf,
		frame:     tf,
		position:  pos,
		f_size:    fsize,
		frame_map: make(map[int]rl.Vector2),
	}

	s.Init_Frame_Map()

	return s
}

func (s *Sprite) Init_Frame_Map() {
	fc := 0

	for v := 0; v < s.vFrames; v++ {
		for h := 0; h < s.hFrames; h++ {
			s.frame_map[fc] = rl.NewVector2(s.f_size.X*float32(h), s.f_size.Y*float32(v))
			fc++
		}

	}
}

func (s *Sprite) Draw() {
	rl.DrawTexturePro(
		s.texture,
		s.Get_Source(),
		s.Get_Dest(),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}

func (s *Sprite) Get_Source() rl.Rectangle {

	frame := s.frame_map[s.frame]

	if s.frame > len(s.frame_map) {
		frame = rl.NewVector2(0, 0)
	}

	fcx := frame.X
	fcy := frame.Y

	return rl.NewRectangle(
		fcx,
		fcy,
		s.f_size.X,
		s.f_size.Y,
	)
}

func (s *Sprite) Get_Dest() rl.Rectangle {
	return rl.NewRectangle(
		s.position.X,
		s.position.Y,
		s.f_size.X*float32(GAME_SIZE),
		s.f_size.Y*float32(GAME_SIZE),
	)
}

func (s *Sprite) Unload() {
	rl.UnloadTexture(s.texture)
}
