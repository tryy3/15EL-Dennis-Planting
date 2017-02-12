package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tryy3/15EL-Dennis-Planting/Äventyr/game"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// monstercheck tar hand om monster,
// när man dödar ett monster så har man en 30% chans på sig att inte skadas,
// annars får man en skada mellan 1-10.
func monstercheck(p *game.Player) {
	fmt.Println("Oh no you found a monster, time to fight.")
	kill := r.Intn(100)
	if kill < 30 {
		fmt.Println("You killed the monster and no damage was made to you.")
		return
	}
	damage := r.Intn(10) + 1
	p.Health -= float64(damage)

	if p.Health <= 0 {
		p.Dead()
	} else {
		fmt.Printf("You killed the monster but the monster damaged you by %d, you now have %.2f health.\n", damage, p.Health)
	}
}

// npccheck tar hand om npc:er,
// det är en 90% chans att en NPC är snäll och ger dig 1-10 i hälsa,
// om npcn är dum så tar den 30-70% av allt ditt guld.
func npccheck(p *game.Player) {
	fmt.Println("You found a npc.")
	kind := r.Intn(100)
	if kind < 90 {
		health := r.Intn(10) + 1
		p.Health += float64(health)
		fmt.Printf("The NPC was really kind and gave you %d health, you now have %.2f health.\n", health, p.Health)
		return
	}

	gold := r.Float64()*0.4 + 0.3 // 0.3-0.7
	m := float64(p.Gold)
	taken := m - (m * gold)
	p.Gold = int(m * gold)
	fmt.Printf("The NPC was rude and stole %.2f from your gold, you now have %d gold.\n", taken, p.Gold)
}

// goldcheck tar hand om guld som du kan få när du går,
// det finns en 50% chans att du får 1 guld,
// 30 % chans att få 5 guld,
// och 20% chans att få 10 guld.
func goldcheck(p *game.Player) {
	value := r.Intn(100)
	amount := 0
	if value < 50 {
		amount = 1
	} else if value < 80 {
		amount = 5
	} else {
		amount = 10
	}

	p.Gold += amount
	fmt.Printf("You found %d gold, you now have %d gold.\n", amount, p.Gold)
}

// pitcheck tar hand om en fall gropp,
// fallgroppen kan vara mellan 1-5m djup och skadar samma mängd som den är djup.
func pitcheck(p *game.Player) {
	depth := rand.Intn(5) + 1
	p.Health -= float64(depth)
	fmt.Printf("You fell down a %d meter pit and got damaged %d, your now have %.2f health.\n", depth, depth, p.Health)
}

func main() {
	player := game.NewPlayer()
	reader := bufio.NewReader(os.Stdin)

	// Lägg till nya Actions som kan hända under ett spel.
	Actions := game.ActionList{}
	if err := Actions.AddAction(&game.Action{Percentage: 1, ActionFunc: monstercheck}); err != nil {
		fmt.Println(err)
		return
	}
	if err := Actions.AddAction(&game.Action{Percentage: 10, ActionFunc: npccheck}); err != nil {
		fmt.Println(err)
		return
	}
	if err := Actions.AddAction(&game.Action{Percentage: 5, ActionFunc: goldcheck}); err != nil {
		fmt.Println(err)
		return
	}
	if err := Actions.AddAction(&game.Action{Percentage: 3, ActionFunc: pitcheck}); err != nil {
		fmt.Println(err)
		return
	}
	player.Actions = Actions

	// Fråga basic frågor.
	fmt.Print("Your name: ")
	text, _ := reader.ReadString('\n')
	player.Name = text

	fmt.Print("Your lastname: ")
	text, _ = reader.ReadString('\n')
	player.Lastname = text

	fmt.Print("Your username: ")
	text, _ = reader.ReadString('\n')
	player.Username = text

	// Ålder är lite speciellt då vi måste verifiera om det är ett nummer eller inte.
	for {
		fmt.Print("Age: ")
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)

		age, err := strconv.Atoi(text)

		if err != nil {
			fmt.Println("Age is not a number.")
			continue
		}
		player.Age = age
		break
	}

	// Starta main funktionen för spelet.
	player.Run()
}
