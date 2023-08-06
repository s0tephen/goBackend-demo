// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"goBackend-demo/gen/orm/model"
)

func newFeedback(db *gorm.DB, opts ...gen.DOOption) feedback {
	_feedback := feedback{}

	_feedback.feedbackDo.UseDB(db, opts...)
	_feedback.feedbackDo.UseModel(&model.Feedback{})

	tableName := _feedback.feedbackDo.TableName()
	_feedback.ALL = field.NewAsterisk(tableName)
	_feedback.Fid = field.NewInt64(tableName, "fid")
	_feedback.FUser = field.NewString(tableName, "fUser")
	_feedback.FMsg = field.NewString(tableName, "fMsg")
	_feedback.FTime = field.NewTime(tableName, "fTime")

	_feedback.fillFieldMap()

	return _feedback
}

type feedback struct {
	feedbackDo

	ALL   field.Asterisk
	Fid   field.Int64  // 反馈ID
	FUser field.String // 反馈者
	FMsg  field.String // 反馈内容
	FTime field.Time   // 反馈时间

	fieldMap map[string]field.Expr
}

func (f feedback) Table(newTableName string) *feedback {
	f.feedbackDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f feedback) As(alias string) *feedback {
	f.feedbackDo.DO = *(f.feedbackDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *feedback) updateTableName(table string) *feedback {
	f.ALL = field.NewAsterisk(table)
	f.Fid = field.NewInt64(table, "fid")
	f.FUser = field.NewString(table, "fUser")
	f.FMsg = field.NewString(table, "fMsg")
	f.FTime = field.NewTime(table, "fTime")

	f.fillFieldMap()

	return f
}

func (f *feedback) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *feedback) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 4)
	f.fieldMap["fid"] = f.Fid
	f.fieldMap["fUser"] = f.FUser
	f.fieldMap["fMsg"] = f.FMsg
	f.fieldMap["fTime"] = f.FTime
}

func (f feedback) clone(db *gorm.DB) feedback {
	f.feedbackDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f feedback) replaceDB(db *gorm.DB) feedback {
	f.feedbackDo.ReplaceDB(db)
	return f
}

type feedbackDo struct{ gen.DO }

func (f feedbackDo) Debug() *feedbackDo {
	return f.withDO(f.DO.Debug())
}

func (f feedbackDo) WithContext(ctx context.Context) *feedbackDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f feedbackDo) ReadDB() *feedbackDo {
	return f.Clauses(dbresolver.Read)
}

func (f feedbackDo) WriteDB() *feedbackDo {
	return f.Clauses(dbresolver.Write)
}

func (f feedbackDo) Session(config *gorm.Session) *feedbackDo {
	return f.withDO(f.DO.Session(config))
}

func (f feedbackDo) Clauses(conds ...clause.Expression) *feedbackDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f feedbackDo) Returning(value interface{}, columns ...string) *feedbackDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f feedbackDo) Not(conds ...gen.Condition) *feedbackDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f feedbackDo) Or(conds ...gen.Condition) *feedbackDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f feedbackDo) Select(conds ...field.Expr) *feedbackDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f feedbackDo) Where(conds ...gen.Condition) *feedbackDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f feedbackDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *feedbackDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f feedbackDo) Order(conds ...field.Expr) *feedbackDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f feedbackDo) Distinct(cols ...field.Expr) *feedbackDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f feedbackDo) Omit(cols ...field.Expr) *feedbackDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f feedbackDo) Join(table schema.Tabler, on ...field.Expr) *feedbackDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f feedbackDo) LeftJoin(table schema.Tabler, on ...field.Expr) *feedbackDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f feedbackDo) RightJoin(table schema.Tabler, on ...field.Expr) *feedbackDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f feedbackDo) Group(cols ...field.Expr) *feedbackDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f feedbackDo) Having(conds ...gen.Condition) *feedbackDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f feedbackDo) Limit(limit int) *feedbackDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f feedbackDo) Offset(offset int) *feedbackDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f feedbackDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *feedbackDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f feedbackDo) Unscoped() *feedbackDo {
	return f.withDO(f.DO.Unscoped())
}

func (f feedbackDo) Create(values ...*model.Feedback) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f feedbackDo) CreateInBatches(values []*model.Feedback, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f feedbackDo) Save(values ...*model.Feedback) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f feedbackDo) First() (*model.Feedback, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Feedback), nil
	}
}

func (f feedbackDo) Take() (*model.Feedback, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Feedback), nil
	}
}

func (f feedbackDo) Last() (*model.Feedback, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Feedback), nil
	}
}

func (f feedbackDo) Find() ([]*model.Feedback, error) {
	result, err := f.DO.Find()
	return result.([]*model.Feedback), err
}

func (f feedbackDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Feedback, err error) {
	buf := make([]*model.Feedback, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f feedbackDo) FindInBatches(result *[]*model.Feedback, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f feedbackDo) Attrs(attrs ...field.AssignExpr) *feedbackDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f feedbackDo) Assign(attrs ...field.AssignExpr) *feedbackDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f feedbackDo) Joins(fields ...field.RelationField) *feedbackDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f feedbackDo) Preload(fields ...field.RelationField) *feedbackDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f feedbackDo) FirstOrInit() (*model.Feedback, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Feedback), nil
	}
}

func (f feedbackDo) FirstOrCreate() (*model.Feedback, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Feedback), nil
	}
}

func (f feedbackDo) FindByPage(offset int, limit int) (result []*model.Feedback, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f feedbackDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f feedbackDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f feedbackDo) Delete(models ...*model.Feedback) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *feedbackDo) withDO(do gen.Dao) *feedbackDo {
	f.DO = *do.(*gen.DO)
	return f
}
