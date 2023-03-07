

//далить i-ый элемент из слайса.

ackage main

mport "fmt"

func main() {
	b := [string{"G", "H", "J", "K", "L"}
	i := 2
copy(b[i:], b[i+1:]) //Совершить сдвиг a[i+1:] влево на один индекс

b[len(b)-1] = "" //Удалить последний элемент

b = b[:len(b)-1] //Усечь срез

	mt.Println(b)
}
