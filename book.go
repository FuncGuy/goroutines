package main

import (
	"fmt"
	_ "fmt"
)

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf("Title:\t\t%q\n"+"Author\t\t%q\n"+"Published\t%v\n",
		b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "Atomic Habits",
		Author:        "James Clear",
		YearPublished: 2020,
		},
	{
		ID:            2,
		Title:         "Elon Musk",
		Author:        "Ashely",
		YearPublished: 2019,
		},
	{
		ID:            3,
		Title:         "Deep Work",
		Author:        "Unknown",
		YearPublished: 2020,
		},
	{
		ID:            4,
		Title:         "Live like a monk",
		Author:        "Jay shetty",
		YearPublished: 2019,
		},
	{
		ID:            5,
		Title:         "C",
		Author:        "Dennis Ritche",
		YearPublished: 1985,
		},
	{
		ID:            6,
		Title:         "C++",
		Author:        "Strostorp",
		YearPublished: 1990,
		}, {
		ID:            7,
		Title:         "Java",
		Author:        "James Gosling",
		YearPublished: 1993,
		}, {
		ID:            8,
		Title:         "Python",
		Author:        "Not getting",
		YearPublished: 1990,
		}, {
		ID:            9,
		Title:         "Hello",
		Author:        "world",
		YearPublished: 2019,
		}, {
		ID:            10,
		Title:         "World",
		Author:        "hello",
		YearPublished: 2019,
		},
}
