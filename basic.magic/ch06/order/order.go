package order

type User struct{
    Uid     int64
    Name    string
    Phone   int64
}

type Order struct {
    OrderSn     int64
    productID   int64
}

type Company struct {
    Cid   int
    Work  string
    CName string
}

/**-------------------------------------------
 *  Line
 **-------------------------------------------*/
func (o *Order)GetOrderSn() int64 {
    return o.OrderSn
}
func (o *Order)GetProductID() int64 {
    return o.productID
}
func (o *Order)SetProductID(pid int64) error {
    o.productID = pid
    return nil
}
/**-------------------------------------------
 *  Circle
 **-------------------------------------------*/
func (c *Company) GetCName() string{
    return c.CName
}
func (c *Company) SetCName(name string){
    c.CName = name
}