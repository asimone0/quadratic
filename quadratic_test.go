package quadratic

import (
    "testing"
)


func TestXYZ(t *testing.T) {
	x1, x2 := Solve(1,3,2)
	t.Log(x1)
	t.Log(x2)
	r1:=real(x1)
	i1:=imag(x1)
	if(!(r1==-1 && i1==0)){
		t.Log("Solve() gave wrong answer")
		t.Fail()
	}
	r2:=real(x2)
	i2:=imag(x2)
	if(!(r2==-2 && i2==0)){
		t.Log("Solve() gave wrong answer")
		t.Fail()
	}
}

