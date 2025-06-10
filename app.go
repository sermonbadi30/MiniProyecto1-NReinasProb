package main

import (
	"context"
	"fmt"
	"math/rand/v2"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func resolverNReinasConLogs(matriz [][]int, cantReinas int, tamano int, fila int, logs *[]string) [][]int {
	if cantReinas == 0 {
		*logs = append(*logs, "Todas las reinas fueron colocadas con éxito.")
		return matriz
	}
	if fila == tamano {
		fila = 0
	}
	lista := crearListaDesordenada(tamano)
	for i := 0; i < tamano; i++ {
		columna := lista[i]
		*logs = append(*logs, fmt.Sprintf("Intentando colocar reina en fila %d, columna %d...", fila, columna))

		if validarReinas(matriz, tamano, fila, columna) {
			if matriz[fila][columna] != 1 {
				*logs = append(*logs, fmt.Sprintf("✅ Colocada reina en [%d][%d]", fila, columna))
				matriz[fila][columna] = 1
				resultado := resolverNReinasConLogs(matriz, cantReinas-1, tamano, fila+1, logs)
				if resultado != nil {
					return resultado
				}
				*logs = append(*logs, fmt.Sprintf("🔁 Quitando reina de [%d][%d]", fila, columna))
				matriz[fila][columna] = 0
			}
		} else {
			*logs = append(*logs, fmt.Sprintf("❌ No se puede colocar reina en [%d][%d]", fila, columna))
		}
	}
	return nil
}

// ResolverLasVegas ejecuta el algoritmo de N-Reinas con el método Las Vegas
type ResultadoNReinas struct {
	Tablero []string `json:"tablero"`
	Logs    []string `json:"logs"`
}

func (a *App) ResolverLasVegas(tamano int, cantReinas int) ResultadoNReinas {
	tablero := crearMatrizNula(tamano)
	logs := []string{fmt.Sprintf("Comenzando resolución con tablero %dx%d y %d reinas.", tamano, tamano, cantReinas)}

	resultado := resolverNReinasConLogs(tablero, cantReinas, tamano, rand.IntN(tamano), &logs)

	if resultado == nil {
		logs = append(logs, "❌ No fue posible colocar todas las reinas en el tablero con la configuración actual.")
		vacio := make([]string, tamano)
		for i := range vacio {
			vacio[i] = ""
		}
		return ResultadoNReinas{
			Tablero: vacio,
			Logs:    logs,
		}
	}

	solucion := []string{}
	for _, fila := range resultado {
		linea := ""
		for _, celda := range fila {
			if celda == 1 {
				linea += "Q"
			} else {
				linea += "."
			}
		}
		solucion = append(solucion, linea)
	}

	return ResultadoNReinas{
		Tablero: solucion,
		Logs:    logs,
	}
}

// --------------------- LÓGICA DEL ALGORITMO ---------------------

func crearMatrizNula(tamano int) [][]int {
	matriz := make([][]int, tamano)
	for i := 0; i < tamano; i++ {
		fila := make([]int, tamano)
		matriz[i] = fila
	}
	return matriz
}

func validarReinas(matriz [][]int, tamano int, fila int, columna int) bool {
	// Revisar columnas y filas
	for i := 0; i < tamano; i++ {
		if i != columna && matriz[fila][i] == 1 {
			return false
		}
		if i != fila && matriz[i][columna] == 1 {
			return false
		}
	}

	// Revisar diagonales
	direcciones := [][2]int{
		{-1, +1}, {-1, -1}, {+1, +1}, {+1, -1},
	}
	for _, d := range direcciones {
		for i, j := fila+d[0], columna+d[1]; i >= 0 && i < tamano && j >= 0 && j < tamano; i, j = i+d[0], j+d[1] {
			if matriz[i][j] == 1 {
				return false
			}
		}
	}
	return true
}

func crearListaDesordenada(n int) []int {
	lista := make([]int, n)
	for i := 0; i < n; i++ {
		lista[i] = i
	}
	return shuffleList(lista)
}

func shuffleList(lista []int) []int {
	for i := len(lista) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		lista[i], lista[j] = lista[j], lista[i]
	}
	return lista
}
