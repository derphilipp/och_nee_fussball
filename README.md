# Och nee, Fussball

Hier gibt es eine Liste aller potentiellen Paarungen (Hashtags) der Deutschen Fussballmannschaften der Männer;

[Liste der Paarungen](https://github.com/derphilipp/och_nee_fussball/blob/paarungen/paarungen.txt)

## Geht automatisiertes Anlegen der Listen?

Ja, geht:

1. Filtertool hier installieren: https://github.com/hiway/mastodon-filter ; Es muss dann in `$PATH` vorhanden sein / ausführbar sein
2. [Liste der Paarungen](https://github.com/derphilipp/och_nee_fussball/blob/paarungen/paarungen.txt) herunterladen und als paarungen.txt abspeichern (alternativ: selbst erstellen)
3. Der Anleitung unter https://github.com/hiway/mastodon-filter nach einen Zugangsschlüssel erhalten
4. Der Anleitung unter https://github.com/hiway/mastodon-filter nach das Tool konfiguieren und mit dem Schlüssel und der Instanz verbinden
5. `./split.sh` ausführen. Das trennt die `paarungen.txt` in Einzeldateien auf (Da die Liste sonst zu lang wird) und ruft `mastodon-filter` Aufruf lauter Einträge an;

Jetzt hast Du eine Liste von Filtern, mit den Namen `Och_nee_Fussball_000` und aufsteigend

## Das ist ja alles super trashig und so!

Stimmt. 
Ich hab keine Lust mich mehr mit Fussball zu beschäftigen - und daher auch dem Projekt hier. Wenn Du Lust hast etwas zu verbessern oder beizutragen - super, immer gerne, die PR sind offen und ich freue mich.

## Was hast Du gegen Fussball?

Nix, ich schau mir auch mal was an mit Freunden, aber mich interessiert es einfach gar nicht. Da ich Freunden nicht die Freude nehmen will und Leute ruhig auf Mastodon darüber quatschen können, will nur ich das nicht; Red mit mir lieber über Apfelkuchen und Löten.
