import axios from "axios";
import { store } from "../redux/store";

const api = axios.create({
    baseURL: "http://localhost:3000/api"
})

api.interceptors.request.use(
    config => {
        const val = store.getState()
        if(store && val.user.accessToken != '') {
            config.headers!['Authorization'] = `Bearer ${val.user.accessToken}`;
        }

        return config
    },
    error => {
        return Promise.reject(error);
    }
)

export default api