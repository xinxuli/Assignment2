import { DefaultApi,Configuration } from "../api";

export let Api = new DefaultApi(new Configuration({
    basePath: "http://localhost:8080"
})) 

export function initApi(apiKey='', basepath = "http://localhost:8080") {
    Api = new DefaultApi(new Configuration({
        basePath: basepath,
        accessToken: apiKey
    }))
}
// 1. 重新生成个新的对象

export const getApi = () => Api;