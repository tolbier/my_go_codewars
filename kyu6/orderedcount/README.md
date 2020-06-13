# Ordered Count
from https://www.codewars.com/kata/57a6633153ba33189e000074/train/go

Count the number of occurrences of each character and return it as a list of tuples in order of appearance. For empty output return an empty list.

Example:
```  
OrderedCount("abracadabra") == []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}

// Where
type Tuple struct {
    Char  rune
    Count int
}
```  
