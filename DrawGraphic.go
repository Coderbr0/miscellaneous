package main

/* • The Ascii Art Colour Program;
filename: DrawGraphic.go; author: Kamal H. Zada; Date: 9 november 2021;

• Ascii-art is a program which consists of receiving a string as an argument and outputting the string in a graphic representation using ASCII.
What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below.

• I must follow the same instructions as in the first subject (Ascii Art) but this time with colours.
The output should manipulate colours using the flag --colour=<colour>, in which --colour is the flag and <colour> is the colour desired by the user.
These colours can be achieved using different notations (colour code systems, like RGB, hsl, ANSI...), it is up to you to choose which one you want to use.
You should be able to choose between colouring a single letter or a set of letters (use your imagination for this one).
If the letter is not specified, the whole string should be coloured.
The flag must have exactly the same format as above, any other formats must return the following usage message:
Usage: go run . [STRING] [OPTION]
Example: go run . "Some String." --colour=<colour>

• option 1: User selects one of 8 standard colours: Black, Red, Green, Yellow, Blue, Magenta, Cyan, White;
Usage: --colour=<black | red | green | yellow | blue | magenta | cyan | white>;<Blink switch>
Example: go run . "Some String"  --colour=magenta       !<Blink switch> = 0 for off OR 1 for Blink On (by default Blink is off);
Example: go run . "Hello World" "--colour=magenta;1"    !<Blink switch> = 0 for off OR 1 for Blink On;
Example: go run . "Hello World" --colour=green\;1       !<Blink switch> = 0 for off OR 1 for Blink On;

• option 2: Next there exists the third argument: for a Letter Range or Word "(<lower bound>:<upper bound>)" OR "(index) for a single letter"
Example: go run . "Hello World" "--colour=cyan" "(7)"      !Selects the single letter: W; displays W in Cyan.
Example: go run . "Hello World" "--colour=blue" "(7:11)"   !Selects the single word: World; displays World in Blue.
parameter Select Word or a single Letter: selWordLetr (the third argument)

• option 3: RGB Palette, RGB - Red Green Blue; Digital 8-bit per channel, (255, 100, 200) OR #FF64C8 (hexadecimal);
#FF64C8 (hexadecimal) means that Red=FF, Green=64, and Blue=C8;
Usage: --colour=<#RedGreenBlue>
Example: go run . "Hello World" "--colour=#FF64C8"     !The Colour is a mixture of: Red amount=255 ; Green amount=100; Blue amount=200
Example: go run . "Hello World" "--colour=#FF64C8;1"   !<Blink switch> = 0 for off OR 1 for Blink On;
Example: go run . "Hello World" --colour=#FF64C8\;1    !<Blink switch> = 0 for off OR 1 for Blink On;

• option 4: RGB Palette, RGB - Red Green Blue; Digital 8-bit per channel, (255, 100, 200) OR #FF64C8 (hexadecimal);
#FF64C8 (hexadecimal) means that Red=FF, Green=64, and Blue=C8;
Usage: --colour=<rgb(255, 100, 200)>
Example: go run . "Hello World" "--colour=rgb(255, 100, 200)"     !The Colour is a mixture of: Red amount=255 ; Green amount=100; Blue amount=200
Example: go run . "Hello World" "--colour=rgb(255, 100, 200);1"   !<Blink switch> = 0 for off OR 1 for Blink On;
Example: go run . "Hello World" --colour=rgb(255, 100, 200)\;1    !<Blink switch> = 0 for off OR 1 for Blink On;
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/* Function Validate HEX Colour Palette;
• option 3: HEX Palette, RGB - Red Green Blue; Digital 8-bit per channel, (255, 100, 200) OR #FF64C8 (hexadecimal);
#FF64C8 (hexadecimal) means that Red=FF, Green=64, and Blue=C8;
Usage: --colour=<#RedGreenBlue>
Example: go run . "Hello World" "--colour=#FF64C8"    !The Colour is a mixture of: Red amount=255 ; Green amount=100; Blue amount=200*/
func ValidatHexColPlt(colourPlt string) byte {
	var idx1 int = 0
	var idx2 int = 0
	var idx3 int = 0
	var redInt int64 = 0
	var redByt byte = 0
	var greenInt int64 = 0
	var greenByt byte = 0
	var blueInt int64 = 0
	var blueByt byte = 0
	var eror1 error
	var blinkOn bool = false
	var retVal byte = 0 /*Return value: Error code; 0 for success*/
	var colPltLen int = 0
	var colourHex string
	var redStr string
	var greenStr string
	var blueStr string

	/*fmt.Println("VH1: Validate HEX Colour Palette Entered; Colour Palette:", colourPlt)*/
	colPltLen = len(colourPlt)
	/*Extract the Hexadecimal value (three bytes)*/
	idx1 = strings.Index(colourPlt, "#")
	idx1++
	idx2 = strings.Index(colourPlt, ";")
	if idx2 == -1 {
		/*there is NO ; */
		idx3 = colPltLen
	} else {
		idx3 = idx2
	} /*if there is no ; */
	/*Extract the required colour (RGB in Hex)*/
	colourHex = colourPlt[idx1:idx3]
	/*Extract and Set the Blink Switch*/
	if idx2 == -1 {
		/*there is NO ; */
		blinkOn = false
	} else {
		idx3++
		if colourPlt[idx3] == 49 {
			/*49 = ascii code for "1"; if Blink switch is On*/
			blinkOn = true
		} else {
			blinkOn = false
		} /*if*/
	} /*if there is no ; */
	retVal = 0
	/*fmt.Println("VH2: Colour String:", colourHex, "; Blink Switch:", blinkOn, "; Idx1:", idx1, "; Idx2:", idx2, "; Idx3:", idx3)*/
	/*Extract the individual base colours: Red, Green and Blue*/
	redStr = colourHex[0:2]
	greenStr = colourHex[2:4]
	blueStr = colourHex[4:6]
	/*fmt.Println("VH3: Colour String:", colourHex, "; Red string:", redStr, "; Green String:", greenStr, "; Blue string:", blueStr)*/
	/*Convert the base colours: Red, Green and Blue to integer values*/
	redInt, eror1 = strconv.ParseInt(redStr, 16, 64)
	if eror1 == nil {
		/*No error*/
		redByt = byte(redInt)
		greenInt, eror1 = strconv.ParseInt(greenStr, 16, 64)
		if eror1 == nil {
			/*No error*/
			greenByt = byte(greenInt)
			/*fmt.Println("Green:", greenByt, redByt)*/
			blueInt, eror1 = strconv.ParseInt(blueStr, 16, 64)
			if eror1 == nil {
				/*No error*/
				retVal = 0
				blueByt = byte(blueInt)
				bytRGBslc[0] = redByt
				bytRGBslc[1] = greenByt
				bytRGBslc[2] = blueByt
				colour = "RGBHex"
				blinkSwitch = blinkOn
				/*fmt.Println("VH4: Colour String:", colourHex, "; Red Byte:", redByt, "; Green Byte:", greenByt, "; Blue Byte:", blueByt)*/
			} else {
				/*Error in Blue string*/
				fmt.Println("Error in Blue String;", eror1, blueStr)
				retVal = 3
			} /*if*/
		} else {
			/*Error in Green string*/
			fmt.Println("Error in Green String;", eror1, greenStr)
			retVal = 2
		} /*if*/
	} else {
		/*Error in Red string*/
		fmt.Println("Error in Red String;", eror1, redStr)
		retVal = 1
	} /*if*/
	return (retVal)
} /*ValidatHexColPlt*/

/* Function Validate RGB Colour Palette;
• option 4: RGB Palette, RGB - Red Green Blue; Digital 8-bit per channel, (255, 100, 200) OR #FF64C8 (hexadecimal);
#FF64C8 (hexadecimal) means that Red=FF, Green=64, and Blue=C8;
Usage: --colour=<rgb(255, 100, 200)>
Example: go run . "Hello World" "--colour=rgb(255, 100, 200)"     !The Colour is a mixture of: Red amount=255 ; Green amount=100; Blue amount=200
Example: go run . "Hello World" "--colour=rgb(255, 100, 200);1"   !<Blink switch> = 0 for off OR 1 for Blink On;
Example: go run . "Hello World" --colour=rgb(255, 100, 200)\;1    !<Blink switch> = 0 for off OR 1 for Blink On;*/
func ValidatRGBColPlt(colourPlt string) byte {
	var idx1 int = 0
	var idx2 int = 0
	var idx3 int = 0
	var redInt int64 = 0
	var redByt byte = 0
	var greenInt int64 = 0
	var greenByt byte = 0
	var blueInt int64 = 0
	var blueByt byte = 0
	var eror1 error
	var blinkOn bool = false
	var retVal byte = 0 /*Return value: Error code; 0 for success*/
	var colPltLen int = 0
	var colourRGB string
	var colourRGBslc []string
	var redStr string
	var greenStr string
	var blueStr string

	/*fmt.Println("VR1: Validate RGB Colour Palette Entered; Colour Palette:", colourPlt)*/
	colPltLen = len(colourPlt)
	/*Extract the RGB value (three bytes)*/
	idx1 = strings.Index(colourPlt, "rgb")
	idx1 = idx1 + 4
	idx2 = strings.Index(colourPlt, ";")
	if idx2 == -1 {
		/*there is NO ; */
		idx3 = colPltLen - 1
	} else {
		idx3 = idx2 - 1
	} /*if there is no ; */
	/*Extract the required colour (RGB in Hex)*/
	colourRGB = colourPlt[idx1:idx3]
	/*Extract and Set the Blink Switch*/
	if idx2 == -1 {
		/*there is NO ; */
		blinkOn = false
	} else {
		idx3 = idx3 + 2
		if colourPlt[idx3] == 49 {
			/*49 = ascii code for "1"; if Blink switch is On*/
			blinkOn = true
		} else {
			blinkOn = false
		} /*if*/
	} /*if there is no ; */
	retVal = 0
	/*Extract the individual base colours: Red, Green and Blue*/
	colourRGBslc = strings.Split(colourRGB, ",")
	/*fmt.Println("VR2: Colour String:", colourRGB, "; Colour String Slice:", colourRGBslc)*/
	redStr = colourRGBslc[0]
	redStr = strings.TrimSpace(redStr)
	greenStr = colourRGBslc[1]
	greenStr = strings.TrimSpace(greenStr)
	blueStr = colourRGBslc[2]
	blueStr = strings.TrimSpace(blueStr)
	/*fmt.Println("VR3: Colour String:", colourRGB, "; Red string:", redStr, "; Green String:", greenStr, "; Blue string:", blueStr)*/
	/*Convert the base colours: Red, Green and Blue to integer values*/
	redInt, eror1 = strconv.ParseInt(redStr, 10, 64)
	if eror1 == nil {
		/*No error*/
		redByt = byte(redInt)
		greenInt, eror1 = strconv.ParseInt(greenStr, 10, 64)
		if eror1 == nil {
			/*No error*/
			greenByt = byte(greenInt)
			blueInt, eror1 = strconv.ParseInt(blueStr, 10, 64)
			if eror1 == nil {
				/*No error*/
				retVal = 0
				blueByt = byte(blueInt)
				bytRGBslc[0] = redByt
				bytRGBslc[1] = greenByt
				bytRGBslc[2] = blueByt
				colour = "RGBDec"
				blinkSwitch = blinkOn
				/*fmt.Println("VR4: Colour String:", colourRGB, "; Red Byte:", redByt, "; Green Byte:", greenByt, "; Blue Byte:", blueByt)*/
			} else {
				/*Error in Blue string*/
				fmt.Println("Error in Blue String;", eror1, blueStr)
				retVal = 3
			} /*if*/
		} else {
			/*Error in Green string*/
			fmt.Println("Error in Green String;", eror1, greenStr)
			retVal = 2
		} /*if*/
	} else {
		/*Error in Red string*/
		fmt.Println("Error in Red String;", eror1, redStr)
		retVal = 1
	} /*if*/
	return (retVal)
} /*ValidatRGBColPlt*/

/*Function Draw Graphic 2 → Ascii-art is a program which consists of receiving a string as an argument
and outputting the string in a graphic representation using ASCII.
DrawGraphic2 is called when the third argument is present;
If there exists the third argument: for a Letter Range or Word "(<lower bound>:<upper bound>)" OR "(index) for a single letter"
Example: go run . "Hello World" "--colour=cyan" "(7)"      !Selects the single letter: W; displays W in Cyan.
Example: go run . "Hello World" "--colour=blue" "(7:11)"   !Selects the single word: World; displays World in Blue.
When the third argument selects a single Letter "(8)", then DrawGraphic2 is called.*/
func DrawGraphic2() byte {
	var err error
	var file *os.File
	var idx1 int = 0
	var idx2 int = 0
	var idx3 int = 0
	var idx4 int = 0
	var idx5 int = 0
	var colour3 int = 0
	var scanner *bufio.Scanner
	var line []string
	var strArt []string
	var colour2 string
	var retVal byte = 0 /*Return value: Error code; 0 for success*/

	colour2 = "white"
	/*fmt.Println("DG2A: Draw Graphic 2 Entered; Colour:", colour, "; Single Letter has been selected:", selSingleLeter, "; Single Letter Index: ", selLeterIdx)*/

	/*Usage: --colour=<Black | Red | Green | Yellow | Blue | Magenta | Cyan | White>;<Blink switch>*/
	switch colour2 {
	case "black":
		if blinkSwitch {
			fmt.Println("\033[30;1;5m")
		} else {
			fmt.Println("\033[30;1;25m")
		} /*if*/
	case "red":
		if blinkSwitch {
			fmt.Println("\033[31;1;5m")
		} else {
			fmt.Println("\033[31;1;25m")
		} /*if*/
	case "green":
		if blinkSwitch {
			fmt.Println("\033[32;1;5m")
		} else {
			fmt.Println("\033[32;1;25m")
		} /*if*/
	case "yellow":
		if blinkSwitch {
			fmt.Println("\033[33;1;5m")
		} else {
			fmt.Println("\033[33;1;25m")
		} /*if*/
	case "blue":
		if blinkSwitch {
			fmt.Println("\033[34;1;5m")
		} else {
			fmt.Println("\033[34;1;25m")
		} /*if*/
	case "magenta":
		if blinkSwitch {
			fmt.Println("\033[35;1;5m")
		} else {
			fmt.Println("\033[35;1;25m")
		} /*if*/
	case "cyan": /*a greenish colour*/
		if blinkSwitch {
			fmt.Println("\033[36;1;5m")
		} else {
			fmt.Println("\033[36;1;25m")
		} /*if*/
	case "white":
		fmt.Println("\033[37;1;25m")
		switch colour {
		case "black":
			colour3 = 30
		case "red":
			colour3 = 31
		case "green":
			colour3 = 32
		case "yellow":
			colour3 = 33
		case "blue":
			colour3 = 34
		case "magenta":
			colour3 = 35
		case "cyan":
			colour3 = 36
		case "white":
			colour3 = 37
		default:
		} /*switch*/
	default:
		/*Error: Non Standard Colour*/
		fmt.Println("Error: Non Standard Colour, please use one of: --colour=<black | red | green | yellow | blue | magenta | cyan | white>.")
		retVal = 1
		return (retVal)
	} /*switch*/

	file, err = os.Open("standard.txt")
	if err != nil {
		log.Fatalf("Error: %s; Is the standard.txt banner file present? Are you sure?", err)
	}
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	line = []string{}
	bannerMap := make(map[int][]string)
	idx1 = 0
	idx2 = 32
	for scanner.Scan() {
		line = append(line, scanner.Text())
		idx1++
		if idx1 == 9 {
			bannerMap[idx2] = line
			line = []string{}
			idx2++
			idx1 = 0
		} /*if*/
	} /*for loop*/
	file.Close()

	strArt = strings.Split(os.Args[1], "\\n")

	for idx3 = 0; idx3 < len(strArt); idx3++ {
		if strArt[idx3] == "" {
			fmt.Printf("\n")
		} else {
			/*fmt.Println("DG2B: Idx3:", idx3, "; Idx4:", idx4, "; Idx5:", idx5)*/
			for idx4 = 1; idx4 <= 8; idx4++ {
				for idx5 = range strArt[idx3] {
					if idx5 == (selLeterIdx - 1) {
						/*fmt.Printf("DG2C: Idx3:%v; Idx4:%v; Idx5:%v", idx3, idx4, idx5)*/
						fmt.Printf("\033[%v;1;25m", colour3)
					} else {
					} /*if*/
					fmt.Printf(bannerMap[int(strArt[idx3][idx5])][idx4])
					if idx5 == (selLeterIdx - 1) {
						/*fmt.Printf("DG2D: Idx3:%v; Idx4:%v; Idx5:%v", idx3, idx4, idx5)*/
						fmt.Printf("\033[37;1;25m")
					} else {
					} /*if*/
				}
				fmt.Printf("\n")
			} /*for loop*/
		}
	} /*for loop*/
	return (retVal)
} /*DrawGraphic2*/

/*Function Draw Graphic 3 → Ascii-art is a program which consists of receiving a string as an argument
and outputting the string in a graphic representation using ASCII.
DrawGraphic3 is called when the third argument is present;
If there exists the third argument: for a Letter Range or Word "(<lower bound>:<upper bound>)" OR "(index) for a single letter"
Example: go run . "Hello World" "--colour=cyan" "(7)"      !Selects the single letter: W; displays W in Cyan.
Example: go run . "Hello World" "--colour=blue" "(7:11)"   !Selects the single word: World; displays World in Blue.
When the third argument selects a Word "(7:11)", then DrawGraphic3 is called.*/
func DrawGraphic3() byte {
	var err error
	var file *os.File
	var idx1 int = 0
	var idx2 int = 0
	var idx3 int = 0
	var idx4 int = 0
	var idx5 int = 0
	var colour3 int = 0
	var scanner *bufio.Scanner
	var line []string
	var strArt []string
	var colour2 string
	var retVal byte = 0 /*Return value: Error code; 0 for success*/

	colour2 = "white"
	/*fmt.Println("DG3A: Draw Graphic 3 Entered; Colour:", colour, "; A Word has been selected:", selSingleLeter, "; Letter Index No 1: ", selLeterIdx, "; Letter Index No 2: ", selLeterIdx2)*/

	/*Usage: --colour=<Black | Red | Green | Yellow | Blue | Magenta | Cyan | White>;<Blink switch>*/
	switch colour2 {
	case "black":
		if blinkSwitch {
			fmt.Println("\033[30;1;5m")
		} else {
			fmt.Println("\033[30;1;25m")
		} /*if*/
	case "red":
		if blinkSwitch {
			fmt.Println("\033[31;1;5m")
		} else {
			fmt.Println("\033[31;1;25m")
		} /*if*/
	case "green":
		if blinkSwitch {
			fmt.Println("\033[32;1;5m")
		} else {
			fmt.Println("\033[32;1;25m")
		} /*if*/
	case "yellow":
		if blinkSwitch {
			fmt.Println("\033[33;1;5m")
		} else {
			fmt.Println("\033[33;1;25m")
		} /*if*/
	case "blue":
		if blinkSwitch {
			fmt.Println("\033[34;1;5m")
		} else {
			fmt.Println("\033[34;1;25m")
		} /*if*/
	case "magenta":
		if blinkSwitch {
			fmt.Println("\033[35;1;5m")
		} else {
			fmt.Println("\033[35;1;25m")
		} /*if*/
	case "cyan": /*a greenish colour*/
		if blinkSwitch {
			fmt.Println("\033[36;1;5m")
		} else {
			fmt.Println("\033[36;1;25m")
		} /*if*/
	case "white":
		fmt.Println("\033[37;1;25m")
		switch colour {
		case "black":
			colour3 = 30
		case "red":
			colour3 = 31
		case "green":
			colour3 = 32
		case "yellow":
			colour3 = 33
		case "blue":
			colour3 = 34
		case "magenta":
			colour3 = 95
		case "cyan":
			/*Bright Cyan is a blueish colour*/
			colour3 = 96
		case "white":
			colour3 = 37
		case "orange":
			/*Orange → Hex RGB: #FF6400 → RGB decimal: (255, 100, 0)*/
			colour3 = 100
		default:
			/*Error: Non Standard Colour*/
			fmt.Println("\033[37;0;25m")
			fmt.Println("Error: Non Standard Colour, please use one of: --colour=<black | red | green | yellow | blue | magenta | cyan | white>.")
			retVal = 1
			return (retVal)
		} /*switch*/
	default:
		/*Error: Non Standard Colour - this code is never called.*/
		fmt.Println("Error: Non Standard Colour, please use one of: --colour=<black | red | green | yellow | blue | magenta | cyan | white>.")
		retVal = 2
		return (retVal)
	} /*switch*/

	file, err = os.Open("standard.txt")
	if err != nil {
		log.Fatalf("Error: %s; Is the standard.txt banner file present? Are you sure?", err)
	}
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	line = []string{}
	bannerMap := make(map[int][]string)
	idx1 = 0
	idx2 = 32
	for scanner.Scan() {
		line = append(line, scanner.Text())
		idx1++
		if idx1 == 9 {
			bannerMap[idx2] = line
			line = []string{}
			idx2++
			idx1 = 0
		} /*if*/
	} /*for loop*/
	file.Close()

	strArt = strings.Split(os.Args[1], "\\n")

	for idx3 = 0; idx3 < len(strArt); idx3++ {
		if strArt[idx3] == "" {
			fmt.Printf("\n")
		} else {
			/*fmt.Println("DG3B: Idx3:", idx3, "; Idx4:", idx4, "; Idx5:", idx5)*/
			for idx4 = 1; idx4 <= 8; idx4++ {
				for idx5 = range strArt[idx3] {
					if idx5 == (selLeterIdx - 1) {
						/*fmt.Printf("DG3C: Idx3:%v; Idx4:%v; Idx5:%v", idx3, idx4, idx5)*/
						if colour3 == 100 {
							/*if colour = orange*/
							/*Orange → Hex RGB: #FFD700 → RGB decimal: (255, 100, 0)*/
							fmt.Printf("\x1b[38;2;%v;%v;%vm", 255, 100, 0)
						} else {
							fmt.Printf("\033[%v;1;25m", colour3)
						} /*if*/
					} else {
					} /*if*/
					fmt.Printf(bannerMap[int(strArt[idx3][idx5])][idx4])
					if idx5 == (selLeterIdx2 - 1) {
						/*fmt.Printf("DG3D: Idx3:%v; Idx4:%v; Idx5:%v", idx3, idx4, idx5)*/
						fmt.Printf("\033[37;1;25m")
					} else {
					} /*if*/
				}
				fmt.Printf("\n")
			} /*for loop*/
		}
	} /*for loop*/
	return (retVal)
} /*DrawGraphic3*/

/*Function Draw Graphic 4 → Ascii-art is a program which consists of receiving a string as an argument
and outputting the string in a graphic representation using ASCII.
DrawGraphic4 is called when the RGB Hex notation for colour is utilised;
• option 3: RGB Palette, RGB - Red Green Blue; Digital 8-bit per channel, (255, 100, 200) OR #FF64C8 (hexadecimal);
#FF64C8 (hexadecimal) means that Red=FF, Green=64, and Blue=C8;
Usage: --colour=<#RedGreenBlue>
Example: go run . "Hello World" "--colour=#FF64C8"     !The Colour is a mixture of: Red amount=255 ; Green amount=100; Blue amount=200
Example: go run . "Hello World" "--colour=#FF64C8;1"   !<Blink switch> = 0 for off OR 1 for Blink On;
Example: go run . "Hello World" --colour=#FF64C8\;1    !<Blink switch> = 0 for off OR 1 for Blink On;
*/
func DrawGraphic4() byte {
	var err error
	var file *os.File
	var idx1 int = 0
	var idx2 int = 0
	var idx3 int = 0
	var idx4 int = 0
	var idx5 int = 0
	var redByt byte = 0
	var greenByt byte = 0
	var blueByt byte = 0
	var scanner *bufio.Scanner
	var line []string
	var strArt []string
	var retVal byte = 0 /*Return value: Error code; 0 for success*/

	/*fmt.Println("DG4A: Draw Graphic 4 Entered; RGB Hex Colour:", bytRGBslc, "; Blink Switch:", blinkSwitch)*/
	redByt = bytRGBslc[0]
	greenByt = bytRGBslc[1]
	blueByt = bytRGBslc[2]
	switch colour {
	case "RGBHex", "RGBDec":
		if blinkSwitch {
			fmt.Printf("\x1b[38;2;%v;%v;%v;5m", redByt, greenByt, blueByt)
		} else {
			fmt.Printf("\x1b[38;2;%v;%v;%vm", redByt, greenByt, blueByt)
		} /*if*/
	default:
		/*Error: Non standard Colour*/
		fmt.Println("Error: Non Standard Colour, please use one of: --colour=<black | red | green | yellow | blue | magenta | cyan | white>.")
		retVal = 1
		return (retVal)
	} /*switch*/

	file, err = os.Open("standard.txt")
	if err != nil {
		log.Fatalf("Error: %s; Is the standard.txt banner file present? Are you sure?", err)
	}
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	line = []string{}
	bannerMap := make(map[int][]string)
	idx1 = 0
	idx2 = 32
	for scanner.Scan() {
		line = append(line, scanner.Text())
		idx1++
		if idx1 == 9 {
			bannerMap[idx2] = line
			line = []string{}
			idx2++
			idx1 = 0
		} /*if*/
	} /*for loop*/
	file.Close()

	strArt = strings.Split(os.Args[1], "\\n")

	for idx3 = 0; idx3 < len(strArt); idx3++ {
		if strArt[idx3] == "" {
			fmt.Printf("\n")
		} else {
			for idx4 = 1; idx4 <= 8; idx4++ {
				for idx5 = range strArt[idx3] {
					fmt.Printf(bannerMap[int(strArt[idx3][idx5])][idx4])
					/*fmt.Println("DG4B: Idx3:", idx3, "; Idx4:", idx4, "; Idx5:", idx5)*/
				}
				fmt.Printf("\n")
			}
		}
	} /*for loop*/

	return (retVal)
} /*DrawGraphic4*/
