
<template>

<!--    <s-card style="max-width: none; height: 600px">-->
    <a href="/shows/Fullmetal Alchemist - Brotherhood">
        <img style="height: 600px" src="/src/assets/fmc.jpg" ></img>
    </a>
    <router-link to="/videoplay/" props:video_path="Fullmetal Alchemist - Brotherhood">
            <h1>Fullmetal Alchemist - Brotherhood</h1>
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