package main 

import "fmt"

func add(x float64, y float64) float64{
	var sum float64 = x + y
	return sum
}

func subtract(x float64, y float64) float64{
	var sum float64 = x - y
	return sum
}

func multiply(x float64, y float64) float64{
	var sum float64 = x * y
	return sum
}

func divide(x float64, y float64) float64{
	var sum float64 = x / y
	return sum
}

func main(){

	var num1, num2 float64

	for i := 0; i < 10; i++ {
		fmt.Println("This is a demo")
	} 

	fmt.Println("Enter in the first number: ") //DEMO FILE
	fmt.Scanf("%f", &num1)

	fmt.Println("Enter in the second number: ") //DEMO FILE
	fmt.Scanf("%f", &num2)

	for i := 0; i <= 4; i++{

		var sum float64

		if i == 1{
			sum = add(num1, num2)
			fmt.Println("The sum of adding", num1, " and ", num2, "is: ", sum)
		}else if i == 2{
			sum = subtract(num1, num2)
			fmt.Println("The sum of subtracting", num1, " and ", num2, "is: ", sum)
		}else if i == 3{
			sum = divide(num1, num2)
			fmt.Println("The sum of dividing", num1, " and ", num2, "is: ", sum)
		}else if i == 4{
			sum = multiply(num1, num2)
			fmt.Println("The sum of multipling", num1, " and ", num2, "is: ", sum)
		}else {
			sum = add(num1, num2)
			fmt.Println("The sum of ", num1, " and ", num2, "is: ", sum)
		}
	}
}