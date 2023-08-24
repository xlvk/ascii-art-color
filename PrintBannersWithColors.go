package asciiArtColor

import (
	"fmt"
	"strings"
)

// Print the full outcome in the triminal
func PrintBannersWithColors(Str, colors string, banners, arr []string) {
	colors = strings.ToLower(colors)
	num := 0
	RGB := "rgb"
	HEX := "#"
	HSL := "hsl"
	ResestColor := "\033[0m"
	var color []string
	//this one for RGP to ansi nums.
	if strings.HasPrefix(colors, RGB) {
		color = RGBtoNum(colors)
		//this one for HEX to ansi nums.
	} else if strings.HasPrefix(colors, HEX) {
		color = HextoRGB(colors)
		//this one for HSL to ansi nums.
	} else if strings.HasPrefix(colors, HSL) {
		color = HSLtoRGB(colors)
		//} else if ......... {
		//for ANSI
	} else {
		//for the left (word: red, green, blue, etc...).
		color = WordColors(colors)
	}
	colors = "\033[38;2;" + color[0] + ";" + color[1] + ";" + color[2] + "m"
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				fmt.Println()
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			if Str == "" {
				for _, j := range ch {
				n := (j-32)*9 + 1
				fmt.Print(colors, arr[int(n)+i])
				}
			} else {
				for _, j := range ch {
					check := true
					n := (j-32)*9 + 1
					for _, kk := range Str {
						if kk == j {
								fmt.Print(colors, arr[int(n)+i])
								check = false
						}
					}
					if check == true {
						fmt.Print(ResestColor, arr[int(n)+i])
					}
					
				}
			}
			
			fmt.Println()
		}
	}
}
