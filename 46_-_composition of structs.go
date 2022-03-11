
//embedded structs

package main
import "fmt"

// We create a struct details
type details struct {
	genre	 string
	genreRating string
	reviews	 string
}

// We create a struct game
type game struct {
	name string
	price string
	details //adding struct details to struct game as an embedded member
}

// this is a method defined
// on the details struct
func (d details) showDetails() {
	fmt.Println("Genre:", d.genre)
	fmt.Println("Genre Rating:", d.genreRating)
	fmt.Println("Reviews:", d.reviews)
}

// this is a method defined
// on the game struct
// this method has access
// to showDetails() as well since
// the details struct is embedded in 
//  the game struct
func (g game) show() {
	fmt.Println("Name: ", g.name)
	fmt.Println("Price:", g.price)
	fmt.Println("Price:", g.genre)
	g.showDetails()
}

func main() {

	// defining a struct
	// object of Type details
	action := details{"Action","18+", "mostly positive"}

	// defining a struct
	// object of Type game
	newGame := game{"batman","$25", action}

	newGame.show()
}

