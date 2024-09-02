# CSV Parser

Diese README.md dient als Anleitung zur Konfiguration des CSV-Parsers mithilfe der `config.yml`-Datei. Diese Datei ermöglicht es, bestimmte Spalten aus einer CSV-Datei auszuwählen und deren Inhalte bei Bedarf anzupassen, indem Präfixe oder Suffixe hinzugefügt werden.

## Konfigurationsformat

Die Konfigurationsdatei `config.yml` definiert eine Liste von Spalten, die der CSV-Parser verarbeiten soll. Jede Spalte wird durch ein Objekt innerhalb der Liste `columns` dargestellt. Die Konfiguration jeder Spalte kann folgende Parameter enthalten:

- `column`: Der Name der Spalte im Ergebnis (erforderlich).
- `name`: Der Name der Spalte, wie sie in der CSV-Datei benannt ist. Entweder `name` oder `index` muss angegeben werden, um die Spalte zu identifizieren.
- `index`: Der Index der Spalte in der CSV-Datei (beginnend bei 1). Alternativ zu `name` kann `index` verwendet werden, um die Spalte zu identifizieren.
- `suffix`: Ein Suffix, das an den Wert der Spalte angehängt wird (optional).
- `prefix`: Ein Präfix, das an den Wert der Spalte vorangestellt wird (optional).
- `output`: Der Name der Ausgabedatei, in die die gefilterten und bearbeiteten Daten geschrieben werden (optional, Standard ist `result.csv`).
- `delimiter`: Das Trennzeichen, das in der Ausgabedatei verwendet werden soll (optional, Standard ist `;`).
- `ignore_header`: Ein boolescher Wert, der angibt, ob die Kopfzeile der CSV-Datei ignoriert werden soll (optional, Standard ist `false`).

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
  - column: "Kosten"
    name: "Ausgehandelter Gesamtbetrag"
    suffix: " €"
output: "result.csv"
delimiter: ","
ignore_header: true
```