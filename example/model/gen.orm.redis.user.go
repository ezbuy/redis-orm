package model

import (
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	redis "gopkg.in/redis.v5"
	"strings"
)

var (
	_ fmt.Formatter
	_ orm.VSet
	_ strings.Reader
)

type _UserRedisMgr struct {
	*orm.RedisStore
}

func UserRedisMgr() *_UserRedisMgr {
	return &_UserRedisMgr{_redis_store}
}

func NewUserRedisMgr(cf *RedisConfig) (*_UserRedisMgr, error) {
	store, err := orm.NewRedisStore(cf.Host, cf.Port, cf.Password, 0)
	if err != nil {
		return nil, err
	}
	return &_UserRedisMgr{store}, nil
}

//! pipeline write
type _UserRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_UserRedisMgr) BeginPipeline() *_UserRedisPipeline {
	return &_UserRedisPipeline{m.Pipeline(), nil}
}

func (m *_UserRedisMgr) Load(db DBFetcher) error {

	if err := m.Clear(); err != nil {
		return err
	}
	return m.AddBySQL(db, "SELECT `id`,`name`,`mailbox`,`sex`,`longitude`,`latitude`,`description`,`password`,`head_url`,`status`,`created_at`, `updated_at` FROM users")

	return nil
}

func (m *_UserRedisMgr) AddBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.Save(obj.(*User)); err != nil {
			return err
		}
	}

	return nil
}
func (m *_UserRedisMgr) DelBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.Delete(obj.(*User)); err != nil {
			return err
		}
	}
	return nil
}

//! redis model read
func (m *_UserRedisMgr) FindOne(unique Unique) (string, error) {
	if relation := unique.UKRelation(); relation != nil {
		return relation.FindOne(unique.Key())
	}
	return "", nil
}

func (m *_UserRedisMgr) Find(index Index) ([]string, error) {
	if relation := index.IDXRelation(); relation != nil {
		return relation.Find(index.Key())
	}
	return nil, nil
}

func (m *_UserRedisMgr) Range(scope Range) ([]string, error) {
	if relation := scope.RNGRelation(); relation != nil {
		return relation.Range(scope.Key(), scope.Begin(), scope.End())
	}
	return nil, nil
}

func (m *_UserRedisMgr) OrderBy(sort OrderBy) ([]string, error) {
	if relation := sort.ORDRelation(); relation != nil {
		return relation.OrderBy(sort.Key(), true)
	}
	return nil, nil
}

func (m *_UserRedisMgr) Fetch(id string) (*User, error) {
	obj := &User{}
	pipe := m.BeginPipeline()
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Id")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Name")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Mailbox")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Sex")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Longitude")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Latitude")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Description")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Password")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "HeadUrl")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Status")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "CreatedAt")
	pipe.HGet(keyOfObject(obj, fmt.Sprint(obj.Id)), "UpdatedAt")
	cmds, err := pipe.Exec()
	if err != nil {
		return nil, err
	}
	str0, err := cmds[0].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str0, &obj.Id); err != nil {
		return nil, err
	}
	str1, err := cmds[1].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str1, &obj.Name); err != nil {
		return nil, err
	}
	str2, err := cmds[2].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str2, &obj.Mailbox); err != nil {
		return nil, err
	}
	str3, err := cmds[3].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str3, &obj.Sex); err != nil {
		return nil, err
	}
	str4, err := cmds[4].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str4, &obj.Longitude); err != nil {
		return nil, err
	}
	str5, err := cmds[5].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str5, &obj.Latitude); err != nil {
		return nil, err
	}
	str6, err := cmds[6].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str6, &obj.Description); err != nil {
		return nil, err
	}
	str7, err := cmds[7].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str7, &obj.Password); err != nil {
		return nil, err
	}
	str8, err := cmds[8].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str8, &obj.HeadUrl); err != nil {
		return nil, err
	}
	str9, err := cmds[9].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(str9, &obj.Status); err != nil {
		return nil, err
	}
	str10, err := cmds[10].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	var val10 string
	if err := m.StringScan(str10, &val10); err != nil {
		return nil, err
	}
	obj.CreatedAt = orm.TimeParse(val10)
	str11, err := cmds[11].(*redis.StringCmd).Result()
	if err != nil {
		return nil, err
	}
	var val11 string
	if err := m.StringScan(str11, &val11); err != nil {
		return nil, err
	}
	obj.UpdatedAt = orm.TimeParse(val11)
	return obj, nil
}

func (m *_UserRedisMgr) FetchByIds(ids []string) ([]*User, error) {
	objs := make([]*User, len(ids))
	for _, id := range ids {
		obj, err := m.Fetch(id)
		if err != nil {
			return objs, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func (m *_UserRedisMgr) Create(obj *User) error {
	return m.Save(obj)
}

func (m *_UserRedisMgr) Update(obj *User) error {
	return m.Save(obj)
}

func (m *_UserRedisMgr) Delete(obj *User) error {
	//! uniques
	uk_key_0 := []string{
		"Mailbox",
		fmt.Sprint(obj.Mailbox),
		"Password",
		fmt.Sprint(obj.Password),
	}
	uk_mgr_0 := MailboxPasswordOfUserUKRelationRedisMgr(m.RedisStore)
	if err := uk_mgr_0.PairRem(strings.Join(uk_key_0, ":")); err != nil {
		return err
	}

	//! indexes
	idx_key_0 := []string{
		"Sex",
		fmt.Sprint(obj.Sex),
	}
	idx_mgr_0 := SexOfUserIDXRelationRedisMgr(m.RedisStore)
	idx_rel_0 := idx_mgr_0.NewSexOfUserIDXRelation(strings.Join(idx_key_0, ":"))
	idx_rel_0.Value = obj.Id
	if err := idx_mgr_0.SetRem(idx_rel_0); err != nil {
		return err
	}

	//! ranges
	rg_key_0 := []string{
		"Name",
		fmt.Sprint(obj.Name),
		"Status",
	}
	rg_mgr_0 := NameStatusOfUserRNGRelationRedisMgr(m.RedisStore)
	rg_rel_0 := rg_mgr_0.NewNameStatusOfUserRNGRelation(strings.Join(rg_key_0, ":"))
	rg_rel_0.Score = orm.ToFloat64(obj.Status)
	rg_rel_0.Value = obj.Id
	if err := rg_mgr_0.ZSetRem(rg_rel_0); err != nil {
		return err
	}

	//! orders

	return m.Del(keyOfObject(obj, fmt.Sprint(obj.Id))).Err()
}

func (m *_UserRedisMgr) Save(obj *User) error {
	pipe := m.BeginPipeline()
	//! fields
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Id", fmt.Sprint(obj.Id))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Name", fmt.Sprint(obj.Name))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Mailbox", fmt.Sprint(obj.Mailbox))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Sex", fmt.Sprint(obj.Sex))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Longitude", fmt.Sprint(obj.Longitude))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Latitude", fmt.Sprint(obj.Latitude))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Description", fmt.Sprint(obj.Description))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Password", fmt.Sprint(obj.Password))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "HeadUrl", fmt.Sprint(obj.HeadUrl))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Status", fmt.Sprint(obj.Status))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "CreatedAt", fmt.Sprint(orm.TimeFormat(obj.CreatedAt)))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "UpdatedAt", fmt.Sprint(orm.TimeFormat(obj.UpdatedAt)))
	if _, err := pipe.Exec(); err != nil {
		return err
	}

	//! uniques
	uk_key_0 := []string{
		"Mailbox",
		fmt.Sprint(obj.Mailbox),
		"Password",
		fmt.Sprint(obj.Password),
	}
	uk_mgr_0 := MailboxPasswordOfUserUKRelationRedisMgr(m.RedisStore)
	uk_rel_0 := uk_mgr_0.NewMailboxPasswordOfUserUKRelation(strings.Join(uk_key_0, ":"))
	uk_rel_0.Value = obj.Id
	if err := uk_mgr_0.PairAdd(uk_rel_0); err != nil {
		return err
	}

	//! indexes
	idx_key_0 := []string{
		"Sex",
		fmt.Sprint(obj.Sex),
	}
	idx_mgr_0 := SexOfUserIDXRelationRedisMgr(m.RedisStore)
	idx_rel_0 := idx_mgr_0.NewSexOfUserIDXRelation(strings.Join(idx_key_0, ":"))
	idx_rel_0.Value = obj.Id
	if err := idx_mgr_0.SetAdd(idx_rel_0); err != nil {
		return err
	}

	//! ranges
	rg_key_0 := []string{
		"Name",
		fmt.Sprint(obj.Name),
		"Status",
	}
	rg_mgr_0 := NameStatusOfUserRNGRelationRedisMgr(m.RedisStore)
	rg_rel_0 := rg_mgr_0.NewNameStatusOfUserRNGRelation(strings.Join(rg_key_0, ":"))
	rg_rel_0.Score = orm.ToFloat64(obj.Status)
	rg_rel_0.Value = obj.Id
	if err := rg_mgr_0.ZSetAdd(rg_rel_0); err != nil {
		return err
	}

	//! orders
	return nil
}

func (m *_UserRedisMgr) Clear() error {
	return nil
}

//! uniques

//! relation
type MailboxPasswordOfUserUKRelation struct {
	Key   string `db:"key" json:"key"`
	Value int32  `db:"value" json:"value"`
}

func (relation *MailboxPasswordOfUserUKRelation) GetClassName() string {
	return "MailboxPasswordOfUserUKRelation"
}

func (relation *MailboxPasswordOfUserUKRelation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *MailboxPasswordOfUserUKRelation) GetStoreType() string {
	return "pair"
}

func (relation *MailboxPasswordOfUserUKRelation) GetPrimaryName() string {
	return "Key"
}

type _MailboxPasswordOfUserUKRelationRedisMgr struct {
	*orm.RedisStore
}

func MailboxPasswordOfUserUKRelationRedisMgr(stores ...*orm.RedisStore) *_MailboxPasswordOfUserUKRelationRedisMgr {
	if len(stores) > 0 {
		return &_MailboxPasswordOfUserUKRelationRedisMgr{stores[0]}
	}
	return &_MailboxPasswordOfUserUKRelationRedisMgr{_redis_store}
}

//! pipeline write
type _MailboxPasswordOfUserUKRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) BeginPipeline() *_MailboxPasswordOfUserUKRelationRedisPipeline {
	return &_MailboxPasswordOfUserUKRelationRedisPipeline{m.Pipeline(), nil}
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) NewMailboxPasswordOfUserUKRelation(key string) *MailboxPasswordOfUserUKRelation {
	return &MailboxPasswordOfUserUKRelation{
		Key: key,
	}
}

//! redis relation pair
func (m *_MailboxPasswordOfUserUKRelationRedisMgr) PairAdd(obj *MailboxPasswordOfUserUKRelation) error {
	return m.Set(pairOfClass(obj.GetClassName(), obj.Key), obj.Value, 0).Err()
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) PairGet(key string) (*MailboxPasswordOfUserUKRelation, error) {
	str, err := m.Get(pairOfClass("MailboxPasswordOfUserUKRelation", key)).Result()
	if err != nil {
		return nil, err
	}

	obj := m.NewMailboxPasswordOfUserUKRelation(key)
	if err := m.StringScan(str, &obj.Value); err != nil {
		return nil, err
	}
	return obj, nil
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) PairRem(key string) error {
	return m.Del(pairOfClass("MailboxPasswordOfUserUKRelation", key)).Err()
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) FindOne(key string) (string, error) {
	return m.Get(pairOfClass("MailboxPasswordOfUserUKRelation", key)).Result()
}

//! indexes

//! relation
type SexOfUserIDXRelation struct {
	Key   string `db:"key" json:"key"`
	Value int32  `db:"value" json:"value"`
}

func (relation *SexOfUserIDXRelation) GetClassName() string {
	return "SexOfUserIDXRelation"
}

func (relation *SexOfUserIDXRelation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *SexOfUserIDXRelation) GetStoreType() string {
	return "set"
}

func (relation *SexOfUserIDXRelation) GetPrimaryName() string {
	return "Key"
}

type _SexOfUserIDXRelationRedisMgr struct {
	*orm.RedisStore
}

func SexOfUserIDXRelationRedisMgr(stores ...*orm.RedisStore) *_SexOfUserIDXRelationRedisMgr {
	if len(stores) > 0 {
		return &_SexOfUserIDXRelationRedisMgr{stores[0]}
	}
	return &_SexOfUserIDXRelationRedisMgr{_redis_store}
}

//! pipeline write
type _SexOfUserIDXRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_SexOfUserIDXRelationRedisMgr) BeginPipeline() *_SexOfUserIDXRelationRedisPipeline {
	return &_SexOfUserIDXRelationRedisPipeline{m.Pipeline(), nil}
}

func (m *_SexOfUserIDXRelationRedisMgr) NewSexOfUserIDXRelation(key string) *SexOfUserIDXRelation {
	return &SexOfUserIDXRelation{
		Key: key,
	}
}

//! redis relation pair
func (m *_SexOfUserIDXRelationRedisMgr) SetAdd(obj *SexOfUserIDXRelation) error {
	return m.SAdd(setOfClass(obj.GetClassName(), obj.Key), obj.Value).Err()
}

func (m *_SexOfUserIDXRelationRedisMgr) SetGet(key string) ([]*SexOfUserIDXRelation, error) {
	strs, err := m.SMembers(setOfClass("SexOfUserIDXRelation", key)).Result()
	if err != nil {
		return nil, err
	}

	objs := make([]*SexOfUserIDXRelation, len(strs))
	for _, str := range strs {
		obj := m.NewSexOfUserIDXRelation(key)
		if err := m.StringScan(str, &obj.Value); err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func (m *_SexOfUserIDXRelationRedisMgr) SetRem(obj *SexOfUserIDXRelation) error {
	return m.SRem(setOfClass(obj.GetClassName(), obj.Key), obj.Value).Err()
}

func (m *_SexOfUserIDXRelationRedisMgr) Find(key string) ([]string, error) {
	return m.SMembers(setOfClass("SexOfUserIDXRelation", key)).Result()
}

//! ranges

//! relation
type NameStatusOfUserRNGRelation struct {
	Key   string  `db:"key" json:"key"`
	Score float64 `db:"score" json:"score"`
	Value int32   `db:"value" json:"value"`
}

func (relation *NameStatusOfUserRNGRelation) GetClassName() string {
	return "NameStatusOfUserRNGRelation"
}

func (relation *NameStatusOfUserRNGRelation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *NameStatusOfUserRNGRelation) GetStoreType() string {
	return "zset"
}

func (relation *NameStatusOfUserRNGRelation) GetPrimaryName() string {
	return "Key"
}

type _NameStatusOfUserRNGRelationRedisMgr struct {
	*orm.RedisStore
}

func NameStatusOfUserRNGRelationRedisMgr(stores ...*orm.RedisStore) *_NameStatusOfUserRNGRelationRedisMgr {
	if len(stores) > 0 {
		return &_NameStatusOfUserRNGRelationRedisMgr{stores[0]}
	}
	return &_NameStatusOfUserRNGRelationRedisMgr{_redis_store}
}

//! pipeline write
type _NameStatusOfUserRNGRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_NameStatusOfUserRNGRelationRedisMgr) BeginPipeline() *_NameStatusOfUserRNGRelationRedisPipeline {
	return &_NameStatusOfUserRNGRelationRedisPipeline{m.Pipeline(), nil}
}

func (m *_NameStatusOfUserRNGRelationRedisMgr) NewNameStatusOfUserRNGRelation(key string) *NameStatusOfUserRNGRelation {
	return &NameStatusOfUserRNGRelation{
		Key: key,
	}
}

//! redis relation zset
func (m *_NameStatusOfUserRNGRelationRedisMgr) ZSetAdd(obj *NameStatusOfUserRNGRelation) error {
	return m.ZAdd(zsetOfClass(obj.GetClassName(), obj.Key), redis.Z{Score: obj.Score, Member: obj.Value}).Err()
}

func (m *_NameStatusOfUserRNGRelationRedisMgr) ZSetRange(key string, min, max int64) ([]*NameStatusOfUserRNGRelation, error) {
	strs, err := m.ZRange(zsetOfClass("NameStatusOfUserRNGRelation", key), min, max).Result()
	if err != nil {
		return nil, err
	}

	objs := make([]*NameStatusOfUserRNGRelation, len(strs))
	for _, str := range strs {
		obj := m.NewNameStatusOfUserRNGRelation(key)
		if err := m.StringScan(str, &obj.Value); err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func (m *_NameStatusOfUserRNGRelationRedisMgr) ZSetRem(obj *NameStatusOfUserRNGRelation) error {
	return m.ZRem(zsetOfClass(obj.GetClassName(), obj.Key), redis.Z{Score: obj.Score, Member: obj.Value}).Err()
}

func (m *_NameStatusOfUserRNGRelationRedisMgr) Range(key string, min, max int64) ([]string, error) {
	return m.ZRange(zsetOfClass("NameStatusOfUserRNGRelation", key), min, max).Result()
}

func (m *_NameStatusOfUserRNGRelationRedisMgr) OrderBy(key string, asc bool) ([]string, error) {
	//! TODO revert
	return m.ZRange(zsetOfClass("NameStatusOfUserRNGRelation", key), 0, -1).Result()
}

//! orders
