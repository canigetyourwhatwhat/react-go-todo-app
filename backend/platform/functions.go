package platform

import "errors"

func New() *ItemList {
	return &ItemList{
		Items: []Item{},
	}
}

func (l *ItemList) Add(item Item){
	l.Items = append(l.Items, item)
}

func (l *ItemList) GetAll() []Item {
	return l.Items
}

func (l *ItemList) GetOneItem(target string) (*Item, error) {
	for index, value := range l.Items{
		if value.ID == target{
			return &l.Items[index], nil
		}
	}
	return nil, errors.New("couldn't find the item'")
}

func (l *ItemList) PutItem(target string) (*Item, error){
	item, err := l.GetOneItem(target)
	if err != nil{
		return item, err
	}

	item.Completed = !item.Completed
	return item, err
}

func (l *ItemList) DeleteItem(target string) (*Item, error){
	for index, item := range l.Items{
		if item.ID == target {
			l.Items = append(l.Items[:index], l.Items[index+1:]...)
			return &item, nil
		}
	}
	return nil, errors.New("couldn't delete the item")
}