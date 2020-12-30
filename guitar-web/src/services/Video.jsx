import {GetHeadersProperties, METHODS} from "./getHeadersProperties";

export default function VideoService() {

    return {
        getAll() {
            return fetch("http://localhost:8080" + '/api/v0/videos', GetHeadersProperties(METHODS.GET));
        },
        upload(video) {
            return fetch("http://localhost:8080" + '/api/v0/upload',  GetHeadersProperties(METHODS.POST));
        }

    }
}