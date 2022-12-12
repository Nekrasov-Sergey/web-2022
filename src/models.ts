export interface IStore {
    UUID: string
    Name: string
    Discount: number
    Price: number
    Quantity: number
    Promo: string[]
    Image: string
}

export interface ICart {
    StoreUUID: string
    Quantity: number
}

export interface IOrder {
    UUID: string
    Store: string
    Quantity: number
    UserUUID: string
    Date: string
    Status: string
}