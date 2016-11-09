package systems

import (
	"fmt"
	"image/color"
	"log"
	"sync"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var (
	basicFont *common.Font
)

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

func (s *ScoreSystem) New(w *ecs.World) {
	basicFont = (&common.Font{URL: "fonts/Roboto-Regular.ttf", Size: 32, FG: color.NRGBA{255, 255, 255, 255}})

	if err := basicFont.CreatePreloaded(); err != nil {
		log.Println("Could not load font:", err)
	}

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
