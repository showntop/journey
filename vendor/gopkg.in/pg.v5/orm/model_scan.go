package orm

import "gopkg.in/pg.v5/types"

type valuesModel struct {
	values []interface{}
}

var _ Model = valuesModel{}

func Scan(values ...interface{}) valuesModel {
	return valuesModel{
		values: values,
	}
}

func (valuesModel) useQueryOne() bool {
	return true
}

func (valuesModel) Reset() error {
	return nil
}

func (m valuesModel) NewModel() ColumnScanner {
	return m
}

func (valuesModel) AddModel(_ ColumnScanner) error {
	return nil
}

func (valuesModel) AfterQuery(_ DB) error {
	return nil
}

func (valuesModel) AfterSelect(_ DB) error {
	return nil
}

func (valuesModel) BeforeInsert(_ DB) error {
	return nil
}

func (valuesModel) AfterInsert(_ DB) error {
	return nil
}

func (valuesModel) BeforeUpdate(_ DB) error {
	return nil
}

func (valuesModel) AfterUpdate(_ DB) error {
	return nil
}

func (valuesModel) BeforeDelete(_ DB) error {
	return nil
}

func (valuesModel) AfterDelete(_ DB) error {
	return nil
}

func (m valuesModel) ScanColumn(colIdx int, _ string, b []byte) error {
	return types.Scan(m.values[colIdx], b)
}
