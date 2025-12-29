package report

type PrintToPDF struct{}

type FileStorage struct{}

func (PrintToPDF) printToPdfFile(report Report) {
	_ = report
}

func (FileStorage) saveToFiles(context string) {
	_ = context
}
