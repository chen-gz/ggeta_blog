
<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import {getDistinct, SearchPostsRequestV4, GetDistinctResponse} from "/apiv4"
import MainWrapper from "@/App.vue";
// import Lists from "@/views/Lists.vue";
let props = defineProps<{
    tag_name: String
}>();
let tags = ref({} as GetDistinctResponse)
let
    searchParam = {} as SearchPostsRequestV4
watch(() => props.tag_name, (old, newe) => {
    console.log("props.tag_name changed")
    searchParam.tags = props.tag_name as string
})
onMounted(() => {
    getDistinct("tags").then((response) => {
        tags.value = response
        // remove empty tag
        tags.value.values = tags.value.values.filter((item) => {
            return item != ""
        })
        if (tags.value.values.length > 0) {
            tags.value.values.sort()
        }
    });
})

</script>


<template>
    <h1>Category </h1>
    <div>
        <div id="tags" class="d-flex flex-wrap">
            <div v-for="(item, index) in tags.values" :key="index">
                <div class="item">
                    <a :href="'/tag/' + item">
                        {{ item }}
                        <span> 1 </span>
                    </a>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
h1 {
    margin-top: 50px;
    margin-bottom: 20px;
}
.item {
    margin: 5px;
    padding: 5px;
    border: 1px solid #ccc;
    border-radius: 5px;
    text-decoration: none;
    color: #000;

    span {
        margin-left: 5px;
        color: #ccc;
    }

    height: 2rem;
    align-items: center;
    align-content: center;

    a {
        display: flex;
        align-items: center;
        align-content: center;
        height: 100%;
        text-decoration: none;
        &:link, &:visited, &:hover, &:active {
            text-decoration: none;
            color: #286fff;
        }
    }
}

#tags {
    flex-wrap: wrap;
    //a {
    //    //margin: 5px;
    //    //padding: 5px;
    //    border: 1px solid #ccc;
    //    border-radius: 5px;
    //    text-decoration: none;
    //    color: #000;
    //    span {
    //        margin-left: 5px;
    //        color: #ccc;
    //    }
    //    height: 80px;
    //}
    //height: 500px;
    //align-items: center;
}
</style>
