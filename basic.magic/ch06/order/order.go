/**
 * 结构体：组合
 * 包内所有成员字段和方法都是可见的
 * 包外首字母大写的成员字段和方法才是可见的
 */
package order

type User struct{
    Uid     int64
    Name    string
    Phone   int64
}

type Order struct {
    OrderSn     int64
    productID   int64   // 私有属性 包外不可见
}
/**-------------------------------------------
 *  User
 **-------------------------------------------*/
func (u *User) GetName() string {
    return u.Name
}
func (u *User) SetName(name string) {
    u.Name = name
}
// 私有方法，包外不可见
func (u *User) save() error{
    return nil
}
/**-------------------------------------------
 *  Order
 **-------------------------------------------*/
func (o *Order)GetOrderSn() int64 {
    return o.OrderSn
}
func (o *Order)SetOrderSn(sn int64) {
    o.OrderSn = sn
}
func (o *Order)GetProductID() int64 {
    return o.productID
}
func (o *Order)SetProductID(pid int64) error {
    o.productID = pid
    return nil
}
