package tempconv

// CToF は摂氏を華氏へ変換します
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は華氏を摂氏へ変換します
func FToC(f Fahrenheit) Celsius { return Celsius((f-32))*5/9 }


