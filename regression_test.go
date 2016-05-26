package regression

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	r := new(Regression)
	r.SetObserved("Murders per annum per 1,000,000 inhabitants")
	r.SetVar(0, "Inhabitants")
	r.SetVar(1, "Percent with incomes below $5000")
	r.SetVar(2, "Percent unemployed")
	dps := []*DataPoint{
		NewDataPoint(11.2, []float64{587000, 16.5, 6.2}),
		NewDataPoint(13.4, []float64{643000, 20.5, 6.4}),
		NewDataPoint(40.7, []float64{635000, 26.3, 9.3}),
		NewDataPoint(5.3, []float64{692000, 16.5, 5.3}),
		NewDataPoint(24.8, []float64{1248000, 19.2, 7.3}),
		NewDataPoint(12.7, []float64{643000, 16.5, 5.9}),
		NewDataPoint(20.9, []float64{1964000, 20.2, 6.4}),
		NewDataPoint(35.7, []float64{1531000, 21.3, 7.6}),
		NewDataPoint(8.7, []float64{713000, 17.2, 4.9}),
		NewDataPoint(9.6, []float64{749000, 14.3, 6.4}),
		NewDataPoint(14.5, []float64{7895000, 18.1, 6}),
		NewDataPoint(26.9, []float64{762000, 23.1, 7.4}),
		NewDataPoint(15.7, []float64{2793000, 19.1, 5.8}),
		NewDataPoint(36.2, []float64{741000, 24.7, 8.6}),
		NewDataPoint(18.1, []float64{625000, 18.6, 6.5}),
		NewDataPoint(28.9, []float64{854000, 24.9, 8.3}),
		NewDataPoint(14.9, []float64{716000, 17.9, 6.7}),
		NewDataPoint(25.8, []float64{921000, 22.4, 8.6}),
		NewDataPoint(21.7, []float64{595000, 20.2, 8.4}),
		NewDataPoint(25.7, []float64{3353000, 16.9, 6.7}),
	}
	r.AddDataPoint(dps)
	r.Train()

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)

	// All vars are known to positively correlate with the murder rate
	for i, c := range r.coeff {
		if i == 0 {
			// This is the offset and not a coeff
			continue
		}
		if c < 0 {
			t.Errorf("Coefficient is negative, but shouldn't be: %.2f", c)
		}
	}

	//  We know this set has an R^2 above 80
	if r.R2 < 0.8 {
		t.Errorf("R^2 was %.2f, but we expected > 80", r.R2)
	}
}
