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
    greeting, err := greetings.Hello("Juan")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(greeting)

    messages, err := greetings.Hellos([]string{"John", "Sally", "Miguel", "Luisa"})
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, greeting := range messages {
        fmt.Println(greeting)
    }
}
```

## Prueba

```go
git clone https://github.com/serranderx/greetings.git
cd greetings
go test -v
```