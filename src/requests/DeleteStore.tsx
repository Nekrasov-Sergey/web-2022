import {deleteStore} from "../modules";


export function DeleteStore(uuid: string) {

    const url = `store`

    function Delete() {
        deleteStore(url, uuid)
    }


    return (
        <form>
            <button onClick={() => Delete()}>Удалить магазин</button>
        </form>
    );

}