package main

import f "fmt"

func main() {
	printBioForms()
}

func printBioForms() {
	forms := [4]bioForm {
		bioForm{"Carcharadon", "Carcharias", true}, bioForm{"Tyrannosaurus", "Rex", false}, Hominin("Sapiens"),Hominin("Habilis")}

	for _, form := range forms {
		f.Println(form.scientificName())
	}
}

type bioForm struct {
	genus string
	species string
	extant bool
}

func (f *bioForm) scientificName() string {
	s := f.genus + " " + f.species + ", extinct: "
	if f.extant {
		s += "not yet!"
	} else {
		s += "yes"
	}
	return s
}

func Hominin(species string) bioForm {
	h := bioForm{genus: "Homo"}
	h.species = species
	if species == "Sapiens" {
		h.extant = true
	} else {
		h.extant = false
	}
	return h
}
