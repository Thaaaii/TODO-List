# TODO-List
TODO-Liste als Webanwendung mit Basic HTML, CSS, Javascript im Frontend und Golang + Gin Framework + SQLite Datenbank im Backend

## Disclaimer
- wirklich schwer getan im Frontend
- kaum bis gar kein vorheriges Wissen zu Webtechnologien vor allem kaum Erfahrung mit Frontend und UI/UX Design
- größtenteils aufgebautes Fundament aus dem Studium genutzt in Kombination mit unten angegebenen Quellen
- Wahl von HTML, CSS, Javascript statt Go Templates
- mehr Zeit reingesteckt als geplant, aber konnte vieles lernen und das war schließlich auch das Ziel der Aufgabe

## Features

## Anleitung
- wie startet man den Server
- Startroute (Login)
- Logindaten eines Dummies
- alternativ: Registrierung und sofort loslegen
- Task einfügen
- abhaken geht immer
- möchte man Titel, Beschreibung oder Kategorien hinzufügen/löschen muss Bearbeitungsmodus ausgewählt werden (Änderungen müssen gespeichert werden)
- komplette Aufgaben können auch gelöscht werden
- Reihenfolge per Drag and Drop

## TODO
- Kommentare und Dokumentation
- Ordnerstruktur und Routen anpassen
- Refactoring

## mögliche Verbesserungen
- Services besser trennen? Docker Container
- Funktionalität des Codes mit Tests belegen
- besseres Datenbankschema (Categories Tabelle überdenken)
- bessere Code-Qualität, Refactoring
- bessere Struktur und Konsistenz, Debugging Code entfernen
- effizientere Lösungen finden, beispielsweise bei Modifikation von Aufgabenelementen
- TODO-Listen teilen
- Anwendung robuster und sicherer machen
- Logging
- Description Felder mit Größe des Inputs wachsen lassen/richtiges Wrapping für Texte
- mehr Error-Nachrichten und Hinweise im Frontend in Abhängigkeit des Error Codes
- Nutzer löschen, Passwort ändern
- hardcodierte (sensible) Daten in eine .env auslagern

## Ablauf
- Gedanken zu möglichen Datenstrukturen gemacht und entsprechendes Datenbankschema erstellt
- Routen zur Erstellung und Modifikation von Daten mit Gin aufgesetzt und getestet (mit cURL)
- Frontend Design der Todo Liste
  - Erstellung von Aufgabenelementen
  - Modifikation der Zustände von Aufgaben (Checkbox, Titel, Beschreibung, Kategorien)
  - Abruf von Daten aus dem Backend
  - Speicherung aller Änderungen im Frontend im Backend (Datenbank)
- Frontend Design des Logins
- Registrierung,Login und Authentifikation mithilfe von JWT als Cookie
- Anpassung der Routen im Front- + Backend und Einführung von Middleware, um Nutzerdaten zu schützen
- Änderung der Reihenfolge von Aufgaben
- Dokumentation und Kommentare von Quelltext
- README vervollständigen

## Quellen
- Icon: https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/GNOME_Todo_icon_2019.svg/1200px-GNOME_Todo_icon_2019.svg.png
- TODO-Listen Design: https://www.youtube.com/watch?v=MkESyVB4oUw
- Check Icons: https://www.youtube.com/watch?v=G0jO8kUrg-I
- Textareas für Description: https://www.youtube.com/watch?v=0xGGe8bCahE
- Tag Icon: https://cdn-icons-png.flaticon.com/512/126/126422.png
- Kategorien: https://www.youtube.com/watch?v=BnXv1dwvebY
- Login Design: https://www.youtube.com/watch?v=L5WWrGMsnpw
- Drag and Drop: https://www.youtube.com/watch?v=jfYWwQrtzzY
- JWT Authentifikation: https://www.sohamkamani.com/golang/jwt-authentication/