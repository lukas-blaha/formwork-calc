package main

type App struct {
	wall         Wall
	formworkList []Formwork
	rows         int
}

func (a *App) SetRows() {
	var maxWidth int
	var maxHeight int
	var biggest Formwork

	for _, f := range a.formworkList {
		if f.Type == "panel" {
			if f.Height >= maxHeight && f.Width >= maxWidth {
				biggest = f
				maxHeight = f.Height
				maxWidth = f.Width
			}
		}
	}

	if biggest.Height > a.wall.Height {
		a.rows = 1
	}
}

func main() {
	fList := NewFormworkList("formwork.csv")

	// w := NewWall()
	w := Wall{
		Type:   "external",
		Width:  1100,
		Height: 305,
		Depth:  25,
	}

	a := App{
		wall:         w,
		formworkList: fList,
	}

	a.SetRows()
}
