package processing

type RowRecord struct {
	row     []string
	current int
}

func NewRowRecord(row []string) *RowRecord {
	rr := &RowRecord{
		row:     make([]string, 0),
		current: 0,
	}
	rr.row = append(rr.row, row...)
	return rr
}

// берет первую не пустую строку из массива после текущей
// если возвращается пустая строка то массив пустой весь
func (r *RowRecord) FirstNoEmpty() string {
	for i, s := range r.row {
		if s != "" {
			r.current = i
			return s
		}
	}
	return ""
}

// берет следующую не пустую строку из массива после текущей
// если возвращается пустая строка то уже нет таких
func (r *RowRecord) NextNoEmpty() string {
	if len(r.row) <= 1 {
		return ""
	}
	// если указатель последняя ячейка массива
	if r.current == (len(r.row) - 1) {
		return ""
	}
	row := r.row[r.current+1:]
	for i, s := range row {
		if s != "" {
			r.current += i + 1
			return s
		}
	}
	return ""
}
