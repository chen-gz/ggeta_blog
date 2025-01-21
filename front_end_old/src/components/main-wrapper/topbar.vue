<script setup lang="ts">
/// dynamic update based on the router path

//watch for the router path change
// import {logined} from "/apiv4.js";
import {logined, newPostV4} from "/apiv4.js";
import {useRouter} from "vue-router";
import {watch, ref} from "vue";

const router = useRouter()

let isLogin = ref(true)
let breadcrumb = ref([]) // name and link
watch(router.currentRoute, (to, from) => {
    // check the loging status
    isLogin.value = logined();
    // Reset breadcrumb and always start with Home
    breadcrumb.value = [{ name: 'Home', link: '/' }];

    // Extract parts of the path to build the breadcrumb
    const pathSegments = to.path.split('/').filter(path => path !== '');

    // get the character back such as "%20"
    pathSegments.forEach((segment, index) => {
        pathSegments[index] = decodeURIComponent(segment);
        // pathSegments[index] = pathSegments[index].charAt(0).toUpperCase() + pathSegments[index].slice(1);
    });
    // if the router is /post_edit/xxx, then the breadcrumb should be /post/xxx and post will post to the page

    let pathSoFar = '';
    pathSegments.forEach(segment => {
        pathSoFar += `/${segment}`;
        //capitalize the first letter for name
        let name = segment.charAt(0).toUpperCase() + segment.slice(1);
        let link  = pathSoFar;
        if (name == "Post_edit") {
            name = "Post"
            // last part of the path is the post id
            let post_id = pathSegments[pathSegments.length - 1]
            link = pathSoFar.replace("post_edit", "post") + `/${post_id}`
        }
        else if (name == "search") {
            // stop the for each loop
            return
        }
        breadcrumb.value.push({ name: name, link: link });
    });})
function new_post() {
    console.log("new post")
    // create a new post and redirect to the new post
    newPostV4().then((response) => {
        console.log(response)
        router.push(`/post_edit/${response.url}`)
    })
}

let search_message = ref("")
function onSearch() {
    console.log("search", search_message.value)
    router.push(`/search/${search_message.value}`)
}

</script>

<template>
    <div id="topbar-wrapper-inner" class="d-flex flex-row">
        <nav id="breadcrumb" aria-label="Breadcrumb">
            <span v-for="(item, index) in breadcrumb" :key="index" style="justify-content: center">
                <a v-if="index < breadcrumb.length - 1" :href="item.link">{{ item.name }}</a>
                <span v-else>{{ item.name }}</span>
            </span>
        </nav>
        <span class="flex-grow-1"></span>
        <input v-model="search_message" v-on:keyup.enter="onSearch"> </input>

        <button v-if="!isLogin" @click="$router.push('/login')">
            <i class="fa-fw fas fa-user" style="font-size: 1.3em;"></i>
				 <!-- fa-user"></i> -->
        </button>
        <button v-else @click="new_post">
            <i class="fa fa-plus"></i>
        </button>
    </div>

</template>

<style scoped lang="sass">
#breadcrumb
    height: 1.5rem
    font-size: 1rem
    line-height: 1.5rem
    align-content: center
    align-self: center
    font-weight: 400
    //letter-spacing: -0.5px
    span
        color: rgb(117, 117, 117)

    a
        text-decoration: none
        color: rgb(0, 86, 178)

button
    margin-left: 1rem

#breadcrumb span:not(:last-child)::after
    content: "›"
    padding-left: 4.8px
    padding-right: 4.8px

#topbar-wrapper-inner
    height: 100%
    align-items: center

input
    min-width: 100px
    width: 200px
    height: 1.6rem
    align-self: center
    border-radius: 1rem
    border: 1px solid #0d6efd
    padding-left: 1rem
//padding-right: 100px
//margin-right: px
//margin-right: 0.7rem

button
    width: 2rem
    height: 2rem
    //display: flex
    //margin-right: 13rem
    text-align: center
    //justify-content: center
    //align-items: center
    align-content: center
    border-width: 0
    background-color: white

    i
        //align-self: center
        font-size: 1.5rem


</style>
