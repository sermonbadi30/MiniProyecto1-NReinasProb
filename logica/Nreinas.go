package main

import "fmt"

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
				fmt.Println("ERROR en vertical")
				valido = false
			}
		}
	}

	// Revisar Horizontalmente
	for i := 0; i < tamano; i++ {
		if i != fila {
			if matriz[i][columna] == 1 {
				fmt.Println("ERROR en horizontal")
				valido = false
			}
		}
	}

	// Diagonal superior derecha
	for i, j := fila-1, columna+1; i >= 0 && j < tamano; i, j = i-1, j+1 {
		if matriz[i][j] == 1 {
			fmt.Println("ERROR en diagonal superior derecha")
			valido = false
		}
	}

	// Diagonal superior izquierda
	for i, j := fila-1, columna-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if matriz[i][j] == 1 {
			fmt.Println("ERROR en Diagonal superior izquierda")
			valido = false
		}
	}

	// Diagonal inferior derecha
	for i, j := fila+1, columna+1; i < tamano && j < tamano; i, j = i+1, j+1 {
		if matriz[i][j] == 1 {
			fmt.Println("ERROR en Diagonal inferior derecha")
			valido = false
		}
	}

	// Diagonal inferior izquierda
	for i, j := fila+1, columna-1; i < tamano && j >= 0; i, j = i+1, j-1 {
		if matriz[i][j] == 1 {
			fmt.Println("ERROR en Diagonal inferior izquierda")
			valido = false
		}
	}

	return valido
}

func main() {
	matrizPrueba := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
	}

	resultado := validarReinas(matrizPrueba, 3, 0, 0)
	fmt.Println("¿Es válido colocar una reina?", resultado)
}
