import React, { useEffect,useState } from "react"
import VideoService from "../services/Video";
import ReactHlsPlayer from "react-hls-player";

export default function Home() {
    const [videos, setVideos] = useState([])
    useEffect(() => {
        VideoService().getAll().then(r => r.json()).then(data => {
            setVideos(data)
        })
    }, [])

    return (
        <div>
            <h1>Videos</h1>
            {videos.map(item => (
                <li>
                    <h2>{item.ID}</h2>
                    <ReactHlsPlayer
                        url={"https://stream.mux.com/"+item.PlaybackIDS[0].ID+".m3u8"}
                        autoplay={false}
                        controls={true}
                        width={500}
                        height={375}
                    />
                </li>
            ))}
        </div>
    )
}