Will be readme here

    package main
    
    import (
        "github.com/xamber/Varis"
        "fmt"
    )
    
    func main() {
    
        varis.DEBUG = true
    
        n := varis.CreateNetwork(2, 3, 1)
    
        fmt.Println("Before training")
        fmt.Println(0.0, 0.0, "-", n.Calculate(0.0, 0.0))
        fmt.Println(1.0, 0.0, "-", n.Calculate(1.0, 0.0))
        fmt.Println(0.0, 1.0, "-", n.Calculate(0.0, 1.0))
        fmt.Println(1.0, 1.0, "-", n.Calculate(1.0, 1.0))
    
        dataset := varis.Dataset{
            varis.Frame{[]float64{0.0, 0.0}, []float64{1.0}},
            varis.Frame{[]float64{1.0, 0.0}, []float64{0.0}},
            varis.Frame{[]float64{0.0, 1.0}, []float64{0.0}},
            varis.Frame{[]float64{1.0, 1.0}, []float64{1.0}},
        }
    
        trainer := varis.Trainer{&n, varis.BackPropagation}
        trainer.TrainByDataset(dataset, 10000)
    
        fmt.Println("After training")
        fmt.Println(0.0, 0.0, "-", n.Calculate(0.0, 0.0))
        fmt.Println(1.0, 0.0, "-", n.Calculate(1.0, 0.0))
        fmt.Println(0.0, 1.0, "-", n.Calculate(0.0, 1.0))
        fmt.Println(1.0, 1.0, "-", n.Calculate(1.0, 1.0))
    }