### Personal opinion.

About code structure:

I Would structure the code a little different, inspired by: https://github.com/golang-standards/project-layout
* main.go in the root of the repo (very minimalistic, only calling _cli.execute()_ )
* internal/hakuna.go as or package/hakuna.go
* (and personally) cmd/hakuna.go with the logic about the cli