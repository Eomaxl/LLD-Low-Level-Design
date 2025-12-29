package report

type Report struct{}

func (Report) generateContent() string {
	return "Content"
}

func (Report) printToPdf() {}

func (Report) saveToFile() {}
