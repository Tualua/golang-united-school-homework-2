package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type NumSides int

const (
	SidesCircle   NumSides = 0
	SidesTriangle NumSides = 3
	SidesSquare   NumSides = 4
)

func CalcSquare(sideLen float64, sidesNum NumSides) (res float64) {

	switch sidesNum {
	case 0:
		{
			res = math.Pi * sideLen * sideLen
		}
	case 3:
		{
			res = math.Sqrt(3) * math.Pow(sideLen, 2) / 4
		}
	case 4:
		{
			res = sideLen * sideLen
		}
	default:
		{
			res = 0
		}
	}

	return res
}
