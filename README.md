# Bill Management

## What we learnt

## Structs and Custom Types

- Used in place of classes as we don't have classes in GO

```go
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

```

## Receiver Function (Methods)

- There are functions which can be only used with predefined types

```go
//Receiver function to format the bill

func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

    //adding the tip
    fs += fmt.Sprintf("%-25v ...$%0.2f", "tip:", b.tip)

	//total
    fs += fmt.Sprintf("\n%-25v ...$%0.2f", "total:", total+b.tip)

    return fs
}

// Update the tip
func (b *bill) updateTip(tip float64){
    b.tip = tip
}

// Adding item to the map
func (b *bill) addItem(name string, price float64){
    b.items[name] = price
}
```

> Note:
>
> By only changing the type of receiver function to pointer, GO automatically handles all other things inside functions
> We don't have to make variables inside the functions as pointer too.

## main.go file for the files above

```go
func main() {
    mybill := newBill("trishan's bill")

    mybill.addItem("samosa", 1.20)
    mybill.addItem("momo", 5.20)
    mybill.addItem("pizza", 24.2)
    mybill.addItem("coke", 4.40)

    mybill.updateTip(20)

    fmt.Println(mybill.format())

}
```

## User Input

- Taking user input from terminal instead of hardcoding

- `bufio` package lets us use io and `os` package lets us use terminal as input

```go
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Create a new bill name:")
    name, _ := reader.ReadString('\n') // Reads from the console

    name = strings.TrimSpace(name) // triming the space around the string

    fmt.Println(name)

```

## Parsing String to float64

- We have to use package called `strconv` (string convert) and use method on that called `parseFloat`
- that method gives two values that is 'value' and 'error'

```go
parsedFloat,err := strconv.ParseFloat(stringInput,64)
```

- the method above gives output in `parsedFloat` (float64 type) variable and if error occurs then it saves that in `err` variable

## Saving result to .txt file

- we use `os` package for this in which we use `WriteFile` method which takes three arguments
- path to save file, file contents and permission

```go
    data := []byte(b.format())
    err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
    if err != nil{
        panic(err)
    }
```
