
<template>
    <div>
        <video width="640" height="480" controls id="video">
            <source :src="VideoPath" type="video/mp4">
            Your browser does not support the video tag.
        </video>
    </div>
</template>

<script lang="ts" setup>
import {GetPresignedUrl, ShowGetItems} from "/apiv4";
import {nextTick, ref, watch} from "vue";
const props = defineProps(["path"]);
let VideoPath = ref("");

// let bucket =  "shows";
// let path = "Fullmetal Alchemist - Brotherhood/Season 1/Fullmetal Alchemist - Brotherhood-S01E01-Fullmetal AlchemistHDTV-1080p.mp4"
console.log("props.path", props.path)
// console.log(props.path)
let bucket = props.path[0];
// path is from 1 to end
let path = props.path.slice(1).join("/");

function getVideoPath() {
    // GetPresignedUrl(props.bucket, props.path).then((response) => {
    GetPresignedUrl(bucket, path).then((response) => {
        VideoPath.value = response.presigned_url;
        console.log(VideoPath.value)
    })
}
getVideoPath();
let subtitle = path.replace("mp4", "zh.srt")
console.log(subtitle)
let SubtitlePath = ref("");
GetPresignedUrl(bucket, subtitle).then((response) => {
    SubtitlePath.value = response.show_url;
    convertSRTtoVTT(SubtitlePath.value)
})


watch(VideoPath, (newPath) => {
    let videoElement = ref(document.getElementById("video") as HTMLVideoElement);
    console.log("videoelemtn.value", videoElement.value)
    if (videoElement.value) {
        videoElement.value.load();
    }
});
watch(SubtitlePath, (newPath) => {
    let videoElement = ref(document.getElementById("video") as HTMLVideoElement);
    console.log("videoelemtn.value", videoElement.value)
    if (videoElement.value) {
        videoElement.value.load();
    }
});


// let SubtitlePath = ref("");

async function convertSRTtoVTT(srtUrl) {
    try {
        // Fetch the SRT file from the URL
        let response = await fetch(srtUrl);
        let srtText = await response.text();

        // Convert SRT to VTT format
        let vttText = "WEBVTT\n\n" + srtText
            .replace(/\r\n|\r|\n/g, "\n") // Normalize line breaks
            .replace(/(\d{2}):(\d{2}):(\d{2}),(\d{3})/g, "$1:$2:$3.$4") // Convert commas to dots in timestamps
            .replace(/\n(\d+)\n/g, "\n\n"); // Ensure double newlines between captions

        // Create a Blob from the converted text
        let blob = new Blob([vttText], { type: "text/vtt" });
        let vttUrl = URL.createObjectURL(blob);

        // Create and append the <track> element
        let track = document.createElement("track");
        track.kind = "subtitles";
        track.label = "English";
        track.srclang = "en";
        track.src = vttUrl;
        track.default = true;

        let video = document.getElementById("video");
        video.appendChild(track);

        console.log("Subtitles loaded successfully!");
    } catch (error) {
        console.error("Error loading subtitles:", error);
    }
}
// const handleClick = (show) => {
//     console.log("get show: " + show)
//     ShowGetItems(show).then((response) => { console.log(response.show_url )
//         ShowPath.value = response.show_url;
//         display_list.value = false;
//     })
//     console.log("get subtitle")
//     let show_subtitle = show.replace("mp4", "zh.srt")
//     console.log(show_subtitle)
//     ShowGetItems(show_subtitle).then((response) => {
//         SubtitlePath.value = response.show_url;
//         convertSRTtoVTT(SubtitlePath.value)
//     })
// }

</script>


<style scoped>
/* General styles */
body {
    font-family: Arial, sans-serif;
    background-color: #f4f4f4;
    color: #333;
    margin: 0;
    padding: 0;
}

/* Container for the video and list */
.container {
    max-width: 800px;
    margin: 20px auto;
    padding: 20px;
    background-color: #fff;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
}

/* List styles */
ul {
    list-style-type: none;
    padding: 0;
}

li {
    padding: 10px;
    margin: 5px 0;
    background-color: #e9e9e9;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;
}

li:hover {
    background-color: #d3d3d3;
}

/* Video styles */
video {
    width: 100%;
    height: auto;
    border-radius: 8px;
    margin-top: 20px;
}

/* Subtitle track styles */
track {
    display: none;
}
</style>