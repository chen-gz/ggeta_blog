<script lang="ts" setup>
import {ref, watch} from "vue";
import {useRouter} from "vue-router";
import {formatDate, SearchPostsRequestV4, SearchPostsResponseV4, searchPostsV4} from "/apiv4";

// let searchParam = SearchPostsRequestV4
const route = useRouter();
// url is tags/:id
// url can be "tag/:id" or "cate/:id" or "post/
// parse the url first
/// watch the url change
let res = ref({} as SearchPostsResponseV4)

function page_init() {
    let tag_cate_post = route.currentRoute.value.path.split("/")[1]
    let url_id = route.currentRoute.value.path.split("/")[2]
    console.log("call list page" + tag_cate_post + " " + url_id)

    let searchParam = {} as SearchPostsRequestV4
    if (tag_cate_post == "tag") {
        searchParam.tags = url_id
        searchParam.sort = "created_at DESC"
    } else if (tag_cate_post == "cate") {
        searchParam.categories = url_id
        searchParam.sort = "created_at DESC"
    } else if (tag_cate_post == "search") {
        searchParam.content = url_id
        searchParam.sort = "created_at DESC"
    } else {
        searchParam.sort = "created_at DESC"
    }

    let params = searchParam
    console.log(params)
    searchPostsV4(params).then((response) => {
        res.value = response
        console.log(res.value)
        if (res.value.number_of_posts == 0) {
            res.value.posts = []
        }
        for (let i = 0; i < res.value.posts.length; i++) {
            res.value.posts[i].created_at = new Date(res.value.posts[i].created_at);
        }
        console.log(res.value)
    })
}

page_init();

watch(route.currentRoute, (to, from) => {
    page_init();
})

</script>

<template>
    <h1/>
    <ul>
        <li v-for="result in res.posts" :key="result.id">
            <a href="#" @click="route.push('/post/' + result.url)">
                {{ result.title }}
            </a>
            <span class="dash">
<!--            add dash to fill the space -->
             </span>

            <time>
                {{ formatDate(result.updated_at) }}
            </time>
        </li>
    </ul>
</template>

<style lang="sass" scoped>
.dash
    margin-left: 5px
    margin-bottom: 10px
    margin-right: 10px
    flex-grow: 1
    border-bottom: 2px dotted currentColor
// ; /* Use a dotted border bottom */

li
    display: inline-flex
    width: 100%
    margin-bottom: 20px

    a
        font-size: 20px
        color: #000
        text-decoration: none

        &:hover
            color: #000
            text-decoration: underline

h1
    margin-top: 50px
    margin-bottom: 20px
</style>