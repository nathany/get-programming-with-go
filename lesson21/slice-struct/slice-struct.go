package main

func main() {
	lats := []float64{-4.5895, -14.5684, -1.9462}
	longs := []float64{137.4417, 175.472636, 354.4734}

	_ = lats
	_ = longs
	type location struct {
		name string
		lat  float64
		long float64
	}

	locations := []location{
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}

	_ = locations
}
