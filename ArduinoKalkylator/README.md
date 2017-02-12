# Arduino Kalkylator
Denna dokumentation är uppdelad i två delar,
fysisk dokumentation och logisk dokumentation,
där fysisk dokumentation är saker så som arduino, kopplingar osv.
Medan logisk dokumentation är programmerings delen.

Jag kände att båda delarna är ganska stora så jag valde att dela upp dem.

Båda delarna har olika versioner/prototyper samt en reflektion över framtida saker man skulle kunna göra.

Projektet är hyfsat enkelt, man har en arduino, några knappar och en LCD.
Jag valde att ha en 16x2 LCD där första linjen används för att visa dem tal och matematiska funktioner som man har skrivit in,
och andra linjen används för att skriva ut resultatet.

Knapparna kan man både göra på ett hyfsat enkelt sätt och ett lite svårare sätt,
beror mest på hur många knappar du vill ha,
jag valde att gå med ett matrix för knapparna eller så kallad keypad,
då jag ville ha många knappar.
Men man kan också enkelt koppla så att varje knapp går till en pin på Arduino.

[Reference till "matrix för knappar".](http://playground.arduino.cc/Main/KeypadTutorial)

## Logisk
Koden dokumenterad där .h filen är en mer överskådlig dokumentation av koden, t.ex. vad varje klass gör och liknande.
Medan .cpp filen är mer specifik, beskriver vad olika funktioner gör, lite kommentarer i varje funktion osv.

Sen finns det också dokumentation om koden nedanför,
som är en mer loggbok/dokumentation där jag beskriver om vad helheten av koden gör,
samt vad jag tycker är bra/dåligt med olika versioner.

.ino filen är en Arduino sketch, jag har inte kvar Arduino sketchen från första versionen av Kalkylator,
men koden från Arduino sketchen ändrades aldrig så mycket,
den stora skillnaden var att andra versionen hade mer funktioner,
men dem största skillnaderna var i själva library och inte i Arduino sketchen.

### Första versionen
Första versionen ligger under LibraryV1.

Det var en ganska enkel version, vad den gör är att den har en funktion som heter addChar som lägger till en char i en buffer,
sen finns det en funktion som heter evaluate.

När evaluate körs, så loopar den igenom hela buffern och bygger en lista av *Values*.

En *Value* är en siffra eller en matematisk funktion (+,-,*,/).

När listan har byggts och inga errors har upptäckts,
så börjar den loopa igenom listan där den kollar efter en siffra och en funktion och beroende på vad funktionen är,
så lägger den ihop siffrorna med funktionen, t.ex. 1+1 så kommer värdet bli 2.

Det var en ganska enkel version, jag gjorde den ganska fort för att kunna komma igång med alla andra delar,
men jag var inte nöjd med koden, den följde inte alla matematiska regler (*/ före +- osv.),
den gav också ut errors där man kanske inte nödvändigtvis behövde osv.

![test1](https://github.com/tryy3/15EL-Dennis-Planting/blob/master/ArduinoKalkylator/bilder/test1.jpg?raw=true)
![test2](https://github.com/tryy3/15EL-Dennis-Planting/blob/master/ArduinoKalkylator/bilder/test2.jpg?raw=true)

Bilderna ovan är några test som jag gjorde,
första bilden visar ett test när jag skulle kolla vilken knapp man klickar på,
andra bilden är ett test där jag testade olika matematiska funktioner.

Så jag påbörjade en ny version.

### Andra versionen
Andra versionen är lite mer komplicerad och är baserad på http://stackoverflow.com/a/26227947/2817262

Vad den gör är att den har 3 huvud funktioner, parseExpression, parseTerm och parseFactor.

parseExpression tar hand om + och - samt att den också kör parseTerm.

parseTerm tar hand om * och / samt att den också kör parseFactor.

parseExpression tar hand om det mesta, (,),^ samt funktioner så som sin/cos/sqrt osv.

Så vad som händer är att man loopar igenom varje char i buffern för att hitta en funktion, siffra osv och sen räkna ut den.

Största skillnaden mellan denna kod och min kod är att denna kod börjar med Faktorer så som paranteser, funktioner, osv. och följer matte regler.
Den är också väldigt flexibel med att den har möjlighet till funktioner, så om man vill så skulle man kunna bygga ut den med fler funktioner.

### Framtiden/Kritik
Jag har inte gjort mycket tester när det kommer till prestanda och liknande.
Andra versionen använder sig av många funktion calls som kan eventuellt leda till dålig prestanda,
så om det finns någon annan lösning så skulle man kunna göra tester och se vilken lösning är best på prestandan.
Men jag är inte så orolig över prestandan då Arduino är ett hyfsat kraftfullt kort och eftersom jag använder mig av en 16x2 LCD så får man inte in så många Char ändå.

Jag gjorde inte så mycket tester med andra versionen då jag hittade koden på stackoverflow och sen ändrade jag koden från java till C++,
så det skulle nog krävas lite mer tester för att hitta eventuella buggar.

Eftersom andra versionen är en ganska flexibel lösning,
så skulle man kunna bygga ut den med fler funktioner som är relaterade till programmering.
Man skulle också kunna koda om den till en mer library baserad kod,
även om jag gjorde den till en arduino library,
så finns det delar man skulle kunna göra mer flexibla så att vem som helst skulle kunna implementera den i sin kod.

Det skulle också vara roligt att bygga på någon historik funktion där man kan se vad man har räknat,
samt någon knapp för typ senaste räknat eller liknande som vissa miniräknare har.

Jag är också inte så säker på hur sin/cos funkar och min lösning på funktionerna kan vara fel,
men det är nog aldrig något jag kommer att använda så det är inte så stor grej för mig.

https://www.arduino.cc/en/Hacking/LibraryTutorial

## Fysisk
### Första prototypen
Första prototypen var ganska enkel,
koppla ihop en LCD med en arduino samt 5x4 button matrix.

![prototyp1](https://github.com/tryy3/15EL-Dennis-Planting/blob/master/ArduinoKalkylator/bilder/prototyp1.jpg?raw=true)

Det var inte några stora problem med hur jag byggde prototypen,
men det var långa sladdar och sladdar hit och ditt,
så den prototypen skulle inte vara så praktisk.

### Andra prototypen
Den andra prototypen var mer seriös,
den har en platta där man kan löda fast alla knappar,
den har också väldigt korta kablar på baksidan som gör att det blir snyggt och man ser vart kablarna går.
Den har också bättre kabel dragning och mer struktur.

Med andra versionen av den logiska biten så expanderade jag knapp matrixen från 4x5 till 5x5 för att kunna ha mer funktioner.

![knapp1](https://github.com/tryy3/15EL-Dennis-Planting/blob/master/ArduinoKalkylator/bilder/knappar1.jpg?raw=true)
![knapp2](https://github.com/tryy3/15EL-Dennis-Planting/blob/master/ArduinoKalkylator/bilder/knappar2.jpg?raw=true)
![knapp3](https://github.com/tryy3/15EL-Dennis-Planting/blob/master/ArduinoKalkylator/bilder/knappar3.jpg?raw=true)

Bilderna ovan är bilder på processen av knapp matrixen.

Med hjälp av fritzing så gjorde jag också en schematic som visar hur kablarna går.

![schematic](https://github.com/tryy3/15EL-Dennis-Planting/blob/master/ArduinoKalkylator/bilder/schematics.jpg?raw=true)

### Framtiden/Kritik
Det finns några saker jag skulle vilja göra i framtiden.
Det första är att göra kablarna mindre,
är inte helt säker på hur jag ska lyckas med det än,
jag försökte flera gånger på olika sätt att komma fram till något enkelt sätt,
att göra kablarna som går från arduino till LCDn kortare men har inte löst det än.

Den andra stora delen jag skulle vilja göra är att göra någon slags box som man kan lägga allt i,
jag lekte runt med idéen att ta några gamla laptop AC/DC adaptrar och använda dem för att lägga allt i.
Eller alternativt göra en box av trä/kartong,
men får se om jag tar tid och gör det,
skulle vara roligt att använda sig av AC/DC adaptrar då det skulle bli unikt, möjligtvis lite klumpigt men ändå unikt.

Det finns också små grejor som jag skulle vilja lägga till,
troligen när jag har en box och har en mer färdig prototyp,
det jag skulle vilja göra är att lägga till ett 9v batteri eller några 1.5v batterier och sen en switch,
så att den blir mer portabel och man kan stänga och sätta på den.