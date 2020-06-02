# Selecto
Interactive filter of stdin

stdin -> selecto -> stdout

# Examples
```bash
$ ll | selecto --stdin | cat
$ cat urls.txt | selecto --stdin | xargs wget
$ cat ~/.ssh/config | grep "Host " | awk '{print $2}' | ./bin/selecto --stdin | xargs -i ssh -tt {}
```

## Gena
Helper for generating string sequences

## Reedo
Helper for testing simple concept

# TODO

- help
- examples
- fuzzy filter by results
