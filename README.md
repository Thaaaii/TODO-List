
# TODO-List
TODO-Liste als Webanwendung mit Basic HTML, CSS, Javascript im Frontend und Golang + Gin Framework + JWT Authentifizierung + SQLite Datenbank im Backend

## Disclaimer
- Wahl von HTML, CSS, Javascript statt Go Templates

## Anleitung

## Features

## TODO
- Kommentare und Dokumentation

## mögliche Verbesserungen
- Services besser trennen? Docker Container
- Funktionalität des Codes mit Tests belegen
- besseres Datenbankschema (Categories Tabelle überdenken)
- bessere Code-Qualität, Refactoring
- TODO-Listen teilen
- Anwendung robuster und sicherer machen
- Logging
- Description Felder mit Größe des Inputs wachsen lassen/richtiges Wrapping für Texte
- mehr Error-Nachrichten und Hinweise im Frontend in Abhängigkeit des Error Codes
- Debugging Code entfernen
- Nutzer löschen, Passwort ändern
- Responsive Design
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
- Registrierung,Login und Authentifikation mithilfe von JWT
- Einbindung von Tokens in das Frontend
- Änderung der Reihenfolge von Aufgaben

## Quellen
- Icon: https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/GNOME_Todo_icon_2019.svg/1200px-GNOME_Todo_icon_2019.svg.png
- TODO-Listen Design: https://www.youtube.com/watch?v=MkESyVB4oUw
- Check Icons: https://www.youtube.com/watch?v=G0jO8kUrg-I
- Textareas für Description: https://www.youtube.com/watch?v=0xGGe8bCahE
- Tag Icon: https://cdn-icons-png.flaticon.com/512/126/126422.png
- Kategorien: https://www.youtube.com/watch?v=BnXv1dwvebY
- Login Design: https://www.youtube.com/watch?v=L5WWrGMsnpw