# Checkmissing

Checkmissing checks if some episodes are missing from the folder passed as an argument. It analyze the folder's name to retrieve the series and all its episodes from the [TVDB](https://www.thetvdb.com/) . Then check which ones are not in the series folder.

## Folders' architecture 

```
Game of thrones 
│
└───Season 1
│   │   Game.of.Thrones.S01E01...
│   │   Game.of.Thrones.S01E05...
│   
└───Season 2
    │   Game.of.Thrones.S02E01...
    │   Game.of.Thrones.S02E06...
    │   ...
```

## Todo 
* ~~check missing seasons~~
* multiple folders in args
* an UI