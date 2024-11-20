package gotodo

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

type CsvRender struct {
	writer io.Writer
}

func NewCsvRender(writer io.Writer) *CsvRender {
	return &CsvRender{writer}
}

func (render *CsvRender) Render(todos *Todos) error {
	tbl := csv.NewWriter(render.writer)
	tbl.Write([]string{"#", "Title", "Completed", "Created At", "Completed At"})
	ts := *todos
	for index, t := range ts.todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Local().Format(time.RFC1123)
			}
		}

		tbl.Write([]string{strconv.Itoa(index), t.Title, completed, t.CreatedAt.Local().Format(time.RFC1123), completedAt})
	}
	// Flush to ensure that all data will be written
	tbl.Flush()

	// To check if errors occur during flush
	if err := tbl.Error(); err != nil {
		return err
	}

	return nil
}
