import {ref} from "vue";

const blogBackendUrl = "http://localhost:2009"

// const blogBackendUrl = "https://blog.ggeta.com"


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
    return await fetch(`${blogBackendUrl}/api/v4/get_post`, {
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
    return await fetch(`${blogBackendUrl}/api/v4/update_post`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
        body: JSON.stringify(request),
    }).then(response => response.json())
}

export interface NewPostResponseV4 {
    status: string
    message: string
    url: string
}

export async function newPostV4(): Promise<NewPostResponseV4> {
    return await fetch(`${blogBackendUrl}/api/v4/new_post`, {
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
    return await fetch(`${blogBackendUrl}/api/v4/get_distinct`, {
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
    return localStorage.getItem("token") != null
}

// export async function searchPostsV4(request: SearchPostsRequestV4){
//
// }

export interface LoginResponse {
    email: string
    message: string
    name: string
    token: string
}

export async function loginV4(email: string, password: string) {
    fetch(`${blogBackendUrl}/api/v4/login`, {
        method: "POST",
        body: JSON.stringify({
            email: email,
            password: password
        }),
    }).then(response => {
            if (!response.ok) {
                showError("Login failed")
                return;
            }
            if (response.status == 200) {
                // convert response to LoginResponse
                response.json().then(response => {
                    const lg_res: LoginResponse = response as LoginResponse
                    localStorage.setItem("token", lg_res.token);
                    localStorage.setItem("userName", lg_res.name);
                    localStorage.setItem("userEmail", lg_res.email);
                })
                showSuccess("Login success");
            } else {
                showError("Login failed")
            }
        }
    );
}


// export async function loginV3(){
//     window.location.href= "https://gitea.ggeta.com/login/oauth/authorize?client_id=4093feeb-ff9b-4103-a091-db2381588ce9&redirect_uri=https://blog.ggeta.com&response_type=code&state=STATE"
// }

export async function verifyToken(): Promise<LoginResponse> {
    return await fetch(`${blogBackendUrl}/api/v4/verify_token`, {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
        },
    }).then(response => response.json())
}


export function logout() {
    localStorage.removeItem("token")
    localStorage.removeItem("userName")
    localStorage.removeItem("userEmail")
}

export const alert = ref({
    show: false,
    type: '',
    message: '',
    color: '',
})

export function showError(msg: string) {
    alert.value = {
        show: true,
        type: 'error',
        message: msg,
        color: 'error'
    }
}

export function showSuccess(msg: string) {
    alert.value = {
        show: true,
        type: 'info',
        message: msg,
        color: 'success'
    }
}

export function formatDate(date: Date) {
    date = new Date(date)
    const day = date.toLocaleString("en-US", {day: '2-digit'})
    const month = date.toLocaleString("en-US", {month: 'short'})
    return day + ' ' + month + ' ' + date.getFullYear();
}


function upload() {
    // Get selected files from the input element.
    var files = document.querySelector("#selector").files;
    for (var i = 0; i < files.length; i++) {
        var file = files[i];
        // Retrieve a URL from our server.
        retrieveNewURL(file, (file, url) => {
            // Upload the file to the server.
            uploadFileToPresignURL(file, url);
        });
    }
}

// return await fetch(`${blogBackendUrl}/api/v4/get_post`, {
//   method: "POST",
//   headers: {
//     "Authorization": `Bearer ${localStorage.getItem("token") || ""}`
//   },
//   body: JSON.stringify(request),
// }).then(response => response.json())

// `retrieveNewURL` accepts the name of the current file and invokes the `/presignedUrl` endpoint to
// generate a pre-signed URL for use in uploading that file:
export interface GetPresignedUrlRequest {
    file_name: string
    hash_crc32: string
}

export interface GetPresignedUrlResponse {
    presigned_url: string
    message: string
    filename: string
}

// import "crc32/lib/crc32.js"

// const crc32 = require("crc32");
// import "crc32/lib/crc32.js"
import CRC32 from "crc-32/crc32.js"

export async function UploadFile(file: File) {
    console.log(file)
    const hash = await file.arrayBuffer().then((buffer) => {
        console.log(buffer)
        const uint8array = new Uint8Array(buffer);
        return (CRC32.buf(uint8array) >>> 0).toString(16);
    });
    const request: GetPresignedUrlRequest = {
        file_name: file.name,
        hash_crc32: hash.toString()
    }
    fetch(`${blogBackendUrl}/api/blog_file/v1/get_presigned_url`, {
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
            const publicUrl = `minio.ggeta.com/blog-public-data/${response.filename}`;
            navigator.clipboard.writeText(publicUrl).then(() => {
                showSuccess("Copied to clipboard")
            });
            uploadFileToPresignURL(file, response.presigned_url).then(response => {
                if (!response.ok) {
                    console.error(response);
                    return;
                }
                fetch(`${blogBackendUrl}/api/blog_file/v1/upload_finish`, {
                    method: "POST",
                    headers: {
                        "Authorization": `Bearer ${localStorage.getItem("token") || ""}`,
                    },
                    body:
                        JSON.stringify(request),
                });
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
