package scan

type Scanner interface {
	Run(report *Report) error
}

func Run() (*Report, error) {
	scanners := []Scanner{
		&DeviceScanner{},
	}

	report := &Report{}
	for _, scanner := range scanners {
		err := scanner.Run(report)
		if err != nil {
			return nil, err
		}
	}
	return report, nil
}
