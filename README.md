This repo contains a scraper for the 2019 NPSL Members Cup Rosters and the DCFC headshots on detcityfc.com. To use, install Colly:

```
go get -u github.com/gocolly/colly/...
```

For rosters: run scraper.go. The resulting JSON file is compatible with the Hooligan Hymnal/Guardbook's FOES MAD feature.

For headshots: run headshots.go. This will output one JPEG per player into the current directory.
