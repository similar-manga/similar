# Similar Manga Recommendations

This repo has both the scraping and matching utilities to find mangas which are close in content to others. The idea is
to create a recommendation system outside MangaDex since there isn't one and thus allow for users to discover other
content. Right now it is pretty rudimentary, but seems to give ok results most of the time. If you want to use it,
please download [Neko](https://github.com/CarlosEsco/Neko) and enable recommendations inside the app or use
the [Chrome Extension](https://chrome.google.com/webstore/detail/mangadex-similar-manga/ghjofcifjnbpgfjbbnbeeipjlhnmbefl).


## Setup / Dependencies

Client was generated using [swagger](https://editor.swagger.io/). You will need to setup
a [golang workspace](https://golang.org/doc/install), and then run the following commands. Only manga need to be
downloaded / scraped from mangadex to be able to perform similar manga identification.

```
go get golang.org/x/oauth2
go get github.com/antihax/optional
go get github.com/james-bowman/nlp
go get github.com/caneroj1/stemmer
go get github.com/PuerkitoBio/goquery
```

## Runtime Instructions


1. Run `./1_cache_mangas` to get local list of all mangas on the site (7 hours)
2. Run `./2_cache_chapters` to download number of chapters available and in what language for each manga (6 hours)
3. Run `./3_calc_similar` to generate similarity index between manga (1.5 hours)
4. Run `./4_calc_mappings` to export external ID to mangadex UUID mapping files (2 minutes)
4. Run `./5_calc_mu_ids` convert MangaUpdate ids into uniform "new" ids for their new api (24 hours)
5. Run `python utils/create_sqlite.py` to generate mapping file for [Neko](https://github.com/CarlosEsco/Neko/tree/master/app/src/main/assets)


## Manga Links Data

<table>
<thead>
<tr>
<th>Key</th>
<th>Related site</th>
<th>URL</th>
<th>URL details</th>
</tr>
</thead>
<tbody><tr>
<td>al</td>
<td>anilist</td>
<td><a href="https://anilist.co/manga/%60%7Bid%7D%60">https://anilist.co/manga/`{id}`</a></td>
<td>Stored as id</td>
</tr>
<tr>
<td>ap</td>
<td>animeplanet</td>
<td><a href="https://www.anime-planet.com/manga/%60%7Bslug%7D%60">https://www.anime-planet.com/manga/`{slug}`</a></td>
<td>Stored as slug</td>
</tr>
<tr>
<td>bw</td>
<td>bookwalker.jp</td>
<td><a href="https://bookwalker.jp/%60%7Bslug%7D%60">https://bookwalker.jp/`{slug}`</a></td>
<td>Stored has "series/{id}"</td>
</tr>
<tr>
<td>mu</td>
<td>mangaupdates</td>
<td><a href="https://www.mangaupdates.com/series.html?id=%60%7Bid%7D%60">https://www.mangaupdates.com/series.html?id=`{id}`</a></td>
<td>Stored has id</td>
</tr>
<tr>
<td>nu</td>
<td>novelupdates</td>
<td><a href="https://www.novelupdates.com/series/%60%7Bslug%7D%60">https://www.novelupdates.com/series/`{slug}`</a></td>
<td>Stored has slug</td>
</tr>
<tr>
<td>kt</td>
<td>kitsu.io</td>
<td><a href="https://kitsu.io/api/edge/manga/%60%7Bid%7D%60">https://kitsu.io/api/edge/manga/`{id}`</a> or <a href="https://kitsu.io/api/edge/manga?filter%5Bslug%5D=%7Bslug%7D">https://kitsu.io/api/edge/manga?filter[slug]={slug}</a></td>
<td>If integer, use id version of the URL, otherwise use slug one</td>
</tr>
<tr>
<td>amz</td>
<td>amazon</td>
<td>N/A</td>
<td>Stored as full URL</td>
</tr>
<tr>
<td>ebj</td>
<td>ebookjapan</td>
<td>N/A</td>
<td>Stored as full URL</td>
</tr>
<tr>
<td>mal</td>
<td>myanimelist</td>
<td><a href="https://myanimelist.net/manga/%7Bid%7D">https://myanimelist.net/manga/{id}</a></td>
<td>Store as id</td>
</tr>
<tr>
<td>raw</td>
<td>N/A</td>
<td>N/A</td>
<td>Stored as full URL, untranslated stuff URL (original language)</td>
</tr>
<tr>
<td>engtl</td>
<td>N/A</td>
<td>N/A</td>
<td>Stored as full URL, official english licenced URL</td>
</tr>
</tbody></table>


