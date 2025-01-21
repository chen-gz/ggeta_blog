<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { getDistinct, SearchPostsRequestV4, GetDistinctResponse } from "/apiv4";
// import Lists from "@/views/Lists.vue";
// let props = defineProps<{
//     tag_name: String;
// }>();
let tags = ref({} as GetDistinctResponse);
let searchParam = {} as SearchPostsRequestV4;
// watch(
//     () => props.tag_name,
//     (old, newe) => {
//         console.log("props.tag_name changed");
//         searchParam.tags = props.tag_name as string;
//     }
// );
onMounted(() => {
    getDistinct("tags").then((response) => {
        tags.value = response;
        // remove empty tag
        tags.value.values = tags.value.values.filter((item) => {
            return item != "";
        });
        if (tags.value.values.length > 0) {
            tags.value.values.sort();
        }
    });
});
</script>

<template>
    <div class="tags-page">
        <h1>Tags</h1>
        <div id="tags" class="tags-container">
            <a v-for="tag in tags.values" :key="tag" :href="'/tag/' + tag" class="tag">
                {{ tag }}
            </a>
        </div>
    </div>
</template>

<style scoped>
.tags-page {
    padding: 1rem;
}

h1 {
    font-size: 2rem;
    margin-bottom: 1rem;
    color: #333;
}

.tags-container {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
}

.tag {
    background-color: #007bff;
    color: white;
    padding: 5px 10px;
    border-radius: 15px;
    font-size: 1rem;
    cursor: pointer;
    text-decoration: none;
    transition: background-color 0.3s;
}

.tag:hover {
    background-color: #0056b3;
}
</style>
