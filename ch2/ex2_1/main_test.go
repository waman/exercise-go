package tempconv

import (
	"fmt"
)

func ExampleKelvin(){
	fmt.Printf("絶対零度は %gK\n", AbsoluteZero)
	fmt.Printf("水の融点は %gK\n", Freezing)
	fmt.Printf("水の沸点は %gK\n", Boiling)
	// Output:
	// 絶対零度は 0K
	// 水の融点は 273.15K
	// 水の沸点は 373.15K
}

func ExampleCelsius(){
	fmt.Printf("絶対零度は %g°C\n", KToC(AbsoluteZero))
	fmt.Printf("水の融点は %g°C\n", KToC(Freezing))
	fmt.Printf("水の沸点は %g°C\n", KToC(Boiling))
	// Output:
	// 絶対零度は -273.15°C
	// 水の融点は 0°C
	// 水の沸点は 100°C
}

func ExampleFahrenheit(){
	fmt.Printf("絶対零度は %g°F\n", KToF(AbsoluteZero))
	fmt.Printf("水の融点は %g°F\n", KToF(Freezing))
	fmt.Printf("水の沸点は %g°F\n", KToF(Boiling))
	// Output:
	// 絶対零度は -459.66999999999996°F
	// 水の融点は 32°F
	// 水の沸点は 212°F
}
