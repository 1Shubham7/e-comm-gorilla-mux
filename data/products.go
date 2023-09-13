package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	// stuff encoded inside `` is called struct tags, they have a specific usecase
	ID          int 		`json:"id"` // now the ID is 'id' during the output, you can call it "hello" or anything
	Name        string 		`json:"name" validate:"required"` // this is  a required field now
	Description string 		`json:"description"`
	Price       float32		`json:"price" validate:"gt=0` // greater than 0
	SKU         string 		`json:"sku"`
	CreatedOn   string 		`json"-"` // t	his will not output, we will use it internally
	UpdatedOn   string 		`json"-"`
	DeletedOn   string 		`json"-"`
}

// instead of json.Marshal we want to use this method instead
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error { //name of function is ToJSON and it returns an error if found any issue
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Product) Validate(r io.Reader) error{
	validate := validator.New()
	
	return validate.Struct(p) // this validates the struct you will enter, we are directly returning it, becuase of this we can add struct tags to the Producct struct
}

// abstracting the products by a function
func GetProducts() Products {
	return productList
}

// POST - this function will add the product we get from POST request to our fake database (productList)
func AddProductToDatabase(p *Product){
	p.ID = getIDForNewProduct()
	productList = append(productList, p)
}

func getIDForNewProduct() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

// PUT 

func UpdateProduct(id int, p *Product) error {
	_, position, err := findProductToUpdate(id)
	if err != nil{
		return err
	}
	p.ID = id
	productList[position] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProductToUpdate(id int) (*Product, int, error) { //takes id as input, gives three outputs
	for i , p := range productList {
		if p.ID == id {
			return p, i, nil // return p and nil error
		}
	}

	return nil, -1, ErrProductNotFound
}


// we just created a data structure - Product
// Let's now create a slice of Product

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Chole Bhature",
		Description: "Tasty",
		Price:       80,
		SKU:         "1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Rajma Chawal",
		Description: "Spicy",
		Price:       100,
		SKU:         "10",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "Pastry",
		Description: "Sweet",
		Price:       50,
		SKU:         "100",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

// productList is a slice of pointers to Product structs. 
