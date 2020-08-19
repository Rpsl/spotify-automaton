# spotify-automaton

At one moment i will wanted to make utility for make some task in my own spotify profile.

One thing of i want it was automatic create playlists by genres from my liked songs.

When i was did first part (download meta info about my library) i lerned that many tracks in spotify don't have genre. 🤷‍♂️

After that i don't have motivation for continue write that utility.


## How to use

### build
```
$ go build -o spotify-automaton .
``` 

### create Spotify Application

- https://developer.spotify.com/dashboard/applications
- create application
- rename `config.example.toml` to `config.toml`
- write spotify_client_id and other into `config.toml` 

### show help
```
$ ./spotify-automaton

Usage: Spotify Automaton COMMAND [arg...]

Utility for automation some tasks in your spotify account

Commands:
  login        get credential for login
  refresh      refresh local database

Run 'Spotify Automaton COMMAND --help' for more information on a command.
```

### get credential
```
$ ./spotify-automaton login 

# follow the instructions. 
# when your browser will be open you need copy code from url and paste into console
```

### create local database of favorite tracks
```
$ ./spotify-automaton refresh

# tracks with meta info will be added into sqlite db ( ./db/base.db )
```