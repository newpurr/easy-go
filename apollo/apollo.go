package apollo

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/storage"
	"reflect"
	"strings"
)

const (
	DefaultNamespace = "application"
)

var (
	client              *agollo.Client
	agolloLoader        *AgolloLoader
	apolloValueRegister = NewValueRegister()
)

func SetDefault(c *agollo.Client) {
	client = c
}

func SetLoader(a *AgolloLoader) {
	agolloLoader = a
}

type ChangeListener struct {
}

func (c *ChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	agolloLoader.Refresh(changeEvent)
}

func (c *ChangeListener) OnNewestChange(event *storage.FullChangeEvent) {}

type ValueRegister struct {
	fieldInitMap map[interface{}]struct{}
	FiledCache   map[string]map[string]reflect.Value
}

func NewValueRegister() *ValueRegister {
	return &ValueRegister{
		map[interface{}]struct{}{},
		map[string]map[string]reflect.Value{},
	}
}

func (r *ValueRegister) Register(ns, subKey string, v reflect.Value) {
	mp := r.getOrNewMap(ns)
	mp[subKey] = v

	r.FiledCache[ns] = mp
}

func (r *ValueRegister) Get(key, subKey string) *reflect.Value {
	mp := r.getOrNewMap(key)
	v, ok := mp[subKey]
	if !ok {
		return nil
	}
	return &v
}

func (r *ValueRegister) getOrNewMap(k string) map[string]reflect.Value {
	mp, ok := r.FiledCache[k]
	if !ok {
		r.FiledCache[k] = map[string]reflect.Value{}
		mp = r.FiledCache[k]
	}

	return mp
}

func (r *ValueRegister) Reflect(target interface{}) {
	_, ok := r.fieldInitMap[target]
	if ok {
		return
	}

	typeOfTarget := reflect.TypeOf(target).Elem()
	valueOfTarget := reflect.ValueOf(target).Elem()
	if numField := valueOfTarget.NumField(); numField > 0 {
		for i := 0; i < numField; i++ {
			f := valueOfTarget.Field(i)
			if !f.CanInterface() {
				continue
			}
			f2 := typeOfTarget.Field(i)
			tag := f2.Tag.Get("apollo-binding")
			if tag == "" {
				continue
			}
			tags := strings.SplitN(tag, "@", 2)
			if len(tags) == 1 {
				tags = []string{DefaultNamespace, tags[0]}
			}
			key := tags[1]
			ns := tags[0]

			r.Register(ns, key, f)
		}
	}
}

type AgolloLoader struct {
	*agollo.Client
	reference map[interface{}]interface{}
}

func NewAgolloLoader(c *agollo.Client) *AgolloLoader {
	return &AgolloLoader{
		c,
		make(map[interface{}]interface{}),
	}
}

func (a *AgolloLoader) Refresh(changeEvent *storage.ChangeEvent) {
	for _, i := range a.reference {
		a.ReadWithRefresh(i)
	}
}

func (a *AgolloLoader) ReadWithRefresh(v interface{}) {
	apolloValueRegister.Reflect(v)

	for ns, m := range apolloValueRegister.FiledCache {
		cache := a.Client.GetConfigCache(ns)
		if cache == nil {
			return
		}
		cache.Range(func(key, v interface{}) bool {
			value, ok := m[key.(string)]
			if ok {
				value.SetString(v.(string))
			}

			return true
		})
	}

	if _, ok := a.reference[v]; !ok {
		a.reference[v] = v
	}

	return
}
