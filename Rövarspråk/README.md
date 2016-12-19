# Rövarspråk
För att köra detta program så navigerar du till denna mapp.

Sen kör du kommandot: go run loop.go

eller go run regex.go

Så kommer den fråga efter en sträng och sen ge dig en rövarspråks version av strängen.

## Flödesdiagram (Loop)
![Flödesdiagram](https://raw.githubusercontent.com/tryy3/15EL-Dennis-Planting/master/R%C3%B6varspr%C3%A5k/Dennis_Planting-R%C3%B6verspr%C3%A5k-Fl%C3%B6desdiagram-Loop.png)

## Pseudokod (Loop)
```
inteRövarspråk string
rövarsträng string
konsonanter string = ["B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Z"]

Fråga om en sträng och spara den under variabeln inteRövarspråk

Loopa varje tecken i strängen inteRövarspråk (for)

    Kolla om tecknet på position T är en konsonant.
        
        Om tecknet är en konsonant, kopiera tecknet till rävarsträng.
        Om tecknet inte är en konsonant, kopera till rävarsträngen, lägg till ett o och kopera tecknet igen.

Skriv ut rävarsträngen.
```

## Flödesdiagram (Regex)
![Flödesdiagram](https://raw.githubusercontent.com/tryy3/15EL-Dennis-Planting/master/R%C3%B6varspr%C3%A5k/Dennis_Planting-R%C3%B6verspr%C3%A5k-Fl%C3%B6desdiagram-Regex.png)

## Pseudokod (Regex)
```
input string
regex Regexp    Någon regex för att kolla konsonanter

Fråga om en sträng och spara den under variabeln input

Skriv ut en sträng där jag byter ut alla konsonanter till tecknet + o + tecknet, med hjälp av Regex
```

## Slutsats
Hade tråkigt så jag ville kolla vilken metod som funkade bäst, så jag körde en benchmark på programmet.

När jag påbörjade Regex delen så trodde jag att den skulle vara mycket snabbare än loopen, men jag var inte säker.
Jag har hört att regex inte alltid är snabb, jag vet att det kan vara väldigt användbart men var inte säker på hastigheten.

Så när jag först körde benchmark så låg Regex ganska högt över Loopen, så jag lekte runt och lyckades få en skillnad på ca 300 ns/op.
Det betyde att loop var 300 ns snabbare än regex varje loop.

Så jag fortsatte att leka runt med Regex, men kom inte långt.

Så jag valde att snygga till mina funktioner lite, för med loopen så gjorde jag tecknet + o + strings.ToLower(tecknet), så att jag inte hade en massa stora bokstäver i meningar.

Och märkte att Golangs regex library har en ReplaceAllStringFunc som gör att jag kan ha min egna funktion när den ska byta ut varje match i regex.

Så jag testade den och fick en snyggare output. Efter det så ville jag köra en sista benchmark så jag la in funktionen i benchmarks och upptäckte att,
när jag använde ReplaceAllStringFunc så gick benchmarken 500ns/op snabbare.

Så slutsatsen är, ReplaceAllStringFunc är den snabbaste varianten av vad jag har testat, efter så kommer loopen och sist kommer en vanlig ReplaceAllString med regex.

```
BenchmarkLoop1-12         300000              5065 ns/op
BenchmarkRegex1-12        300000              5319 ns/op
BenchmarkRegex2-12        300000              5312 ns/op
BenchmarkRegex3-12        300000              4598 ns/op
```

Benchmarks och tests finns i tests/default_test.go, du kan köra dem genom att byta directory till tests och köra kommandot: go test -bench=.

Även om Loopen hade varit snabbare än alla regex funktioner, så tror jag att jag föredrar regex i denna situation, då regex bara krävde ca 4 linjer (några fler om du ska printa ut saker.)
2 linjer för att läsa input och 2 linjer för regex.

Medans Loopen krävs en funktion för att kolla om tecknet finns i en slice plus en loop för att kolla en massa saker och liknande.
