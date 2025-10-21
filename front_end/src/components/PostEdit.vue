<script lang="ts" setup>
import { useRouter } from "vue-router";
import { getPostV4, savePost, UploadFile, V4PostData } from "/apiv4";
import { onMounted, ref } from "vue";
import { useEditor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Placeholder from '@tiptap/extension-placeholder'

let router = useRouter();
let url = router.currentRoute.value.params.id as string;
console.log(url);

let post = ref({} as V4PostData);
let editor_shows = ref("content");
const editorWrapper = ref<HTMLDivElement | null>(null);

const editor = useEditor({
  content: '',
  extensions: [
    StarterKit,
    Placeholder.configure({
      placeholder: 'Write something …',
    }),
  ],
})

getPostV4(url, true).then((response) => {
    console.log(response);
    post.value = response.post;
    if (editor.value) {
      editor.value.commands.setContent(post.value.content);
    }
});

document.addEventListener("keydown", function (e) {
    // control + 'S' to save or (command + 'S' on mac)
    if (
        (window.navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey) &&
        e.key === "s"
    ) {
        e.preventDefault();
        console.log("ctrl+s");
        if (!editor.value) return;

        if (editor_shows.value === 'meta') {
            try {
                post.value = JSON.parse(editor.value.getText());
            } catch (error) {
                alert("Invalid JSON, cannot save.");
                return;
            }
        } else {
            post.value.content = editor.value.getHTML();
        }

        savePost(post.value).then(
          (response) => {
              if (response.status == "success") {
                  alert("Post saved")
                  console.log("Post saved")
              } else {
                  alert("failed to save post, login is required");
                  router.push("/login");
              }
          })

      // push to new url
        router.push("/post_edit/" + post.value.url);
    }
    // F2 to toggle meta view
    if (e.key === "F2") {
        e.preventDefault();
        console.log("F2");
        if (!editor.value) return;

        if (editor_shows.value === "content") {
            editor_shows.value = "meta";
            // Update content from editor before switching
            post.value.content = editor.value.getHTML();
            editor.value.commands.setContent(JSON.stringify(post.value, null, 4));
            // Tiptap doesn't have language modes, so we just show the text.
        } else {
            editor_shows.value = "content";
            try {
                const updatedPost = JSON.parse(editor.value.getText());
                post.value = updatedPost;
                editor.value.commands.setContent(post.value.content);
            } catch (error) {
                alert("Invalid JSON, cannot switch back to content view.");
            }
        }
    }
});

onMounted(() => {
    if (editorWrapper.value) {
      editorWrapper.value.addEventListener("drop", function (e: DragEvent) {
          e.preventDefault();
          e.stopPropagation();
          if (e.dataTransfer && e.dataTransfer.files.length > 0) {
            let file = e.dataTransfer.files[0];
            console.log(file);
            UploadFile(file, post.value.id);
          }
      });
    }
});

</script>
<template>
    <div ref="editorWrapper">
      <editor-content :editor="editor" />
    </div>
</template>

<style>
.ProseMirror {
  height: 75vh;
  overflow-y: scroll;
  border: 1px solid #ccc;
  padding: 10px;
}
</style>
