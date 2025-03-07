// Author : Cassiopée Gossin

package main

import (
	"bufio"
	"fmt"
	"os"
)

var filename string = "input"

func main() {
	fmt.Println("partie 1:")
	fmt.Println(pb1(filename))
	fmt.Println("partie 2:")
	fmt.Println(pb2(filename))
}

// solve the first part of the problem
func pb1(filename string) int {
	f, _ := os.Open(filename)
	sc := bufio.NewScanner(f)
	var tab [][][]bool
	tab = initTab(tab)
	for sc.Scan() {
		line := sc.Text()
		tab = processLineMatrix(line, tab)
	}
	taille := sumBool(tab)
	return taille
}

// structure to represent a block
type Pave struct {
	xMin  int
	xMax  int
	yMin  int
	yMax  int
	zMin  int
	zMax  int
	value bool
}

// solve the second part of the problem
func pb2(filename string) int {
	f, _ := os.Open(filename)
	sc := bufio.NewScanner(f)
	var tab []Pave
	for sc.Scan() {
		line := sc.Text()
		tab = processLine2(line, tab)
	}
	taille := sumPave(tab)
	return taille
}

//
// *************** First part with a matrix of bool *********************
//

// executed for one line of the input
func processLineMatrix(line string, tab [][][]bool) [][][]bool {
	var (
		mot  string
		xMin int
		xMax int
		yMin int
		yMax int
		zMin int
		zMax int
	)

	// read the line
	fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &mot, &xMin, &xMax, &yMin, &yMax, &zMin, &zMax)

	if inArea(xMin, xMax, yMin, yMax, zMin, zMax) {
		// modify values to be in the initialization procedure area and translation to stay in init area with the matrix
		xMin = restrictValues(xMin) + 50
		xMax = restrictValues(xMax) + 50
		yMin = restrictValues(yMin) + 50
		yMax = restrictValues(yMax) + 50
		zMin = restrictValues(zMin) + 50
		zMax = restrictValues(zMax) + 50

		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				for k := zMin; k <= zMax; k++ {
					if mot == "off" {
						tab[i][j][k] = false
					} else {
						tab[i][j][k] = true
					}
				}
			}
		}
	}
	return tab
}

// returns the total sum of the number of true cells in the matrix
func sumBool(tab [][][]bool) int {
	nb := 0
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[0]); j++ {
			for k := 0; k < len(tab[0][0]); k++ {
				if tab[i][j][k] {
					nb++
				}
			}
		}
	}
	return nb
}

// initialise the matrix with all the cells set to false
func initTab(tab [][][]bool) [][][]bool {
	for i := 0; i <= 100; i++ {
		var new_mat [][]bool
		tab = append(tab, new_mat)
		for j := 0; j <= 100; j++ {
			var new_tab []bool
			tab[i] = append(tab[i], new_tab)
			for k := 0; k <= 100; k++ {
				new_val := false
				tab[i][j] = append(tab[i][j], new_val)
			}
		}
	}
	return tab
}

// limit values to the initialization area
func restrictValues(value int) int {
	if value < (-50) {
		return -50
	} else if value > 50 {
		return 50
	} else {
		return value
	}
}

// check if the values are in the initialization area
func inArea(xMin int, xMax int, yMin int, yMax int, zMin int, zMax int) bool {
	return xMin <= 50 && xMax >= -50 && yMin <= 50 && yMax >= -50 && zMin <= 50 && zMax >= -50
}

//
// *************** Second part with a tab of Pave *********************
//

// executed for one line of the input
func processLine2(line string, tab []Pave) []Pave {
	var (
		mot     string
		xMin    int
		xMax    int
		yMin    int
		yMax    int
		zMin    int
		zMax    int
		newPave Pave
	)

	// read the line
	fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &mot, &xMin, &xMax, &yMin, &yMax, &zMin, &zMax)
	// create a block
	if mot == "on" {
		newPave = Pave{xMin, xMax, yMin, yMax, zMin, zMax, true}
	} else {
		newPave = Pave{xMin, xMax, yMin, yMax, zMin, zMax, false}
	}
	for i := range tab {
		pave := tab[i]
		intersection := intersectionPave(pave, newPave)
		if !(intersection.xMin == 1 && intersection.xMax == -1 && !intersection.value) { // if there is an intersection
			tab = append(tab, intersection)
		}
	}
	if mot == "on" {
		tab = append(tab, newPave)
	}
	return tab
}

// finds the intersection between the read block and the block in the block table, if there is no any intersection it returns an impossible block
func intersectionPave(pave Pave, newPave Pave) Pave {
	var inter Pave
	inArea := pave.xMin <= newPave.xMax && pave.xMax >= newPave.xMin && pave.yMin <= newPave.yMax && pave.yMax >= newPave.yMin && pave.zMin <= newPave.zMax && pave.zMax >= newPave.zMin
	if inArea {
		// find intersection values
		inter.xMin = greatest(pave.xMin, newPave.xMin)
		inter.xMax = smallest(pave.xMax, newPave.xMax)
		inter.yMin = greatest(pave.yMin, newPave.yMin)
		inter.yMax = smallest(pave.yMax, newPave.yMax)
		inter.zMin = greatest(pave.zMin, newPave.zMin)
		inter.zMax = smallest(pave.zMax, newPave.zMax)
		inter.value = !(pave.value)
	} else {
		inter = Pave{1, -1, 1, -1, 1, -1, false} // impossible block to have a null pave
	}
	return inter
}

// find the greatest value between who integer
func greatest(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}

// find the smallest value between who integer
func smallest(val1 int, val2 int) int {
	if val1 > val2 {
		return val2
	} else {
		return val1
	}
}

// returns the total sum of the number of true values in all blocks
func sumPave(tab []Pave) int {
	nb := 0
	for _, pave := range tab {
		if pave.value {
			nb += nbValuePave(pave)
		} else {
			nb -= nbValuePave(pave)
		}
	}
	return nb
}

// find the number of values in one block
func nbValuePave(pave Pave) int {
	return (pave.xMax - pave.xMin + 1) * (pave.yMax - pave.yMin + 1) * (pave.zMax - pave.zMin + 1)
}
