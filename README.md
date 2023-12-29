# Fedorov

Fedorov is a local library microservice to get, sync, serve metadata, digital artifacts from your LitRes account. Can be used as a CLI app.

Fedorov is named in honor of [Ivan Fedorov](https://en.wikipedia.org/wiki/Ivan_Fyodorov_(printer)) ([Иван Фёдоров](https://ru.wikipedia.org/wiki/Иван_Фёдоров)).

## Installation

The recommended way to install `fedorov` is with docker-compose:

- create a `docker-compose.yaml` fil (this minimal example omits common settings like network, restart, etc):
- NOTE: cold storage signifies resources used less frequently
- NOTE: hot storage signifies resources used on most page loads

```yaml
version: '3'
services:
  fedorov:
    container_name: fedorov
    image: ghcr.io/beauxarts/fedorov:latest
    environment:
      # - FV_SERVE_ADMIN-USERNAME=ADMIN-USERNAME
      # - FV_SERVE_ADMIN-PASSWORD=ADMIN-PASSWORD
      # - FV_SERVE_SHARED-USERNAME=SHARED-USERNAME
      # - FV_SERVE_SHARED-PASSWORD=SHARED-PASSWORD
      # - FV_WEBHOOK-URL=http://FEDOROV-ADDRESS/prerender
    volumes:
      # backups (cold storage)
      - /docker/fedorov/backups:/var/lib/fedorov/backups
      # metadata dir (hot storage)
      - /docker/fedorov/metadata:/var/lib/fedorov/metadata
      # input dir (cold storage)
      - /docker/fedorov:/var/lib/fedorov/input
      # output dir (cold storage)
      - /docker/fedorov:/var/lib/fedorov/output
      # covers dir (hot storage)
      - /docker/fedorov/covers:/var/lib/fedorov/covers
      # downloads dir (cold storage)
      - /docker/fedorov/downloads:/var/lib/fedorov/downloads
      # imported dir (cold storage)
      - /docker/fedorov/_imported:/var/lib/fedorov/_imported
      # sharing timezone from the host
      - /etc/localtime:/etc/localtime:ro
    ports:
      # https://en.wikipedia.org/wiki/Ivan_Fyodorov_(printer)#Biography
      - "1510:1510"
```
- (move it to location of your choice, e.g. `/docker/fedorov` or remote server or anywhere else)
- while in the directory with that config - pull the image with `docker-compose pull`
- start the service with `docker-compose up -d`

## Getting started

After you've installed `fedorov`, you need to authenticate your LitRes.ru username / password.
Please note - your credentials are not stored by `fedorov` and only used to get session cookies from LitRes.ru -
exactly the same way you would log in to a website.

To do that you'll need to import your cookies from existing browser session. To do that you need to create `cookies.txt` in the `root state dir` folder (see [docker installation](#Installation)),
then follow [instructions here](https://github.com/boggydigital/coost#copying-session-cookies-from-an-existing-browser-session) to copy `litres.ru` cookies into that file. When you run `fedorov` for the first time, it'll import the cookie header value and split individual parameters.

Regardless of how you do it, the content of `cookies.txt` should look like this:

```text
litres.ru
 cookie-header=...
```

If you want to de-authorize `fedorov` from accessing your LitRes.ru data - delete the `cookies.txt` file. All the account specific data you'll have accumulated until that point will be preserved. 