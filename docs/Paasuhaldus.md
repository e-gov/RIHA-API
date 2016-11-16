# RIHA pääsuhaldus

spetsifikatsioon

v0.1, 16.11.2016

Sisukord

- [1 Ülevaade](#1-käsitlusala)
- [2 Mõisted ja lühendid](#2-mõisted-ja-lühendid)
- [3 Olulised viited](#3-olulised-viited)
- [4 Vajadus](#4-vajadus)
- [5 Lahendus]

## 1 Käsitlusala

Pääsuhalduse korraldamine RIHA kesksüsteemis. Ei käsitleta pääsuhalduse korraldust lokaal-RIHAdes.

## 2 Mõisted ja lühendid

| mõiste | seletus |
|--------|---------|
| _RIHA kesksüsteem_ | RIA taristus käitatav, andmebaasist, REST API-st ja mitmetest kasutajaliidesega moodulitest RIHA hajuslahenduse keskne komponent |
| _lokaal-RIHA_ | RIHA kesksüsteemist eraldi paigutatud, asutuse enda kontrolli all olevad, _kirjeldused_ või neid kirjeldusi välja andvad API-d; neid API-sid realiseerivad tarkvararakendused; eraldi paigaldatud RIHA kirjeldusmoodul vm moodulid (kui neid peaks tulevikus tekkima) |
| _kirjeldus_ | RIHA kirjeldusstandardile vastav riigi infosüsteemi komponendi (nt infosüsteemi, teenuse, klassifikaatori vm) masintöödeldav kirjeldus |

## 3 Olulised viited

[RIHA üldvaade](https://raw.githubusercontent.com/e-gov/RIHA-API/master/docs/YLDVAADE.md)

## 4 Vajadus

1. RIHA üldpõhimõte on teabe avatus. Siiski on vaja juurdepääsu piirata järgmises toimingutes:
  1. _kirjelduse_ koostamine, muutmine ja kustutamine
  2. _kooskõlastuse_ andmine
  3. piiratud juurdepääsuga teabe vaatamisel
  4. teiste kasutajate pääsuõiguste haldamisel (andmisel ja äravõtmisel)
  
## 5 Lahendus

Joonis 1 

![](RIHA-SSO.png)


  
