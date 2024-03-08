# Materials for the Let's Go book

This is a web application for creating and viewing code snippets, similar to GitHub gists

## Project structure
- `cmd` - contains applications within this project. For example, `web` contains the web application,
  could also have a `cli` folder if this contained a CLI application as well
- `internal` - non-application specific code such as reusable internal libraries or SQL files. This is 
  a special folder name in Go which only allows packages within this directory to be imported by Go code 
  within this project. So if this project were distributed as a package, people who install the package 
  could not import anything from `internal`.
- `html` - HTML templates
- `static` - CSS and JavaScript

## Web application
Run the application:
```bash
go run ./cmd/web/
```
