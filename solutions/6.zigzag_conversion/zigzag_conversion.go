package zigzagconversion

func convert(s string, numRows int) string {
	// initializing matrix
	m := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		m[i] = make([]string, 0)
	}

	col := 0
	for i := 0; i < len(s); {
		// full column filling
		for row := 0; row < numRows && i < len(s); row++ {
			m[row] = append(m[row], string(s[i]))
			i++
		}

		// diagonal column filling
		for row := numRows - 2; row > 0 && i < len(s); row-- {
			col++
			m[row] = append(m[row], string(s[i]))
			i++
		}
		col++
	}

	// generating the result
	result := ""
	for _, row := range m {
		for _, element := range row {
			if element != "" {
				result += element
			}
		}
	}

	return result
}
