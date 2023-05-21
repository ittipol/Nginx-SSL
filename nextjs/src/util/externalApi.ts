import axios from "axios";

const externalApi = axios.create({
    baseURL: process.env["API_URL"]
})

externalApi.interceptors.request.use(
    config => {

        config.headers!['Content-Type'] = "application/json"
        config.headers!['Origin'] = "http://abc:3000"

        return config
    },
    error => {
        return Promise.reject(error);
    }
)

export default externalApi