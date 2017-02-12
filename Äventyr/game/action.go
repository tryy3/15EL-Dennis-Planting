package game

import "errors"

// Action är en enkel struct för varje händelse som
// kan hända när en spelare förflyttar sig.
type Action struct {
	ActionFunc func(p *Player)
	Percentage int
}

// ActionList är en lista av flera Actions
type ActionList struct {
	Actions []*Action
}

// AddAction är en funktion för att lägga till nya Actions.
// Vad funktionen gör att den kollar om det finns tidigare Action,
// och om det finns en tidigare Action så hämtar den procenten från den,
// och lägger in det i Weight, så när man loopar igenom ActionList,
// så kommer procenten att följa med.
func (l *ActionList) AddAction(a *Action) error {
	weight := 0

	if len(l.Actions) > 0 {
		weight = l.Actions[len(l.Actions)-1].Percentage
	}

	if a.Percentage+weight > 100 {
		return errors.New("Can't add more then 100%")
	}

	a.Percentage += weight
	l.Actions = append(l.Actions, a)
	return nil
}
