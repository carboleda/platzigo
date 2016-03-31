/** (Reto)
* Range: Es una funcion que itera sobre las estrcturas de datos como arrays, slice y maps
*/

package main

import "fmt"

func main() {
	numeros := []int{2,4,6,3};
	suma := 0;

	for _, num := range numeros {
		suma += num;
	}

	fmt.Printf("Suma: %d\n", suma);

	for i, numero := range numeros {
		if numero == 3 {
			fmt.Println("Index: ", i);
		}
	}

	algo := map[string]string{"a": "auto", "b": "bebe"};
	for key, value := range algo {
		fmt.Printf("Key: %s, Value: %s\n", key, value);
	}

}