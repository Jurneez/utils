package export

func ExportExcel(headers []string, in interface{}, path string) error {
	excel, err := NewExcel(headers, in, path)
	if err != nil {
		return err
	}
	err = excel.Export()
	if err != nil {
		return err
	}

	return nil
}
