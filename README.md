# CSV Parser

Diese README.md dient als Anleitung zur Konfiguration des CSV-Parsers mithilfe der `config.yml`-Datei. Diese Datei ermöglicht es, bestimmte Spalten aus einer CSV-Datei auszuwählen und deren Inhalte bei Bedarf anzupassen, indem Präfixe oder Suffixe hinzugefügt werden.

## Konfigurationsformat

Die Konfigurationsdatei `config.yml` definiert eine Liste von Spalten, die der CSV-Parser verarbeiten soll. Jede Spalte wird durch ein Objekt innerhalb der Liste `columns` dargestellt. Die Konfiguration jeder Spalte kann folgende Parameter enthalten:

| Parameter       | Beschreibung                                                                              | Optional | Standard     |
| --------------- | ----------------------------------------------------------------------------------------- | -------- | ------------ |
| `column`        | Der Name der Spalte im Ergebnis.                                                          | Nein     | -            |
| `name`          | Der Name der Spalte in der CSV-Datei. Entweder `name` oder `index` muss angegeben werden. | Ja       | -            |
| `index`         | Der Index der Spalte in der CSV-Datei (beginnend bei 1). Alternativ zu `name`.            | Ja       | -            |
| `suffix`        | Ein Suffix, das an den Wert der Spalte angehängt wird.                                    | Ja       | -            |
| `prefix`        | Ein Präfix, das an den Wert der Spalte vorangestellt wird.                                | Ja       | -            |
| `replace`       | Eine Map von Werten, die ersetzt werden sollen.                                           | Ja       | -            |
| `output`        | Der Name der Ausgabedatei.                                                                | Ja       | `result.csv` |
| `delimiter`     | Das Trennzeichen in der Ausgabedatei.                                                     | Ja       | `;`          |
| `ignore_header` | Ob die Kopfzeile der CSV-Datei ignoriert werden soll.                                     | Ja       | `false`      |

## Beispiel-Konfiguration

Nachfolgend ist ein Beispiel für eine `config.yml`-Datei, die zeigt, wie die verschiedenen Parameter verwendet werden können:

```yaml
columns:
  - column: "Datum"
    name: "Versanddatum"
  - column: "Trackingnummer"
    index: 5
  - column: "Empfänger"
    index: 41
    replace:
      "Herr": "Mr."
      "Frau": "Ms."
  - column: "Kosten"
    name: "Ausgehandelter Gesamtbetrag"
    suffix: " €"
output: "result.csv"
delimiter: ","
ignore_header: true
```