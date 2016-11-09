#Grafik
##Inledning
Detta spel är baserad på Java tutorial https://www.youtube.com/playlist?list=PLpkWX5olvmC_DbuidMtQb3pZG30JsZjUZ alla funktioner är inte med än och jag har tänkt ändra det lite.

Golang har inte något officiellt grafik paket så som Java/C#/C++ och liknande språk har,
då språket är väldigt nytt och har mest använts i områden som folk har behövts,
i detta fall server miljö och inte desktop miljö.

Så jag var tvungen att lära mig något library som var inofficiellt,
jag testade ett flertal libraries från http://awesome-go.com/#gui men några funkade inte bra på windows,
medans andra var väldigt komplicerade för vad jag ville göra.

Så jag började kolla på mer spel baserade libraries http://awesome-go.com/#game-development.
Men det var ändå väldigt komplicerade, t.ex. att jag behövde lära mig OpenGL eller liknande,
men efter lite bråk och testning, så hittade jag https://github.com/EngoEngine/engo.

Det såg ganska bra ut, så efter att jag satt i några dagar med att försöka förstå det,
så lyckades jag modifera deras pong demo https://github.com/EngoEngine/engo/tree/master/demos/pong för att få resultatet jag ville ha.
Det tog ett tag för att förstå dem flesta delarna men jag lyckades förstå vad det mesta gör och hur allt funkar.

##Installering
För att köra spelet, så är det bara att ha Grafik.exe och assets mappen i samma mapp (assets mappen innehåller olika bilder och fonts) och sen köra Grafik.exe.

Om man ska bygga ihop programmet själv, så krävs det lite extra jobb, först måste man installera Go.
Sen kan man följa första och andra steget på sidan https://github.com/EngoEngine/engo#practice-getting-it-to-run
Därefter är det bara att gå in i denna map med cmd/terminal och skriva go build.

Jag har inte hunnit att testa att bygga ihop programmet på någon test dator ännu,
så är inte säker på att installations stegen funkar eller inte.

##Bugar
Just nu vet jag bara om en bug och det är att hastigheten ändrar sig lite random, borde vara enkelt att fixa.

##TODO
* Snygga till kodningen (just nu är det mer eller mindre en kopia av Pong koden, jag vill lägga till min egna style och enklare vy av alla komponenter i koden)
* Dokumentera koden (Jag kunde inte dokumentera samtidigt som jag kodade, då jag inte riktigt viste vad allt gjorde)
* Lägg till fler funktioner så som små kuber som bollen kan förstöra.

##System
Detta spel har flera olika system, som hanterar olika saker.

Genom att seperara spelet i olika entities/system så skapar det en enklare vy av vad allt gör.
ScoreSystem hanterar poäng, BounceSystem hanterar studsningar på väggarna, ControlSystem hanterar kontroller osv.
Då vet man vart allt händer och om man ska ändra något så behöver man inte gå igenom all kod, man behöver bara hitta rätt system.

###Ball
När spelet startar för första gången,
så laddar den in assets/textures/ball.png
och skapar en boll i spelet.

Den läggs också till i SpeedSystem och BounceSystem,
så att den kan åka runt och om den åker på Paddle så studsar den.

###Paddle
När spelet start för första gången,
så laddar den in assets/textures/paddle.png
och skapar en paddle i spelet.

Den läggs också till i ControlSystem för att man ska kunna kontrollera att paddle ska gå mot höger och vänster.
Paddle kontrolleras med höger och vänster piltangent,
det skulle gå att ändra vilken knapp som används om man vill.

###ScoreSystem
ScoreSystem hanterar rendering av texten som håller reda på hur många poäng man har,
hur många gånger man har dött,
vilken highscore man har osv.

Om jag skulle lägga till att highscoren sparas efter att man stänger ner spelet,
så skulle man göra det i detta system.

Den gör inte så mycket,
då den för det mesta bara sitter och väntar på inkommande ScoreMessage medelanden från andra system.

###BounceSystem
BounceSystem hanterar när bollen studsar på väggarna,
t.ex. om bollen åker under Paddle så förlorar man,
medans om bollen åker på sidorna eller högst upp på skärmen,
så studsar den.

Detta system gör inte så mycket mer än att håller koll på bollen inte åker utanför skärmen och att om den åker längst ner på skärmen så förlorar man.

###ControlSystem
ControlSystem är ett ganska enkelt system,
den väntar på input från registrerade tangenter (i detta fall vänster och höger piltangent),
och flyttar paddle mot vänster och höger beroende på vilken piltangent man klickade.
Den kollar också så att man inte åker utanför skärmen.

###SpeedSystem
SpeedSystem hanterar rörelsen av bollen,
den har i uppgift att flytta bollen varje gång systemet får en uppdatering (varje frame).

Om jag vill att bollen ska succesivt öka hastighet,
så skulle det vara i detta system jag skulle ändra det.