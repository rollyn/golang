package libraries

import (
  "fmt"
  "../domain"
)
func World() string {
 p := &domain.Product{ProductID: 333}  
 return fmt.Sprintf(" world Id:%d", p.ProductID)
}