package main

/* • The Ascii Art Colour Program;
filename: main.go; author: Kamal H. Zada; Date: 3 november 2021;

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

• option 3: HEX (RGB) Palette, RGB - Red Green Blue; Digital 8-bit per channel, (255, 100, 200) OR #FF64C8 (hexadecimal);
#FF64C8 (hexadecimal) means that Red=FF, Green=64, and Blue=C8;
Usage: --colour=<#RedGreenBlue>
Example: go run . "Hello World" "--colour=#FF64C8"    !The Colour is a mixture of: Red amount=255 ; Green amount=100; Blue amount=200

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

/*Global Variables*/
var (
	blinkSwitch    bool   = false /*Blink Switch: Blink On = true; Blink Off = false*/
	selSingleLeter bool   = true  /*Is a single letter being selected.*/
	selLeterIdx    int    = 0     /*Selected Single Letter Index no 1 (Lower Bound)*/
	selLeterIdx2   int    = 0     /*Selected Single Letter Index no 2 (Upper Bound); used when there is a Range of Letters*/
	bytRGBslc      []byte         /*RGB (Red Green Blue) Byte Slice, RGB Palette*/
	colour         string /*Colour: one of 8 standard colours: Black, Red, Green, Yellow, Blue, Magenta, Cyan, White */)

/*Function Validate the Colour Argument → --colour=<colour>
• option 1: User selects one of 8 standard colours: Black, Red, Green, Yellow, Blue, Magenta, Cyan, White;
Usage: --colour=<black | red | green | yellow | blue | magenta | cyan | white>;<Blink switch>
Example: go run . "Some String"  --colour=magenta       !<Blink switch> = 0 for off OR 1 for Blink On (by default Blink is off);
Example: go run . "Hello World" "--colour=magenta;1"    !<Blink switch> = 0 for off OR 1 for Blink On;
• option 2: Next there exists the third argument: for a Letter Range or Word "(<lower bound>:<upper bound>)" OR "(index) for a single letter"
Example: go run . "Hello World" "--colour=cyan" "(7)"      !Selects the single letter: W; displays W in Cyan.
Example: go run . "Hello World" "--colour=blue" "(7:11)"   !Selects the single word: World; displays World in Blue.
• option 3: RGB Palette, RGB - Red Green Blue; Digital 8-bit per channel, (255, 100, 200) OR #FF64C8 (hexadecimal);
#FF64C8 (hexadecimal) means that Red=FF, Green=64, and Blue=C8;
Usage: --colour=<#RedGreenBlue>
Example: go run . "Hello World" "--colour=#FF64C8"    !The Colour is a mixture of: Red amount=255 ; Green amount=100; Blue amount=200
parameter Colour Palette: colourPlt */
func ValidatColPlt(colourPlt string) byte {
	var colPltValid bool = false
	var blinkOn bool = false
	var colourStr string
	var idx1 int = 0
	var idx2 int = 0
	var idx3 int = 0
	var idx4 int = 0 /*Index for the RGB option 3*/
	var idx5 int = 0 /*Index for the RGB option 4*/
	var colPltLen int = 0
	var retVal byte = 0  /*Return value: Error code; 0 for success*/
	var retVal2 byte = 0 /*Return value from Validate Hex Palette: Error code; 0 for success*/

	colPltLen = len(colourPlt)
	/*fmt.Println("VCL1: Validate Colour Palette Entered; Colour Palette:", colourPlt, "; Colour Palette Length:", colPltLen)*/
	colPltValid = strings.Contains(colourPlt, "--colour=")
	if colPltValid {
		/*Colour Palette is valid*/
		idx1 = strings.Index(colourPlt, "=")
		idx1++
		idx3 = strings.Index(colourPlt, ";")
		idx4 = strings.Index(colourPlt, "#")
		idx5 = strings.Index(colourPlt, "rgb")
		if idx4 == -1 {
			/*there is NO # */
		} else {
			/*there is a #, so it's Validate HEX Palette */
			retVal2 = ValidatHexColPlt(colourPlt)
			retVal2 = retVal2 + 2
			return (retVal2)
		} /*if*/
		if idx5 == -1 {
			/*there is NO rgb */
		} else {
			/*there is "rgb", so it's Validate RGB in decimal Palette */
			retVal2 = ValidatRGBColPlt(colourPlt)
			retVal2 = retVal2 + 2
			return (retVal2)
		} /*if*/
		if idx3 == -1 {
			/*there is NO ; */
			idx2 = colPltLen
		} else {
			idx2 = idx3
		} /*if there is no ; */
		/*Extract the required colour*/
		colourStr = colourPlt[idx1:idx2]
		/*Extract and Set the Blink Switch*/
		if idx3 == -1 {
			/*there is NO ; */
			blinkOn = false
		} else {
			idx2++
			if colourPlt[idx2] == 49 {
				/*49 = ascii code for "1"; if Blink switch is On*/
				blinkOn = true
			} else {
				blinkOn = false
			} /*if*/
		} /*if there is no ; */
		colour = colourStr
		retVal = 0
		/*fmt.Println("VCL2: Colour String:", colourStr, "; Blink Switch:", blinkOn, "; Idx1:", idx1, "; Idx2:", idx2, "; Idx3:", idx3)*/
	} else {
		/*Error: Colour Palette is not valid*/
		colour = ""
		retVal = 1
		fmt.Println("Error: Colour Palette is not valid; Usage: --colour=<Black | Red | Green | Yellow | Blue | Magenta | Cyan | White>")
		fmt.Println("Usage: go run . [STRING] [OPTION]\nEX: go run . something --colour=<colour>")
	} /*if*/
	blinkSwitch = blinkOn
	return (retVal)
} /*ValidatColPlt*/

/*Function Validate the Word or a single Letter Argument → "(n:m)"
option 1: User selects one of 8 standard colours: Black, Red, Green, Yellow, Blue, Magenta, Cyan, White;
Usage: --colour=<black | red | green | yellow | blue | magenta | cyan | white>;<Blink switch>
Example: go run . "Some String"  --colour=magenta       !<Blink switch> = 0 for off OR 1 for Blink On (by default Blink is off);
Example: go run . "Hello World" "--colour=magenta;1"    !<Blink switch> = 0 for off OR 1 for Blink On;
Next there exists the third argument: for a Letter Range or Word "(<lower bound>:<upper bound>)" OR "(index) for a single letter"
Example: go run . "Hello World" "--colour=cyan" "(7)"      !Selects the single letter: W; displays W in Cyan.
Example: go run . "Hello World" "--colour=blue" "(7:11)"   !Selects the single word: World; displays World in Blue.
parameter Select Word or a single Letter: selWordLetr (the third argument) */
func ValidatWrdSel(selWordLetr string) byte {
	var retVal byte = 0 /*Return value: Error code; 0 for success*/
	/*var selWrdLen int = 0*/
	var idx1 int = 0
	var idx2 int = 0
	var letrIdx1 int = 0
	var letrIdx2 int = 0
	var inpStrLen int = 0
	var numStr1 string
	var numStr2 string
	var eror1 error
	var singleLetr bool = false

	/*selWrdLen = len(selWordLetr)*/
	/*fmt.Println("VWS1: Validate Select Word or a single Letter Entered; Select Word/Letter:", selWordLetr, "; Select Word/Letter Length:", selWrdLen)*/
	inpStrLen = len(os.Args[1])
	idx1 = strings.Index(selWordLetr, ":")
	if idx1 == -1 {
		/*there is NO ':'  so a single letter is being selected.*/
		singleLetr = true
		/*Extract the Letter number (index)*/
		idx2 = strings.Index(selWordLetr, ")")
		if idx2 == -1 {
			/*there is no ')' so it's an error*/
			retVal = 1
			fmt.Println("Error: Invalid third argument, should be (n).")
		} else {
			/*there is a ')' so we can continue*/
			/*Extract the Letter number (index)*/
			if selWordLetr[0] == 40 && selWordLetr[idx2] == 41 {
				/*40 is Ascii code for '('  and 41 is Ascii code for ')' */
				/*Extract the Letter number (index), numStr1*/
				numStr1 = selWordLetr[1:idx2]
				/*fmt.Println("VWS2: Number String:", numStr1, "; Idx1:", idx1, "; Idx2:", idx2)*/
				letrIdx1, eror1 = strconv.Atoi(numStr1)
				if eror1 == nil {
					/*No error*/
					/*fmt.Println("VWS3: Letter Index:", letrIdx1, "; Error:", eror1)*/
					/* Now check that the Letter index is valid */
					if (letrIdx1 > 0) && (letrIdx1 <= inpStrLen) {
						/*No error so set Global var: Selected Letter Index*/
						selLeterIdx = letrIdx1
					} else {
						retVal = 2
						fmt.Println("Error: Single Letter Index:", letrIdx1, "is greater than length of Input String OR is Less than 1.")
					} /*if*/
				} else {
					log.Fatalf("Error: Select a single Letter is not numeric; %s", eror1)
				} /*if*/
			} else {
				/*Error*/
				retVal = 3
				fmt.Println("Error: Select a single Letter should be in parenthesis (n).")
			} /*if*/
		} /*if*/
	} else {
		/*a Letters range is being specified, e.g. (5:11)*/
		singleLetr = false
		/*Extract the Letters Range (index1:index2)*/
		idx2 = strings.Index(selWordLetr, ")")
		if idx2 == -1 {
			/*there is no ')' so it's an error*/
			retVal = 4
			fmt.Println("Error: Invalid third argument, should be (n1:n2).")
		} else {
			/*there is a ')' so we can continue*/
			if selWordLetr[0] == 40 && selWordLetr[idx2] == 41 {
				/*40 is Ascii code for '('  and 41 is Ascii code for ')' */
				/*Extract the first Letter number (index1), numStr1*/
				numStr1 = selWordLetr[1:idx1]
				/*fmt.Println("VWS4: Number String no 1:", numStr1, "; Idx1:", idx1, "; Idx2:", idx2)*/
				letrIdx1, eror1 = strconv.Atoi(numStr1)
				if eror1 == nil {
					/*No error*/
					/*fmt.Println("VWS5: Letter Index No 1:", letrIdx1, "; Error:", eror1)*/
					/* Now check that the Letter index is valid */
					if (letrIdx1 > 0) && (letrIdx1 <= inpStrLen) {
						/*No error so set Global var: Selected Letter Index No 1*/
						selLeterIdx = letrIdx1
						/*Now extract the Letter Index no 2*/
						idx1++
						numStr2 = selWordLetr[idx1:idx2]
						/*fmt.Println("VWS6: Number String no 2:", numStr2, "; Idx1:", idx1, "; Idx2:", idx2)*/
						letrIdx2, eror1 = strconv.Atoi(numStr2)
						if eror1 == nil {
							/*No error*/
							/*fmt.Println("VWS7: Letter Index No 2:", letrIdx2, "; Error:", eror1)*/
							/* Now check that the Letter index is valid */
							if (letrIdx2 > 0) && (letrIdx2 <= inpStrLen) && (letrIdx1 < letrIdx2) {
								/*No error so set Global var: Selected Letter Index No 2*/
								selLeterIdx2 = letrIdx2
							} else {
								retVal = 6
								fmt.Println("Error: Single Letter Upper Bound:", letrIdx2, "is greater than length of Input String OR is equal to Lower Bound.")
							} /*if*/
						} else {
							log.Fatalf("Error: Select a Letter Range, second number is not numeric; %s", eror1)
						} /*if*/
					} else {
						retVal = 5
						fmt.Println("Error: Single Letter Lower Bound:", letrIdx1, "is greater than length of Input String OR is less than 1.")
					} /*if*/
				} else {
					log.Fatalf("Error: Select a Letter Range, first number is not numeric; %s", eror1)
				} /*if*/
			} else {
				/*Error*/
				retVal = 7
				fmt.Println("Error: Select a Letters Range should be in parenthesis (n1:n2).")
			} /*if*/
		} /*if*/
	} /*if there is no ; */
	selSingleLeter = singleLetr
	return (retVal)
} /*ValidatWrdSel*/

/*Function Draw Graphic → Ascii-art is a program which consists of receiving a string as an argument
and outputting the string in a graphic representation using ASCII.
Colour Palette: colourP */
func DrawGraphic() byte {
	var err error
	var file *os.File
	var idx1 int = 0
	var idx2 int = 0
	var idx3 int = 0
	var idx4 int = 0
	var idx5 int = 0
	var scanner *bufio.Scanner
	var line []string
	var strArt []string
	var retVal byte = 0 /*Return value: Error code; 0 for success*/

	/*fmt.Println("DG1: Draw Graphic Entered; Colour:", colour, "; Blink Switch:", blinkSwitch)*/
	/*Usage: --colour=<Black | Red | Green | Yellow | Blue | Magenta | Cyan | White>;<Blink switch>*/
	switch colour {
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
		if blinkSwitch {
			fmt.Println("\033[37;1;5m")
		} else {
			fmt.Println("\033[37;1;25m")
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
					/*fmt.Println("DG2: Idx3:", idx3, "; Idx4:", idx4, "; Idx5:", idx5)*/
					fmt.Printf(bannerMap[int(strArt[idx3][idx5])][idx4])
				}
				fmt.Printf("\n")
			}
		}
	} /*for loop*/
	return (retVal)
} /*DrawGraphic*/

/*Function Main - Entry point*/
func main() {
	var noOfArgs int = 0
	var erorCod byte = 0
	var poundSign bool = false
	var arguments []string
	var colourP string    /*Colour Palette*/
	var wordOrLetr string /*Select a single Letter or a word (Range of Letters)*/

	noOfArgs = len(os.Args)
	arguments = os.Args
	bytRGBslc = []byte{0, 0, 0} /*Initialise the Global var: RGB Byte slice*/

	/*fmt.Println("M1: Main Entered; Number of Arguments:", noOfArgs, "; Arguments:", arguments)*/
	switch noOfArgs {
	case 0:
		fmt.Println("Error: No Arguments.")
		return
	case 1:
		fmt.Println("Error: No Arguments.")
		return
	case 2:
		fmt.Println("Error: Only One Argument\nUsage: go run . [STRING] [OPTION]\nEX: go run . something --colour=<colour>")
		return
	case 3:
		/*Correct Number of Arguments, so we can continue.*/
		/*Two Arguments allowed: [STRING] [OPTION] */
		/*fmt.Println("Error: There is one and only one argument allowed - a", "\"string inside quotes\"")*/
		if arguments[1] == "" {
			fmt.Println("Error: First Argument is empty.")
			return
		} /*if*/
		if arguments[1] == "\\n" {
			fmt.Println("Error: Newline in First Argument (Input String).")
			return
		} /*if*/
		poundSign = strings.Contains(arguments[1], "£")
		if poundSign {
			fmt.Println("Error: Pound sign (£) in Input String.")
			return
		} /*if £ sign in input string*/
		colourP = arguments[2]
		erorCod = ValidatColPlt(colourP)
		if erorCod == 0 {
			/*Error code = 0, means success*/
			erorCod = DrawGraphic()
		} else if erorCod == 2 {
			/*the RGB Hex Notation is being used*/
			erorCod = DrawGraphic4()
		} /*if error*/
	case 4:
		/*Correct Number of Arguments, so we can continue.*/
		/*Three Arguments allowed: [STRING] [OPTION] [Select Word or Letter option]*/
		wordOrLetr = arguments[3]
		/*fmt.Println("M2: Three Arguments are allowed; Third argument:", wordOrLetr)*/
		if arguments[1] == "" {
			fmt.Println("First Argument is empty.")
			return
		} /*if*/
		if arguments[1] == "\\n" {
			fmt.Println("Newline in First Argument.")
			return
		} /*if*/
		colourP = arguments[2]
		erorCod = ValidatColPlt(colourP)
		if erorCod == 0 {
			/*Error code = 0, means success*/
			erorCod = ValidatWrdSel(wordOrLetr)
			if erorCod == 0 {
				/*Error code = 0, means success*/
				if selSingleLeter {
					erorCod = DrawGraphic2()
				} else {
					erorCod = DrawGraphic3()
				} /*if*/
			} else {
			} /*if*/
		} else {
		} /*if error*/
		return
	default:
		/*Too many arguments*/
		fmt.Println("Error: Too many Arguments.")
		return
	} /*switch*/
} /*main*/
