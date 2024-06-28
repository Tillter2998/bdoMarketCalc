package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

func main() {
	salePrice := flag.String("salePrice", "", "Specify the sale price to calculate for")

	flag.Parse()

	salePriceNoComma := strings.ReplaceAll(*salePrice, ",", "")
	saleInt, err := strconv.ParseFloat(salePriceNoComma, 64)
	if err != nil {
		fmt.Println(err)
	}

	baseAmount := saleInt * (0.65 * (1 + 0.005))
	valuePack := saleInt * (0.65 * (1 + 0.3 + 0.005))
	merchantsRing := saleInt * (0.65 * (1 + 0.05 + 0.005))
	vpRmr := saleInt * (0.65 * (1 + 0.3 + 0.05 + 0.005))

	outputTemplate := `Collected Silver:
    Base Amount:              %v
    With Value Pack:          %v
    With Rich Merchants Ring: %v
    With VP and RMR:          %v`

	fmt.Printf(outputTemplate, insertCommas(baseAmount), insertCommas(valuePack), insertCommas(merchantsRing), insertCommas(vpRmr))
}

func insertCommas(f float64) string {
	integer := int(f)
	tmpl := `{{. | comma}}`

	funcMap := template.FuncMap{
		"comma": func(v int) string {
			stringVal := fmt.Sprintf("%d", v)
			n := len(stringVal)
			if n <= 3 {
				return stringVal
			}

			var result string
			for i, c := range stringVal {
				if (n-i)%3 == 0 && i != 0 {
					result += ","
				}
				result += string(c)
			}
			return result
		},
	}

	t := template.Must(template.New("test").Funcs(funcMap).Parse(tmpl))
	var buf bytes.Buffer
	err := t.Execute(&buf, integer)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return buf.String()
}
