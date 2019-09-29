This repo contains a scraper for the opponent rosters and the DCFC headshots on detcityfc.com. To use, install Colly:

```
go get -u github.com/gocolly/colly/...
```

For 2019 NPSLMC rosters: run scraper.go. The resulting JSON file is compatible with the Hooligan Hymnal/Guardbook's FOES MAD feature.

For 2019 UWS rosters: run scraper.go. This is also for the FOES MAD feature, and is a prototype for the 2020 season.

For headshots: run headshots.go. This will output one JPEG per player into the current directory.
