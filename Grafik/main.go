package main

import (
	"image/color"
	"log"

	"github.com/tryy3/15EL-Dennis-Planting/Grafik/systems"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type GrafikGame struct{}

type Ball struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
	systems.SpeedComponent
}

type Paddle struct {
	ecs.BasicEntity
	systems.ControlComponent
	common.CollisionComponent
	common.RenderComponent
	common.SpaceComponent
}

type Score struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (grafik *GrafikGame) Preload() {
	err := engo.Files.Load("fonts/Roboto-Regular.ttf", "textures/ball.png", "textures/paddle.png")
	if err != nil {
		log.Println(err)
	}
}

func (grafik *GrafikGame) Setup(w *ecs.World) {
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.CollisionSystem{})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&systems.ControlSystem{})
	w.AddSystem(&systems.SpeedSystem{})
	w.AddSystem(&systems.BounceSystem{})
	w.AddSystem(&systems.ScoreSystem{})

	basicFont := (&common.Font{URL: "fonts/Roboto-Regular.ttf", Size: 32, FG: color.NRGBA{255, 255, 255, 255}})

	if err := basicFont.CreatePreloaded(); err != nil {
		log.Println("Could not load font:", err)
	}

	ballTexture, err := common.LoadedSprite("textures/ball.png")
	if err != nil {
		log.Println("Could not load texture:", err)
	}

	ball := Ball{BasicEntity: ecs.NewBasic()}
	ball.RenderComponent = common.RenderComponent{
		Drawable: ballTexture,
		Scale:    engo.Point{2, 2},
	}

	ball.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{(engo.GameWidth() - ballTexture.Width()) / 2, (engo.GameHeight() - ballTexture.Height()) / 2},
		Width:    ballTexture.Width() * ball.RenderComponent.Scale.X,
		Height:   ballTexture.Height() * ball.RenderComponent.Scale.Y,
	}

	ball.CollisionComponent = common.CollisionComponent{
		Main:  true,
		Solid: true,
	}
	ball.SpeedComponent = systems.SpeedComponent{Point: engo.Point{300, 1000}}

	// Add our entity to the appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&ball.BasicEntity, &ball.RenderComponent, &ball.SpaceComponent)
		case *common.CollisionSystem:
			sys.Add(&ball.BasicEntity, &ball.CollisionComponent, &ball.SpaceComponent)
		case *systems.SpeedSystem:
			sys.Add(&ball.BasicEntity, &ball.SpeedComponent, &ball.SpaceComponent)
		case *systems.BounceSystem:
			sys.Add(&ball.BasicEntity, &ball.SpeedComponent, &ball.SpaceComponent)
		}
	}

	engo.Input.RegisterAxis("move", engo.AxisKeyPair{engo.ArrowLeft, engo.ArrowRight})
	paddleTexture, err := common.LoadedSprite("textures/paddle.png")
	if err != nil {
		log.Println(err)
	}

	paddle := Paddle{BasicEntity: ecs.NewBasic()}
	paddle.RenderComponent = common.RenderComponent{
		Drawable: paddleTexture,
		Scale:    engo.Point{4, 2},
	}

	paddle.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{(engo.GameWidth() - (paddle.RenderComponent.Scale.X * paddleTexture.Width())) / 2, engo.GameHeight() - 32},
		Width:    paddle.RenderComponent.Scale.X * paddleTexture.Width(),
		Height:   paddle.RenderComponent.Scale.Y * paddleTexture.Height(),
	}
	paddle.CollisionComponent = common.CollisionComponent{
		Main:  false,
		Solid: true,
	}
	// Add our entity to the appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&paddle.BasicEntity, &paddle.RenderComponent, &paddle.SpaceComponent)
		case *common.CollisionSystem:
			sys.Add(&paddle.BasicEntity, &paddle.CollisionComponent, &paddle.SpaceComponent)
		case *systems.ControlSystem:
			sys.Add(&paddle.BasicEntity, &paddle.ControlComponent, &paddle.SpaceComponent)
		}
	}

	score := Score{BasicEntity: ecs.NewBasic()}

	score.RenderComponent = common.RenderComponent{Drawable: basicFont.Render(" ")}
	score.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{100, 100},
		Width:    100,
		Height:   100,
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&score.BasicEntity, &score.RenderComponent, &score.SpaceComponent)
		case *systems.ScoreSystem:
			sys.Add(&score.BasicEntity, &score.RenderComponent, &score.SpaceComponent)
		}
	}
}

func (*GrafikGame) Type() string {
	return "GrafikGame"
}

func main() {
	opts := engo.RunOptions{
		Title:         "Grafik Game",
		Width:         1920 * 0.8,
		Height:        1080 * 0.8,
		ScaleOnResize: true,
	}
	engo.Run(opts, &GrafikGame{})
}
