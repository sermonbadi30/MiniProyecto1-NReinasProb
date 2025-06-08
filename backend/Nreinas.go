package main

import (
	"fmt"
	"math/rand/v2"
)

func crearMatrizNula(tamano int) [][]int {
	var matriz [][]int
	var temp []int
	for i := 0; i < tamano; i++ {
		for j := 0; j < tamano; j++ {
			temp = append(temp, 0)
		}
		matriz = append(matriz, temp)
		temp = []int{}
	}
	return matriz
}

func validarReinas(matriz [][]int, tamano int, fila int, columna int) bool {
	var valido = true
	// Revisar verticalmente
	for i := 0; i < tamano; i++ {
		if i != columna {
			if matriz[fila][i] == 1 {
				
				valido = false
			}
		}
	}

	// Revisar Horizontalmente
	for i := 0; i < tamano; i++ {
		if i != fila {
			if matriz[i][columna] == 1 {
				
				valido = false
			}
		}
	}

	// Diagonal superior derecha
	for i, j := fila-1, columna+1; i >= 0 && j < tamano; i, j = i-1, j+1 {
		if matriz[i][j] == 1 {
			
			valido = false
		}
	}

	// Diagonal superior izquierda
	for i, j := fila-1, columna-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if matriz[i][j] == 1 {
			
			valido = false
		}
	}

	// Diagonal inferior derecha
	for i, j := fila+1, columna+1; i < tamano && j < tamano; i, j = i+1, j+1 {
		if matriz[i][j] == 1 {
			
			valido = false
		}
	}

	// Diagonal inferior izquierda
	for i, j := fila+1, columna-1; i < tamano && j >= 0; i, j = i+1, j-1 {
		if matriz[i][j] == 1 {
		
			valido = false
		}
	}

	return valido
}

func resolverNReinas(tamano int, cantReinas int) [][]int{
	var matriz [][]int = crearMatrizNula(tamano)
	return resolverNReinasAux(matriz, cantReinas, tamano)
}

func resolverNReinasAux(matriz [][]int, cantReinas int, tamano int) [][]int{
	if cantReinas==0{return matriz}
	fila:=tamano-cantReinas
	lista:= crearListaDesordenada(tamano)
	
	for i:=0; i<tamano; i++{
		columna:= lista[i]

		if validarReinas(matriz, tamano, fila, columna){
			matriz[fila][columna] = 1
			resultado:= resolverNReinasAux(matriz, cantReinas-1, tamano)
			if resultado!=nil{
				return resultado
			}
			matriz[fila][columna]=0
		}
	}
	return nil
}

func main(){
	fmt.Print(resolverNReinas(1,1))
}

func crearListaDesordenada(n int) []int{
	lista := make([]int, n)
	for i:=0;i<n; i++{
		lista[i]=i
	} 
	return shuffleList(lista)
}

func shuffleList(lista []int) []int{
	largo:= len(lista)
    for i := largo - 1; i > 0; i-- {
        j := rand.IntN(i + 1) 
        lista[i], lista[j] = lista[j], lista[i] 
    }
	return lista
}