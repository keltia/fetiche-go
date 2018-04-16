# fetiche / fetiched

fetiched is a web-scrapping/monitoring daemon for full/partial pages.

## Goals

- having a list of http/https endpoints
- for each one, get the page headers and/or content
- notify of every change whole page/partial ?
- diff ?
- report mail/rss/… ?

## Building blocks

- list of URLs with some parameters
- site w/ multiple urls or just list of urls?
- cache of web pages
- CLI interface
- Web interface?

### CLI

- url   add/info/rm/list
- site  info/rm/list (add is implicit)
- cache info/list/rm/clear
- report
- import site list? Bookmarks?

### fetiched

- endpoint on lcoalhost for cli submission
- demon with goroutines
- rss support?
- queuing system to handle scans
- report generation
- web interface?
- rss generation?

## Data structures

### Site

- name
- freq
- retention

### URL

- site (site:name)
- url
- login/password
- css/xpath?        (partial content)
- fetch + cache (db? fs?)
- sha2 of content, store that? (no diff on whole or partial page)
- fetch page, strip non essential, hash that?

### Cache

- ID        unique/key
- url
- date
- content

## Tools

- bbolt + storm?
- badger?
- pg?


