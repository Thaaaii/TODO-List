# TODO-List
TODO-Liste als Webanwendung mit Basic HTML, CSS, Javascript im Frontend und Golang + Gin Framework + SQLite Datenbank im Backend

## Anmerkungen
Diese Anwendung wurde innerhalb einer Woche entwickelt als Aufgabe für ein Praktika, um meine aus dem Studium erworbenen Fähigkeiten zu testen.
Ich habe viel mehr Zeit reingesteckt als eigentlich geplant war, aber ich konnte viele Sachen und Konzepte lernen, was auch schließlich das Ziel der Aufgabe war.
Problematisch für mich war größenteils das Frontend und im Allgemeinen fehlendes/mangelhaftes Wissen zu Webtechnologien, welche ich mithilfe der Aufgabe zumindest ein bisschen
selbst erarbeiten konnte. Viel Zeit ist ins recherchieren, ausprobieren, testen von Funktionalität und beheben von Fehlern gegangen und
viele Sachen habe ich zum ersten Mal gesehen, d.h. Best-Practices, eine gute Struktur und Code-Qualität sind nicht gewährleistet. Da ich schlechte Erfahrung mit der Erstellung von Frontends
machen musste und nicht wirklich ein UI/UX Experte bin, habe ich mir Designs aus Videos abgeschaut, welche in den Quellen unten angegeben wurden. Außerdem findet man dort Bilder und weitere Tutorials,
welche ich genutzt und für meine Zwecke angepasst habe.

## Features
- Erstellung von Nutzern und Login
- erstellen neuer Aufgaben durch Eingabe eines Titels
- Auflistung aller Aufgaben
- Änderung des Titels, der Beschreibung, der Kategorien und des Erledigungsstatus
- Löschen von Aufgaben
- verschieben von Aufgaben

## Anleitung
- Anwendung starten, indem main.go ausgeführt wird
```bash
go run main.go 
```
- im Browser die [Anwendung](http://localhost:8080/login) öffnen (http://localhost:8080/login)
- Logindaten eines Dummies eingeben (Nutzer: Thai, Passwort: 123)
- alternativ: Registrierung eines neuen Nutzers und sofort loslegen
- Task einfügen, indem ein Titel eingegeben wird und mit "Aufgabe hinzufügen" bestätigt wird
- Aufgaben kann man immer als "Erledigt" markieren, indem die Checkbox oben links in der Ecke eines Aufgabenelements angeklickt wird
- möchte man Titel, Beschreibung oder Kategorien hinzufügen/löschen muss der Bearbeitungsmodus durch betätigen von "Bearbeiten" ausgewählt werden (Hinweis: Änderungen müssen gespeichert werden durch "Speichern")
- Kategorien können per Komma getrennt eingegeben werden, um mehrere Label hinzuzufügen und müssen mit der Enter-Taste bestätigt und anschließend gespeichert werden
- Kategorien können nur im Bearbeitungsmodus gelöscht werden, indem auf das Kreuz neben dem Label geklickt wird
- komplette Aufgaben können auch gelöscht werden durch betätigen von "Löschen"
- die Reihenfolge kann per Drag and Drop geändert werden, indem die gezogenen Elemente ober- bzw. unterhalb von anderen Elementen gezogen werden
- durch betätigen des Logout Buttons wird man wieder zurück zum Login Fenster geführt

## mögliche Verbesserungen
- Services besser trennen? eventuell Nutzung von Docker Container
- Funktionalität des Codes mit Tests belegen
- besseres Datenbankschema (Categories Tabelle überdenken); effizientere Datenstrukturen zum Austausch von Daten
- bessere Code-Qualität; Refactoring
- bessere Struktur und Konsistenz; Debugging Code entfernen
- effizientere Lösungen finden, beispielsweise bei Modifikation von Aufgabenelementen und Datenbankanfragen
- Anwendung robuster und sicherer machen
- Logging
- Description Felder mit Größe des Inputs wachsen lassen; richtiges Wrapping für Texte
- mehr Error-Nachrichten und Hinweise im Frontend in Abhängigkeit des Error Codes
- Nutzer löschen; Passwort ändern
- hardcodierte "sensible" Daten in eine .env auslagern
- TODO-Listen teilen

## Ablauf
- Gedanken zu möglichen Datenstrukturen gemacht und entsprechendes Datenbankschema erstellt
- Routen zur Erstellung und Modifikation von Daten mit Gin aufgesetzt und getestet (mit cURL)
- Frontend Design der Todo Liste
  - Erstellung von Aufgabenelementen
  - Modifikation der Zustände von Aufgaben (Checkbox, Titel, Beschreibung, Kategorien)
  - Abruf von Daten aus dem Backend
  - Speicherung aller Änderungen im Frontend im Backend (Datenbank)
- Frontend Design des Logins
- Registrierung, Login und Authentifikation mithilfe von JWT als Cookie
- Anpassung der Routen im Front- + Backend und Einführung von Middleware, um Nutzerdaten zu schützen
- Änderung der Reihenfolge von Aufgaben
- Kommentare im Quelltext hinzugefügt
- README vervollständigt

## Quellen
- [Icon](https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/GNOME_Todo_icon_2019.svg/1200px-GNOME_Todo_icon_2019.svg.png)
- [TODO-Listen Design](https://www.youtube.com/watch?v=MkESyVB4oUw)
- [Check Icons](https://www.youtube.com/watch?v=G0jO8kUrg-I)
- [Textareas für Description](https://www.youtube.com/watch?v=0xGGe8bCahE)
- [Tag Icon](https://cdn-icons-png.flaticon.com/512/126/126422.png)
- [Kategorien](https://www.youtube.com/watch?v=BnXv1dwvebY)
- [Login Design](https://www.youtube.com/watch?v=L5WWrGMsnpw)
- [Drag and Drop](https://www.youtube.com/watch?v=jfYWwQrtzzY)
- [JWT Authentifikation](https://www.sohamkamani.com/golang/jwt-authentication/)
