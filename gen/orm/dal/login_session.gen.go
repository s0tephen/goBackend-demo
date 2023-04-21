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

	"index_Demo/gen/orm/model"
)

func newLoginSession(db *gorm.DB, opts ...gen.DOOption) loginSession {
	_loginSession := loginSession{}

	_loginSession.loginSessionDo.UseDB(db, opts...)
	_loginSession.loginSessionDo.UseModel(&model.LoginSession{})

	tableName := _loginSession.loginSessionDo.TableName()
	_loginSession.ALL = field.NewAsterisk(tableName)
	_loginSession.ID = field.NewInt32(tableName, "id")
	_loginSession.Token = field.NewString(tableName, "token")
	_loginSession.UID = field.NewInt32(tableName, "uid")
	_loginSession.LoginIP = field.NewString(tableName, "login_ip")

	_loginSession.fillFieldMap()

	return _loginSession
}

type loginSession struct {
	loginSessionDo

	ALL     field.Asterisk
	ID      field.Int32  // token_Id
	Token   field.String // 令牌
	UID     field.Int32  // 用户ID
	LoginIP field.String // 登陆IP

	fieldMap map[string]field.Expr
}

func (l loginSession) Table(newTableName string) *loginSession {
	l.loginSessionDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l loginSession) As(alias string) *loginSession {
	l.loginSessionDo.DO = *(l.loginSessionDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *loginSession) updateTableName(table string) *loginSession {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.Token = field.NewString(table, "token")
	l.UID = field.NewInt32(table, "uid")
	l.LoginIP = field.NewString(table, "login_ip")

	l.fillFieldMap()

	return l
}

func (l *loginSession) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *loginSession) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 4)
	l.fieldMap["id"] = l.ID
	l.fieldMap["token"] = l.Token
	l.fieldMap["uid"] = l.UID
	l.fieldMap["login_ip"] = l.LoginIP
}

func (l loginSession) clone(db *gorm.DB) loginSession {
	l.loginSessionDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l loginSession) replaceDB(db *gorm.DB) loginSession {
	l.loginSessionDo.ReplaceDB(db)
	return l
}

type loginSessionDo struct{ gen.DO }

func (l loginSessionDo) Debug() *loginSessionDo {
	return l.withDO(l.DO.Debug())
}

func (l loginSessionDo) WithContext(ctx context.Context) *loginSessionDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l loginSessionDo) ReadDB() *loginSessionDo {
	return l.Clauses(dbresolver.Read)
}

func (l loginSessionDo) WriteDB() *loginSessionDo {
	return l.Clauses(dbresolver.Write)
}

func (l loginSessionDo) Session(config *gorm.Session) *loginSessionDo {
	return l.withDO(l.DO.Session(config))
}

func (l loginSessionDo) Clauses(conds ...clause.Expression) *loginSessionDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l loginSessionDo) Returning(value interface{}, columns ...string) *loginSessionDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l loginSessionDo) Not(conds ...gen.Condition) *loginSessionDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l loginSessionDo) Or(conds ...gen.Condition) *loginSessionDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l loginSessionDo) Select(conds ...field.Expr) *loginSessionDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l loginSessionDo) Where(conds ...gen.Condition) *loginSessionDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l loginSessionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *loginSessionDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l loginSessionDo) Order(conds ...field.Expr) *loginSessionDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l loginSessionDo) Distinct(cols ...field.Expr) *loginSessionDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l loginSessionDo) Omit(cols ...field.Expr) *loginSessionDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l loginSessionDo) Join(table schema.Tabler, on ...field.Expr) *loginSessionDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l loginSessionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *loginSessionDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l loginSessionDo) RightJoin(table schema.Tabler, on ...field.Expr) *loginSessionDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l loginSessionDo) Group(cols ...field.Expr) *loginSessionDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l loginSessionDo) Having(conds ...gen.Condition) *loginSessionDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l loginSessionDo) Limit(limit int) *loginSessionDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l loginSessionDo) Offset(offset int) *loginSessionDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l loginSessionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *loginSessionDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l loginSessionDo) Unscoped() *loginSessionDo {
	return l.withDO(l.DO.Unscoped())
}

func (l loginSessionDo) Create(values ...*model.LoginSession) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l loginSessionDo) CreateInBatches(values []*model.LoginSession, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l loginSessionDo) Save(values ...*model.LoginSession) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l loginSessionDo) First() (*model.LoginSession, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginSession), nil
	}
}

func (l loginSessionDo) Take() (*model.LoginSession, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginSession), nil
	}
}

func (l loginSessionDo) Last() (*model.LoginSession, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginSession), nil
	}
}

func (l loginSessionDo) Find() ([]*model.LoginSession, error) {
	result, err := l.DO.Find()
	return result.([]*model.LoginSession), err
}

func (l loginSessionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LoginSession, err error) {
	buf := make([]*model.LoginSession, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l loginSessionDo) FindInBatches(result *[]*model.LoginSession, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l loginSessionDo) Attrs(attrs ...field.AssignExpr) *loginSessionDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l loginSessionDo) Assign(attrs ...field.AssignExpr) *loginSessionDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l loginSessionDo) Joins(fields ...field.RelationField) *loginSessionDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l loginSessionDo) Preload(fields ...field.RelationField) *loginSessionDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l loginSessionDo) FirstOrInit() (*model.LoginSession, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginSession), nil
	}
}

func (l loginSessionDo) FirstOrCreate() (*model.LoginSession, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginSession), nil
	}
}

func (l loginSessionDo) FindByPage(offset int, limit int) (result []*model.LoginSession, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l loginSessionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l loginSessionDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l loginSessionDo) Delete(models ...*model.LoginSession) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *loginSessionDo) withDO(do gen.Dao) *loginSessionDo {
	l.DO = *do.(*gen.DO)
	return l
}
