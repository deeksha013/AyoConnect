package HouseRobber

import("fmt"
       "math")


func collect(sli []float64)float64{
  if len(sli)==0{
    return 0
  }
  if(len(sli)==1){
    return sli[0]
  }
  if len(sli)==2{
    return math.Max(sli[0],sli[1])
  }
  sli[1]=math.Max(sli[0],sli[1])
  for i:=2;i<len(sli);i++{
    sli[i] = math.Max(sli[i]+sli[i-2],sli[i-1])
  }
  return sli[len(sli)-1]
}

func robbedMoney(sli []float64)float64{
  if len(sli)==0{
    return 0
  }
  if(len(sli)==1){
    return sli[0]
  }
  if len(sli)==2{
    return math.Max(sli[0],sli[1])
  }
  slicopy:=make([]float64,len(sli))
  copy(slicopy, sli)
  return math.Max(collect(sli[0:len(sli)-1]),collect(slicopy[1:len(slicopy)]))
}

func RunHouseRobber(){
 fmt.Println("[2,3,2] :", robbedMoney([]float64{2,3,2}))
 fmt.Println("[1,2,3,1] : ",robbedMoney([]float64{1,2,3,1}))
 fmt.Println("[] : ",robbedMoney([]float64{}))
}
