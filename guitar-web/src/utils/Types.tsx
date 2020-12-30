


export interface IAsset {
    test:   boolean;
    status:   string;
    playbackIDS:   IPlaybackIDS[];
    mp4support: string;
    masterAccess:  string;
    id:   string;
    created_at : string
}

export interface IPlaybackIDS {
    policy: string;
    playbackID : string;
}