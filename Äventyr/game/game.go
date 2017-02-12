package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Om man ska lägga till nya kommandon så lägger man till det här.
const (
	SUICIDE int = iota // SUICIDE är ett kommando för att ta självmord.
	LEFT
	RIGHT
	UP
	DOWN
)

// NewPlayer skapar en ny spelare då guld och hälsa har default values.
func NewPlayer() *Player {
	return &Player{
		Gold:   10,
		Health: 100,
	}
}

// Player är en struct för spelet/spelare
type Player struct {
	Name     string
	Lastname string
	Username string
	Age      int
	Gold     int
	Health   float64
	x        int
	y        int
	steps    int
	Actions  ActionList
}

// Move tar hand om förflyttningen på en spelare, ändra positionen och öka steps
func (p *Player) Move(direction int) {
	switch direction {
	case LEFT:
		p.x--
	case RIGHT:
		p.x++
	case UP:
		p.y++
	case DOWN:
		p.y--
	default:
		return
	}
	p.steps++
	p.Health -= 0.1
}

// Dead är en enkel funktion för att ta död på en spelare
func (p *Player) Dead() {
	p.Health = 0
	fmt.Println("You are now dead.")
}

// Resurrect återupplivar spelaren, hälsan kommer att motsvara 30% av den mängd guld
// som spelaren hade, när den dog.
func (p *Player) Resurrect() bool {
	if p.Gold <= 0 {
		return false
	}
	health := float64(p.Gold) * 0.3
	if health <= 0 {
		return false
	}
	p.Health = health
	p.Gold = 0
	fmt.Printf("You were revived with 0 Gold and %.2f health.\n", p.Health)
	return true
}

// Run är huvud funktionen för spelet, här händer det mesta
func (p *Player) Run() {
	reader := bufio.NewReader(os.Stdin)

	// Start på spelet, infinite for loop som frågar om vilken väg man ska gå.
	for {
		fmt.Print("Which direction (w,a,s,d): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		// Kolla vilket håll som skrevs
		switch strings.ToLower(text) {
		case "w":
			p.Move(UP)
		case "s":
			p.Move(DOWN)
		case "a":
			p.Move(LEFT)
		case "d":
			p.Move(RIGHT)
		case "suicide":
			p.Dead()
			return
		default:
			fmt.Println("You did not supply a valid function.")
		}

		// Kolla om spelaren dog
		if p.Health <= 0 {
			return
		}

		fmt.Printf("You moved to position %d,%d, you have now moved %d steps and you have %.2f health.\n", p.x, p.y, p.steps, p.Health)

		// Skapa en ny procentenhet, loopa sen igenom alla Actions
		// till den hittar en korrekt Action och sen kör ActionFunc.
		percentage := r.Intn(100)
		for _, a := range p.Actions.Actions {
			if percentage < a.Percentage {
				a.ActionFunc(p)
				break
			}
		}

		if p.Health <= 0 {
			if !p.Resurrect() {
				return
			}
		}
	}
}
