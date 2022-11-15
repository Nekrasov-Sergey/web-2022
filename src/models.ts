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
    Store: string
    Quantity: number
}