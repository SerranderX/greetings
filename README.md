# SALUDOS EN GO

Este es un pequeño paquete que contiene una función para generar un saludo aleatorio.

## Instalación

```go

go get github.com/serranderx/greetings
```

## Uso

```go
package main

import (
	"fmt"
	"github.com/serranderx/greetings"
)

func main() {
	fmt.Println(greetings.Hello("Juan"))
}
```

## Prueba

```go
git clone https://github.com/serranderx/greetings.git
cd greetings
go test -v
```