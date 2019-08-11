# Checkmissing

It checks the missing episodes from a TV show folder passed as an argument. It analyzes the folder's name to retrieve the show and all its episodes from the [TVDB](https://www.thetvdb.com/). 
To retrieve the show's info from the TVDB, you need to have credentials to communicate with TVDB API. Those credentials then need to be set as environment variables and for this you have 2 choices:
* load from .env file
* load from classic env var

In root project, you'll find the `.env.dist` as an example of those variables.

## Folders' architecture 

The series folder must follow this specific structure : 
```
Game of thrones 
│
└───Season 1 (Or Saison 1 if you're french like me ^^)
│   │   Game.of.Thrones.S01E01...
│   │   Game.of.Thrones.S01E05...
│   
└───Season 2
    │   Game.of.Thrones.S02E01...
    │   Game.of.Thrones.S02E06...
    │   ...
```

E.g it will output :

```
Missing episodes are :
* S01E8
* S01E2
* S01E3
* S01E4
* S01E6
* S01E7
* S01E9
* S01E10
* S02E2
* S02E3
* S02E4
* S02E5
* S02E7
* S02E8
* S02E9
* S02E10
Missing seasons are :
* Season 3
* Season 4
* Season 5
* Season 6
* Season 7
* Season 8
```
## Contributing

This is my first project as a go developer. I've got a lot to learn but if you want to help feel free to fork and create merge requests to improve it and I will take a look.

## Todo 

* ~~check missing seasons~~
* multiple folders in args
* an UI