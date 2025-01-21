<script lang="ts" setup>
import loader from "@monaco-editor/loader";
import { useRouter } from "vue-router";
import { getPostV4, savePost, UploadFile, V4PostData } from "/apiv4";
import { onMounted, ref } from "vue";

let router = useRouter();
let url = router.currentRoute.value.params.id as string;
console.log(url);

let post = ref({} as V4PostData);

let editor: any = null;
getPostV4(url, true).then((response) => {
    console.log(response);
    post.value = response.post;
    loader.init().then((monaco) => {
        editor = monaco.editor.create(document.getElementById("code_editor"), {
            value: post.value.content,
            language: "markdown",
            theme: "one-light",
            wrappingColumn: 80,
            wordWrap: "on",
            scrollBeyondLastLine: false,
        });
    });
});
let editor_shows = ref("content");
document.addEventListener("keydown", function (e) {
    // control + 'S' to save or (command + 'S' on mac)
    if (
        (window.navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey) &&
        e.key === "s"
    ) {
        e.preventDefault();
        console.log("ctrl+s");
        // save the post
        if (editor) {
            if (editor_shows.value === "meta")
                post.value = JSON.parse(editor.getValue());
            else post.value.content = editor.getValue();
        }
        savePost(post.value).then((response) => {
            // console.log(response)
        });
        // push to new url
        router.push("/post_edit/" + post.value.url);
    }
    // control + 'E' to edit or (command + 'E' on mac)
    // if ((window.navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey) && e.key === 'e') {
    if (e.ctrlKey && e.key === "e") {
        e.preventDefault();
        console.log("ctrl+e");
        // if editor is showing content, switch to meta
        if (editor_shows.value === "content") {
            editor_shows.value = "meta";
            editor.setValue(JSON.stringify(post.value, null, 4));
            // disable word wrap and set language to json
            editor.updateOptions({ wordWrap: "off", language: "json" });
        } else {
            editor_shows.value = "content";
            editor.setValue(post.value.content);
            // editor.setModelLanguage(editor.getModel(), 'markdown')
            // enable word wrap
            editor.updateOptions({ wordWrap: "on", language: "markdown" });
        }
    }
});
// monitor drop event on the editor (id: code_editor) after mount

onMounted(() => {
    document
        .getElementById("code_editor")
        .addEventListener("drop", function (e) {
            e.preventDefault();
            e.stopPropagation();
            // get the file
            let file = e.dataTransfer.files[0];
            console.log(file);
            // upload the file
            UploadFile(file, post.value.id);
        });
});
</script>
<template>
    <div id="code_editor"></div>
</template>

<style lang="sass" scoped>

#code_editor
    width: calc(100% - 2px)
    height: calc(100% - 2px)
    border: 1px solid black
    overflow: hidden
    position: relative
    flex-grow: 1
</style>
