# Kirjeldusmoodul. Ülevaade

- Kirjeldusmoodul on RIHA üks nähtavamaid komponente. Mooduli võimalused:
  - inimkasutaja saab leida, vaadata, lisada, muuta ja hallata (ümber nimetada, kustutada) RIHA andmekirjeldusi
  -	kirjelduste lugemine RIHA andmebaasist ja sinna tagasi kirjutamine REST liidese abil
  - kirjelduste lugemine ja kirjutamine ka failist (faili)
  - kirjeldamise tüürimine mallide abil
  -	andmete vaatamine külalise režiimis; muud tegevused vastava õiguse olemasolul.
  
  - soovitavaks omaduseks on ka RIHA kesksüsteemist eraldi paigaldatavus.
 
 ## Hinnang tehtule
 
 - Tehnoloogia
   - valitud on aktuaalne, üheks parimaks peetav tehnoloogiapinu (_stack_)
     - React, _Javascripti-põhine UI raamistik_
     - Gulp, _ehitamise automatiseerimisvahend_
     - Babel, _Javascripti versiooniteisendaja_
   - tuleb arvestada, et valitud tehnoloogiad, sh Javascript arenevad väga kiirelt. Seetõttu vajab tarkvara kasutatavuse ja paindlikkuse tagamiseks iga-aastast hooldust
   - mallidega tüürimine on teostatud; mallid tehakse kahes järgus, kõigepealt Exceli tabelina; seejärel Perl rakendus tõlgib Excelist JSON-isse
     - mallid osutusid mahukaks (5000+ JSON rida); soovi korral saaks JSON malle otse redigeerida, jättes Exceli järgu ära
     - mallide muutmist (versioneerimist) ei ole katsetatud; mallidesse puutuv on ka ebapiisavalt dokumenteeritud
 
 ## Tööde plaan

Märkus. Tööde plaanis on ka töid (nt URL-de teema), mida võib liigitada ka teiste moodulite alla.

|     | Täitja  | Eeldused  |
|-----|---------|-----------|
| Dokumentatsiooni täiendamine         |    |   |
| - malli (JSON, ja Excel) struktuur vajab selget dokumenteerimist |  |  |
| Mallikäsitluse analüüs, projekteerimine, dokumenteerimine ja testimine - _Kuidas malli muutmise protsess käib? Mida see tähendab? Mida arvestada, et asjad ei läheks katki?_ |  |  |
| URL-de süsteemi väljatöötamine - _Vaja on süsteemi RIHA objektidele identifikaatorite omistamiseks. Kokku lepitud on, et identifikaatorid võiksid olla URL-id; millised, seda tuleb analüüsida ja leida lihtne süsteem_ | | |
| Avaldamine                             |    |   |
| - dokumentatsiooni avaldamine (GitHub) |    |   |
| - koodi avaldamine (GitHub)            |    |   |
| Eraldipaigaldatavuse saavutamine     |    |   |
| - analüüs ja projekteerimine          |    |   |
| - teostus                            |    |   |
| Kujunduse vahetamine                 |    | Uus kujunduslahendus |
 
 ## Arhitektuurijoonis
 
 <img src='Kirjeldusmoodul01.PNG' width='75%'>
 
