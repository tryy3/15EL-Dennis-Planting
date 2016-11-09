# Minecraft2D
## Inledning
Jag ville ha något större projekt som jag kunde arbeta med för att lära mig mer om Golang.

Men jag ville inte göra små projekt så som att göra något enkelt chat system, statisk webbsida osv.
Och jag gillar inte att hålla på med frontend så mycket, så game dev är inte super kul för mig.

Så jag försökte komma på något projekt som jag skulle tycka vara roligt eller ha någon nytta av i dagliga livet.

Men kunde inte komma på något, så jag försökte komma på saker som jag kunde replikera på en mindre skala och som jag kan mycket om.

När jag började lära mig att koda så höll jag på mycket med Minecraft Servrar och jag höll på med dem i flera år.
Så jag kan redan innan jag börjar med detta projekt om hur Mojang kodade Minecraft, hur dem generar världen i flera chunks, hur sevrar funkar osv.
Så mycket av arkitekten på hur allt funkar, kan jag redan.

Så jag bestämde mig för att göra något enkelt där jag fokuserar mer på server miljön än på desktop/frontend miljön.
Med det menar jag, jag gör ett enkelt spel i 2D och fokuserar mycket på hur datan lagrar, hur datan skickas osv. istället.

Så jag bestämde mig för att göra Minecraft i 2D där jag bara lägger till några enkla block, en gubbe och att man kan placera och ta bort block.
Sen lägger jag mesta dels av mitt arbete på hur världen lagras, multiplayer kommunikation, funktioner så som placera/ta bort blocks osv.

## Systemen
Eftersom jag skriver detta innan jag har påbörjat projektet, så vet jag inte riktigt hur slut resultatet kommer att se ut.
Men jag har lärt mig att det är bra att planera innan hur strukturen av koden kommer att vara.
Så jag tänkte att jag kunde göra detta som ett mer seriöst projekt, där jag lägger mycket krut på att göra mycket ordentligt från början.
Istället för att programmera om allt flera gånger.

Så jag har delat in spelet i 2 stora delar, frontend och backend.
Jag delar också in varje del i flera delar, även om jag har delat in saker i olika delar så kan vissa arbeta med varandra och ibland ligger dem i samma service.
T.ex. så delar jag in server och lagring i två olika delar, men servern hanterar lagring, anledningen till detta är för att jag ville skriva ner mina tankar om hur jag ska lagra grejor
och det blev en hel del text, så jag la den i sin egna del.

### Frontend
När jag gjorde Grafik spelet så märkte jag att det Golang inte hade så mycket support för desktop miljöer.
Det fanns inte så många tutorials på hur man gör GUI och libraries var mer eller mindre bara bindings för OpenGL/Qmt/GTK osv. eller en framework på dem.

Så istället för att lägga mycket krut på att försöka komma på hur frontend ska funka, vilken library som funkar bäst för mig osv.
Så bestämde jag mig för att frontend ska vara väldigt simpelt, så jag gör en webbsida istället, där jag kan använda HTML5 som jag kan relativt mycket om från att ha använt NodeJS i 2-3 år (och innan så använde jag PHP i något år).

Jag vet inte om jag kommer använda någon HTML5 engine eller om jag bara kommer göra allt själv, det återstå att se.

Men i princip så kommer grafiken komma från HTML5, sen använder jag någon kommunikations sätt, antingen en REST API eller någon kommunikation så som WebSockets eller liknande.
Jag tror WebSockets kommer att bli bäst, då jag kan optimera hur mycket data som kommer att skickas, t.ex. att jag inte skickar hela världen istället skickar specifika chunks och generar det i HTML5.
Det kommer också göra att jag kan enkelt kommunicera multiplayer funktioner så som att man har 2 personer, en förstör ett block så får alla klienter en uppdatering osv.

Frontend kommer inte göra så mycket, mer än att uppdatera och skicka data fram och tillbaka, det kan hända att jag delar in frontend i olika delar när jag påbörjar projektet.
T.ex. att kommunkationen är en del, hur personen rör sig är en annan, placera/ta-bort blocks osv.

### Backend
Då jag fokuserar mest på server delen så har jag tänkt mest på hur jag löser alla delar bäst på server delen istället.

Jag har delat in server delen i några olika delar.
Webbservern, server, lagring.

#### Webbservern
Jag valde att webbservern och server skulle vara två olika delar.
Då det gör saker mer flexibelt, jag valde att använda HTML5 som frontend, men när projektet är färdigt så skulle jag kunna bygga på t.ex. en desktop klient.
Som fungerar likadant som HTML5 klienten men i desktop klienten, om jag valde att sätta ihop dem, så skulle det möjligtvis göra att saker bli komplicerade.

Vissa saker funkar bara med HTML5 medans andra saker bara funkar på desktop klienten.
Om jag gör det på ett mer OOP/flexibelt sätt, så jag kan lägga till fler klienter, mobil klienter, desktop klienter osv. utan att ändra server miljön, då dem kommunicerar på samma sätt.

Så Webbservern har inte så stor betydelse, allt den gör är levera HTML5 klienten (html, js, css, bilder osv.)
Medans server är sin egna service, så webbservern kan ligga på en dator, medans server ligger på en annan.

#### Server
En server kommer att ha några enklar delar.
Den hanterar lagring för olika världar, kommunikationen mellan klienter och möjligtvis några fler funktioner.

Jag tror att jag kommer att använda WebSockets eller sockets av något slag, då jag kan använda mig av global broadcast och fram och tillbaka kommunkation.

Om jag använder mig av t.ex. REST API så kan bara klienten prata med servern och inte tillbaka, samt att det kan krävas lite extra bandwidth.

Men om jag använder mig av WebSockets, så kan jag skicka specefika data när saker uppdateras, t.ex. om en person förflyttar sig, så kan jag skicka dem nya kordinaterna och deras UUID.

Den kommer också ha i uppgift att spara filer, så som världen/chunks och liknande.

#### Lagring
Jag är fortfarande lite osäker på hur mycket jag kommer att lagra.

Men tanken är att jag kommer att lagra världen i binary format.
Så varje värld kommer att vara i sin egna mapp, den har varje värld olika chunks, som kommer att vara t.ex. 16x128 eller 16x256 (längd x höjd), jag har inte bestämt hur stora chunks ska vara ännu (Minecraft är 16x16x256).

Jag är inte säker på om varje chunk ska vara i separata filer eller om chunks ska vara i en stor fil.
Om jag gör det i en stor fil, så kan det riskera att filen blir väldigt stor och det kommer att ta tid att läsa igenom filen varje gång osv.
Om jag gör det i flera filer, så kan det bli en hel del filer, jag tror inte det är något problem, men jag har hört att datorer inte är så förtjusta i en massa små filer, t.ex. anti-virus skydd kan börja sakta ner och liknande.
Jag kommer att behöva läsa lite om vilken metod som är bäst för denna situation, jag tror att separera chunks i olika filer, kommer vara bäst då vi har väldigt bra datorer idag.

Jag kommer nog att använda flera filer, då filerna blir mindre och snabbare att läsa/spara, så att jag kan läsa bara det jag behöver och inte en massa andra.
Det kommer också att göra genereringen av nya chunks lättare, jag kollar vilket chunk som ska skapas, om det inte existerar, skapa den.
Medans om jag hade en stor fil, så hade jag behövt gå igenom hela filen för att kolla om chunken finns och liknande.

Jag hade tänkt spara varje chunk som 16x128 bytes, där varje block är endast en byte, men efter att ha skrivit denna plan, så kom jag på att det kommer att göra att jag bara kan ha 256 blocks.
Eftersom 1 byte = 8 bitar = 256 olika kombinationer (0-255).
Istället kommer jag att ha en separation mellan varje block ID, så block ID kan bestå av flera bytes, som sen resulterar i ett ID.
Så om jag 2 blocks i ett chunk, så kommer filen att vara 3 bytes total.
Första byten kanske är en etta, medans sista byten är en tvåa, sen andra byten är separations byten, som antagligen blir en nolla.
Så t.ex. så kan ett block ID vara 2 bytes som ger den kombinationerna 65,534 olika kombinationer (1-65,535).

Jag har testat lite innan projektet och att översätta binary till och från int, är väldigt enkelt i Golang.

Detta är hur jag har tänkt, när det kommer till hur jag ska lagra varje block, så varje block lagras med ett specefikt block ID och ingenting mer.

I framtiden om jag vill lagra mer, så som custom namn, kistor med blocks i osv. så kommer jag att spara dem i separata filer.
Anledningen till detta är för det mesta hastigheten, i varje chunk så behöver jag generera varje block och spara dem.
Men det är inte säkert att varje block som har extra funktioner behöver spara sin information.
T.ex. om en kista är tom, så behöver den inte ta upp data och sakta ner hastigheten på att skicka data fram och tillbaka.

Jag kommer också behöva spara olika spelares data, så som användarnamn, deras placering och andra saker.
Jag tror att jag kommer att spara detta i en json fil eller i någon databas, eller separera dem.
Så att jag har något registrerings funktion som spara lösenord, användarnamn/email och liknande i en databas, samt en UUID av något slag.
Och sen sparar jag deras placering i världen med UUID i någon fil av något slag.

Jag tror att jag kommer att använda json eller liknande för spelare och "extra" funktioner, då det kan bli svårt att använda binary,
när man håller på med extra data och jag är inte säker på att jag optimera det tillräckligt för att det ska ge någon fördel.

## Att tänka på
Det är några saker som jag vill skriva upp, som jag kommer att behöva tänka på medans jag kodar, så jag ville bara skriva ner dem.

* Säkerheten
  * När jag genererar en ny chunk så måste jag kolla att spelaren är i närheten, så att någon inte gör något skript som generar en massa chunks.
    Möjligtvis att jag sätter någon slags limit på hur många chunks en värld kan ha.
* Flexibelt
  * Jag ska arbeta med services/microservices istället för en stor del, t.ex. att klient sidan och multiplayer kommunkationen/lagringen är separata
    så att jag kan göra nya klienter utan att ändra några större delar.
  * Block ID ska inte bli hard kodade, det ska finnas något register över vilka ID som finns och liknande, men sen är det upp till klienterna att implementera varje Block.
    Det vill säga, att servern ska inte hantera utdelningen av grafiken, den ska bara leverera IDs och data, sen hur datan hanteras är upp till varje klient.
    Möjligtvis att servern har någon slags version, så att klienten är tvungen att vara på rätt version.

## Slutsats
Tanken är att detta ska vara ett större projekt, men inte ett jätte stort projekt och att jag lägger mest energi på backend funktioner, så som lagring.
Än att hålla på med frontend/klienter och att grafiken ska bli snygg.

Så tanken är att i början, så ska det inte finnas mångra funktioner.
Det ska finnas några blocks att välja mellan, man ska kunna bygga och förstöra, multiplayer funktion.

Sen ska det också finnas en server funktion, där servern hanterar multiplayer funktioner så som, kommunkationen mellan klienterna.
Samt att den ska hantera lagring av data, så som block ID och liknande.
Det är också serverns uppgift att generara nya chunks.
Jag kommer nog inte lägga in någon avancerad världs generator, så den kommer nog bara generara t.ex. 2 lager med jord och resten sten eller liknande.

Det kan hända att jag lägger till någon slags registrerings funktion.
Så att man går till en webbsida och registrerar sig, jag är inte säker på om jag kommer att göra detta på en gång eller när jag håller på att bli klar.

Beroende på hur lång tid det tar att göra detta, så kanske jag lägger till fler funktioner.
* Så som fler blocks med extra data t.ex. kistor.
* Flera världar och att man kan flytta mellan dem.
* Ett chatt system (detta kommer nog bli en ganska hör prioritet efter alla huvud funktioner är färdiga)
* Fler klienter (mobil klient, desktop klient osv.)