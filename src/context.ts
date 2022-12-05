import {ICart, IStore} from "./models";

export let StoreContext: IStore = {
    UUID: "",
    Name: "",
    Discount: 0,
    Price: 0,
    Quantity: 0,
    Promo: [],
    Image: ""
}

export let CartContext: ICart = {
    StoreUUID: "",
    Quantity: 0
}