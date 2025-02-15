
<template>
    <div v-if="display_list">
        <ul>
            <li v-for="show in show_list_array" :key="show" @click="handleClick(show)">
<!--                {{ show.split('/').pop() }}-->
<!--                <br>-->
<!--                {{show}}-->
<!--                <router-link to="/videoplay/shows/ + show">-->
                <router-link :to="'/videoplay/shows/' + show">
                    {{ show.split('/').pop() }}
<!--                    <img style="height: 200px" src="/src/assets/video_posters/the day after tomorrow.jpg" ></img>-->
                </router-link>

            </li>
        </ul>
    </div>
    <div v-else>
        <div>
        <video width="640" height="480" controls id="video">
            <source :src="ShowPath" type="video/mp4">
            Your browser does not support the video tag.
        </video>
        </div>
    </div>

</template>

<script lang="ts" setup>
import {ShowGetItems, ShowGetList} from "/apiv4";
import {useRouter} from "vue-router";
import {ref} from "vue";
import {routerKey} from "vue-router";

// defineProps(['show_name']);
const props = defineProps(["show_name"]);
let router = useRouter();

let display_list = ref(true)
let ShowPath = ref("");
let ShowSubtitlePath = ref("");

let show_list_array = ref("")
// parse the url first; get the show name
// define prons show name



// ShowGetList("Fullmetal Alchemist - Brotherhood").then((response) => {
ShowGetList(props.show_name).then((response) => {
    // if failed; go to login page
    console.log("response: ", response.status)
    if (response.shows.length == 0) {
        // alert("failed to get show list, login is required");
        router.push("/login");
        return;
    }
    // put response to show list
    show_list_array.value = response.shows;
    // process show_list_array only reserved the mp4
    show_list_array.value = show_list_array.value.filter((item) => {
        return item.includes("mp4");
    });
    // for (let i = 0; i < show_list_array.length; i++) {
    //   show_list.innerHTML += "<a href='/shows/"+show_list_array[i]+"'>"+show_list_array[i]+"</a><br>";
    // }
})
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
const handleClick = (show) => {
    console.log("get show: " + show)
    ShowGetItems(show).then((response) => { console.log(response.show_url )
        ShowPath.value = response.show_url;
        display_list.value = false;
    })
    console.log("get subtitle")
    let show_subtitle = show.replace("mp4", "zh.srt")
    console.log(show_subtitle)
    ShowGetItems(show_subtitle).then((response) => {
        ShowSubtitlePath.value = response.show_url;
        convertSRTtoVTT(ShowSubtitlePath.value)
    })
    // let's convert srt subtitle to vtt
    // get the srt file from ShowSubtitlePath
    // convert it to vtt
    // let srt = ShowSubtitlePath.value


}

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