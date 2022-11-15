import {ENDPOINT} from "./App";
import axios from "axios";

export function getFromBackend(url: string) {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}




