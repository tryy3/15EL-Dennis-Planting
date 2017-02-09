# Arduino Kalkylator
## Första versionen
Första versionen ligger under LibraryV1.

Den var ganska enkel, vad den gör att den har en funktion som heter addChar som lägger till en char i en buffer,
sen finns det en funktion som heter evaluate.

När evaluate körs, så loopar den igenom hela buffern och bygger en lista av Values.

En Value är en siffra eller en matematiskt funktion (+,-,*,/).

När listan har byggts och inga errors har upptäckts,
så börjar den loopa igenom listan där den kollar efter en siffra och en funktion och beroende på vad funktionen är,
lägger den ihop siffrorna med funktionen, t.ex. 1+1 så kommer värdet bli 2.

Det var en ganska enkel version, jag gjorde den ganska fort för att kunna komma igång med alla andra delar,
men jag var inte nöjd med koden, den följde inte alla matematiska regler (*/ före +- osv.),
den gav också ut errors där man kanske inte nödvändigtvis behövde osv.

Så jag påbörjade en ny version.

## Andra versionen
Andra versionen är lite mer komplicerad och är baserad på http://stackoverflow.com/a/26227947/2817262

Vad den gör är att den har 3 huvud funktioner, parseExpression, parseTerm och parseFactor.

parseExpression tar hand om + och - samt att den också kör parseTerm.

parseTerm tar hand om * och / samt att den också kör parseFactor.

parseExpression tar hand om det mesta, (,),^ samt funktioner så som sin/cos/sqrt osv.

Så vad som händer är att man loopar igenom varje char i buffern för att hitta en funktion, siffra osv och sen räkna ut den.

Största skillnaden mellan denna kod och min kod är att denna kod börjar med Faktorer 

https://www.arduino.cc/en/Hacking/LibraryTutorial