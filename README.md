<img src="https://raw.githubusercontent.com/StephanSchmidt/erro/master/ErroLogo.png" width="300">

# Erro: Faster and easier debugging in case of an error

![Build state](https://github.com/StephanSchmidt/erro/actions/workflows/test.yml/badge.svg)  ![Go Version](https://img.shields.io/github/go-mod/go-version/StephanSchmidt/erro) ![Version](https://img.shields.io/github/v/tag/StephanSchmidt/erro?include_prereleases)  ![Issues](https://img.shields.io/github/issues/StephanSchmidt/erro) ![Report](https://goreportcard.com/badge/github.com/StephanSchmidt/erro)

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-71%25-brightgreen.svg?longCache=true&style=flat)</a>

*Forked from [snwfdhmp/errlog](https://github.com/snwfdhmp/errlog)*

Erro (pronounced *'arrow'*) is a simple libary that helps finding the reasons for errors in your Golang code. When an error is created through **erro**, the source of the failing code is shown together with the stack trace and the variables that led to the error.

![Erro example outpuit](https://raw.githubusercontent.com/StephanSchmidt/erro/master/ErroExample.png)


## Introduction

Use Erro to help understanding errors and **speed up debugging while you create amazing code** :

- Highlight source code
- **Detect and point out** which func call is causing the fail
- Pretty stack trace
- **No-op mode** for production

## Get started

### Install

```shell
go get github.com/StephanSchmidt/erro
```

## License information

MIT License

## Contributors

### Major

- [Stephan Schmidt](https://github.com/StephanSchmidt): Author and maintainer
- Martin Joly wrote the original [errlog](https://github.com/snwfdhmp/errlog) which this code is based on
