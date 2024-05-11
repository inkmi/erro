<img src="https://raw.githubusercontent.com/StephanSchmidt/erro/master/ErroLogo.png" width="300">

# Erro: Faster and easier Go debugging in case of an error

![Build state](https://github.com/StephanSchmidt/erro/actions/workflows/test.yml/badge.svg)  ![Go Version](https://img.shields.io/github/go-mod/go-version/StephanSchmidt/erro) ![Version](https://img.shields.io/github/v/tag/StephanSchmidt/erro?include_prereleases)  ![Issues](https://img.shields.io/github/issues/StephanSchmidt/erro) ![Report](https://goreportcard.com/badge/github.com/StephanSchmidt/erro)

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-79%25-brightgreen.svg?longCache=true&style=flat)</a>

*Initially forked from [snwfdhmp/errlog](https://github.com/snwfdhmp/errlog)*

Erro (pronounced *'arrow'*) is a simple tool that helps finding the reasons for errors in your code.

Currently supported languages:
* Golang

Piping your structured log into `erro`

```
yourProgram 2>&1 | erro
```

turns

![Erro example structured log](https://raw.githubusercontent.com/inkmi/erro/master/ErrorStructured.png)

into

![Erro example pipe output](https://raw.githubusercontent.com/inkmi/erro/master/ErroPipe.png)

In an IDE like JetBrains Golang, the line where the error
comes from is clickable and takes you to that line.

![Erro example Jetbrains output](https://raw.githubusercontent.com/inkmi/erro/master/ErroPipeIdea.png)

## Install

```shell
go install github.com/Inkmi/erro
```

The log data needs to be structured as JSON.

The log data needs to include caller information, for example

`{ "caller": "testerro.go:22" }`

Configuration depends on your logging library, for Zerolog it's

`logger := log.With().Caller().Logger()`

After `make build` you can test `erro` with

```shell
./bin/testerro 2>&1| ./bin/erro
```

## License information

MIT License

## Contributors

### Major

- [Stephan Schmidt](https://github.com/StephanSchmidt): Author and maintainer
- Martin Joly wrote the original [errlog](https://github.com/snwfdhmp/errlog) which this code is based on
