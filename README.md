# GoSort

GoSort est un petit outil en ligne de commande écrit en Go qui range automatiquement tes fichiers en fonction de leur extension. Il scanne un dossier (par exemple ton dossier Téléchargements) et déplace chaque fichier vers un dossier dédié selon son type : images, vidéos, documents, scripts, archives, applications, audios, etc.

## Pourquoi

Fini les Téléchargements en bazar : GoSort trie tout automatiquement en une commande, sans avoir à le faire à la main.

## Utilisation

```bash
go run ./cmd/gosort
```

Le programme te demandera confirmation avant de déplacer les fichiers.

## Configuration

Les extensions gérées et les dossiers de destination sont définis dans `internal/config/config.go`. Tu peux librement ajouter, retirer ou modifier des catégories selon tes besoins.

## Structure du projet

```
gosort/
├── cmd/gosort/       — point d'entrée
├── internal/config/  — extensions & chemins
├── internal/design/  — bannière & couleurs
└── internal/sort/    — logique de tri des fichiers
```
