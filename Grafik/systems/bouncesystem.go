package systems

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

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
