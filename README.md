# Selecto

## deprecated
This project is reinvention of [fzf](https://github.com/junegunn/fzf)

Interactive filter of stdin

stdin -> selecto -> stdout

# Examples
```bash
# download file from list of urls
$ cat urls.txt | selecto --stdin | xargs wget
# show logs of some docker container
$ docker ps | tr -s ' ' | cut -d ' ' -f 1,2 | selecto --stdin | awk '{printf("%02d",$1)}' | xargs -0 -I {} docker logs {} -f
```

## Gena
Helper for generating string sequences

## Reedo
Helper for testing simple concept

# TODO

- help
- examples
- fuzzy filter by results
