package ordcol

import "sort"

type item struct {
    key int
    value int
}

type collection struct {
    items map[int]Item
    mas []int
}

type iterator struct {
    items map[int]Item
    mas []int
    index int
}

func NewItem(key int, value int) *item {
    m := item{
        key: key,
        value: value,
    }
    return &m
}

func NewCollection() *collection {
    m := collection{
        items: make(map[int]Item),
        mas: make([]int, 0),
    }
    return &m
}

func (c *collection) NewIterator(order IterationOrder) *iterator {
    if order == ByInsertion {
        m := iterator{
            items: c.items,
            mas: c.mas,
            index: 0,
        }
        return &m
    } else {
        newm := c.mas
        sort.Ints(newm)
        m := iterator {
            items: c.items,
            mas: newm,
            index: 0,
        }
        return &m
    }
}

func (m *item) Key() int {
    return m.key
}

func (m *item) Value() int {
    return m.value
}

func (c *collection) Add(item Item) error {
    _, ok := c.items[item.Key()]
    if ok {
        return ErrDuplicateKey
    }
    c.items[item.Key()] = item
    c.mas = append(c.mas, item.Key())
    return nil
}

func (c *collection) At(key int) (Item, bool) {
    m, ok := c.items[key]
    if !ok {
        return nil, false
    }
    return m, true
}

func (i *iterator) HasNext() bool {
    if i.index == len(i.mas) {
        return false
    }
    return true
}


func (i *iterator) Next() (Item, error) {
    if !i.HasNext() {
        return nil, ErrEmptyIterator
    }
    item := i.items[i.mas[i.index]]
    i.index++
    return item, nil

}

func (c *collection) IterateBy(order IterationOrder) Iterator {
    if (order != ByKey && order != ByInsertion) {
        panic("bad key")
    }
    return c.NewIterator(order)
}

