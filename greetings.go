package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// HELLO
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	return randomGreeting(name), nil
}

func Hellos(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, errors.New("names cannot be empty")
	}

	messages := make(map[string]string)
	for _, name := range names {
		message := randomGreeting(name)
		messages[name] = message
	}
	return messages, nil
}

func randomGreeting(name string) string {
	random := []string{
		"¡Hello, %s!",
		"¡Hi, %s!",
		"¡Hey, %s!",
		"¡Hola, %s!",
		"¡Que bueno verte %s!",
		"¡Bienvenido %s!",
		"¡Un gusto saludarte de nuevo %s!",
	}

	return fmt.Sprintf(random[rand.Intn(len(random))]+"\n", name)
}
