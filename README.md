<img src="https://raw.githubusercontent.com/StephanSchmidt/erro/master/ErroLogo.png" width="300">

# Erro

*Forked from [snwfdhmp/errlog](https://github.com/snwfdhmp/errlog)*

Erro is a simple libary that helps finding the reasons for errors in your Golang code. When an error is created through **erro**, the source of the failing code is shown together with the stack trace and the variables that led to the error.

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
