
<template>

<!--    <s-card style="max-width: none; height: 600px">-->
    <a href="/shows/Fullmetal Alchemist - Brotherhood">
        <img style="height: 300px" src="/src/assets/fmc.jpg" ></img>
    </a>
    <router-link to="/videoplay/movie/The Day After Tomorrow (2004)/The Day After Tomorrow (2004) Bluray-1080p AAC.mkv">
        <img style="height: 300px" src="/src/assets/video_posters/the day after tomorrow.jpg" ></img>
    </router-link>
<!--    </s-card>-->
</template>

<script lang="ts" setup>
import {ShowGetItems, ShowGetList} from "/apiv4";
import {useRouter} from "vue-router";
import {ref} from "vue";
let router = useRouter();


</script>


<style scoped>
</style>