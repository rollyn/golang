package conversion

import ("strconv"
)
 
func DecimalToBinary(decimal int64) string {
    binary := strconv.FormatInt(decimal, 2)
    return binary
}