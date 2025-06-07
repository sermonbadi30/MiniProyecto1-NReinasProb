package main
//import "fmt"

func crearMatrizNula(tamano int) [][]int{
	var matriz[][] int
	var temp[]int
	for i:=0; i<tamano; i++{
		for j:=0; j<tamano; j++{
			temp=append(temp, 0)
		}
		matriz=append(matriz, temp)
		temp=[]int{}
	}
	return matriz
}

