    package main

    import (
        "time"

        "github.com/common-nighthawk/go-figure"
    )

    func main() {
        myFigure := figure.NewColorFigure("evgeny-goldin is Awesome!!!!", "larry3d", "yellow", true)
        myFigure.Print()
        time.Sleep(10 * time.Hour)
    }
