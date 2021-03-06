package helper

import (
	"github.com/tealeg/xlsx"
	"strings"
)

type XlsxFile struct {
	file *xlsx.File

	sheets []TableSheet
}

func (self *XlsxFile) Sheets() (ret []TableSheet) {

	return self.sheets
}

func (self *XlsxFile) Save(filename string) error {
	return self.file.Save(filename)
}

func (self *XlsxFile) Load(filename string) error {

	file, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}

	self.FromXFile(file)

	return nil
}

func (self *XlsxFile) FromXFile(file *xlsx.File) {
	self.file = file

	for _, sheet := range file.Sheets {
		self.sheets = append(self.sheets, newXlsxSheet(sheet))
	}
}

func NewXlsxFile() TableFile {

	self := &XlsxFile{}

	return self
}

type XlsxSheet struct {
	*xlsx.Sheet
}

func (self *XlsxSheet) Name() string {
	return self.Sheet.Name
}

func (self *XlsxSheet) MaxColumn() int {
	return self.Sheet.MaxCol
}

func (self *XlsxSheet) IsFullRowEmpty(row int) bool {

	for col := 0; col < self.Sheet.MaxCol; col++ {

		data := self.GetValue(row, col)

		if data != "" {
			return false
		}
	}

	return true
}

func (self *XlsxSheet) GetValue(row, col int) (ret string) {
	c := self.Sheet.Cell(row, col)

	// 取列头所在列和当前行交叉的单元格
	return strings.TrimSpace(c.Value)
}

func (self *XlsxSheet) GetValueEx(row, col int, isFloat bool) (ret string) {

	c := self.Sheet.Cell(row, col)

	// 浮点数单元格按原样输出
	if isFloat {
		ret, _ = c.GeneralNumeric()
		ret = strings.TrimSpace(ret)
	} else {
		// 取列头所在列和当前行交叉的单元格
		ret = strings.TrimSpace(c.Value)
	}

	return
}

func (self *XlsxSheet) WriteRow(valueList ...string) {
	row := self.Sheet.AddRow()

	for _, value := range valueList {

		cell := row.AddCell()
		cell.SetValue(value)
	}
}

func newXlsxSheet(sheet *xlsx.Sheet) TableSheet {
	return &XlsxSheet{
		Sheet: sheet,
	}
}
