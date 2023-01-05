**DATA TYPES**
- elementary (int, float, bool, string) 
- structured ( struct, array , slice, map, channel) 
- interfaces (they describe the behaviour of the type) 


**Character type**
there is no type called character type in go . the characters are a special case of integers. The byte type is an alias for uint8  and this is okay for the traditional ASCII-encoding 

**Strings** 
they are sequence of utf-8 characters 
in go strings are variable length from 1-4 bytes whereas in languages like java they are always fixed size 2 bytes 

the usual comparison operators work in string by using(==, != , <= >= ) by comparing byte by byte in memory 
len(str) 

***strings and strconv package**
- **strings.HasPrefix(s, prefix string) bool** 
- **strings.HasSuffix(s, suffix string) bool** 
- **strings.Contains(s, str string) int** 
- **strings.LastIndex(s, str string) int**

strings.Replace 
strings.Count
strings.Repeat
strings.ToLower
strings.ToUpper
strings.TrimSpace
strings.

strings.Fields(s) - splits the string by whitespace 


Reading from a string 
- strings.NewReader(str) -> this produces an pointer to a reader value that provides amongst others the following functions 
- Read() to read a []byte 
- ReadByte() to read the next byte from the string
- ReadRune() to read the next rune from the string 


for conversions to and from a string use the strconv package 
strconv.IntSize
strconv.Itoa(i int) string 
strconv.FormatFloat(f float64, fmt byte, prec int , bitSize int) strin g

strconv.Atoi(s string) (i int , error) 
strconv.ParseFloat(s string, bitsize Int)(f float64, err error) 

### times and dates 
- time package
- time.Now() and then parts can be obtained by doing t.Day(), t.Hour()

### fallthrough keyword with case 


switch var0 {
	case 0,2,4: 
		fallthrough 
	case 1: 
}

in the above case if the case 0 does not match it will execute al the remaining case branches until a branch without fallthrough is encountered 


switch with initialization 


switch a, b:=0,1 {

}


for x, pos := range str2 {
	println(x) 
	println(pos)
}

**function overrloadding is not allowed in go**

type binOp func(int, int) int 
** all primitive values int, float, boolean , string and its variants and array are passed by value 
only referenced types like slices, maps, channels are passed by reference 
