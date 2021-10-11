# General

This project follows the project layout from https://github.com/golang-standards/project-layout
At first glance, maybe seems like a overkill, but a good base is always better
for the lifetime of a project

# Development

For development this project uses [air](https://github.com/cosmtrek/air).
This is not mandatory but helps a lot because the use of live reloading in the web
server.

```sh
$ make run-development-with-air # but you must have the `air` binary in your PATH
$ # or you can use the plain go command with:
$ make run-development # but you have to reaload manually the server
```
