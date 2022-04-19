package platform

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T){
	repo := New()

	repo.Add(Item{"1", "it is detail", false})

	if len(repo.Items) != 1{
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T){
	repo := New()

	repo.Add(Item{"1", "task 1", false})
	repo.Add(Item{"2", "task 2", false})
	repo.Add(Item{"3", "task 3", false})

	if len(repo.GetAll()) != 3{
		t.Errorf("Not all items are displayed")
	}
}

func TestGetOne(t *testing.T){
	repo := New()

	item := Item{"2", "task 2", false}

	repo.Add(Item{"1", "task 1", false})
	repo.Add(item)
	repo.Add(Item{"3", "task 3", false})

	result, err := repo.GetOneItem("2")

	if err != nil {
		t.Errorf("Can't get the item")
	}else if reflect.DeepEqual(result, item) {
		t.Errorf("Picked item is not correct")
	}
}

func TestPutItem(t *testing.T){
	repo := New()

	repo.Add(Item{"1", "task 1", false})
	repo.Add(Item{"2", "task 2", false})
	repo.Add(Item{"3", "task 3", false})

	result, err := repo.PutItem("2")
	expected := Item{"2", "task 2", true}

	if err != nil {
		t.Errorf("Can't get the item")
	}else if reflect.DeepEqual(result, expected) {
		t.Errorf("Picked item is not modified")
	}
}

func TestDeleteItem(t *testing.T) {
	repo := New()

	repo.Add(Item{"1", "task 1", false})
	repo.Add(Item{"2", "task 2", false})
	repo.Add(Item{"3", "task 3", false})

	_, err := repo.DeleteItem("2")

	if err != nil {
		t.Errorf("Can't get the item")
	}else if len(repo.Items) != 2 {
		t.Errorf("Picked item is not modified")
	}
}
