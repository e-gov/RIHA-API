# RIHA andmete masinloetavate vormingute spetsifikatsioon

Versioon 2.0, 29.04.2016

Tellija: Riigi Infosüsteemi Amet; Täitja: Girf OÜ, Degeetia OÜ, Mindstone OÜ

## Sisukord
- [Muutelugu](#muutelugu)
- [Ülevaade](#ulevaade)
- [Seotud dokumendid](#seotud dokumendid)
- [Mõisted](#moisted)
- [Andmete automaatesitamine RIHA jaoks](#andmete-automaatesitamine-riha-jaoks)
- [Andmete tehnilise esituse saamine RIHA-st](#andmete-tehnilise-esituse-saamine-riha-st)
- [Põhi-JSON-veebileht](#pohi-json-veebileht)
- [Objektide identifitseerimine URI-de abil](#objektide-identifitseerimine-uri-de-abil)
- [Seotud dokumendid: inimloetavad failid](#seotud-dokumendid-inimloetavad-failid)
- [Andmevormingu spetsifikatsioonid](#andmevormingu spetsifikatsioonid)

## Ülevaade
Selles dokumendis kirjeldame, kuidas esitada RIHA-sse automaatselt loetaval ja masintöödeldaval kujul informatsiooni infosüsteemide, teenuste jms kohta ning kuidas RIHA sama informatsiooni välja annab.

Dokument selgitab info esitamise struktuuri ja põhimõtteid ning toob seejärel spetsifikatsioonina konkreetsed, detailsete kommentaaridega varustatud näited.

Dokument on mõeldud eelkõige välistele osapooltele, kes RIHAsse andmeid sisestavad ja sealt loevad, eeskätt aga nende osapoolte IT-arendajatele.

Samuti on dokumenti mõeldud RIA arhitektidele, RIHA töögrupi liikmetele ja uue RIHA arendajatele.

## Seotud dokumendid
RIHA välimist ja sisemist infovahetust kirjeldab üldisem dokument, mis ei ole käesoleva dokumendi mõistmiseks aga otseselt vajalik:
•	"RIHA andmete masinloetavate vormingute põhimõtted.docx"

Tehnilised lisad masinloetavate andmete kirjeldamise konkreetsele spetsifikatsioonile:
•	infosystem.js – infosüsteemi kirjelduse import-ekspordifaili näide,
•	service.js – teenuse kirjelduse import-ekspordifaili näide,
•	document.js – dokumendi ekspordifaili näide,
•	comment.js – kommentaari ekspordifaili näide.

Dokumendi teise poole mõistmiseks on vaja omada arusaamist SQL süntaksist ja JSON formaadist: konkreetseid dokumente nende standardtehnoloogiate jaoks siinkohal ei esitata.
•	Kasutatav JSON standard: JSON Data Interchange Format (ECMA-404, http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf).

## Mõisted
•	RIHA: vaata https://riha.eesti.ee/
•	JSON: vaata https://en.wikipedia.org/wiki/JSON ja lisaks eelmist peatükki
•	SQL: vaata https://en.wikipedia.org/wiki/SQL
•	WSDL: vaata https://en.wikipedia.org/wiki/Web\_Services\_Description\_Language
•	URI: vaata https://en.wikipedia.org/wiki/SQL
•	URN: vaata https://en.wikipedia.org/wiki/URN

## Andmete automaatesitamine RIHA jaoks
RIHA-sse saab hetkel esitada nelja sorti infot, millest olulisem on esimene:
•	Infosüsteem koos tema andmebaaside ja teenuste detailkirjeldusega,
•	Valdkonnad
•	Valdkonnasõnastikud,
•	Klassifikaatorid,
•	XML varad.

Iga sellise objekti jaoks (kui asutus vastavat sorti objekti haldab) luuakse asutuse poolt eraldi JSON-veebileht. Kui asutuse poolt mõnda neist tüüpidest (näiteks valdkonnasõnastikke või XML varasid) ei hallata, siis vastavat JSON-lehte ka ei looda.
RIHA andmed esitatakse üldjuhul JSON objektidena, kuid täiendavalt kasutatakse ka muid formaate:
•	SQL lihtsama alternatiivina JSON-ile andmebaasi tabelite struktuuri esitamisel,
•	WSDL SOAP-teenuste struktuuri esitamisel.

Andmete esitamiseks tuleb:
•	luua asutuse vastavat tüüpi objekti põhi-JSON-veebileht, mille sisuks on tehniline JSON esitus, mis sisaldab kogu vajalikku põhi-informatsiooni,
•	luua võimalikud täiendavad veebilehed / dokumendid, millele põhi-veebilehest viidatakse,
•	paigaldada kõik need lehed / failid veebiserverisse kättesaadavaks
•	kui osadele lehtedele / failidele on ligipääs piiratud, siis - vajadusel - luua RIHA jaoks veebilehtedele/failidele ligipääsuõigused kasutajatunnuse ja parooli abil,
•	sisestada põhi-veebilehe URL (ja vajadusel kasutajatunnus ja parool) RIHAs andmekogu/sõnastike/klassifikaatorite/XML varade kirjelduse vastavatele väljadele.

Seepeale hakkab RIHA rakendus regulaarselt automaatselt lugema veebilehte, temast viidatavaid veebilehti / faile ja nende abil automaatselt uuendama RIHAs andmestiku kirjeldust.

Asutus ise omal initsiatiivil RIHA-sse andmeid ei postita. Küll aga näeme ette, et
•	RIHAs realiseeritakse testsüsteem, mille kaudu andmekogu arendaja saab ise saata JSON-kirjeldusi nende õigsuse automaatkontrolliks.
•	JSON-veebilehe lugemise ja töötlemise järel salvestab RIHA tulemi korrektsuse kinnituse või veateate vastava infosüsteemi või muu andmestiku juurde kõigile osapooltele RIHA kasutajaliidesest leitaval kujul.

## Andmete tehnilise esituse saamine RIHA-st
Kõiki eelmainitud sorti andmeid saab RIHA-st vabalt tõmmata, kui neile ei ole seatud piirangut, mis nõuab eelnevalt autentimist.
Andmete tõmbamiseks on RIHAS üks konkreetne URL (ei ole hetkel veel määratud), millele antakse objekti URI parameetrina ette ja mis väljastab sisestamiseks sobivale kujule analoogilise JSON-struktuuri.
Lisaks objekti URI-le saab nimetatud URL-le anda ette parameetreid soovitud formaadi valikuks, näiteks SQL või JSON, hierarhia sügavus jms. Nimetatud parameetrid ei ole hetkel veel määratud.

## Põhi-JSON-veebileht
Üldjuhul tuleks automaatkirjelduse JSON-veebilehtede ja muude tehniliste failide koostamine anda andmekogu arendaja ülesandeks. Põhi-JSON-veebilehe koostamine eeldab seejuures koostööd arendaja ja andmekogu ärilise haldaja vahel: viimane teab teatud hulka informatsiooni, mida arendaja ei tea.

Konkreetne põhi-JSON-veebilehe koostamine ja serverisse kättesaadavaks tegemine on igal juhul arendaja ülesanne.
Põhi-JSON-veebilehe muutmiseks on lihtne realiseerida ka eraldi väike rakendus, mis võimaldab lehe välju lihtsalt vormide kaudu muuta. RIHA realisatsiooni käigus luuaksegi taoline standardlahendus, mida asutus võib kasutada kas keskse RIHA koosseisus või installeerida oma infosüsteemi.

Põhi-JSON-veebilehe struktuur esitakse lihtsalt JSON objektina kujul, kus kasutatakse samu väljanimesid, mis on kasutusel RIHA baasis ja kus väärtused on omakorda JSON objektid: lihtsamatel juhtudel tekstid, keerukamatel sisemise struktuuriga objektid. Väljanimedest arusaamiseks ja nende täitmise võimaldamiseks luuakse eraldi juhend ja annoteeritud JSON-template ja näited.

Väljad jaotuvad kolme gruppi:
•	Otsese väärtusega väljad, mis kirjeldavad infosüsteemi/teenust/jne ja mis võivad olla lihtsalt tekstilised väärtused või väiksemad loendid/objektid.
•	Erinevad seotud dokumendid, mis reeglina esitatakse URLi või URLe sisaldavate objektide kujul, mis viitavad veebist kättesaadavatele dokumentidele.
•	Erinevad seotud detailkirjelduste loetelud, mida võib esitada nii* ◦◦konkreetse sügava struktuurina kui
•	* ◦◦URLide loeteluna, mis viitavad objekti (näiteks andmebaas või teenus) tehnilisele kirjeldusele.

Andmebaaside ja teenuste loetelus võib kasutada ka otseseid väärtusi JSON kujul (vaata allpool detailkirjeldust), kuid see on suure mahu tõttu üldjuhul ebapraktiline.

### Objektide identifitseerimine URI-de abil

Infosüsteemidel, teenustel, sõnastikel ja klassifikaatoritel on RIHA-s unikaalne, versioonist sõltumatu identifikaatori tekstilise URI kujul. URI on info haldaja valikul kas:
•	Harilik URL (näiteks http://amet.ee/systeemX/) mis kirjeldab antud süsteemi või klassifikaatorit vms, mitte ei ole näiteks asutuse enda URL
•	Või URN kujul "urn:fdc:riha.eesti.ee:2016:infosystem:lyhikood" infosüsteemide jaoks (klassifikaatorite, sõnastike jms jaoks kasuta "classifier", "dictionary" jms) kus alguses on alati "urn:fdc:riha.eesti.ee:aastanr" kus aastanr on selle konkreetse URN-i esimese kasutuselevõtu aasta.

### Seotud dokumendid: inimloetavad failid
Põhi-veebileht võib viidata ametlikele ja seletavatele dokumentidele, mis on lihtsalt näiteks PDF või wordi formaadis, ning mida RIHA süsteem automaatselt ei töötle. Need tuleb lihtsalt veebiserverisse paigutada ja lisada põhi-JSON-lehel URL-d.

## Andmevormingu spetsifikatsioonid
Infosüsteemi objekti andmevormingu näited on lisatud eraldi dokumentidena, mis on järjest toodud ka siinse dokumendi lõpus:
•	js – infosüsteemi kirjelduse näide
•	js - valdkonna kirjelduse näide
•	js - valdkonnasõnastiku kirjelduse näide
•	js - klassifikaatori kirjelduse näide
•	js - XML vara kirjelduse näide

### Infosüsteemi detailkirjeldus
Esitame infosüsteemi objekti formaadi spetsifikatsiooni näitena, kuhu on lisatud spetsifitseerivad kommentaarid (viimased ei ole JSON struktuuri osad):

```json
/* object(infosystem): Selles struktuuris hoitakse infosüsteemi põhiinfot.
Infosüsteemid hoitakse tabelis versioonidena, iga infosüsteemi versioon on eraldi kirje. Sealjuures lähtutakse põhimõttest, et kõige esimene tekitatud infosüsteemi kirje kajastab alati infosüsteemi jooksvat hetkeseisu ning hiljem tekitatud kirjed on infosüsteemi kirjelduse vanemad versioonid. */
"infosystem": {
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
    "uri": "urn:fdc:riha.eesti.ee:2016:infosystem:riha",
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Infosüsteemi inimloetav nimi */
    "name": "Riigi infosüsteemi haldussüsteem",
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Infosüsteemi omanik */
    "owner": "70006317",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi lühinimi */
    "short_name": "riha",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimloetav versiooni nr, ei kasutata identifikaatorina */
    "version": "1.2",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Alaminfosüsteemi korral üleminfosüsteemi URI */
    "parent_uri": "string",
    /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
    "access_restriction": 0,
    /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Infosüsteemi kirjelduse olek */
    "state": "C",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
    "start_date": "1998-11-01T00:00:00",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
    "end_date": "2010-11-01T00:00:00",
    /* Andmetüüp: array(object(translation)), Piirangud: -, Kirjeldus: Võõrkeelsed nimed */
    "names": [{"en":"Management system of state information system"}],
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi kirjeldus. Lühike vabatekst */
    "description": "RIHA on riigi infosüsteemi kataloog ning ühtlasi keskkond koosvõime, andmekaitse jmt nõuete tagamise menetlustele. RIHA on X-tee, klassifikaatorite süsteemi ja eesti.ee tugisüsteemiks. RIHA on riigi infosüsteemi juhtimise vahend. RIHA on Euroopa mastaabis unikaalne, rahvusvahelist tunnustust võitnud süsteem.
ISKE audit viidi läbi 13. märts 2015  3. juuni 2015.ISKE on rakendatud, auditeeritud märkustega või soovitustega. Tuvastati keskmiselt olulisi puuduseid ISKE rakendamisel ja anti soovitusi puuduste likvideerimiseks.",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi asutamise eesmärk */
    "purpose": "Riigi infosüsteemi haldussüsteemi eesmärk on riigi infosüsteemi haldamise läbipaistvuse tagamine, riigi infohalduse planeerimine ning riigi, kohaliku omavalitsuse ja avalikke ülesandeid täitvate eraõiguslike isikute andmekogude koosvõime toetamine ning andmekogude nõuetele vastavuse kontrollfunktsiooni võimaldamine.",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi ülesanne */
    "task": "RIHA ülesanne on pakkuda platvormi infosüsteemide ja andmekogude asutamiseks, riigi infosüsteemi nõuete kontrollimiseks, teenuste registreerimiseks, X-teega liitumiseks, klassifikaatorite haldamiseks, XML skeemide haldamiseks ja muudele riigi infosüsteemi järjepidevaks arenguks ja optimaalseks toimimiseks vajalikele protsessidele. RIHA ülesanne on olla riigi infosüsteemi terviklik ja üksikasjalik kaardistus.",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Infosüsteemi kasutusele võtmise kuupäev */
    "take_into_use_date": "1998-11-01T00:00:00",
    /* Andmetüüp: string, Piirangud: enum(infosystem_type), Kirjeldus: Infosüsteemi tüüp */
    "infosystem_type": "kindlustav_systeem",
    /* Andmetüüp: string, Piirangud: enum(infosystem_category), Kirjeldus: Infosüsteemi kategooria */
    "infosystem_category": "seadus_ylesanded",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Infosüsteemi lõpetamise kuupäev */
    "finishing_date": "2016-01-01T00:00:00",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi lõpetamise kirjeldus */
    "finishing_description": "string",
    /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas sisaldab isikuandmeid? */
    "personal_data": True,
    /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas sisaldab delikaatseid isikuandmeid? */
    "sensitive_personal_data": False,
    /* Andmetüüp: string, Piirangud: enum(manner_of_archive_keeping), Kirjeldus: Arhiivi pidamise viis */
    "manner_of_archive_keeping": "elektrooniline",
    /* Andmetüüp: number, Piirangud: -, Kirjeldus: Vanimate andmete aeg, aastaarv (mis ajast on kõige vanemad andmed kogutud) */
    "older_data": 1998,
    /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas andmete säilitamisaeg on alaline? */
    "data_retention_period_eternal": False,
    /* Andmetüüp: integer, Piirangud: -, Kirjeldus: Andmete säilitamise aeg aastates */
    "data_retention_period": 25,
    /* Andmetüüp: string, Piirangud: enum(regularity_of_updating), Kirjeldus: Andmete uuendamise regulaarsus */
    "regularity_of_updating": "pidev",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kasutatavate standardite kirjeldus */
    "standards": "string",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Andmehõive protsessi kirjeldus */
    "data_acquisition_process": "string",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi asukoht internetis (URL) */
    "url": "https://riha.eesti.ee/",
    /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas tegemist on standardlahendusega? Jah/ei */
    "is_standard_solution": False,
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Erisused standardlahendusest */
    "differences_from_standard_solution": "string",
    /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas tegemist on eeskujulikult kirjeldatud infosüsteemiga? Jah/ei */
    "excellent": true,
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi viitenumber */
    "reference_number": "IS000161",
    /* Andmetüüp: array(object(linked_infosystem)), Piirangud: -, Kirjeldus: Seosed teiste infosüsteemidega */
    "linked_infosystems": [{
        /* Andmetüüp: string, Piirangud: enum(infosystem_link_type), Kirjeldus: seose tüüp */
        "type": "string",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: seotud infosüteemi uri */
        "uri": "string",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seose kirjeldus */
        "description": "string"
    }],
    /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Infosüsteemi valdkonnad */
    "areas": ["Infoühiskond","Riigi infosüsteem","Kindlustavad süsteemid"],
    /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Infosüsteemi grupid */
    "groups": ["Majandus- ja Kommunikatsiooniministeeriumi valitsemisala"],
    /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Kasutatavad klassifikaatorid */
    "classifiers": ["EHAK 2015v1"],
    /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Semantika. Terminite tekstiline list, mis esitavad veergude/välja tähendust (inimene, auto, ...) */
    "semantics": ["string"],
    /* Andmetüüp: array(object(linked_organization)), Piirangud: -, Kirjeldus: Infosüsteemi vastutavad ja volitatud töötlejad */
    "organizations": [{"organization":"70006317","type":"vastutav_tootleja"}],
    /* Andmetüüp: array(object(linked_person)), Piirangud: -, Kirjeldus: Infosüsteemi kontaktisikute info */
    "contacts": [{
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud isiku isikukood */
        "person": "string",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse alguskuupäev */
        "date_from": "2016-01-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse lõppkuupäev */
        "date_to": "2016-01-01T00:00:00",
        /* Andmetüüp: string, Piirangud: enum(person_link_type), Kirjeldus: Seose tüüp */
        "type": "string",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Ametikoht */
        "occupation": "string",
        /* Andmetüüp: string, Piirangud: enum(person_contact_level), Kirjeldus: Kontakti tase */
        "contact_level": "string",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
        "description": "string"
    }],
    /* Andmetüüp: object(linked_person), Piirangud: -, Kirjeldus: Isik, kes vastutab isikuandmete töötlemise eest */
    "responsible_for_personal_data": {
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud isiku isikukood */
        "person": "string",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse alguskuupäev */
        "date_from": "2016-01-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse lõppkuupäev */
        "date_to": "2016-01-01T00:00:00",
        /* Andmetüüp: string, Piirangud: enum(person_link_type), Kirjeldus: Seose tüüp */
        "type": "string",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Ametikoht */
        "occupation": "string",
        /* Andmetüüp: string, Piirangud: enum(person_contact_level), Kirjeldus: Kontakti tase */
        "contact_level": "string",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
        "description": "string"
    },
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: X-teega liitumise kuupäev */
    "xroad_join_date": "2005-01-03T00:00:00",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: X-tee lõpetamise kuupäev */
    "xroad_finishing_date": "2016-01-01T00:00:00",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: X-tee lõpetamise kirjeldus */
    "xroad_finishing_description": "string",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: X-tee sertifikaadi kehtivuse lõppkuupäev */
    "xroad_sertificate_validity_deadline": "2016-01-01T00:00:00",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: IKSE turvaklass */
    "iske_security_class": "K2T2S1",
    /* Andmetüüp: string, Piirangud: enum(iske_audit_status), Kirjeldus: ISKE rakendamise staatus */
    "status_of_applying_iske": "audit_markustega",
    /* Andmetüüp: array(object(iske_audit)), Piirangud: -, Kirjeldus: ISKE auditite andmed */
    "iske_audit": [{
        /* Andmetüüp: string, Piirangud: enum(iske_audit_status), Kirjeldus: ISKE auditi staatus */
        "iske_audit_status": "string",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: ISKE auditi kuupäev */
        "iske_audit_date": "2016-01-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: ISKE auditi tähtaeg */
        "iske_audit_deadline": "2016-01-01T00:00:00",
        /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: ISKE auditi sisu/kokkuvõte */
        "document": [{
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
            "url": "http://asutus.ee/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
            "name": "Määrus X",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
            "filename": "/opt/documents/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
            "mime": "application/word",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
            "state": "C",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
            "start_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
            "end_date": "2010-11-01T00:00:00",
            /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
            "content": "YmFzZTY0Cg==",
            /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
            "metadata": "object"
        }],
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas ISKE audit on avalik või varjatud */
        "access_restriction": true
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Infosüsteemi tehniline kirjeldus (failid ja viited) */
    "technical_specifications": [{"name":"RIHA kontseptuaalne ülevaade", "filename":"RIHA kontseptuaalne ülevaade.rtf", "mime":"application/rtf", "content":"YmFzZTY0Cg=="}],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Sisemiselt kasutatav kirjeldus (failid ja viited) */
    "internal_description": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Muu lisatav dokument või kirjeldus (failid ja viited) */
    "other_description": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: string, Piirangud: valik sagedustest, Kirjeldus: ADS aadressiandmete värskendamise sagedus */
    "ads_update_frequency": "string",
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Kirjeldus, kas ja kuidas säilitatakse aadresside muutmise ajalugu */
    "ads_change_history": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Turvameetmete kirjeldus */
    "security_measures": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Andmete turva- ja riskianalüüs */
    "data_security_analysis": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Isikuandmete kasutamise printsiibid */
    "pesonal_data_usage_principles": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Säilitustähtaegade dokumendid */
    "data_retention_period_specification": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Andmete eraldmise funktsionaalsuse kirjeldus */
    "data_extraction_principles": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Teabele juurdepääsu kirjeldus */
    "data_access_principles": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Infosüsteemi alusdokumendid */
    "infosystem_source_document": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid infosüsteemi kohta (need ei moodusta infosüsteemi ametlikku kirjeldust) */
    "comment": [{"date":"2016-04-17T11:32:00", "person":"37012062719", "content":"Kommentaar infosüsteemi kirjelduse kohta"}],
    /* Andmetüüp: array(object(development)), Piirangud: -, Kirjeldus: Infosüsteemi arenduspartneri kohta käiv info ja finantsinfo */
    "is_developments": [{
        /* Andmetüüp: number, Piirangud: -, Kirjeldus: arendustööde teostamise aasta */
        "year": 9.9,
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: arenduspartneri nimi */
        "developer": "string",
        /* Andmetüüp: number, Piirangud: -, Kirjeldus: tööde rahaline maht */
        "budget": 9.9
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Eesti teabeväravaga liidestamise kirjeldus */
    "eesti_ee_interface_principles": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Andmesubjektile enda andmetele juurdepääsu kirjeldus */
    "subject_data_access_principles": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
        "url": "http://asutus.ee/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
        "name": "Määrus X",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
        "filename": "/opt/documents/mydoc.docx",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
        "mime": "application/word",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
        "content": "YmFzZTY0Cg==",
        /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
        "metadata": "object"
    }],
    /* Andmetüüp: array(object(function)), Piirangud: -, Kirjeldus: Infosüsteemi eesmärkide/funktsioonide kirjeldus */
    "functions": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Infosüsteemi funktsiooni versioonist sõltumatu URI */
        "uri": "urn:fdc:riha.eesti.ee:2016:function:325131",
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Infosüsteemi funktsiooni nimi */
        "name": "Riigi infosüsteemi haldamine",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Andmeobjekti olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: array(object(translation)), Piirangud: -, Kirjeldus: Nimi eri keeltes */
        "names": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge eesti keeles */
            "et": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge inglise keeles */
            "en": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge läti keeles */
            "lv": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge leedu keeles */
            "lt": "string"
        }],
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
        "description": "string",
        /* Andmetüüp: string, Piirangud: enum(archival_type), Kirjeldus: arhiiviväärtuse/säilitustähtaja määrang fikseeritud loeteluna */
        "archival_type": "hiljem",
        /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid andmeobjekti kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
        "comments": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
            "organization": "70006317",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
            "main_resource_uri": "https://riha.eesti.ee/",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
            "data_object_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
            "document_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
            "comment_uri": "string",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
            "state": "A",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
            "creator": "36505160211",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
            "modifier": "36505160211",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
            "creation_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
            "modified_date": "2011-11-01T00:00:00",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
            "content": "siin on kommentaar"
        }],
        /* Andmetüüp: number, Piirangud: -, Kirjeldus: Andmeobjekti säilitustähtaeg */
        "retention_period": 75,
        /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: Funktsiooniga seotud informatiivsed dokumendid */
        "documents": {
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
            "url": "http://asutus.ee/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
            "name": "Määrus X",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
            "filename": "/opt/documents/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
            "mime": "application/word",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
            "state": "C",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
            "start_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
            "end_date": "2010-11-01T00:00:00",
            /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
            "content": "YmFzZTY0Cg==",
            /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
            "metadata": "object"
        },
        /* Andmetüüp: array(object(function)), Piirangud: -, Kirjeldus: Alamfunktsioonid */
        "functions": [{
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Infosüsteemi funktsiooni versioonist sõltumatu URI */
            "uri": "urn:fdc:riha.eesti.ee:2016:function:325131",
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Infosüsteemi funktsiooni nimi */
            "name": "Riigi infosüsteemi haldamine",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Andmeobjekti olek */
            "state": "C",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
            "start_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
            "end_date": "2010-11-01T00:00:00",
            /* Andmetüüp: array(object(translation)), Piirangud: -, Kirjeldus: Nimi eri keeltes */
            "names": [{
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge eesti keeles */
                "et": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge inglise keeles */
                "en": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge läti keeles */
                "lv": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge leedu keeles */
                "lt": "string"
            }],
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
            "description": "string",
            /* Andmetüüp: string, Piirangud: enum(archival_type), Kirjeldus: arhiiviväärtuse/säilitustähtaja määrang fikseeritud loeteluna */
            "archival_type": "hiljem",
            /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid andmeobjekti kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
            "comments": [{
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
                "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
                "organization": "70006317",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
                "main_resource_uri": "https://riha.eesti.ee/",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
                "data_object_uri": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
                "document_uri": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
                "comment_uri": "string",
                /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                "access_restriction": 0,
                /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
                "state": "A",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
                "creator": "36505160211",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
                "modifier": "36505160211",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
                "creation_date": "1998-11-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
                "modified_date": "2011-11-01T00:00:00",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
                "content": "siin on kommentaar"
            }],
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Andmeobjekti säilitustähtaeg */
            "retention_period": 75,
            /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: Funktsiooniga seotud informatiivsed dokumendid */
            "documents": {
                /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
                "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
                "url": "http://asutus.ee/mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
                "name": "Määrus X",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
                "filename": "/opt/documents/mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
                "mime": "application/word",
                /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                "access_restriction": 0,
                /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
                "state": "C",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
                "start_date": "1998-11-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
                "end_date": "2010-11-01T00:00:00",
                /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
                "content": "YmFzZTY0Cg==",
                /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
                "metadata": "object"
            },
            /* Andmetüüp: array(object(function)), Piirangud: -, Kirjeldus: Alamfunktsioonid */
            "functions": [{}]
        }]
    }],
    /* Andmetüüp: array(object(entity)), Piirangud: -, Kirjeldus: Infosüsteemi andmekoosseisu kirjeldus */
    "entities": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Andmeobjekti versioonist sõltumatu URI */
        "uri": "urn:fdc:riha.eesti.ee:2016:entity:alusdokument",
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Andmeobjekti nimi */
        "name": "alusdokument",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Andmeobjekti olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: array(object(translation)), Piirangud: -, Kirjeldus: Nimed eri keeltes */
        "names": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge eesti keeles */
            "et": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge inglise keeles */
            "en": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge läti keeles */
            "lv": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge leedu keeles */
            "lt": "string"
        }],
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
        "description": "Infosüsteemi alusdokumendid",
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: tehniliste klassifikaatorite list, kust väljal võivad väärtused olla */
        "classifiers": ["class11:document"],
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: terminite tekstiline list, mis esitavad veergude/välja tähendust */
        "semantics": ["dokument"],
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Tegemist on andmekogu põhiandmetega */
        "main_data": true,
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Tegemist on delikaatsete isikuandmetega */
        "sensitive_personal_data": true,
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Tegemist on isikuandmetega */
        "personal_data": true,
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: RIHA viitenumber */
        "reference_nr": "AO00636532",
        /* Andmetüüp: string, Piirangud: enum(archival_type), Kirjeldus: arhiiviväärtuse/säilitustähtaja määrang fikseeritud loeteluna */
        "archival_type": "hiljem",
        /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Andmete töötlemise alus - viit aktile ning selle paragrahvile */
        "data_processing_basis": [{
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
            "url": "http://asutus.ee/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
            "name": "Määrus X",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
            "filename": "/opt/documents/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
            "mime": "application/word",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
            "state": "C",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
            "start_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
            "end_date": "2010-11-01T00:00:00",
            /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
            "content": "YmFzZTY0Cg==",
            /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
            "metadata": "object"
        }],
        /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid andmeobjekti kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
        "comments": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
            "organization": "70006317",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
            "main_resource_uri": "https://riha.eesti.ee/",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
            "data_object_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
            "document_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
            "comment_uri": "string",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
            "state": "A",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
            "creator": "36505160211",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
            "modifier": "36505160211",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
            "creation_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
            "modified_date": "2011-11-01T00:00:00",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
            "content": "siin on kommentaar"
        }],
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Lipp, mis märgib, et antud andmeobjekti näol on tegemist aadressiandmetega */
        "address_data": true,
        /* Andmetüüp: number, Piirangud: -, Kirjeldus: Andmeobjekti säilitustähtaeg */
        "retention_period": 9.9,
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Infosüsteemide funktsioonide nimed, mille alla antud andmeobjekt kuulub. */
        "functions_names": ["Riigi infosüsteemi haldamine"],
        /* Andmetüüp: array(object(entity)), Piirangud: -, Kirjeldus: Alam-andmeobjektid */
        "entities": [{"name":"Nimi", "description":"Alusdokumendi nimi"}]
    }],
    /* Andmetüüp: array(object(database)), Piirangud: -, Kirjeldus: Infosüsteemi andmebaaside kirjeldus */
    "databases": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Andmebaasi versioonist sõltumatu URI */
        "uri": "urn:fdc:riha.eesti.ee:2016:database:riha",
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Andmebaasi nimi */
        "name": "RIHA",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Andmeobjekti olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Andmebaasi tehniline kommentaar */
        "comment": "$"Riigi infosüsteemi haldamine"",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimkasutuseks arusaadav nimi */
        "business_name": "string",
        /* Andmetüüp: array(object(translation)), Piirangud: -, Kirjeldus: inimkasutuseks arusaadavad nimed eri keeltes */
        "business_names": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge eesti keeles */
            "et": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge inglise keeles */
            "en": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge läti keeles */
            "lv": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge leedu keeles */
            "lt": "string"
        }],
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
        "description": "string",
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Andmeobjektide nimed, mida antud andmebaas katab */
        "entities_names": ["string"],
        /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid andmeobjekti kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
        "comments": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
            "organization": "70006317",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
            "main_resource_uri": "https://riha.eesti.ee/",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
            "data_object_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
            "document_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
            "comment_uri": "string",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
            "state": "A",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
            "creator": "36505160211",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
            "modifier": "36505160211",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
            "creation_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
            "modified_date": "2011-11-01T00:00:00",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
            "content": "siin on kommentaar"
        }],
        /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: SQL fail, millest tuleb sisse parsida andmebaasi tabelite kirjeldused */
        "sql_schema": {"url:"http://riha.eesti.ee/data/riha.sql"},
        /* Andmetüüp: array(object(table)), Piirangud: -, Kirjeldus: Andmebaasis olevate tabelite kirjeldused */
        "tables": [{
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Andmetabeli versioonist sõltumatu URI */
            "uri": "urn:fdc:riha.eesti.ee:2016:table:alusdokument",
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Andmetabeli nimi */
            "name": "alusdokument",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Andmeobjekti olek */
            "state": "C",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
            "start_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
            "end_date": "2010-11-01T00:00:00",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Andmetabeli tehniline kommentaar */
            "comment": "sisaldab #dokument, !Alusdokument, @class11:document",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimkasutuseks arusaadav nimi */
            "business_name": "Infosüsteemi alusdokumendid",
            /* Andmetüüp: array(object(translation)), Piirangud: -, Kirjeldus: inimkasutuseks arusaadavad nimed eri keeltes */
            "business_names": [{
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge eesti keeles */
                "et": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge inglise keeles */
                "en": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge läti keeles */
                "lv": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge leedu keeles */
                "lt": "string"
            }],
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
            "description": "string",
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Andmekoosseisu andmeobjektide nimed, mida antud tabelis hoitakse */
            "entities_names": ["Alusdokument"],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: tehniliste klassifikaatorite list, kust väljal võivad väärtused olla */
            "classifiers": ["class11:document"],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: terminite tekstiline list, mis esitavad veergude/välja tähendust */
            "semantics": ["dokument"],
            /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid andmeobjekti kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
            "comments": [{
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
                "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
                "organization": "70006317",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
                "main_resource_uri": "https://riha.eesti.ee/",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
                "data_object_uri": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
                "document_uri": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
                "comment_uri": "string",
                /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                "access_restriction": 0,
                /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
                "state": "A",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
                "creator": "36505160211",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
                "modifier": "36505160211",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
                "creation_date": "1998-11-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
                "modified_date": "2011-11-01T00:00:00",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
                "content": "siin on kommentaar"
            }],
            /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: SQL fail, millest tuleb sisse parsida andmetabeli kirjeldus */
            "sql_schema": {
                /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
                "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
                "url": "http://asutus.ee/mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
                "name": "Määrus X",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
                "filename": "/opt/documents/mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
                "mime": "application/word",
                /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                "access_restriction": 0,
                /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
                "state": "C",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
                "start_date": "1998-11-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
                "end_date": "2010-11-01T00:00:00",
                /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
                "content": "YmFzZTY0Cg==",
                /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
                "metadata": "object"
            },
            /* Andmetüüp: array(object(field)), Piirangud: -, Kirjeldus: Andmebaasitabeli atribuudid */
            "fields": [{
                /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Atribuudi versioonist sõltumatu URI */
                "uri": "urn:fdc:riha.eesti.ee:2016:field:alusdokument:nimi",
                /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Atribuudi tehniline nimi */
                "name": "nimi",
                /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                "access_restriction": 0,
                /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Andmeobjekti olek */
                "state": "C",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
                "start_date": "1998-11-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
                "end_date": "2010-11-01T00:00:00",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Andmetüüp */
                "type": "varchar(100)",
                /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Andmeelemendi piirangute tekstiline list */
                "constraints": ["not null"],
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Andmeelemendi tehniline kommentaar */
                "comment": "sisaldab #"dokumendi nimi", !"Alusdokumendi nimi"",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimkasutuseks arusaadav nimi */
                "business_name": "Alusdokumendi nimi",
                /* Andmetüüp: array(object(translation)), Piirangud: -, Kirjeldus: inimkasutuseks arusaadavad nimed eri keeltes */
                "business_names": [{
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge eesti keeles */
                    "et": "string",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge inglise keeles */
                    "en": "string",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge läti keeles */
                    "lv": "string",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tõlge leedu keeles */
                    "lt": "string"
                }],
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
                "description": "string",
                /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Andmekoosseisu andmeobjektide nimed, mida antud tabeli väljal hoitakse */
                "entities_names": ["Alusdokumendi nimi"],
                /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: tehniliste klassifikaatorite list, kust väljal võivad väärtused olla */
                "classifiers": ["string"],
                /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: terminite tekstiline list, mis esitavad veergude/välja tähendust */
                "semantics": ["dokumendi nimi"],
                /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid andmeobjekti kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
                "comments": [{
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
                    "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
                    "organization": "70006317",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
                    "main_resource_uri": "https://riha.eesti.ee/",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
                    "data_object_uri": "string",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
                    "document_uri": "string",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
                    "comment_uri": "string",
                    /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                    "access_restriction": 0,
                    /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
                    "state": "A",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
                    "creator": "36505160211",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
                    "modifier": "36505160211",
                    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
                    "creation_date": "1998-11-01T00:00:00",
                    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
                    "modified_date": "2011-11-01T00:00:00",
                    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
                    "content": "siin on kommentaar"
                }]
            }]
        }]
    }],
    /* Andmetüüp: array(object(service)), Piirangud: -, Kirjeldus: Infosüsteemi poolt pakutavad teenused */
    "services": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:service:xroad:riha.legacyRIHA.v1",
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse inimloetav nimi */
        "name": "RIHA legacy autentimine",
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse omanik */
        "owner": "70006317",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse lühinimi */
        "short_name": "legacyRIHA",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimloetav versiooni nr, ei kasutata identifikaatorina */
        "version": "v1",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi URI, mille alla antud teenuse kirjeldus kuulub */
        "parent_uri": "urn:fdc:riha.eesti.ee:2016:infosystem:riha",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Teenuse olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenusele ligipääsu URL */
        "url": "string",
        /* Andmetüüp: string, Piirangud: enum(service_type), Kirjeldus: Teenuse tüüp */
        "service_type": "xtee_teenus",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse kood (x-tee teenuse korral andmekogu ja teenuse koodnimi) */
        "service_code": "riha.legacyRIHA",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenusele RIHA poolt omistatud unikaalne kood */
        "reference_number": "TE001550",
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusega töödeldakse delikaatseid isikuandmeid */
        "sensitive_personal_data": False,
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusega töödeldakse isikuandmeid */
        "personal_data": True,
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusel on infosüsteemiga samad kontaktisikud */
        "use_infosystem_contacts": False,
        /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: Teenuse WSDL kirjeldus */
        "wsdl": {
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
            "url": "http://asutus.ee/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
            "name": "Määrus X",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
            "filename": "/opt/documents/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
            "mime": "application/word",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
            "state": "C",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
            "start_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
            "end_date": "2010-11-01T00:00:00",
            /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
            "content": "YmFzZTY0Cg==",
            /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
            "metadata": "object"
        },
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse kirjeldus */
        "description": "RIHA legacy autentimine",
        /* Andmetüüp: string, Piirangud: enum(service_user_group), Kirjeldus: Teenust kasutav kasutajagrupp */
        "user_group": "xtee_oigusega",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: IKSE turvaklass */
        "iske_security_class": "K2T2S1",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse lõpetamisinfo */
        "end_comment": "string",
        /* Andmetüüp: array(object(data_object)), Piirangud: -, Kirjeldus: Teenuse sisendis olevate andmeobjektide kirjeldused. Võivad olla ka viited infosüsteemi andmekoosseisu vastavatele obektidele (URI kaudu) */
        "inputs": [{"type":"string", "name":"Teenuse nimi"}],
        /* Andmetüüp: array(object(data_object)), Piirangud: -, Kirjeldus: Teenuse väljundis olevate andmeobjektide kirjeldused. Võivad olla ka viited infosüsteemi andmekoosseisu vastavatele obektidele (URI kaudu) */
        "outputs": [{"type":"string", "name":"Teenuse URL"}],
        /* Andmetüüp: object(service_level), Piirangud: -, Kirjeldus: Teenustaseme parameetrid */
        "service_level": {
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenustaseme parameetrite unikaalne URI */
            "uri": "string",
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Teenustaseme parameetrid samad infosüsteemiga */
            "same_as_infosystem": true,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Intsidentidele reageerimise aeg (h) */
            "incident_initial_response_time": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Teenuse nõuetele reageerimise aeg (h) */
            "service_request_initial_response_time": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Kuu keskmine käideldavus (%) */
            "monthly_average_service_availability": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Maksimaalne ühekordne katkestuse kestvus (h) */
            "max_duration_of_service_interruption": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Standardpäringu aeg (s) */
            "acceptable_response_time": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Usaldusväärsus (katkestust/kuus) */
            "monthly_average_service_reliability": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Maksimaalne koormus (kasutajat/minutis) */
            "service_load_level": 9.9,
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Teenuse osutamise alustamise kellaaeg */
            "service_offering_start_time": "2016-01-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Teenuse osutamise lõpetamise kellaaeg */
            "service_offering_end_time": "2016-01-01T00:00:00",
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Tööaeg 24x7 */
            "service_time_24x7": true,
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tööaja märkused */
            "service_time_comments": "string"
        },
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud märgendite nimekiri */
        "tags": ["#tag1","!tag2"],
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud semantilised terminid (sisend/väljund, termini RDF id) */
        "semantics": ["string"],
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud valdkonnad */
        "classifiers": ["Infoühiskond"],
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenuse kasutajad */
        "users": ["70008747"],
        /* Andmetüüp: array(object(linked_person)), Piirangud: -, Kirjeldus: Teenusega seotud kontaktisikud */
        "contact_persons": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud isiku isikukood */
            "person": "string",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse alguskuupäev */
            "date_from": "2016-01-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse lõppkuupäev */
            "date_to": "2016-01-01T00:00:00",
            /* Andmetüüp: string, Piirangud: enum(person_link_type), Kirjeldus: Seose tüüp */
            "type": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Ametikoht */
            "occupation": "string",
            /* Andmetüüp: string, Piirangud: enum(person_contact_level), Kirjeldus: Kontakti tase */
            "contact_level": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
            "description": "string"
        }],
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenuses kasutatavate autentimismeetodite nimekiri */
        "auth_methods": ["string"],
        /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid teenuse kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
        "comment": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
            "organization": "70006317",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
            "main_resource_uri": "https://riha.eesti.ee/",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
            "data_object_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
            "document_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
            "comment_uri": "string",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
            "state": "A",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
            "creator": "36505160211",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
            "modifier": "36505160211",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
            "creation_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
            "modified_date": "2011-11-01T00:00:00",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
            "content": "siin on kommentaar"
        }],
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Teenus on avalikult indekseeritav */
        "indexable": true,
        /* Andmetüüp: array(object(service)), Piirangud: -, Kirjeldus: Teenuse alamteenused */
        "services": [{
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:service:xroad:riha.legacyRIHA.v1",
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse inimloetav nimi */
            "name": "RIHA legacy autentimine",
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse omanik */
            "owner": "70006317",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse lühinimi */
            "short_name": "legacyRIHA",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimloetav versiooni nr, ei kasutata identifikaatorina */
            "version": "v1",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi URI, mille alla antud teenuse kirjeldus kuulub */
            "parent_uri": "urn:fdc:riha.eesti.ee:2016:infosystem:riha",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Teenuse olek */
            "state": "C",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
            "start_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
            "end_date": "2010-11-01T00:00:00",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenusele ligipääsu URL */
            "url": "string",
            /* Andmetüüp: string, Piirangud: enum(service_type), Kirjeldus: Teenuse tüüp */
            "service_type": "xtee_teenus",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse kood (x-tee teenuse korral andmekogu ja teenuse koodnimi) */
            "service_code": "riha.legacyRIHA",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenusele RIHA poolt omistatud unikaalne kood */
            "reference_number": "TE001550",
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusega töödeldakse delikaatseid isikuandmeid */
            "sensitive_personal_data": False,
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusega töödeldakse isikuandmeid */
            "personal_data": True,
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusel on infosüsteemiga samad kontaktisikud */
            "use_infosystem_contacts": False,
            /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: Teenuse WSDL kirjeldus */
            "wsdl": {
                /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
                "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
                "url": "http://asutus.ee/mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
                "name": "Määrus X",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
                "filename": "/opt/documents/mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
                "mime": "application/word",
                /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                "access_restriction": 0,
                /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
                "state": "C",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
                "start_date": "1998-11-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
                "end_date": "2010-11-01T00:00:00",
                /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
                "content": "YmFzZTY0Cg==",
                /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
                "metadata": "object"
            },
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse kirjeldus */
            "description": "RIHA legacy autentimine",
            /* Andmetüüp: string, Piirangud: enum(service_user_group), Kirjeldus: Teenust kasutav kasutajagrupp */
            "user_group": "xtee_oigusega",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: IKSE turvaklass */
            "iske_security_class": "K2T2S1",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse lõpetamisinfo */
            "end_comment": "string",
            /* Andmetüüp: array(object(data_object)), Piirangud: -, Kirjeldus: Teenuse sisendis olevate andmeobjektide kirjeldused. Võivad olla ka viited infosüsteemi andmekoosseisu vastavatele obektidele (URI kaudu) */
            "inputs": [{"type":"string", "name":"Teenuse nimi"}],
            /* Andmetüüp: array(object(data_object)), Piirangud: -, Kirjeldus: Teenuse väljundis olevate andmeobjektide kirjeldused. Võivad olla ka viited infosüsteemi andmekoosseisu vastavatele obektidele (URI kaudu) */
            "outputs": [{"type":"string", "name":"Teenuse URL"}],
            /* Andmetüüp: object(service_level), Piirangud: -, Kirjeldus: Teenustaseme parameetrid */
            "service_level": {
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenustaseme parameetrite unikaalne URI */
                "uri": "string",
                /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Teenustaseme parameetrid samad infosüsteemiga */
                "same_as_infosystem": true,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Intsidentidele reageerimise aeg (h) */
                "incident_initial_response_time": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Teenuse nõuetele reageerimise aeg (h) */
                "service_request_initial_response_time": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Kuu keskmine käideldavus (%) */
                "monthly_average_service_availability": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Maksimaalne ühekordne katkestuse kestvus (h) */
                "max_duration_of_service_interruption": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Standardpäringu aeg (s) */
                "acceptable_response_time": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Usaldusväärsus (katkestust/kuus) */
                "monthly_average_service_reliability": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Maksimaalne koormus (kasutajat/minutis) */
                "service_load_level": 9.9,
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Teenuse osutamise alustamise kellaaeg */
                "service_offering_start_time": "2016-01-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Teenuse osutamise lõpetamise kellaaeg */
                "service_offering_end_time": "2016-01-01T00:00:00",
                /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Tööaeg 24x7 */
                "service_time_24x7": true,
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tööaja märkused */
                "service_time_comments": "string"
            },
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud märgendite nimekiri */
            "tags": ["#tag1","!tag2"],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud semantilised terminid (sisend/väljund, termini RDF id) */
            "semantics": ["string"],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud valdkonnad */
            "classifiers": ["Infoühiskond"],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenuse kasutajad */
            "users": ["70008747"],
            /* Andmetüüp: array(object(linked_person)), Piirangud: -, Kirjeldus: Teenusega seotud kontaktisikud */
            "contact_persons": [{
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud isiku isikukood */
                "person": "string",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse alguskuupäev */
                "date_from": "2016-01-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse lõppkuupäev */
                "date_to": "2016-01-01T00:00:00",
                /* Andmetüüp: string, Piirangud: enum(person_link_type), Kirjeldus: Seose tüüp */
                "type": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Ametikoht */
                "occupation": "string",
                /* Andmetüüp: string, Piirangud: enum(person_contact_level), Kirjeldus: Kontakti tase */
                "contact_level": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
                "description": "string"
            }],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenuses kasutatavate autentimismeetodite nimekiri */
            "auth_methods": ["string"],
            /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid teenuse kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
            "comment": [{
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
                "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
                "organization": "70006317",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
                "main_resource_uri": "https://riha.eesti.ee/",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
                "data_object_uri": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
                "document_uri": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
                "comment_uri": "string",
                /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                "access_restriction": 0,
                /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
                "state": "A",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
                "creator": "36505160211",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
                "modifier": "36505160211",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
                "creation_date": "1998-11-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
                "modified_date": "2011-11-01T00:00:00",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
                "content": "siin on kommentaar"
            }],
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Teenus on avalikult indekseeritav */
            "indexable": true,
            /* Andmetüüp: array(object(service)), Piirangud: -, Kirjeldus: Teenuse alamteenused */
            "services": [{}]
        }]
    }],
    /* Andmetüüp: array(object(service)), Piirangud: -, Kirjeldus: Infosüsteemi poolt kasutatavad teenused */
    "used_services": [{
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse unikaalne URI, mis on kasutaja määrata */
        "uri": "urn:fdc:riha.eesti.ee:2016:service:xroad:riha.legacyRIHA.v1",
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse inimloetav nimi */
        "name": "RIHA legacy autentimine",
        /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse omanik */
        "owner": "70006317",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse lühinimi */
        "short_name": "legacyRIHA",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimloetav versiooni nr, ei kasutata identifikaatorina */
        "version": "v1",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi URI, mille alla antud teenuse kirjeldus kuulub */
        "parent_uri": "urn:fdc:riha.eesti.ee:2016:infosystem:riha",
        /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
        "access_restriction": 0,
        /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Teenuse olek */
        "state": "C",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
        "start_date": "1998-11-01T00:00:00",
        /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
        "end_date": "2010-11-01T00:00:00",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenusele ligipääsu URL */
        "url": "string",
        /* Andmetüüp: string, Piirangud: enum(service_type), Kirjeldus: Teenuse tüüp */
        "service_type": "xtee_teenus",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse kood (x-tee teenuse korral andmekogu ja teenuse koodnimi) */
        "service_code": "riha.legacyRIHA",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenusele RIHA poolt omistatud unikaalne kood */
        "reference_number": "TE001550",
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusega töödeldakse delikaatseid isikuandmeid */
        "sensitive_personal_data": False,
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusega töödeldakse isikuandmeid */
        "personal_data": True,
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusel on infosüsteemiga samad kontaktisikud */
        "use_infosystem_contacts": False,
        /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: Teenuse WSDL kirjeldus */
        "wsdl": {
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
            "url": "http://asutus.ee/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
            "name": "Määrus X",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
            "filename": "/opt/documents/mydoc.docx",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
            "mime": "application/word",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
            "state": "C",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
            "start_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
            "end_date": "2010-11-01T00:00:00",
            /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
            "content": "YmFzZTY0Cg==",
            /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
            "metadata": "object"
        },
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse kirjeldus */
        "description": "RIHA legacy autentimine",
        /* Andmetüüp: string, Piirangud: enum(service_user_group), Kirjeldus: Teenust kasutav kasutajagrupp */
        "user_group": "xtee_oigusega",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: IKSE turvaklass */
        "iske_security_class": "K2T2S1",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse lõpetamisinfo */
        "end_comment": "string",
        /* Andmetüüp: array(object(data_object)), Piirangud: -, Kirjeldus: Teenuse sisendis olevate andmeobjektide kirjeldused. Võivad olla ka viited infosüsteemi andmekoosseisu vastavatele obektidele (URI kaudu) */
        "inputs": [{"type":"string", "name":"Teenuse nimi"}],
        /* Andmetüüp: array(object(data_object)), Piirangud: -, Kirjeldus: Teenuse väljundis olevate andmeobjektide kirjeldused. Võivad olla ka viited infosüsteemi andmekoosseisu vastavatele obektidele (URI kaudu) */
        "outputs": [{"type":"string", "name":"Teenuse URL"}],
        /* Andmetüüp: object(service_level), Piirangud: -, Kirjeldus: Teenustaseme parameetrid */
        "service_level": {
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenustaseme parameetrite unikaalne URI */
            "uri": "string",
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Teenustaseme parameetrid samad infosüsteemiga */
            "same_as_infosystem": true,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Intsidentidele reageerimise aeg (h) */
            "incident_initial_response_time": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Teenuse nõuetele reageerimise aeg (h) */
            "service_request_initial_response_time": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Kuu keskmine käideldavus (%) */
            "monthly_average_service_availability": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Maksimaalne ühekordne katkestuse kestvus (h) */
            "max_duration_of_service_interruption": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Standardpäringu aeg (s) */
            "acceptable_response_time": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Usaldusväärsus (katkestust/kuus) */
            "monthly_average_service_reliability": 9.9,
            /* Andmetüüp: number, Piirangud: -, Kirjeldus: Maksimaalne koormus (kasutajat/minutis) */
            "service_load_level": 9.9,
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Teenuse osutamise alustamise kellaaeg */
            "service_offering_start_time": "2016-01-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Teenuse osutamise lõpetamise kellaaeg */
            "service_offering_end_time": "2016-01-01T00:00:00",
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Tööaeg 24x7 */
            "service_time_24x7": true,
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tööaja märkused */
            "service_time_comments": "string"
        },
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud märgendite nimekiri */
        "tags": ["#tag1","!tag2"],
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud semantilised terminid (sisend/väljund, termini RDF id) */
        "semantics": ["string"],
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud valdkonnad */
        "classifiers": ["Infoühiskond"],
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenuse kasutajad */
        "users": ["70008747"],
        /* Andmetüüp: array(object(linked_person)), Piirangud: -, Kirjeldus: Teenusega seotud kontaktisikud */
        "contact_persons": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud isiku isikukood */
            "person": "string",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse alguskuupäev */
            "date_from": "2016-01-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse lõppkuupäev */
            "date_to": "2016-01-01T00:00:00",
            /* Andmetüüp: string, Piirangud: enum(person_link_type), Kirjeldus: Seose tüüp */
            "type": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Ametikoht */
            "occupation": "string",
            /* Andmetüüp: string, Piirangud: enum(person_contact_level), Kirjeldus: Kontakti tase */
            "contact_level": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
            "description": "string"
        }],
        /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenuses kasutatavate autentimismeetodite nimekiri */
        "auth_methods": ["string"],
        /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid teenuse kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
        "comment": [{
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
            "organization": "70006317",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
            "main_resource_uri": "https://riha.eesti.ee/",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
            "data_object_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
            "document_uri": "string",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
            "comment_uri": "string",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
            "state": "A",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
            "creator": "36505160211",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
            "modifier": "36505160211",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
            "creation_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
            "modified_date": "2011-11-01T00:00:00",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
            "content": "siin on kommentaar"
        }],
        /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Teenus on avalikult indekseeritav */
        "indexable": true,
        /* Andmetüüp: array(object(service)), Piirangud: -, Kirjeldus: Teenuse alamteenused */
        "services": [{
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse unikaalne URI, mis on kasutaja määrata */
            "uri": "urn:fdc:riha.eesti.ee:2016:service:xroad:riha.legacyRIHA.v1",
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse inimloetav nimi */
            "name": "RIHA legacy autentimine",
            /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Teenuse omanik */
            "owner": "70006317",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse lühinimi */
            "short_name": "legacyRIHA",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimloetav versiooni nr, ei kasutata identifikaatorina */
            "version": "v1",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi URI, mille alla antud teenuse kirjeldus kuulub */
            "parent_uri": "urn:fdc:riha.eesti.ee:2016:infosystem:riha",
            /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
            "access_restriction": 0,
            /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Teenuse olek */
            "state": "C",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
            "start_date": "1998-11-01T00:00:00",
            /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
            "end_date": "2010-11-01T00:00:00",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenusele ligipääsu URL */
            "url": "string",
            /* Andmetüüp: string, Piirangud: enum(service_type), Kirjeldus: Teenuse tüüp */
            "service_type": "xtee_teenus",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse kood (x-tee teenuse korral andmekogu ja teenuse koodnimi) */
            "service_code": "riha.legacyRIHA",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenusele RIHA poolt omistatud unikaalne kood */
            "reference_number": "TE001550",
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusega töödeldakse delikaatseid isikuandmeid */
            "sensitive_personal_data": False,
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusega töödeldakse isikuandmeid */
            "personal_data": True,
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas teenusel on infosüsteemiga samad kontaktisikud */
            "use_infosystem_contacts": False,
            /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: Teenuse WSDL kirjeldus */
            "wsdl": {
                /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Dokumendi unikaalne URI, mis on kasutaja määrata */
                "uri": "urn:fdc:riha.eesti.ee:2016:document:asutus.ee:mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi URL, mille kaudu see on loetav */
                "url": "http://asutus.ee/mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Idokumendi inimloetav nimi */
                "name": "Määrus X",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi failinimi */
                "filename": "/opt/documents/mydoc.docx",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Dokumendi sisu MIME tüüp */
                "mime": "application/word",
                /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                "access_restriction": 0,
                /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Dokumendi olek */
                "state": "C",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: dokumendi kehtivuse algus */
                "start_date": "1998-11-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Dokumendi kehtivuse lõpp */
                "end_date": "2010-11-01T00:00:00",
                /* Andmetüüp: base64, Piirangud: -, Kirjeldus: Dokumendi sisu binaarkujul konverditud stringiks */
                "content": "YmFzZTY0Cg==",
                /* Andmetüüp: object, Piirangud: -, Kirjeldus: Dokumendi metaandmed */
                "metadata": "object"
            },
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse kirjeldus */
            "description": "RIHA legacy autentimine",
            /* Andmetüüp: string, Piirangud: enum(service_user_group), Kirjeldus: Teenust kasutav kasutajagrupp */
            "user_group": "xtee_oigusega",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: IKSE turvaklass */
            "iske_security_class": "K2T2S1",
            /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenuse lõpetamisinfo */
            "end_comment": "string",
            /* Andmetüüp: array(object(data_object)), Piirangud: -, Kirjeldus: Teenuse sisendis olevate andmeobjektide kirjeldused. Võivad olla ka viited infosüsteemi andmekoosseisu vastavatele obektidele (URI kaudu) */
            "inputs": [{"type":"string", "name":"Teenuse nimi"}],
            /* Andmetüüp: array(object(data_object)), Piirangud: -, Kirjeldus: Teenuse väljundis olevate andmeobjektide kirjeldused. Võivad olla ka viited infosüsteemi andmekoosseisu vastavatele obektidele (URI kaudu) */
            "outputs": [{"type":"string", "name":"Teenuse URL"}],
            /* Andmetüüp: object(service_level), Piirangud: -, Kirjeldus: Teenustaseme parameetrid */
            "service_level": {
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Teenustaseme parameetrite unikaalne URI */
                "uri": "string",
                /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Teenustaseme parameetrid samad infosüsteemiga */
                "same_as_infosystem": true,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Intsidentidele reageerimise aeg (h) */
                "incident_initial_response_time": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Teenuse nõuetele reageerimise aeg (h) */
                "service_request_initial_response_time": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Kuu keskmine käideldavus (%) */
                "monthly_average_service_availability": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Maksimaalne ühekordne katkestuse kestvus (h) */
                "max_duration_of_service_interruption": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Standardpäringu aeg (s) */
                "acceptable_response_time": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Usaldusväärsus (katkestust/kuus) */
                "monthly_average_service_reliability": 9.9,
                /* Andmetüüp: number, Piirangud: -, Kirjeldus: Maksimaalne koormus (kasutajat/minutis) */
                "service_load_level": 9.9,
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Teenuse osutamise alustamise kellaaeg */
                "service_offering_start_time": "2016-01-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Teenuse osutamise lõpetamise kellaaeg */
                "service_offering_end_time": "2016-01-01T00:00:00",
                /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Tööaeg 24x7 */
                "service_time_24x7": true,
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Tööaja märkused */
                "service_time_comments": "string"
            },
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud märgendite nimekiri */
            "tags": ["#tag1","!tag2"],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud semantilised terminid (sisend/väljund, termini RDF id) */
            "semantics": ["string"],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenusega seotud valdkonnad */
            "classifiers": ["Infoühiskond"],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenuse kasutajad */
            "users": ["70008747"],
            /* Andmetüüp: array(object(linked_person)), Piirangud: -, Kirjeldus: Teenusega seotud kontaktisikud */
            "contact_persons": [{
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud isiku isikukood */
                "person": "string",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse alguskuupäev */
                "date_from": "2016-01-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: Seotuse lõppkuupäev */
                "date_to": "2016-01-01T00:00:00",
                /* Andmetüüp: string, Piirangud: enum(person_link_type), Kirjeldus: Seose tüüp */
                "type": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Ametikoht */
                "occupation": "string",
                /* Andmetüüp: string, Piirangud: enum(person_contact_level), Kirjeldus: Kontakti tase */
                "contact_level": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
                "description": "string"
            }],
            /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Teenuses kasutatavate autentimismeetodite nimekiri */
            "auth_methods": ["string"],
            /* Andmetüüp: array(object(comment)), Piirangud: -, Kirjeldus: Erinevate osapoolte kommentaarid teenuse kohta. Ei moodusta infosüsteemi ametlikku kirjeldust. */
            "comment": [{
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi unikaalne URI, mis on kasutaja määrata */
                "uri": "urn:fdc:riha.eesti.ee:2016:comment:324131",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Infosüsteemi omanik */
                "organization": "70006317",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud infosüsteemi või teenuse URI */
                "main_resource_uri": "https://riha.eesti.ee/",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud andmeobjekti URI */
                "data_object_uri": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud dokumendi URI */
                "document_uri": "string",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seotud kommentaari URI */
                "comment_uri": "string",
                /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
                "access_restriction": 0,
                /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Kommentaari olek */
                "state": "A",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: algne sisestaja */
                "creator": "36505160211",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: viimane muutja */
                "modifier": "36505160211",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: algne sisestamisaeg */
                "creation_date": "1998-11-01T00:00:00",
                /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: viimase muutmise aeg */
                "modified_date": "2011-11-01T00:00:00",
                /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kommentaari tekst */
                "content": "siin on kommentaar"
            }],
            /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Teenus on avalikult indekseeritav */
            "indexable": true,
            /* Andmetüüp: array(object(service)), Piirangud: -, Kirjeldus: Teenuse alamteenused */
            "services": [{}]
        }]
    }]
}
```

### Andmekogu detailkirjeldus
Andmekogu detailkirjeldus on eeltoodud infosüsteemi kirjelduse osa (infosüsteemis võib olla ka mitu andmekogu) mida võib esitada kas:
•	eraldi JSON-lehtedena, millele infosüsteemi kirjeldus viitab,
•	või sügava hierarhilise struktuurina infosüsteemi kirjelduse JSON-struktuuri sees.

Andmekogu detailkirjeldamiseks on ette nähtud mitu alternatiivset viisi, mille vahel andmekogu arendaja võib valida endale sobivama:
•	Detailkirjeldus andmebaasi skeemi koostamise SQL-lausetena, mille kommentaaridesse on kodeeritud täiendav vajalik informatsioon RIHA jaoks.
•	Detailkirjeldus etteantud formaadis JSON-struktuurina.
•	Perspektiivis näeme ette ka muude formaatide lisamise.

Hakatuseks võib detailkirjelduse realiseerida staatilise SQL-failina, mis pannakse serverisse.

Hea mõte on seejärel realiseerida detailkirjeldust selliselt, et arendaja koostab väikese veebirakenduse, mis teeb andmebaasi päringu ja küsib sealt automaatselt baasi skeemi SQL lausetena. Selle tulemusena on igal hetkel kättesaadav värske, reaalne andmebaasi struktuur, mida ei ole vaja käsitsi täiendada.

Oluline osa detailkirjelduse loomisel on kommentaaride sisestamine andmebaasi kirjeldusele. Seda saab teha kahel moel:
•	Kirjutada kommentaarid otse skeemi loomise lausetesse. Miinus on see, et enamus andmebaase ei suuda neid kommentaare päringu peale väljastada.
•	Sisestada kommentaarid SQL-s comment on …. lausetena. Sel juhul saab neid automaatselt päringuga küsida.

### Tähenduste kodeerimine andmebaasi ja teenuste kommentaarides
Andmekogu või teenuse struktuurist ja kasutusest arusaamiseks on äärmiselt kasulik täiendada välju inimloetavate kommentaaridega. Neid kommentaare aga ei suuda RIHA otsirakendused adekvaatselt kasutada. Seepärast on vaja lisada olulisematele tabelitele ja väljadele täiendavat infot, mida RIHA rakendus sealt automaatselt leiab ja suudab automaatselt töödelda ja otsitavaks muuta. Sama täiendav info on oluliseks abiks ka arendajatele, kes saavad sealt näiteks välja lugeda võimalikke väärtus-kodeeringuid ja nende tähendusi.

Tagi asukoht kommentaaris ei ole oluline, põhimõtteliselt võib tagi-märke lisada ka inimkeelse kommentaari olemasolevatele sõnadele.

Tage võib ja tuleb lisada nii tabelite kui väljade kommentaaridele. Näiteks on põhiobjektid tüüpiliselt esitatud tabelites ja siis tuleks põhiobjekti tag lisada tabeli kommentaarile, klassifikaatorid aga vastupidi, on reeglina seotud väljadega, ja nad tuleb lisada vastava välja (teenuse korral sisend- või väljundvälja) kommentaarile. Vältida tuleks seda, et lisada terve tabeli kommentaarile näiteks klassifikaatori tage lihtsalt seepärast, et tabelis on mõnel väljal kasutatud klassifikaatorit.

Lisatavad automaattöödeldavad tähendused jagunevad kolme gruppi, kusjuures mõnel baasiväljal ei pruugi olla ühtegi, ja mõnel võib olla mitu.

Tagisõnades ei eristata väike- ja suurtähti. Võib kasutada mõlemaid segiläbi.

Tagisõnades ei ole vaja kasutada jutumärke, kui tegu on ühe sõnaga, mitmesõnalisel objektil peavad olema ümber jutumärgid " või apostroofid '

Neli eraldi liiki tähendusi on järgmised:
•	põhiobjektid infosüsteemi määruse mõttes lisatakse kommentaari tagidena, mille alguses on hüüumärk, näiteks !klient, !tootja, !"hooned ja ehitised", kus põhiobjektide nimistu leiab arendaja andmekogu põhimäärusest. Arusaadavalt on arendajal vaja koostada endale konkreetne nimistu põhimääruses kirjeldatavare põhiobjektide sõnadest või fraasidest, et vältida sama põhiobjekti kirjeldamist mitme erineva tagiga.
•	infosüsteemi funktsioonid/eesmärgid lisatakse kommentaari tagidena, mille alguses on sümbol $.
•	semantika valdkonnasõnastike / hierarhia mõttes lisatakse kommentaari tagidena, mille alguses on # trellid, näiteks #isik, #isikukood, #asutus, kusjuures:
•	* ◦◦Selleks, et näidata, millise sõnastiku sõnaga on tegemist, tuleb panna # ja sõna vahele prefiksina sõnastiku nimi ja koolon, näiteks #meieasutus:isik, #meieasutus:salajasus kus meieasutus on konkreetse sõnastiku nimi.
•	* ◦◦Sõnastike loomine, otsimine ja haldamine ei ole käesoleva juhendi skoobis: selleks luuakse eraldi juhend.
•	* ◦◦Kasutusel on ka üks nö universaalsõnastik schema.org: selle sõnade kasutamisel kasutatakse prefiksit so, näiteks #so:person, #so:thing
•	* ◦◦Prefiksit ei ole vaja, kui andmetabelile on määratud selle tabeli standardsõnastik: selleks tuleb panna tabeli kommentaari tag #vocabulary:sõnastikunimi
•	klassifikaatorid näitavad ära, mis on välja võimalikud (kodeeritud) väärtused ja kuidas neist aru saada: tegemist siis võimalike väärtuste/seletuste paaride loeteluga.
•	* ◦◦Klassifikaatori lisamiseks tuleb klassifikaatori nime ette panna @ at märk, näiteks @sugu, @organisatsiooni_tyyp mis tähistavad konkreetseid võimalikke liste nagu 
Unknown macro: {"M"} 
.
•	* ◦◦Selleks, et näidata, millise klassifikaatorite-komplektiga on tegemist, tuleb panna @ ja sõna vahele prefiksina klassifikaatori-komplekti nimi ja koolon, näiteks @klassif_12:sugu kus kus klassif_12 on konkreetse klassifikaatori-komplekti nimi.
•	* ◦◦Klassifikaatorite loomine, otsimine ja haldamine ei ole käesoleva juhendi skoobis: selleks luuakse eraldi juhend.
•	* ◦◦Prefiksit ei ole vaja, kui andmetabelile on määratud selle tabeli standardklassifikaatori-komplekt: selleks tuleb panna tabeli kommentaari tag @classifiers:klassifikaatorikomplekti-nimi
•	* ◦◦
Järgnevas näitena üks kommenteeritud väljadega tabel, kus kommentaarid sisaldavadki eelnevalt kirjeldatud tage:

```sql
create table clients ( -- klientide tabel, sisaldab #isik !klient kasutab @classifiers:standardklassif
  id integer primary key default nextval('clients_id_seq'), 
  username varchar(100) not null unique, -- kliendi kasutajanimi a la facebook:0013230
  personcode varchar(50),  -- ametlik #isikukood ala 36505160211
  fullname varchar(100), -- kliendi #täisnimi a la Jaan Tamm
  phone varchar(50),
  email varchar(100),
  address varchar(200),
  remarks varchar(200),  
  gender char(1), -- M: mees, N: naine #sugu @sugu
  religion varchar(50), -- #tundlikud_isikuandmed religioon
  status char(1) default 'A', --- A: active, D: deactivated !meieklassif:aktiivsus
  created_at timestamp default now(), 
  updated_at timestamp default now(),
  updated_by varchar(100) -- username or systemname creating/updating
);
```

Kuidas sisestada kommentaare näiteks postgres ja oracle baasides:

```sql
COMMENT ON TABLE clients IS &#39;klientide tabel, sisaldab #isik !klient&#39;;

COMMENT ON COLUMN clients.fullname IS &#39;kliendi #täisnimi a la Jaan Tamm&#39;;
```

### Andmebaasi kirjeldamine JSON-struktuurina

Järgnevas näites on toodud alternatiiv andmekogu kõigi andmebaaside ja tabelite kirjelduse andmiseks JSON kujul. Tegemist on hierarhilise esitusega, mille tipmine tase on kõigi andmebaaside loend, igaühe sees on kõigi tabelite loend, nende sees omakorda kõigi väljade loend.

Kui kasutatakse dokument-andmebaase või sisemisi JSON-struktuure, võib selle kirjelduse struktuur olla ka sügavam.

RIHA jaoks masinloetav struktuur ei pea seejuures koosnema ühest universaalsest failist, mis kirjeldab kõiki andmebaase, vaid neid võib esitada ka eraldi veebilehtedes iga andmebaasi ja iga tabeli jaoks, vt allpool näites
```
http://meieasutus.ee/datadesc/describe?products"
…
http://meieasutus.ee/datadesc/describe?prices"
```

Üldpõhimõttena võib iga objekt allpool sisaldada kas otseseid väärtusi või siis olla stringina URLi kujul " http://mingiurl.ee/" kus " http://mingiurl.ee/" sisu esitabki objekti kogu sisu.

RIHA jaoks masinloetav JSON, mis ei sisalda kõiki võimalikke välju, vaid tüüpilisi põhivälju:

```
"databases": [{
        "name": "RIHA",
        "comment": "$"Riigi infosüsteemi haldamine"",
        "sql_schema": {"url:"http://riha.eesti.ee/data/riha.sql"},
        "tables": [{
            "name": "alusdokument",
            "comment": "sisaldab #dokument, !Alusdokument, @class11:document",
            "business_name": "Infosüsteemi alusdokumendid",
            "entities_names": ["Alusdokument"],
            "classifiers": ["class11:document"],
            "semantics": ["dokument"],
            "fields": [{
                "name": "nimi",
                "type": "varchar(100)",
                "constraints": ["not null"],
                "comment": "sisaldab #"dokumendi nimi", !"Alusdokumendi nimi"",
                "business_name": "Alusdokumendi nimi",
                "entities_names": ["Alusdokumendi nimi"],
                "semantics": ["dokumendi nimi"],

            }]
        }]
    }],
```    

### Teenuste kirjeldamine JSON-struktuuriga

Andmekogu poolt realiseeritud teenuse kirjeldus koosneb üldjuhul nii URLi kaudu viidatavast WSDL-failist kui JSON tekstist, mis näitab ära konkreetsed teenused, nende metaandmed ja täiendava info väljadele.

Teenuseväljade täiendav info on sama, kui andmebaasi väljadel, sisaldades üldjuhul inimloetavat kommentaari koos tagidega, mis määravad ära välja tähenduse ja klassifikaatori.

Soovitav lahendus välja kommentaari, tähenduse ja klassifikaatori esitamiseks on sisestada xtee:notes väljale kommentaar koos tagidega vastavalt eeltoodud tagimise juhendile.

Alternatiivina saab kommentaari ja tagid lisada JSON struktuuris.

Ühe konkreetse teenuse esitamise näide andmekogu teenuste loendis:
```
    "services": [{
        "uri": "urn:fdc:riha.eesti.ee:2016:service:xroad:riha.legacyRIHA.v1",
        "name": "RIHA legacy autentimine",
        "owner": "70006317",
        "short_name": "legacyRIHA",
        "version": "v1",
        "service_type": "xtee_teenus",
        "service_code": "riha.legacyRIHA",
        "reference_number": "TE001550",
        "sensitive_personal_data": False,
        "personal_data": True,
        "use_infosystem_contacts": False,
        "description": "RIHA legacy autentimine",
        "user_group": "xtee_oigusega",
        "iske_security_class": "K2T2S1",
        "inputs": [{"type":"string", "name":"Teenuse nimi"}],
        "outputs": [{"type":"string", "name":"Teenuse URL"}],
        "tags": ["#tag1","!tag2"],
        "classifiers": ["Infoühiskond"],
        "users": ["70008747"]
}]
```
"inputs" ja "outputs" loendid võivad olla esitamata, kui sisendite/väljundite tage sisaldavad kommentaarid on olemas WSDL failis väljadel xtee:notes.
Valdkonnasõnastike kirjeldamine JSON-struktuuriga
Sõnastiku kirjeldus sisaldab
•	Standardsel kujul üldinfot
•	Loetelu objektidest, mis sisaldavad mis sisaldavad kas kas otse sõnastiku esitust rdfs kujul või objekte, mis sisaldavad nende URL-e.
RIHA edasises arendamises soovitame tungivalt töötada välja sõnastike uus struktureeritud esitus, mis oleks lihtsam, selgem ja kergemini hallatav, kui praegune.

/* object(vocabulary): Valdkonna sõnastikud */
"vocabulary": {
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Sõnastiku unikaalne URI */
    "uri": "urn:fdc:riha.eesti.ee:2016:vocabulary:413423212",
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Sõnastiku nimetus */
    "name": "Äriregistri terminid",
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Sõnastiku omanik */
    "owner": "70006317",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Sõnastiku lühinimi */
    "short_name": "string",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Sõnastiku versiooni nr */
    "version": "Versioon 2. (19.07.2010)",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: versioonist sõltumatu hierarhia viit - sõnastiku korral ei kasutata */
    "parent_uri": "string",
    /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
    "access_restriction": 0,
    /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Sõnastiku olek */
    "state": "C",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
    "start_date": "1998-11-01T00:00:00",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
    "end_date": "2010-11-01T00:00:00",
    /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: Masinloetav dokument */
    "machine_readable_document": {"name":"Masinloetav dokument", "url":"https://riha.eesti.ee/riha/onto/ettevotlus.ariregister/2010/r2"},
    /* Andmetüüp: object(document), Piirangud: -, Kirjeldus: Inimloetav dokument */
    "human_readable_document": {"name":"Inimloetav dokument", "filename":"humanr.rtf", "mime":"application/rtf", "content":"YmFzZTY0Cg=="},
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Versiooni kirjeldus */
    "version": "Versioon 2. (19.07.2010)",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
    "description": "string",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Valdkond, mille alla antud sõnastik kuulub */
    "area_name": "Äriregister"
}
Klassifikaatorite kirjeldamine JSON-struktuuriga
Klassifikaatori kirjeldus sisaldab
•	Standardsel kujul üldinfot
•	Loetelu objektidest, mis sisaldavad mis sisaldavad nende URL-e.
RIHA edasises arendamises soovitame tungivalt töötada välja klassifikaatorite unifitseeritud struktureeritud esitus, kust on võimalik konkreetsed väärtused ja kommentaarid eraldi välja eristada.

/* object(classifier): Infosüsteemi ja/või andmeobjekti puhul kasutatav klassifikaator */
"classifier": {
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Klassifikaatori unikaalne URI */
    "uri": "urn:fdc:riha.eesti.ee:2016:classifier:413423212",
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Klassifikaatori inimloetav nimi */
    "name": "Ametite klassifikaator",
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: Klassifikaatori haldaja asutus */
    "owner": "70006317",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Klassifikaatori lühinimi */
    "short_name": "AK",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimloetav versiooni nr, ei kasutata identifikaatorina */
    "version": "v2",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Versioonist sõltumatu hiearhia viit - klassifikaatori korral ei kasutata */
    "parent_uri": "string",
    /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
    "access_restriction": 0,
    /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: Klassifikaatori olek */
    "state": "C",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
    "start_date": "1998-11-01T00:00:00",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
    "end_date": "2010-11-01T00:00:00",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Alusklassifikaatori nimi */
    "base_classifier": "International Standard Classification of Occupations, 1988 (ISCO-88)",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Õiguslik alus */
    "legal_basis": "string",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Uuendamissagedus */
    "update_frequency": "Vastavalt rahvusvahelise klassifikaatori muutmisele",
    /* Andmetüüp: array(object(linked_classifier)), Piirangud: -, Kirjeldus: Sidusklassifikaator, eelnev klassifikaator, versioon */
    "related_classifier": [{
        /* Andmetüüp: string, Piirangud: enum(classifier_link_type), Kirjeldus: seose tüüp */
        "type": "string",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: seotud klassifikaatori uri */
        "uri": "string",
        /* Andmetüüp: string, Piirangud: -, Kirjeldus: Seose kirjeldus */
        "description": "string"
    }],
    /* Andmetüüp: array(object(document)), Piirangud: -, Kirjeldus: Klassifikaatori struktuur dokumendina */
    "structure": [{"name":"Klassifikaatori struktuur", "filename":"struktuur.rtf", "mime":"application/rtf", "content":"YmFzZTY0Cg=="}],
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Lühiiseloomustus */
    "short_description": "Klassifikaatori kaks põhimõistet on: sooritatava tegevuse laad/iseloom ehk TÖÖ ning vajalik väljaõpe ja tööoskus ehk KVALIFIKATSIOON ehk ametioskus",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Lisainformatsioon */
    "additional_information": "string",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Lõpetamise kirjeldus */
    "finishing_description": "string",
    /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Valdkondade nimed, mis on antud klassifikaatoriga seotud */
    "areas_names": ["Ametid"],
    /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: Kas on eeskujulikult kirjeldatud klassifikaator */
    "excellent": false,
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Objekti viitenumber RIHAs */
    "reference_number": "KL000511"
}
XML varade kirjeldamine JSON-struktuuriga
XML vara kirjeldus sisaldab
•	Standardsel kujul üldinfot
•	Loetelu objektidest, mis sisaldavad kas otse konkreetseid XML-tekste või objekte, mis sisaldavad nende URL-e.
XML varade edasine kirjeldamine on küsitava väärtusega, soovitame teostada selle vajaduse analüüs.

/* object(xmlasset): XML varad */
"xmlasset": {
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: XML vara unikaalne URI */
    "uri": "urn:fdc:riha.eesti.ee:2016:classifier:413423212",
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: XML vara inimloetav nimi */
    "name": "Dokumendihalduse metaandmed",
    /* Andmetüüp: string, Piirangud: not null, Kirjeldus: XML vara haldav asutus */
    "owner": "70006317",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: XML vara lühinimi */
    "short_name": "string",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: inimloetav versiooni nr, ei kasutata identifikaatorina */
    "version": "3.0",
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Versioonist sõltumatu hiearhia viit - xml vara korral ei kasutata */
    "parent_uri": "string",
    /* Andmetüüp: integer, Piirangud: enum(access_restriction), Kirjeldus: Juurdepääsupirang */
    "access_restriction": 0,
    /* Andmetüüp: string, Piirangud: enum(state), Kirjeldus: XML vara olek */
    "state": "C",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse algus */
    "start_date": "2013-02-01T00:00:00",
    /* Andmetüüp: datetime, Piirangud: -, Kirjeldus: versiooni kehtivuse lõpp */
    "end_date": "2016-01-01T00:00:00",
    /* Andmetüüp: object(translation), Piirangud: -, Kirjeldus: Võõrkeelsed nimed */
    "names": {"en":"Records Management Metadata"},
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Kirjeldus */
    "description": "Dokumendihalduse metaandmete loend ja skeemid on mõeldud rakendamiseks eelkõige Eesti avaliku sektori asutustes, kuid sobivad kasutamiseks ka teistes organisatsioonides. Eesmärgiks on ühtlustada metaandmeid erinevates dokumente haldavates süsteemides, et: 1) tagada dokumentide, nende konteksti ja haldamise ajaloo nõuetekohane kirjeldamine, 2) edendada elektroonilist dokumendivahetust, 3) ühtlustada dokumentide avalikustamisel kasutatavaid metaandmeid, 4) võimaldada statistilist aruandlust, 5) lihtsustada dokumendiliikide andmekirjelduste koostamist ning 6) lihtsustada dokumentide üleandmist Rahvusarhiivi. Metaandmete loendis on elementidele lisatud viited sama tähendusega elementidele, mida sisaldavad mõned teiste riikide või rahvusvahelised metaandmemudelid. Versioonis 3.0 on võetud arvesse uute ja uuendatud õigusaktide nõudeid ja vahepealseid rahvusvahelisi arenguid.",
    /* Andmetüüp: object(translation), Piirangud: -, Kirjeldus: Võõrkeelsed kirjeldused */
    "descriptions": {"en":"The Records Management Metadata set (the XML schemas and the list of elements) is designed to be applied primarily in the public sector agencies of Estonia, but is also suitable for other organisations. The objective of application of the metadata set is to harmonise the metadata in various systems managing documents and records in order to: 1) ensure proper describing of records, their context, and their management through time, 2) enhance electronic exchange of documents and records, 3) harmonise the metadata used in the disclosure of records, 4) enable statistical reporting, 5) simplify the development of data descriptions for record types, and 6) simplify the transfer of records to the National Archives. The list of metadata elements contains references to equivalent elements found in some foreign or international metadata models. In version 3.0, the requirements of new and amended legislative acts and international developments have been taken into account."},
    /* Andmetüüp: boolean, Piirangud: -, Kirjeldus: On arhiiviväärtusega */
    "archival_value": false,
    /* Andmetüüp: number, Piirangud: -, Kirjeldus: Säilitustähtaeg aastates */
    "retention_period": 5,
    /* Andmetüüp: string, Piirangud: enum(xmlasset_type), Kirjeldus: Liik */
    "type": "standardsed_metaandmed",
    /* Andmetüüp: array(string), Piirangud: -, Kirjeldus: Valdkondade nimed, mis on antud XML varaga seotud */
    "areas_names": ["Avalik haldus", "Riigi infosüsteem"],
    /* Andmetüüp: string, Piirangud: -, Kirjeldus: Objekti viitenumber RIHAs */
    "reference_number": "XV000240"
}
Veateated
Veateateid andmevahetuses ei liigu. Küll aga on võimalik RIHA rakenduse kasutajaliidesest saada informatsiooni kirjelduslehtede lugemise ja töötlemise tulemuse staatusest: millal viimati teostati, kas õnnestus või mis veateated salvestati. Samuti on võimalik testida kirjeldusi vastu RIHA test-veebilehte, mis annab võimalikud veateated kohe tagasi.

## Muutelugu
| Versioon, kuupäev	| Muudatus | 
|-------------------|----------|
| 0.9, 26.09.2016   | Täitja üleantud dokument viidud Markdowni, omistatud versiooninumber 0.9 |
