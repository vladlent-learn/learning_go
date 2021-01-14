package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	fourWheel bool
	vehicle
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		fourWheel: true,
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: false,
	}

	// Print struct with field names
	fmt.Printf("%+v\n", t)

	// Inner field gets promoted
	fmt.Println(t.color)
	fmt.Println(t.doors)
	fmt.Println(t.fourWheel)

	fmt.Println(s.color)
	fmt.Println(s.doors)
	fmt.Println(s.luxury)

	anon := struct {
		desc string
	}{
		desc: "This is an anonymous struct",
	}

	fmt.Println(anon)
}
