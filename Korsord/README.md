# Korsord
## Beskrivning
Poängen med detta program är att göra korsord i ett program t.ex. i terminalen eller i någon slags GUI.
Först så sätter man upp storleken på layouten av korsordet (hur många kolumer och rader den ska ha).
Sen frågar den efter vilken position och siffra, så att man sätter ut en siffra över hela korsordet.
Varje ruta behöver inte ha en siffra, om den inte har någon siffra så kommer det att räknas som ett "block", där man inte kan sätta en bokstav.

När layouten är färdig så börjar programmet att fråga efter en siffra och bokstav.
Det är här korsordet börjar att lösas, siffrorna som man skrev ut tidigare kommer att bli ersatta av bokstäver.
Varje siffra kommer att ha en bokstav i sig, så om en siffra finns på två eller mer olika ställen så kommer varje position att ha samma bokstav.

Då golang inte riktigt har någon officiell eller hyfsat enkel/bra UI library så bestämde jag mig för att göra det enkelt för mig
och använda CUI (Console User Interface).

Så att jag har 3 olika rutor. En main ruta där korsords layouten syns och blir uppdaterad varje gång något nytt händer.
En ruta på sidan där varje siffra med en bokstav listas upp, så att man kan se vilka siffror man har fyllt i och liknande.
Sen en ruta längst ner där man skriver in saker och den frågar efter olika saker och liknande.

Jag kommer att spara information på två olika sätt.

Jag kommer att använda mig av pointers.
För det jag tänker mig, är att jag har en 2D slice där jag sparar en pointer till en siffra och en bokstav.
Samt att jag har en enkel slice där jag har en Key som pointer till en siffra och en pointer till en bokstav som motsvarar 2D slicen.

Anledning till detta är, i teorin om jag ska ändra en bokstav eller liknande, så loopar jag igenom den enkla slicen och sen ändrar jag bokstaven, så borde 2D slicen uppdateras automatiskt.

Libraries för terminaler
https://github.com/jroimartin/gocui
https://github.com/gizak/termui

Saker som jag kan göra för att expandera programmer är att lägga till save/load funtion, så att du kan fortsätta på ett korsord om du inte vart färdig med det.

Möjligtvis att jag också lägger in så man kan göra färdiga korsord som man kan skicka till kompisar och liknande.
Så att dem har en färdig layout och möjligtvis vissa "hints" som redan är färdiga, sen i samma fil så finns det ett redan löst korsord,
så när man listat ut alla rätta bokstäver så har man klarat korsordet.

## Installation
För att göra installation lättare så använder jag mig av [Glide](https://github.com/Masterminds/glide).
Så först så måste man installera Glide.

Sen efter att du har klonat detta projekt (Korsord) så öppnar du en terminal och navigerar till projektet och kör kommandot
- glide install

Sen kan du bygga detta projekt med Go
- go build

Sista steget är att köra projekter
- ./Korsord

Sen följer du bara instruktionerna på terminalen.