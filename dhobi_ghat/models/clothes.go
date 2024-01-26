package models 
import "fmt"

type cloth struct{
    name      string  `json:"name"`
    id        int      `json:"id"`
    clothType string `json:"cloth_type"`
    averageWeight float64 `json:"average_weight"`
    price float64 `json:"price"`
    expressTag bool `json:"express_tag"`
}
