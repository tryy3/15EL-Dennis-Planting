package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"sync"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type GrafikGame struct{}

var (
	basicFont *common.Font
)

type Ball struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
	SpeedComponent
}

type Score struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type Paddle struct {
	ecs.BasicEntity
	ControlComponent
	common.CollisionComponent
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
	w.AddSystem(&ControlSystem{})
	w.AddSystem(&SpeedSystem{})
	w.AddSystem(&BounceSystem{})
	w.AddSystem(&ScoreSystem{})

	basicFont = (&common.Font{URL: "fonts/Roboto-Regular.ttf", Size: 32, FG: color.NRGBA{255, 255, 255, 255}})
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
	ball.SpeedComponent = SpeedComponent{Point: engo.Point{300, 1000}}

	// Add our entity to the appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&ball.BasicEntity, &ball.RenderComponent, &ball.SpaceComponent)
		case *common.CollisionSystem:
			sys.Add(&ball.BasicEntity, &ball.CollisionComponent, &ball.SpaceComponent)
		case *SpeedSystem:
			sys.Add(&ball.BasicEntity, &ball.SpeedComponent, &ball.SpaceComponent)
		case *BounceSystem:
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
		case *ControlSystem:
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
		case *ScoreSystem:
			sys.Add(&score.BasicEntity, &score.RenderComponent, &score.SpaceComponent)
		}
	}
}

func (*GrafikGame) Type() string {
	return "GrafikGame"
}

type scoreEntity struct {
	*ecs.BasicEntity
	*common.RenderComponent
	*common.SpaceComponent
}

type ScoreSystem struct {
	entities                  []scoreEntity
	points, deaths, highscore int
	upToDate                  bool
	scoreLock                 sync.RWMutex
}

func (s *ScoreSystem) New(*ecs.World) {
	s.upToDate = false
	engo.Mailbox.Listen("ScoreMessage", func(message engo.Message) {
		scoreMessage, isScore := message.(ScoreMessage)
		if !isScore {
			return
		}

		s.scoreLock.Lock()
		if !scoreMessage.Death {
			s.points += 1

			log.Println("The score is now", s.points)
		} else {
			if s.points >= s.highscore {
				s.highscore = s.points
			}
			s.points = 0
			s.deaths += 1

			log.Println("The highscore is now", s.highscore, "and deaths", s.deaths)
		}

		s.upToDate = false
		s.scoreLock.Unlock()
	})
}

func (s *ScoreSystem) Add(basic *ecs.BasicEntity, render *common.RenderComponent, space *common.SpaceComponent) {
	s.entities = append(s.entities, scoreEntity{basic, render, space})
}

func (s *ScoreSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range s.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

func (s *ScoreSystem) Update(dt float32) {
	for _, e := range s.entities {
		if !s.upToDate {
			s.scoreLock.RLock()
			label := fmt.Sprintf("Points: %v, Deaths: %v, Highscore: %v", s.points, s.deaths, s.highscore)
			s.upToDate = true
			s.scoreLock.RUnlock()

			e.RenderComponent.Drawable.Close()

			e.RenderComponent.Drawable = basicFont.Render(label)

			width, _, _ := basicFont.TextDimensions(label)

			e.SpaceComponent.Position.X = (engo.GameWidth() - float32(width)) / 2
		}
	}
}

type ScoreMessage struct {
	Death bool
}

func (ScoreMessage) Type() string {
	return "ScoreMessage"
}

type bounceEntity struct {
	*ecs.BasicEntity
	*SpeedComponent
	*common.SpaceComponent
}

type BounceSystem struct {
	entities []bounceEntity
}

func (b *BounceSystem) Add(basic *ecs.BasicEntity, speed *SpeedComponent, space *common.SpaceComponent) {
	b.entities = append(b.entities, bounceEntity{basic, speed, space})
}

func (b *BounceSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range b.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		b.entities = append(b.entities[:delete], b.entities[delete+1:]...)
	}
}

func (b *BounceSystem) Update(dt float32) {
	for _, e := range b.entities {

		if e.SpaceComponent.Position.Y >= engo.GameHeight()-e.SpaceComponent.Height {
			engo.Mailbox.Dispatch(ScoreMessage{true})

			e.SpaceComponent.Position.X = (engo.GameWidth() / 2) - 16
			e.SpaceComponent.Position.Y = (engo.GameHeight() / 2) - 16
			e.SpeedComponent.X = engo.GameWidth() * rand.Float32()
			e.SpeedComponent.Y = engo.GameHeight() * rand.Float32()
		}

		if e.SpaceComponent.Position.X >= engo.GameWidth()-e.SpaceComponent.Width {
			e.SpaceComponent.Position.X = engo.GameWidth() - e.SpaceComponent.Width
			e.SpeedComponent.X *= -1
		}
		if e.SpaceComponent.Position.X < 0 {
			e.SpaceComponent.Position.X = 0
			e.SpeedComponent.X *= -1
		}

		if e.SpaceComponent.Position.Y < 0 {
			e.SpaceComponent.Position.Y = 0
			e.SpeedComponent.Y *= -1
		}
	}
}

type SpeedComponent struct {
	engo.Point
}

type speedEntity struct {
	*ecs.BasicEntity
	*SpeedComponent
	*common.SpaceComponent
}

type SpeedSystem struct {
	entities []speedEntity
}

func (s *SpeedSystem) New(*ecs.World) {
	engo.Mailbox.Listen("CollisionMessage", func(message engo.Message) {
		log.Println("collision")

		collision, isCollision := message.(common.CollisionMessage)
		if isCollision {
			for _, e := range s.entities {
				if e.ID() == collision.Entity.BasicEntity.ID() {
					e.SpeedComponent.Y *= -1

					engo.Mailbox.Dispatch(ScoreMessage{false})
				}
			}
		}
	})
}

func (s *SpeedSystem) Add(basic *ecs.BasicEntity, speed *SpeedComponent, space *common.SpaceComponent) {
	s.entities = append(s.entities, speedEntity{basic, speed, space})
}

func (s *SpeedSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range s.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

func (s *SpeedSystem) Update(dt float32) {
	speedMultiplier := float32(100)

	for _, e := range s.entities {
		e.SpaceComponent.Position.X += e.SpeedComponent.X * dt
		e.SpaceComponent.Position.Y += e.SpeedComponent.Y * dt

		var direction float32
		/*if e.SpeedComponent.Y < 0 {
			direction = 1.0
		} else {
			direction = -1.0
		}*/

		e.SpeedComponent.X += speedMultiplier * dt * direction
		e.SpeedComponent.Y += speedMultiplier * dt * direction
	}
}

type ControlComponent struct {
}

type controlEntity struct {
	*ecs.BasicEntity
	*ControlComponent
	*common.SpaceComponent
}

type ControlSystem struct {
	entities []controlEntity
}

func (c *ControlSystem) New(w *ecs.World) {

}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, control *ControlComponent, space *common.SpaceComponent) {
	c.entities = append(c.entities, controlEntity{basic, control, space})
}

func (c *ControlSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}

func (c *ControlSystem) Update(dt float32) {
	for _, e := range c.entities {
		speed := engo.GameWidth() * dt

		vert := engo.Input.Axis("move")
		e.SpaceComponent.Position.X += speed * vert.Value()

		if (e.SpaceComponent.Width + e.SpaceComponent.Position.X) > engo.GameWidth() {
			e.SpaceComponent.Position.X = engo.GameWidth() - e.SpaceComponent.Width
		} else if e.SpaceComponent.Position.X < 0 {
			e.SpaceComponent.Position.X = 0
		}
	}
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
