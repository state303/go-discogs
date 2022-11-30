# go-discogs

### Discogs Dump Batch Updater with Go

### Examples

Run with shorthand and full flag with config

```shell
go-discogs --update-marker -y 2010 --month 3 -f db-config.yaml
```

You would most likely wish to keep the dsn variable to ENV.
If that is the case...

```shell
export GO_DISCOGS_DSN=$my_dsn # or something else...
go-discogs --update \         # will read from ENV
     -y 2010 \
     -m 3 
```

If all you need is to update artist AND label...

```shell
go-discogs --update \           # optional if you wish to just use local marker
     -t artist,labels \         # you call artist or artists... potato, potato. just the same.
     -y 2010 \                  # required
     -m 3 \                     # required
     -s postgres://username:password@localhost:5432/database_name   # required
```

### Commands

| FLAG        | HAS_VALUE | DEFAULT                      | NOTE                           |
|-------------|-----------|------------------------------|--------------------------------|
| --chunk -b  | O         | 5000                         | Chunk size for batch insertion |
| --config -c | O         | $HOME/go-discogs/config.yaml | Config file location           |
| --data -d   | O         | $HOME/go-discogs/            | File directory                 |
| --dsn -s    | O         | X                            | Required                       |
| --year -y   | O         | Current Year                 | Target Year                    |
| --month -m  | O         | Current Month                | Target Month                   |
| --types -t  | O         | artists                      | Batch these types              |
| --update -u | X         | false                        | Update data dump records       |
| --purge -p  | X         | false                        | Keep files after batch         |
| --new -n    | X         | false                        | Keep files after batch         |

### 💾 Files

#### Dump XML.GZ files

By default, it uses a local file to read and dispatch insertions from data.discogs.com.
You can decide where to store them via --data (or just -d) flag.
It will not delete such data, as their size are large enough to be considered costly work.
If you insist to delete them after batch job, please use --purge (or just -p) option.

####   

```shell
# first run
# disabling marker will prevent usage or creation of local marker
go-discogs --no-marker -f config.yaml

# second run, but there is no local marker
# initially updates and stores marker; because it cannot find any local marker
go-discogs -f config.yaml

# third run will not update, but simply use the local marker
go-discogs -f config.yaml

# fourth run will update local marker
go-discogs -f config.yaml --update-marker
```

### Database Connection

Database connection can be done with following options.
