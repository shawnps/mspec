package examples

import "errors"

// this file holds the example models under test.
//

const normalColor = "brown"

type paint struct {
	color      string
	iswashable bool
}

type dog struct {
	color        string
	paint        *paint
	washed       bool
	timesWashed  int
	passedAway   bool
	hasBrokenLeg bool
}

func BirthDog() *dog {
	return &dog{
		color:       normalColor,
		paint:       nil,
		washed:      false,
		timesWashed: 0,
	}
}

func (d *dog) Paint(p *paint) {
	if p == nil {
		return
	}
	d.color = p.color
	d.paint = p
	d.washed = false
}

func (d *dog) Wash() error {

	if d.paint.iswashable {
		d.color = normalColor
		d.paint = nil
		d.washed = true
		d.timesWashed++
		return nil
	}

	return errors.New("The paint is not washable!")
}
