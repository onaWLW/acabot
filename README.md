# acabot

## Déploiement

- Cloner le projet
- Avoir d'installé go 1.25.5 et sqlite3
- mettre à la racine du projet un .env

exemple de .env :
```
SQLITE_DATABASE_PATH=scores.db
DISCORD_TOKEN=xyz
```

- `go build`
- `./acabot`
