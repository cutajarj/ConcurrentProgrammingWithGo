package listing10_13

import (
    "fmt"
    "time"
)

const (
    ovenTime           = 5
    everyThingElseTime = 2
)

func PrepareTray(trayNumber int) string {
    fmt.Println("Preparing empty tray", trayNumber)
    time.Sleep(everyThingElseTime * time.Second)
    return fmt.Sprintf("tray number %d", trayNumber)
}

func Mixture(tray string) string {
    fmt.Println("Pouring cupcake Mixture in", tray)
    time.Sleep(everyThingElseTime * time.Second)
    return fmt.Sprintf("cupcake in %s", tray)
}

func Bake(mixture string) string {
    fmt.Println("Baking", mixture)
    time.Sleep(ovenTime * time.Second)
    return fmt.Sprintf("baked %s", mixture)
}

func AddToppings(bakedCupCake string) string {
    fmt.Println("Adding topping to", bakedCupCake)
    time.Sleep(everyThingElseTime * time.Second)
    return fmt.Sprintf("topping on %s", bakedCupCake)
}

func Box(finishedCupCake string) string {
    fmt.Println("Boxing", finishedCupCake)
    time.Sleep(everyThingElseTime * time.Second)
    return fmt.Sprintf("%s boxed", finishedCupCake)
}
