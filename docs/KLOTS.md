
## Absoluutne miinimum (MUST HAVE) / abivahendid (NICE TO HAVE)

RIHA funktsionaalsuste liigitus absoluutselt vajalikeks (MUST HAVE) ja abivahenditeks (NICE TO HAVE)

|  #   |  funktsionaalsus | M(ust have) või N(ice to have)? |  mida selleks on vaja?     |
|------|--------|---|-------|
| 1.1  | Infosüsteemi haldaja: kirjeldab ise oma infosüsteemi, teeb seda oma asutuses, on ise oma kirjelduse omanik | M |   |
| 1.2  | ...kasutades lihtsat kirjelduskeelt     | M | ➔ RIHA kirjeldusstandard  |
| 1.3  | ...kirjeldus on masintöödeldav  | M  | ➔ RIHA kirjeldusstandard   |
| 1.4  | ...samas ka inimloetav ja vajadusel -töödeldav | M | inim- ja masinloetavust ühitav vorming (nt JSON, YAML, Markdown) |
| 1.5  | ...kirjelduse võib malli alusel koostada tavalise tekstiredaktoriga | M | ➔ RIHA kirjeldusstandard |
| 1.6  | samamoodi kirjeldatakse klassifikaatoreid, e-teenuseid, XML-skeeme jm riigi infosüsteemi varasid | M | ➔ RIHA kirjeldusstandard |
| 2.1  | publitseerib ise oma kirjelduse(d), asutuse veebilehel vm asutuse kontrolli all olevas, internetile avatud keskkonnas | M | võimekus kirjeldusi avalikult üles panna või avalikku API-t luua |
| 2.2  | teatab RIHA kesksüsteemile oma kirjelduse asukoha | M | kirjelduse asukoha teatamise funktsionaalsus |
| 2.3  | kirjeldused on ülesleitavad | M | RIHA kesksüsteem peab nimekirja kirjelduste asukohtadest ➔ Lihtne kataloog |
| 3.1  | kirjeldused on kogumina kiiresti töödeldavad | M | RIHA kesksüsteem kogub kirjeldused regulaarselt kesksüsteemi kokku ➔ Koguja ("Andmekrabaja", "Reha") |
| 4.1  | ... ja teeb kõigile kättesaadavaks nii masinloetaval | M | ➔ RIHA kesksüsteemi REST API |
| 4.2  | ... kui ka inimloetaval kujul | M | ➔ Lihtne kuvamisvõimalus |
| 5.1  | kirjeldusi saab hinnata ja kooskõlastada | M | Kooskõlastaja annab hinnanguid ja kooskõlastusi RIHA kesksüsteemi kokkukogutud kirjeldustele ➔ Kooskõlastusmoodul |
| 6.1  | Infosüsteemi haldaja võib kirjeldamiseks kasutada tööriista | N | ➔ Kirjeldusmoodul |
| 6.2  | - kirjeldustööriista saab paigaldada oma asutusse | N | ➔ Kirjeldusmoodul |
| 6.3  | - kirjeldustööriista saab kasutada veebiteenusena | N | ➔ Kirjeldusmoodul | 
| 6.4  | Ise publitseerimise asemel saab kirjeldusi hoida RIHA kesksüsteemis | N | Kirjelduste RIHA kesksüsteemis hoidmise lahendus|
| 6.5  | Senised kirjeldused teisendatakse (pool)automaatselt uude vormingusse | N | Andmemigratsiooni lahendus |
| 6.6  | võib tööriista abil kontrollida kirjelduste kvaliteeti | N | ➔ Automaatkontrollide moodul | 
| 7.1  | võib oma töötajatele õiguste andmiseks kasutada RIHA kesksüsteemi pääsuhalduslahendust | N | ➔ RIHA pääsuhalduslahendus |   
| 8.1  | Kokkukogutud teabe kasutaja saab teavet visualiseeritult | N | ➔ Visualiseerimismoodul | 
| 8.2  | saab koondandmeid, statistikat, aruandeid | N | ➔ Avaandmete moodul |
| 9.1  | 	RIHA kirjeldusstandardi semantika on selge | N | ➔ Riigi infosüsteemi mõistete süsteem; Sõnastike moodul |
| 10.1 | RIHA toimib platvormina mitmesugustele võimalikele riigi infosüsteemi haldusprotsessidele, nt X-teega liitumine | N | ➔ X-tee liitumismoodul |
| 10.2 | RIHAst leiab XML jms varade spetsialiseeritud mooduli | N | ➔ XML-varade moodul |


