# RIHA andmete masinloetavate vormingute põhimõtted

versioon 0.9

Tellija: Riigi Infosüsteemi Amet

Täitja: OÜ Girf, OÜ Degeetia, OÜ Mindstone
 
##Sisukord
- Ülevaade	1
- Sihtgrupp	2
- Seotud dokumendid	2
- Mõisted	3
- Andmekogu/andmete tehniline automaatkirjeldamine välise osapoole poolt	3
- Tähenduste kodeerimine andmebaasi ja teenuste kommentaarides	4
- Andmekogu/andmete esitamine välisele osapoolele	5
- RIHA sisemine andmevahetus põhiobjektide osas: eksport ja import.	5
- Jooksev andmevahetus RIHA moodulite vahel	5

## Muutelugu
| Versioon, kuupäev	| Muudatus | 
|-------------------|----------|
| 0.9, 26.09.2016   | Täitja üleantud dokument viidud Markdowni, omistatud versiooninumber 0.9 |
|-------------------|----------|

##Ülevaade
RIHA andmete masinloetava vormingu abil kirjeldatakse nii RIHAsse masintöödeldaval kujul esitatav info kui RIHA enda poolt välistele süsteemidele esitatav info kui ka RIHA enda sisemine andmevahetus. Info esitamisel on seejuures neli peamist eesmärki, millel kõigil on oma spetsiifika ja erisused:
•	Andmekogu/klassifikaatori/sõnastiku vms põhiobjekti kirjeldamine RIHA jaoks nii, et RIHA veebilehe kaudu ei ole vaja midagi sisestada, vaid RIHA loeb kirjeldust regulaarselt ise andmekogu jaoks etteantud veebilehelt (import).
•	Andmekogu/klassifikaatori/sõnastiku vms põhiobjekti jaoks RIHAs olemasoleva informatsiooni eksport, eeskätt võimaldamaks kergemini esitada infot RIHA impordi jaoks sobival kujul.
•	RIHA sisemine andmevahetus põhiobjektide osas: eksport ja import.
•	Jooksev andmevahetus RIHA moodulite vahel, eeskätt kasutajaliidese ja andmehoidla vahel: RIHA API.

Kuivõrd põhiobjekte kirjeldavate väljade hulk RIHA arenduse ja edasise kasutuse käigus regulaarselt muutub, siis ei fikseeri me siin põhimõtte-dokumendis kogu konkreetset väljahulka – seda tehakse lisatud näidetes ja regulaarselt uuendatavates eraldi dokumentides – vaid andmete esituse põhimõtteid.

## Sihtgrupp

Dokument on mõeldud RIA arhitektidele, RIHA töögrupi liikmetele ja uue RIHA arendajatele. Dokument ei ole mõeldud välistele osapooltele.

Dokument võib harvadel juhtudel olla abiks andmekogu haldajale, kuid neile on käesoleva dokumendi asemel ette nähtud eraldi detaildokument “RIHA andmete masinloetavate vormingute spetsifikatsioon.docx” RIHA arendajatele on oluline veel täiendav dokument „RIHA API spetsifikatsioon.docx”

## Seotud dokumendid

Vajalik käesoleva dokumendi mõistmiseks:
•	„RIHA andmebaasi kontseptuaalne mudel.docx“ selgitab andmebaasi struktuuri põhimõtteid koos näidetega.

Masinloetavate andmete kirjeldamise ja komponentide vahelise andmevahetuse API konkreetsed spetsifikatsioonid kirjeldavad kahte erinevat aspekti ja RIHA andmevahetuses ning on mõeldud erinevatele sihtgruppidele:
•	“RIHA andmete masinloetavate vormingute spetsifikatsioon.docx” kirjeldab eeskätt andmevahetust väliste süsteemidega.
•	“RIHA API spetsifikatsioon.docx” kirjeldab eeskätt andmevahetust RIHA komponentide vahel.
•	“riha_andmedetailid.xlsx” on detailne JSON-struktuuri tasemel põhitabelite kirjeldus eeskätt RIHA arendajatele.
•	"riha_vorming_spetsifikatsioon.js" on eelmises punktis toodud exceli faili baasilt koostatud JSON-struktuuri täismahulise ülesehitusega näide, kus on kommentaaridena iga välja juures kirjas andmetüüp, nimetus, väärtuse piirangud jms.
•	"riha_vorming_json-schema.js" on failile "riha_andmedetailid.xlsx" vastav täpne JSON-struktuuri spetsifikatsioon esitatuna JSON-SCHEMA standardi kohaselt (vt http://json-schema.org/)

Tehnilised lisad masinloetavate andmete kirjeldamise konkreetsele spetsifikatsioonile: 
•	infosystem.js – infosüsteemi kirjelduse näide
•	area.js - valdkonna kirjelduse näide
•	vocabulary.js - valdkonna sõnastiku kirjelduse näide
•	classifier.js - klassifikaatori sõnastiku kirjelduse näide
•	xmlasset.js - XML vara kirjelduse näide

Dokumendi mõistmiseks on vaja omada arusaamist SQL süntaksist ja JSON formaadist: konkreetseid dokumente nende standardtehnoloogiate jaoks siinkohal ei esitata.
•	Kasutatav JSON standard: JSON Data Interchange Format (ECMA-404, http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf). 

## Mõisted
- RIHA: vaata https://riha.eesti.ee/
- JSON: vaata https://en.wikipedia.org/wiki/JSON ja lisaks eelmist peatükki
- SQL: vaata https://en.wikipedia.org/wiki/SQL
- WSDL: vaata https://en.wikipedia.org/wiki/Web_Services_Description_Language

## Andmekogu/andmete tehniline automaatkirjeldamine välise osapoole poolt
RIHA põhiandmestikuks on :
•	infosüsteem koos tema andmebaaside ja teenuste detailkirjeldusega,
•	Valdkonnad
•	Valdkonnasõnastikud,
•	Klassifikaatorid,
•	XML varad.

Üldjuhul on tegemist hierarhiliste struktuuridega, mis seovad andmeid mitmetest erinevatest tabelitest. Automaatkirjeldamise juures integreeritakse need erinevatesse tabelitesse jaotuvad andmed üheks hierarhiliseks JSON-failiks, mida võib omakorda jagada erinevateks URL-ga viidatud veebilehtedeks.

Igaühe jaoks neist neljast luuakse asutuse poolt eraldi JSON-veebileht. Kui asutuse poolt mõnda neist tüüpidest (näiteks valdkonnasõnastikke või XML varasid) ei hallata, siis vastavat JSON-lehte ka ei looda.
RIHA andmed esitatakse üldjuhul JSON objektidena, kuid täiendavalt kasutatakse ka muid formaate:
•	SQL soovitava alternatiivina JSON-ile tabelite struktuuri esitamisel,
•	WSDL SOAP-teenuste struktuuri esitamisel.
Seejuures kasutatakse nimelt nö harilikku JSON-i, mitte JSON-i põhiseid eriformaate nagu JSON-LD (vaata https://en.wikipedia.org/wiki/JSON-LD ja https://www.w3.org/TR/json-ld/ ja http://json-ld.org/ ).
JSON-LD lisab „harilikule“ JSON-ile mõned eritähendusega võtmed nagu näiteks @context, @id ja @type, mis hõlbustab JSON-formaadis esitatud andmete konverteerimist RDF kujule. Näide Wikipediast:

```json
{
   "@context": {
    "name": "http://xmlns.com/foaf/0.1/name",
    "homepage": {
      "@id": "http://xmlns.com/foaf/0.1/workplaceHomepage",
      "@type": "@id"
    },
    "Person": "http://xmlns.com/foaf/0.1/Person"
  },
  "@id": "http://addcolor.it",
  "@type": "Person",
  "name": "Olgiati Davide Stefano",
  "homepage": "http://www.addcolor.it/"
}
```

Olulise põhimõttena näeb JSON-LD ette @context sisu esitamise JSON tekstist sõltumatu, viidatava failina, tsitaat JSON-LD standardist: „JSON documents can be interpreted as JSON-LD without having to be modified by referencing a context via an HTTP Link Header as described in section 6.8 Interpreting JSON as JSON-LD. It is also possible to apply a custom context using the JSON-LD API [JSON-LD-API].“

Seega on soovi korral võimalik edaspidi käsitleda RIHA andmete JSON-kirjeldusi JSON-LD võtmes, kui töötada välja vastav viidatav fail ja lisada lehe päisesse vastav viide.

RIHA projekti kontekstis ei näe me RDF kujule konverteerimise hõlbustamises eelist: kui tekib soov konverteerida RIHA JSON-kuju RDF kujule (mida me ei pea kuigi tõenäoliseks), tuleb selleks igal juhul esitada vastav konverter, mida on samas küllalt lihtne teha ka ilma JSON-LD kasutuselevõtuta.

Negatiivse poole pealt lisaks JSON-LD nõue RIHA formaadi kasutajatele täiendava keerukusena ülesande nimetatud formaadist ja selle eripäradest aru saada, mis ei ole meie hinnangul põhjendatud.

## Tähenduste kodeerimine andmebaasi ja teenuste kommentaarides

Tähenduste kodeerimise spetsifikatsioon ja näited on detailselt esitatud dokumendis “RIHA andmete masinloetavate vormingute spetsifikatsioon.docx”.

Põhimõttena kodeeritakse nii  funktsioonid/eesmärgid, andmekoosseisu andmeobjektid, klassifikaatorid kui terminid sõnastikus tagidena andmetabeli, andmevälja või teenuse sisendi/väljundi vabatekstilistes kommentaarides vastavalt $, !, # või @-prefiksiga tekstina, näiteks “klientide tabel, sisaldab #isik !klient kasutab  classifiers:standardklassif”. 

See konkreetne esitus on sõltumatu klassifikaatorite või sõnastike haldamise ja esitamise viisidest.

Sõnastike osas soovitame edaspidi põhisõnastikuna kasutada http://schema.org sõnastikku tõlkimata kujul. Seda sõnastikku on soovitav laiendada Eesti sõnastikega vastavalt schema.org laiendamispõhimõtetele ja lähtudes samadest põhimõtetest. 

Schema.org (vaata https://en.wikipedia.org/wiki/Schema.org ja  http://schema.org/ ja ) on kõige laiema toega ja levikuga tehniline universaalsõnastik, mis on samas avatud. Tsitaat schema.org esilehelt: “Schema.org is sponsored by Google, Microsoft, Yahoo and Yandex. The vocabularies are developed by an open community process, using the public-schemaorg@w3.org mailing list and through GitHub.”

Klassifikaatorite, sõnastike ja XML-varade haldamise ja rakendamise täpsemad põhimõtted tuleks edasistes etappides eraldi põhjalikult läbi töötada.

## Andmekogu/andmete esitamine välisele osapoolele

Selleks luuakse RIHA-sse vastav URL, mille taga olev rakendus kuvab esitatud URI-ga objekti võimalike lisaparameetritega määratud formaadivariandis (üldjuhul siis JSON, SQL või hoopis mõni konkreetne fail).

Selliselt esitatud objekti saab väline osapool kasutada ka oma objekti kirjelduse näitena.

## RIHA sisemine andmevahetus põhiobjektide osas: eksport ja import.

Selleks luuakse RIHA-sse vastav URL, mille taga olev rakendus kuvab esitatud URI-ga objekti võimalike lisaparameetritega määratud formaadivariandis (üldjuhul siis JSON, SQL või hoopis mõni konkreetne fail), kuid erineb kuigivõrd väliseks kasutuseks ettenähtud variandist:
•	Lisab kõik tehnilised väljad, kaasa arvatud versioonitundliku id välja, mida väliseks infovahetuseks üldjuhul ei kasutata.
•	Ei teosta kodeeringu lihtsustusi/teisendusi.
•	Ei esita andmeid tingimata sügavate hierarhiatena.
•	Võimaldab esitada ka neid andmeid, mida väliseks kasutuseks ei esitatagi.

Konkreetne vahe sisemise import/ekspordi vahel ning järgmisena kirjeldatud API vahel realiseerub eksport/impordi ja API realiseerimise käigus: detailseid nüansse kui välistesse osapooltele mittepuutuvaid ei ole hetkel otstarbekas spetsifitseerida.

## Jooksev andmevahetus RIHA moodulite vahel

Jooksev andmevahetus toimub üldjuhul tabeli valitud ridade kaupa, mitte aga terviklike hierarhiliste struktuuride kaupa. Samuti peab olema võimaldatud objektide osakaupa muutmine, kustutamine jne.

Selleks realiseeritakse RIHA-s API.

API funktsioneerib REST põhimõtete alusel, kasutades andmeformaadiks JSON-it, ning lisab nö traditsioonilistele REST parameetritele erinevaid praktilisi laiendusi. 

API kaudu esitatav andmete esitus kasutab samuti JSON-i ning neidsamu väljanimesid, mis on kasutuses eelkirjeldatud avalikes andmekirjeldustes ja on samased andmebaasis kasutatavate väljanimedega.

API on kirjeldatud eraldi dokumendis „RIHA API spetsifikatsioon.docx“
