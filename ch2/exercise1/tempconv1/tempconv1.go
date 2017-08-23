// tempconv1 は摂氏・華氏・絶対温度の温度変換を行います。
package tempconv1

import "fmt"

type Kelvin     float64
type Celsius    float64
type Fahrenheit float64

const(
	AbsoluteZero Kelvin = 0
  Freezing     Kelvin = 273.15
	Boiling      Kelvin = 373.15
)

func (t Kelvin)     String() string { return fmt.Sprintf("%gK", t) }
func (c Celsius)    String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }