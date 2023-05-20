import axios from "axios";
import { store } from "../redux/store";
import { refreshToken } from "@/redux/features/user/userSlice";

const IGNORED_PATHS: Array<string> = ['token/refresh']
const REQUEST_ATTEMP_TIMES = 1


const api = axios.create({
    // baseURL: "http://testhost:3000/api"
    baseURL: "/api"
})

api.interceptors.request.use(
    config => {
        console.log('request interceptors %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%')
        console.log(config.url)        

        // set auth token
        const val = store.getState()
        if(store && val.user.accessToken != '') {
            config.headers!['Authorization'] = `Bearer ${val.user.accessToken}`
            // console.log(`Set Authorization Bearer ${val.user.accessToken}`)
        }

        return config
    },
    error => {
        return Promise.reject(error);
    }
)

api.interceptors.response.use(
    response => {
        return response
    },
    async err => {    
        const originalConfig = err.config;  
        const accessToken = store.getState().user.accessToken

        // console.log('Response interceptors error ===============================')
        // console.log(`accessToken: ${accessToken}`)
        // console.log(originalConfig._retryAttempt)

        const isIgnored = IGNORED_PATHS.some(path => originalConfig.url.includes(path))

        console.log(`Gen New Token: [${(err.response.status === 401 && !isIgnored && accessToken !== '')}]`)

        if(err.response.status === 401 && !isIgnored && accessToken !== '') {
            
            if(!originalConfig._retryAttempt) {
                originalConfig._retryAttempt = 0                
            }
            
            if(++originalConfig._retryAttempt <= REQUEST_ATTEMP_TIMES) {
                const res = await store.dispatch(refreshToken())

                if (res.meta.requestStatus === "fulfilled") {
                    return api(originalConfig);
                } else {
                    return Promise.reject(err);
                }

            }
        }

        return Promise.reject(err);
    }
)

export default api