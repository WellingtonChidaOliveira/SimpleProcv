package data

func MatchColumns(value, templatePathFirst, sheet string) ([]int, error) {
	f, err := openFile(templatePathFirst)
	if err != nil {
		return nil, err
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	matchedColumnsPrefix := make([]int, 0)

	for i, row := range rows {
		for _, cell := range row {
			if value == string(cell) {
				matchedColumnsPrefix = append(matchedColumnsPrefix, i+1)
			}

		}
	}
	return matchedColumnsPrefix, nil
}
