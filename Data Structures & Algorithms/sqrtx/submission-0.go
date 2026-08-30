func mySqrt(x int) int {
	l:=1
	h:=x/2
	if(x==1){
		return 1
	}

	for l<=h{
		mid:=l+(h-l)/2
		r:=mid*mid
		if r==x{
			return mid
		} else if(r<x){
		
			l=mid+1
		} else{
			h=mid-1
		}
	}
	return l-1
}
