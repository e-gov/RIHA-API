# "Sõnastikuhoidja"

Arhitektuuri kontrollnimistu

- [x] 1. Kas rakendusel on üks, selge funktsioon?
- [x] 2. Kas rakendus pakub kõiki oma andmeid REST API kaudu?
  - [x] 3. Kas REST API on formaalselt kirjeldatud (Swagger/Open API YAML)
- [x] 4. Kas rakenduse lähtekood on avalik?
- [x] 5. Kas rakenduse dokumentatsioon on avalik?
- [x] 6. Kas lähtekoodile rakendatakse versioonihaldust?
- [x] 7. Kas pääsuhalduseks kasutatakse taristulist identiteedi- ja pääsulahendust v välist autentimis- ja autoriseerimissüsteemi?
- [x] 8. Kas kasutajad saavad teatada vigu ja teha ettepanekuid?
  - [x] 9. Kas vigade ja ettepanekute teave on avalik?
  - [x] 10. Kas ettepanekuid saab kommenteerida?
- [x] 11. Kas rakenduse ja selle andmete varundamine on lahendatud?
- [x] 12. Kas rakendus on nii väike, selle saab vajadusel mõistliku kuluga asendada funktsionaalselt samaväärse, teisel tehnoloogial teostatud rakendusega (_Throwaway_ põhimõte)?

## Põhifunktsionaalsus

1 Sõnastiku täiendamine
  - GitHub
    - GitHub-i redigeerimis-UI (inimkasutaja liides)
      - CRUD (sõnade lisamine, muutmine, eemaldamine)

2 Sõnastiku hoidmine
  - GitHub
    - sõnad hoitakse YAML, JSON v Markdown failidena, kaustadesse jagatult

3  Sõnastiku esitamine inimloetaval kujul
  - GitHub Pages, kasutades:
    - Jekyll
      - Liquid
    - HTML5
    - CSS
    - Javascript
    - jQuery
    - Bootstrap
    - Material Design Lite
    - vajadusel muid Javascripti teeke

4 Sõnastiku esitamine APi kaudu
  - GET sõnade loetelu
  - GET üksiku sõna kirjeldus *pole selge, kas ja kuidas on teostatav
  - API formaalne kirjeldus
    - Swagger 2.0 (Open API Initiative)
      - YAML süntaks

## Täiendav funktsionaalsus

5 Pääsuhaldus
  - GitHubi pääsuhaldus

6 Vearaporteerimine ja ettepanekute haldus
  - GitHub issues

7 Versioneerimine
  - Git (GitHub)
  
8 Arendusdokumentatsioon
  - GitHubi repo

