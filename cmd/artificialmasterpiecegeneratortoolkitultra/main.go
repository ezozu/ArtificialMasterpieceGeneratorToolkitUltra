// cmd/artificialmasterpiecegeneratortoolkitultra/main.go
package main

import (
"flag"
"log"
"os"

"artificialmasterpiecegeneratortoolkitultra/internal/artificialmasterpiecegeneratortoolkitultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmasterpiecegeneratortoolkitultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
