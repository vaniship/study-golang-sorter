package qsort

func qsort(data []int)  {
  if len(data) <= 1 {
    return
  }
  mid, i := data[0], 1
  head, tail := 0, len(data)-1
  for head < tail {
    if data[i] > mid {
      data[i], data[tail] = data[tail], data[i]
      tail--
    } else {
      data[i], data[head] = data[head], data[i]
      head++
      i++
    }
  }
  data[head] = mid
  qsort(data[:head])
  qsort(data[head+1:])
}

func quickSort(values []int, left, right int) {
  temp:=values[left]
  p:=left
  i,j:=left,right

  for i<=j{
    for j>=p && values[j]>=temp{
      j--
    }
    if j>=p {
      values[p] = values[j]
      p=j
    }

    if values[i]<=temp && i<=p{
      i++
    }

    if i<=p{
      values[p]=values[i]
      p=i
    }
  }

  values[p]=temp
  if p-left>1{
    quickSort(values,left,p-1)
  }
  if right-p>1{
    quickSort(values,p+1,right)
  }
}

func QuickSort(values []int) {
	//quickSort(values, 0, len(values)-1)
  qsort(values)
}
