// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package database

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"Vadek/model/entity"
)

func newThemeSetting(db *gorm.DB, opts ...gen.DOOption) themeSetting {
	_themeSetting := themeSetting{}

	_themeSetting.themeSettingDo.UseDB(db, opts...)
	_themeSetting.themeSettingDo.UseModel(&entity.ThemeSetting{})

	tableName := _themeSetting.themeSettingDo.TableName()
	_themeSetting.ALL = field.NewAsterisk(tableName)
	_themeSetting.ID = field.NewInt32(tableName, "id")
	_themeSetting.CreateTime = field.NewTime(tableName, "create_time")
	_themeSetting.UpdateTime = field.NewTime(tableName, "update_time")
	_themeSetting.SettingKey = field.NewString(tableName, "setting_key")
	_themeSetting.ThemeID = field.NewString(tableName, "theme_id")
	_themeSetting.SettingValue = field.NewString(tableName, "setting_value")

	_themeSetting.fillFieldMap()

	return _themeSetting
}

type themeSetting struct {
	themeSettingDo themeSettingDo

	ALL          field.Asterisk
	ID           field.Int32
	CreateTime   field.Time
	UpdateTime   field.Time
	SettingKey   field.String
	ThemeID      field.String
	SettingValue field.String

	fieldMap map[string]field.Expr
}

func (t themeSetting) Table(newTableName string) *themeSetting {
	t.themeSettingDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t themeSetting) As(alias string) *themeSetting {
	t.themeSettingDo.DO = *(t.themeSettingDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *themeSetting) updateTableName(table string) *themeSetting {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.CreateTime = field.NewTime(table, "create_time")
	t.UpdateTime = field.NewTime(table, "update_time")
	t.SettingKey = field.NewString(table, "setting_key")
	t.ThemeID = field.NewString(table, "theme_id")
	t.SettingValue = field.NewString(table, "setting_value")

	t.fillFieldMap()

	return t
}

func (t *themeSetting) WithContext(ctx context.Context) *themeSettingDo {
	return t.themeSettingDo.WithContext(ctx)
}

func (t themeSetting) TableName() string { return t.themeSettingDo.TableName() }

func (t themeSetting) Alias() string { return t.themeSettingDo.Alias() }

func (t *themeSetting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *themeSetting) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.ID
	t.fieldMap["create_time"] = t.CreateTime
	t.fieldMap["update_time"] = t.UpdateTime
	t.fieldMap["setting_key"] = t.SettingKey
	t.fieldMap["theme_id"] = t.ThemeID
	t.fieldMap["setting_value"] = t.SettingValue
}

func (t themeSetting) clone(db *gorm.DB) themeSetting {
	t.themeSettingDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t themeSetting) replaceDB(db *gorm.DB) themeSetting {
	t.themeSettingDo.ReplaceDB(db)
	return t
}

type themeSettingDo struct{ gen.DO }

func (t themeSettingDo) Debug() *themeSettingDo {
	return t.withDO(t.DO.Debug())
}

func (t themeSettingDo) WithContext(ctx context.Context) *themeSettingDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t themeSettingDo) ReadDB() *themeSettingDo {
	return t.Clauses(dbresolver.Read)
}

func (t themeSettingDo) WriteDB() *themeSettingDo {
	return t.Clauses(dbresolver.Write)
}

func (t themeSettingDo) Session(config *gorm.Session) *themeSettingDo {
	return t.withDO(t.DO.Session(config))
}

func (t themeSettingDo) Clauses(conds ...clause.Expression) *themeSettingDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t themeSettingDo) Returning(value interface{}, columns ...string) *themeSettingDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t themeSettingDo) Not(conds ...gen.Condition) *themeSettingDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t themeSettingDo) Or(conds ...gen.Condition) *themeSettingDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t themeSettingDo) Select(conds ...field.Expr) *themeSettingDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t themeSettingDo) Where(conds ...gen.Condition) *themeSettingDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t themeSettingDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *themeSettingDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t themeSettingDo) Order(conds ...field.Expr) *themeSettingDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t themeSettingDo) Distinct(cols ...field.Expr) *themeSettingDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t themeSettingDo) Omit(cols ...field.Expr) *themeSettingDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t themeSettingDo) Join(table schema.Tabler, on ...field.Expr) *themeSettingDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t themeSettingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *themeSettingDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t themeSettingDo) RightJoin(table schema.Tabler, on ...field.Expr) *themeSettingDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t themeSettingDo) Group(cols ...field.Expr) *themeSettingDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t themeSettingDo) Having(conds ...gen.Condition) *themeSettingDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t themeSettingDo) Limit(limit int) *themeSettingDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t themeSettingDo) Offset(offset int) *themeSettingDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t themeSettingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *themeSettingDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t themeSettingDo) Unscoped() *themeSettingDo {
	return t.withDO(t.DO.Unscoped())
}

func (t themeSettingDo) Create(values ...*entity.ThemeSetting) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t themeSettingDo) CreateInBatches(values []*entity.ThemeSetting, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t themeSettingDo) Save(values ...*entity.ThemeSetting) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t themeSettingDo) First() (*entity.ThemeSetting, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ThemeSetting), nil
	}
}

func (t themeSettingDo) Take() (*entity.ThemeSetting, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ThemeSetting), nil
	}
}

func (t themeSettingDo) Last() (*entity.ThemeSetting, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ThemeSetting), nil
	}
}

func (t themeSettingDo) Find() ([]*entity.ThemeSetting, error) {
	result, err := t.DO.Find()
	return result.([]*entity.ThemeSetting), err
}

func (t themeSettingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.ThemeSetting, err error) {
	buf := make([]*entity.ThemeSetting, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t themeSettingDo) FindInBatches(result *[]*entity.ThemeSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t themeSettingDo) Attrs(attrs ...field.AssignExpr) *themeSettingDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t themeSettingDo) Assign(attrs ...field.AssignExpr) *themeSettingDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t themeSettingDo) Joins(fields ...field.RelationField) *themeSettingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t themeSettingDo) Preload(fields ...field.RelationField) *themeSettingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t themeSettingDo) FirstOrInit() (*entity.ThemeSetting, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ThemeSetting), nil
	}
}

func (t themeSettingDo) FirstOrCreate() (*entity.ThemeSetting, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ThemeSetting), nil
	}
}

func (t themeSettingDo) FindByPage(offset int, limit int) (result []*entity.ThemeSetting, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t themeSettingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t themeSettingDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t themeSettingDo) Delete(models ...*entity.ThemeSetting) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *themeSettingDo) withDO(do gen.Dao) *themeSettingDo {
	t.DO = *do.(*gen.DO)
	return t
}
