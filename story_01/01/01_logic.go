package quest01

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_stories/utils"
)

// ========================
// PART I
// ========================

func I(file string) int {
	rows := parseInput(file)
	max  := 0

	for _, row := range rows {
		sum := ENI(row.A, row.X, row.M) +
				ENI(row.B, row.Y, row.M) +
				ENI(row.C, row.Z, row.M)

				
		if sum > max {
			max = sum
		}
	}

	return max
}

// ========================
// ENI
// ========================

func ENI(n, exp, mod int) int {
    score := 1

    var digits []int
    for range exp {
        score = (score * n) % mod
        digits = append([]int{score}, digits...)
    }

    var fullValue strings.Builder
    for _, v := range digits {
		fullValue.WriteString(strconv.Itoa(v))
    }

	result, _ := strconv.Atoi(fullValue.String())
    return result
}

// ========================
// PARSER
// ========================

type Row struct {
	A,B,C,X,Y,Z,M int
}

func parseInput(file string) []Row{
	data := utils.ReadFile(file)
	rows := make([]Row,0)

	for _, line := range strings.Split(data, "\n") {
		var row Row

		for _, value := range utils.QuickMatch(line, `[A-Z]=\d+`) {
			parts := strings.Split(value, "=")
			num, _ := strconv.Atoi(parts[1])
			
			switch parts[0] {
            case "A":
                row.A = num
            case "B":
                row.B = num
            case "C":
                row.C = num
            case "X":
                row.X = num
            case "Y":
                row.Y = num
            case "Z":
                row.Z = num
			case "M":
                row.M = num
            }
		}
		rows = append(rows, row)
	}
	return rows
}