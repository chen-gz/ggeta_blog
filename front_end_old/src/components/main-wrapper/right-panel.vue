<script lang="ts" setup>
import {getPostV4,searchPostsV4, SearchPostsRequestV4, GetFileList} from "../../../apiv4.js";
import {ref, watch} from "vue";
import {useRouter} from "vue-router";

let param = {} as SearchPostsRequestV4
param.limit = {start: 0, size: 5}
param.sort = "updated_at DESC"
let posts = ref([])
searchPostsV4(param).then((res) => {
    posts.value = res.posts
})
let editor = ref(false)
// watch url

let router = useRouter()
let path = router.currentRoute.value.path
let file_list = ref([])
watch(useRouter().currentRoute, (to, from) => {
    // if post_edit in url
    if (to.path.includes("post_edit")) {
        editor.value = true
    } else {
        editor.value = false
    }
    let url = to.path.split("/")
    console.log(url)
    let url2 = url[url.length - 1] as string
    // convert url2 to string from html encoded
    url2 = decodeURIComponent(url2)
    console.log(url2)

    // get post
    let current_post = ref({})
    getPostV4(url2, false).then((res) => {
        current_post.value = res.post
        console.log("current post", current_post.value)

        console.log("to path", to.path)
        if (editor.value) {
            // get files list from the server
            console.warn("todo: get files list")
            GetFileList(current_post.value.id).then((res) => {
                file_list.value = res

                // console.log(res)
                // console.log("files list", res)
            })
        }

    })

})

</script>

<template>
    <div id="right-panel-wrapper-inner">
        <div id="recently-updated">
            <h4>Recently Updated</h4>
            <ul>
                <li v-for
                        ="p in posts" :key="p.id">
                    <a :href="'/post/' + p.url">{{ p.title }}</a>
                </li>
            </ul>
        </div>
        <div v-if="editor">
            <h4>Files</h4>
            <ul v-if="file_list.length > 0">
                <li v-for="f in file_list" :key="f">
                    <a :href="'/file/' + f.id">{{ f}}</a>
                </li>
            </ul>
        </div>
        <div id="toc-wrapper">
            <h4 style="margin-bottom=3px;"> Contents</h4>
            <div id="toc"></div>
        </div>
    </div>
</template>

<style lang="sass" scoped>
#right-panel-wrapper-inner
    display: flex
    flex-grow: 0
    flex-shrink: 0
    padding-top: 48px
    // padding-left: 15px
    flex-direction: column
    // add vertical spacing between each items
    > div
        margin-bottom: 20px

#recently-updated
    margin-bottom: 20px
    border-left: 2px solid #f1f1f1
    padding-left: 10px
    ul
        margin:  0 0 10px 0
        //margin-bottom: 10px
        font-size: 13.6px
        line-height: 17px
        list-style: none
        li
            margin-bottom: 6.4px
            margin-top: 6.4px
            text-align: left
            box-sizing: content-box
        a
            padding-left: 17px
            padding-bottom: 3.2px
            padding-top: 3.2px


#toc-wrapper
    padding-left: 10px


</style>
