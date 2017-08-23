package tempconv1

// KToC は絶対温度を摂氏へ変換します
func KToC(t Kelvin) Celsius { return Celsius(t - 273.15) }

// CToK は摂氏を絶対温度へ変換します
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// CToF は摂氏を華氏へ変換します
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は華氏を摂氏へ変換します
func FToC(f Fahrenheit) Celsius { return Celsius((f-32))*5/9 }

// KToF は絶対温度を華氏へ変換します
func KToF(t Kelvin) Fahrenheit { return CToF(KToC(t)) }

// FToK は華氏を絶対温度へ変換します
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }



