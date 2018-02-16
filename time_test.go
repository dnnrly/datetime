package datetime

import (
	"fmt"
	"time"
)

func Example_fomattingTime() {
	t := time.Date(2018, 1, 5, 20, 43, 22, 123456789, time.Local)
	fmt.Println(t.Format("2006/01/02 15:04:05.999 -07"))
	// Output: 2018/01/05 20:43:22.123 +00
}
