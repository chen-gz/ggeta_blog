import {ref, watch} from "vue";
import {useRouter} from "vue-router";
// const crc32 = require("crc32");
// import "crc32/lib/crc32.js"
// import CRC32 from "crc-32/crc32.js"
// import router from "@/router";

// const blogBackendUrl = "http://localhost:2009"

const blogBackendUrl = "https://ggeta.com"

export let is_logined = ref(false)

export interface V4PostData {
    id: number
    title: string
    author: string
    author_email: string
    url: string
    is_draft: boolean
    is_deleted: boolean
    content: string
    content_rendered: string
    summary: string
    tags: string
    category: string
    cover_image: string
    created_at: Date
    updated_at: Date
    view_groups: string[]
    edit_groups: string[]
}

export interface GetPostResponseV3 {
    status: string
    message: string
    post: V4PostData
    html: string
}

export async function getPostV4(url: string, rendered: boolean): Promise<GetPostResponseV3> {
    const request = {
        url: url,
        rendered: rendered
    }
    return await fetch(`${blogBackendUrl}/api/post/get`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
        body: JSON.stringify(request),
    }).then(response => response.json())
}


export interface UpdatePostResponseV4 {
    status: string
    message: string
    post: V4PostData
    html: string
}

export async function updatePostV4(request: V4PostData): Promise<UpdatePostResponseV4> {
    console.log(localStorage.getItem("token"))
    return await fetch(`${blogBackendUrl}/api/post/update`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
        body: JSON.stringify(request),
    }).then(response => response.json())
}

export async function deletePost(post: V4PostData) {
    post.is_deleted = true
    return updatePostV4(post);//.then(response => {
        // if (response.status == "success") {
            // showSuccess("Delete success")
            // router.push({path: '/posts'})
        // } else {
            // showError("Delete failed")
        // }
    // })
}
//
export async function savePost(post: V4PostData) {
    return updatePostV4(post)
    // .then(
    //     (response) => {
    //         if (response.status == "success") {
    //             // showSuccess("Post saved")
    //             alert("Post saved")
    //             // router.push({path: '/posts/edit/' + response.post.url})
    //         } else {
    //             // showError("Failed to save post")
    //             alert("failed to save post, login is required");
    //         }
    //     })
}

export interface NewPostResponseV4 {
    status: string
    message: string
    url: string
}

export async function newPostV4(): Promise<NewPostResponseV4> {
    return await fetch(`${blogBackendUrl}/api/post/new`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
    }).then(response => response.json())
}

export interface GetDistinctResponse {
    status: string
    message: string
    values: string[]
    length: number
}

export interface GetDistinctRequest {
    field: string
}

export async function getDistinct(col: string): Promise<GetDistinctResponse> {
    return await fetch(`${blogBackendUrl}/api/post/get_distinct`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
        body: JSON.stringify({field: col}),
    }).then(response => response.json())
}

export interface SearchPostsRequestV4 {
    author: string
    title: string
    limit: {
        start: number
        size: number
    }
    sort: string
    rendered: boolean
    counts_only: boolean
    content: string
    tags: string
    categories: string
    private_level: number
    is_draft: boolean
    is_deleted: boolean
}

export interface SearchPostsResponseV4 {
    status: string
    message: string
    posts: V4PostData[]
    number_of_posts: number
}

export async function searchPostsV4(request: SearchPostsRequestV4): Promise<SearchPostsResponseV4> {
    return fetch(`${blogBackendUrl}/api/v4/search_posts`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
        body: JSON.stringify(request),
    }).then(response => response.json())
}

export function logined() {
    is_logined.value = localStorage.getItem("token") != null
    return is_logined.value
    // return localStorage.getItem("token") != null
}

export interface LoginResponse {
    email: string
    message: string
    name: string
    token: string
}

export async function loginV4(email: string, password: string) : Promise<boolean> {
    return fetch(`${blogBackendUrl}/api/login`, {
        method: "POST",
        body: JSON.stringify({
            email: email,
            password: password
        }),
    }).then(response => {
            if (!response.ok) {
                // showError("Login failed");
                return false;
            }
            if (response.status == 200) {
                // convert response to LoginResponse
                response.json().then(response => {
                    const lg_res: LoginResponse = response as LoginResponse;
                    localStorage.setItem("token", lg_res.token);
                    localStorage.setItem("userName", lg_res.name);
                    localStorage.setItem("userEmail", lg_res.email);
                });
                // showSuccess("Login success");
                return true;
            } else {
                // showError("Login failed");
                return false;
            }

        }
    );
}

export async function verifyToken(): Promise<LoginResponse> {
    return await fetch(`${blogBackendUrl}/api/verify_token`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
    }).then(response => {
            if (response.ok) {
                return response.json()
            } else {
                logout()
            }
        }
    )
}


export function logout() {
    localStorage.removeItem("token")
    localStorage.removeItem("userName")
    localStorage.removeItem("userEmail")
}

// export const alert = ref({
//     show: false,
//     type: '',
//     message: '',
//     color: '',
// })

// export function showError(msg: string) {
//     alert.value = {
//         show: true,
//         type: 'error',
//         message: msg,
//         color: 'error'
//     }
// }
//
// export function showSuccess(msg: string) {
//     alert.value = {
//         show: true,
//         type: 'info',
//         message: msg,
//         color: 'success'
//     }
// }

export function formatDate(date: Date) {
    date = new Date(date)
    const day = date.toLocaleString("en-US", {day: '2-digit'})
    const month = date.toLocaleString("en-US", {month: 'short'})
    return day + ' ' + month + ' ' + date.getFullYear();
}


// `retrieveNewURL` accepts the name of the current file and invokes the `/presignedUrl` endpoint to
// generate a pre-signed URL for use in uploading that file:
export interface GetPresignedUrlRequest {
    file_name: string
    post_id: number
    hash_crc32: string

}

export interface GetPresignedUrlResponse {
    presigned_url: string
    message: string
    filename: string
    file_url: string
}


export async function UploadFile(file: File, post_id: number) {
    console.log(file)
    const hash = await file.arrayBuffer().then(async (buffer) => {
        console.log(buffer)
        const uint8array = new Uint8Array(buffer);
        let digits = await window.crypto.subtle.digest("SHA-256", uint8array);
        // get the first 4 bytes of the hash
        let hash = new Uint32Array(digits)[0];
        return hash.toString(16);
    });
    const request: GetPresignedUrlRequest = {
        file_name: file.name,
        post_id: post_id,
        hash_crc32: hash.toString()
    }
    fetch(`${blogBackendUrl}/api/v1/blog_file/get_presigned_url`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`,
        },
        body:
            JSON.stringify(request),
    }).then((response) => {
        if (!response.ok) {
            console.error(response);
            return;
        }
        response.json().then((response) => {
            response as GetPresignedUrlResponse
            console.log(response)
            navigator.clipboard.writeText(response.file_url).then(() => {
                // showSuccess("Copied to clipboard")
            });
            uploadFileToPresignURL(file, response.presigned_url).then(response => {
                if (!response.ok) {
                    console.error(response);
                    return;
                }
            })
        })
    }).catch((e) => {
        console.error(e);
    });
}

function uploadFileToPresignURL(file: File, presignedURL: string): Promise<Response> {
    return fetch(presignedURL, {
        method: 'PUT',
        body: file
    })
}

export interface RenderResponse {
    format: string
    message: string
    rendered: string
}
export async function getRenderedContent(content: string): Promise<string> {
    return await fetch(`${blogBackendUrl}/api/v1/post/render`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`,
        },
        body: JSON.stringify({
            content: content,
            format: "markdown"
        }),
    }).then((response) => {
        if (!response.ok) {
            console.error(response);
            return response.json()
        }
        return response.json().then((response) => {
            response as RenderResponse
            return response.rendered
        })
    })
}

// type UploadFileResponse struct {
//     Filenames []string `json:"filenames"`
//     FileUrl   []string `json:"file_url"` // the file url with be updated by the server
//     Message   string   `json:"message"`
// }
export interface GetFileListResponse  {
    filenames: string[]
    file_url: string[]
    message: string
}

// r.GET("/api/blog_file/v1/get_file_lists/:id", func(c *gin.Context) {
export async function GetFileList(post_id: number): Promise<string[]> {
    console.log(`${blogBackendUrl}/api/v1/blog_file/get_file_lists/${post_id}`)
    return await fetch(`${blogBackendUrl}/api/v1/blog_file/get_file_lists/${post_id}`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`,
        },
    }).then((response) => {
        if (!response.ok) {
            console.error(response);
            // return response.json()
        }
        console.log(response)
        return response.json().then((response) => {
            response as GetFileListResponse
            return response.filenames
        })
    })
}



// GET localhost:2009/api/shows/v1/get_list

// type response struct {
//     Shows []string `json:"shows"`
// }
export interface GetShowListRequest {
    show_name: string
}
export interface GetShowListResponse {
    shows: string[]
}
export async function ShowGetList(show: string): Promise<GetShowListResponse> {
    const request = {show_name: show}
    return await fetch(`${blogBackendUrl}/api/v1/shows/get_list`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
        body: JSON.stringify(request),
    }).then(response => response.json())

}
export interface GetShowItemsResponse {
    show_url: string
}
export interface GetShowItemsRequest {
    path: string
}

export async function ShowGetItems(path: string): Promise<GetShowItemsResponse> {
    // POST https://ggeta.com/api/shows/v1/get_presigned_url

    const request = { path: path}
    return await fetch(`${blogBackendUrl}/api/v1/shows/get_presigned_url`, {
        method: "POST",
        mode: "cors",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
        body: JSON.stringify(request),
    }).then(response => response.json())
}

export async function GetPresignedUrl(bucket:string, path: string): Promise<GetPresignedUrlResponse> {
    // POST https://ggeta.com/api/shows/v1/get_presigned_url
    const request = { bucket: bucket, path: path}
    return await fetch(`${blogBackendUrl}/api/v1/get_presigned_url_new`, {
        method: "POST",
        mode: "cors",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
        body: JSON.stringify(request),
    }).then(response => response.json())
}