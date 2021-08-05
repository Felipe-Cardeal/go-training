package main

func main() {
    //fmt.println ("soma", soma(1,1))
    //fmt.println ("subtracao", subtracao(1.1,1))
    //fmt.println ()
    //fmt.println ()
    //fmt.println ()
}
func soma (x float64 , y float64 ) float64 {
  return x + y
}
func subtracao (x float64  , y float64 ) float64 {
    return x - y
}

func divisao (x float64 , y float64) float64 {
    return x / y
}

func multiplicacao (x float64 , y float64) float64 {
    return x * y
}

func adicionarPercentual (valor float64 , porcentagem float64) float64 {
    return valor + (valor * porcentagem / 100)
}

func descobrirPercentual (valor float64 , porcentagem float64) float64 {
    return (valor / porcentagem) * 100
}


