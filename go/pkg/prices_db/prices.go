// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package prices_db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Price is an object representing the database table.
type Price struct {
	Timestamp time.Time         `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Currency  string            `boil:"currency" json:"currency" toml:"currency" yaml:"currency"`
	Usd       types.NullDecimal `boil:"usd" json:"usd,omitempty" toml:"usd" yaml:"usd,omitempty"`

	R *priceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L priceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PriceColumns = struct {
	Timestamp string
	Currency  string
	Usd       string
}{
	Timestamp: "timestamp",
	Currency:  "currency",
	Usd:       "usd",
}

var PriceTableColumns = struct {
	Timestamp string
	Currency  string
	Usd       string
}{
	Timestamp: "prices.timestamp",
	Currency:  "prices.currency",
	Usd:       "prices.usd",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertypes_NullDecimal struct{ field string }

func (w whereHelpertypes_NullDecimal) EQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_NullDecimal) NEQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_NullDecimal) LT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_NullDecimal) LTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_NullDecimal) GT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_NullDecimal) GTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpertypes_NullDecimal) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_NullDecimal) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}

var PriceWhere = struct {
	Timestamp whereHelpertime_Time
	Currency  whereHelperstring
	Usd       whereHelpertypes_NullDecimal
}{
	Timestamp: whereHelpertime_Time{field: "\"prices\".\"timestamp\""},
	Currency:  whereHelperstring{field: "\"prices\".\"currency\""},
	Usd:       whereHelpertypes_NullDecimal{field: "\"prices\".\"usd\""},
}

// PriceRels is where relationship names are stored.
var PriceRels = struct {
}{}

// priceR is where relationships are stored.
type priceR struct {
}

// NewStruct creates a new relationship struct
func (*priceR) NewStruct() *priceR {
	return &priceR{}
}

// priceL is where Load methods for each relationship are stored.
type priceL struct{}

var (
	priceAllColumns            = []string{"timestamp", "currency", "usd"}
	priceColumnsWithoutDefault = []string{"timestamp", "currency"}
	priceColumnsWithDefault    = []string{"usd"}
	pricePrimaryKeyColumns     = []string{"timestamp", "currency"}
	priceGeneratedColumns      = []string{}
)

type (
	// PriceSlice is an alias for a slice of pointers to Price.
	// This should almost always be used instead of []Price.
	PriceSlice []*Price
	// PriceHook is the signature for custom Price hook methods
	PriceHook func(context.Context, boil.ContextExecutor, *Price) error

	priceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	priceType                 = reflect.TypeOf(&Price{})
	priceMapping              = queries.MakeStructMapping(priceType)
	pricePrimaryKeyMapping, _ = queries.BindMapping(priceType, priceMapping, pricePrimaryKeyColumns)
	priceInsertCacheMut       sync.RWMutex
	priceInsertCache          = make(map[string]insertCache)
	priceUpdateCacheMut       sync.RWMutex
	priceUpdateCache          = make(map[string]updateCache)
	priceUpsertCacheMut       sync.RWMutex
	priceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var priceAfterSelectHooks []PriceHook

var priceBeforeInsertHooks []PriceHook
var priceAfterInsertHooks []PriceHook

var priceBeforeUpdateHooks []PriceHook
var priceAfterUpdateHooks []PriceHook

var priceBeforeDeleteHooks []PriceHook
var priceAfterDeleteHooks []PriceHook

var priceBeforeUpsertHooks []PriceHook
var priceAfterUpsertHooks []PriceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Price) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range priceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Price) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range priceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Price) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range priceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Price) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range priceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Price) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range priceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Price) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range priceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Price) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range priceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Price) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range priceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Price) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range priceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPriceHook registers your hook function for all future operations.
func AddPriceHook(hookPoint boil.HookPoint, priceHook PriceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		priceAfterSelectHooks = append(priceAfterSelectHooks, priceHook)
	case boil.BeforeInsertHook:
		priceBeforeInsertHooks = append(priceBeforeInsertHooks, priceHook)
	case boil.AfterInsertHook:
		priceAfterInsertHooks = append(priceAfterInsertHooks, priceHook)
	case boil.BeforeUpdateHook:
		priceBeforeUpdateHooks = append(priceBeforeUpdateHooks, priceHook)
	case boil.AfterUpdateHook:
		priceAfterUpdateHooks = append(priceAfterUpdateHooks, priceHook)
	case boil.BeforeDeleteHook:
		priceBeforeDeleteHooks = append(priceBeforeDeleteHooks, priceHook)
	case boil.AfterDeleteHook:
		priceAfterDeleteHooks = append(priceAfterDeleteHooks, priceHook)
	case boil.BeforeUpsertHook:
		priceBeforeUpsertHooks = append(priceBeforeUpsertHooks, priceHook)
	case boil.AfterUpsertHook:
		priceAfterUpsertHooks = append(priceAfterUpsertHooks, priceHook)
	}
}

// One returns a single price record from the query.
func (q priceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Price, error) {
	o := &Price{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "prices_db: failed to execute a one query for prices")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Price records from the query.
func (q priceQuery) All(ctx context.Context, exec boil.ContextExecutor) (PriceSlice, error) {
	var o []*Price

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "prices_db: failed to assign all query results to Price slice")
	}

	if len(priceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Price records in the query.
func (q priceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: failed to count prices rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q priceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "prices_db: failed to check if prices exists")
	}

	return count > 0, nil
}

// Prices retrieves all the records using an executor.
func Prices(mods ...qm.QueryMod) priceQuery {
	mods = append(mods, qm.From("\"prices\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"prices\".*"})
	}

	return priceQuery{q}
}

// FindPrice retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPrice(ctx context.Context, exec boil.ContextExecutor, timestamp time.Time, currency string, selectCols ...string) (*Price, error) {
	priceObj := &Price{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"prices\" where \"timestamp\"=$1 AND \"currency\"=$2", sel,
	)

	q := queries.Raw(query, timestamp, currency)

	err := q.Bind(ctx, exec, priceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "prices_db: unable to select from prices")
	}

	if err = priceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return priceObj, err
	}

	return priceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Price) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("prices_db: no prices provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(priceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	priceInsertCacheMut.RLock()
	cache, cached := priceInsertCache[key]
	priceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			priceAllColumns,
			priceColumnsWithDefault,
			priceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(priceType, priceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(priceType, priceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"prices\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"prices\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "prices_db: unable to insert into prices")
	}

	if !cached {
		priceInsertCacheMut.Lock()
		priceInsertCache[key] = cache
		priceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Price.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Price) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	priceUpdateCacheMut.RLock()
	cache, cached := priceUpdateCache[key]
	priceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			priceAllColumns,
			pricePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("prices_db: unable to update prices, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"prices\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pricePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(priceType, priceMapping, append(wl, pricePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: unable to update prices row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: failed to get rows affected by update for prices")
	}

	if !cached {
		priceUpdateCacheMut.Lock()
		priceUpdateCache[key] = cache
		priceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q priceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: unable to update all for prices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: unable to retrieve rows affected for prices")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PriceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("prices_db: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"prices\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pricePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: unable to update all in price slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: unable to retrieve rows affected all in update all price")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Price) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("prices_db: no prices provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(priceColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	priceUpsertCacheMut.RLock()
	cache, cached := priceUpsertCache[key]
	priceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			priceAllColumns,
			priceColumnsWithDefault,
			priceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			priceAllColumns,
			pricePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("prices_db: unable to upsert prices, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pricePrimaryKeyColumns))
			copy(conflict, pricePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"prices\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(priceType, priceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(priceType, priceMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "prices_db: unable to upsert prices")
	}

	if !cached {
		priceUpsertCacheMut.Lock()
		priceUpsertCache[key] = cache
		priceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Price record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Price) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("prices_db: no Price provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pricePrimaryKeyMapping)
	sql := "DELETE FROM \"prices\" WHERE \"timestamp\"=$1 AND \"currency\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: unable to delete from prices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: failed to get rows affected by delete for prices")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q priceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("prices_db: no priceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: unable to delete all from prices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: failed to get rows affected by deleteall for prices")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PriceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(priceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"prices\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pricePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: unable to delete all from price slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "prices_db: failed to get rows affected by deleteall for prices")
	}

	if len(priceAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Price) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPrice(ctx, exec, o.Timestamp, o.Currency)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PriceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PriceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"prices\".* FROM \"prices\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pricePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "prices_db: unable to reload all in PriceSlice")
	}

	*o = slice

	return nil
}

// PriceExists checks if the Price row exists.
func PriceExists(ctx context.Context, exec boil.ContextExecutor, timestamp time.Time, currency string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"prices\" where \"timestamp\"=$1 AND \"currency\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, timestamp, currency)
	}
	row := exec.QueryRowContext(ctx, sql, timestamp, currency)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "prices_db: unable to check if prices exists")
	}

	return exists, nil
}