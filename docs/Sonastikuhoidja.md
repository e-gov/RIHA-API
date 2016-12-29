# "Sõnastikuhoidja"

Kontrollnimistu

[x] Kas rakendusel on üks selge funktsioon?
[x] Kas rakendus pakub kõiki oma andmeid REST API kaudu?
  [x] Kas REST API on formaalselt kirjeldatud (Swagger/Open API YAML)
[x] Kas rakenduse lähtekood on avalik?
[x] Kas lähtekoodile rakendatakse versioonihaldust?
[x] Kas pääsuhalduseks kasutatakse mõnd taristut v välist süsteemi
[x] Kas kasutajad saavad teatada vigu ja teha ettepanekuid?
  [x] Kas vigade ja ettepanekute teave on avalik?
  [x] Kas ettepanekuid saab kommenteerida?
[x] Kas rakenduse ja selle andmete varundamine on lahendatud?  

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

