/*A computer store sells diskettes at a price of $1 each for small orders.
The store sells them at a price of 70 cents each for orders of 25+ units.
Furthermore, if you are a member of Club Z, the store will give you an extra discount of 2%.
Write an algorithm that determines the total price for a purchase.*/

package main

import (
	"fmt"
)

func main(){

		var myNumber int
		for  {

			fmt.Println("Enter the number of disquette or type exit: ")

			n, err := fmt.Scanln(&myNumber)

			if  n<1 || err != nil {
				fmt.Println("exit!")
				fmt.Println("bye bye , see you soon !")
				return
			}


			var price float32
			if  myNumber <25 {
				price  = 1.00
			}else {
				price  = 0.70
			}

			var club string ="N"

			fmt.Println("are you member of clubZ? enter 'Y' for yes, 'N' for no")
			n, err = fmt.Scanln(&club)
			fmt.Println(n)
			if  n<1 || err != nil {
				fmt.Println("invalid input")
				return
			}


			var totalBill float32

           if club=="N" {
			   totalBill = price * float32(myNumber)
			   s := fmt.Sprintf("%.2f", totalBill)
			   fmt.Println("totalBill  : " + s + "$")
		   }else if club=="Y" {
				totalBill = price * float32(myNumber)  - (.02)*(price * float32(myNumber))
				s := fmt.Sprintf("%.2f", totalBill) //
				fmt.Println("totalBill  : " + s + "$")
			}

		}

	}






