package main

import (
	"fmt"
	"time"
)

type MaErro struct {
	when time.Time
	what string
}

func (m *MaErro) Error() string {
	return fmt.Sprintf("at %v, %s", m.when, m.what)
}

func run() error {
	return &MaErro{
		time.Now(), "some broken",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

}
