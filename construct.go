package main

type Item struct {
	NoOption   string
	Parameter1 string
	Parameter2 int
}

type option func(item *Item)

func NewItem(opts ...option) *Item {
	i := &Item{
		NoOption:   "usual",
		Parameter1: "default",
		Parameter2: 42,
	}

	for _, opt := range opts {
		opt(i)
	}

	return i
}

func Option1(option1 string) option {
	return func(i *Item) {
		i.Parameter1 = option1
	}
}

func Option2(option2 int) option {
	return func(i *Item) {
		i.Parameter2 = option2
	}

}
