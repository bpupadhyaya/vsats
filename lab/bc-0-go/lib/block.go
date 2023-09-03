package block_pkg
import(
"fmt"
)

func CreateBlock(Header, Body string) {
    fmt.Println(Header, "\n", Body)
}