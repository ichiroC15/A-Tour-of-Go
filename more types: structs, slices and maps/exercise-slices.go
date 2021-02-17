package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// まず容量dy, 各要素が配列[]uint8となる配列を作成
	pic := make([][]uint8, dy)

	// 配列picの各要素に対してループを回す
	for y := range pic {
		// 容量dx, 各要素がuint8の配列を作成
		pic[y] = make([]uint8, dx)
		// インデックスx, yの値を適当にいじった値をunit8に変換して、ループを回している空配列に入れる
		for x := range pic[y] {
			pic[y][x] = uint8((x * y) / 2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
