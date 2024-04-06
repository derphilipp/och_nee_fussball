package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type VereineListe struct {
	Vereine []map[string][]string
}

func main() {
	// Einlesen der JSON-Datei
	data, err := os.ReadFile("maenner.json")
	if err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
		return
	}

	var vereine VereineListe
	err = json.Unmarshal(data, &vereine)
	if err != nil {
		fmt.Println("Fehler beim Parsen der JSON-Daten:", err)
		return
	}

	// Extraktion der Kürzel
	var kuerzel []string
	for _, verein := range vereine.Vereine {
		for _, v := range verein {
			kuerzel = append(kuerzel, v...)
		}
	}

	// Erstellung der Paarungen
	paarungenMap := make(map[string]bool)
	for i, k1 := range kuerzel {
		for j, k2 := range kuerzel {
			if i != j && !paarungenMap[k1+k2] {
				paarungenMap[k1+k2] = true
				paarungenMap[k2+k1] = true // Verhindert Duplikate in umgekehrter Reihenfolge
			}
		}
	}

	// Umwandlung der Map zurück in eine Slice
	var paarungen []string
	for paarung := range paarungenMap {
		paarungen = append(paarungen, paarung)
	}

	// Sortierung der Paarungen
	sort.Strings(paarungen)

	// Ausgabe der Paarungen in eine Datei
	file, err := os.Create("paarungen.txt")
	if err != nil {
		fmt.Println("Fehler beim Erstellen der Datei:", err)
		return
	}
	defer file.Close()

	for _, paarung := range paarungen {
		_, err := file.WriteString(paarung + "\n")
		if err != nil {
			fmt.Println("Fehler beim Schreiben in die Datei:", err)
			return
		}
	}

	fmt.Println("Paarungen erfolgreich in paarungen.txt gespeichert.")
}
