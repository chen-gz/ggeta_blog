<script setup>
import {ref} from 'vue'
import {useRouter} from 'vue-router'

const username = ref('')
const password = ref('')

const router = useRouter()

import {loginV4} from "/apiv4";

// const login = () => {
function login() {
    console.log('login')
    loginV4(username.value, password.value).then((response) => {
        if (response) {
            router.push('/')
            console.log('login success')
        } else {
            // alert('Invalid credentials')
            console.log('login failed')
        }
    })
}

</script>
<template>
    <div id="login">
        <form @submit.prevent="login">
            <input type="text" placeholder="Username" v-model="username">
            <input type="password" placeholder="Password" v-model="password">
            <button type="submit">Login</button>
        </form>
    </div>
</template>


<style scoped lang="sass">
#login
    display: flex
    justify-content: center
    padding-bottom: 90px
    align-items: center
    height: calc(100% - 90px)

    form
        display: flex
        flex-direction: column
        width: 260px

        input
            margin: 10px
            padding: 10px
            border: 1px solid #ccc
            border-radius: 5px
            //width: 200px

        button
            margin: 10px
            padding: 10px
            border: 1px solid #ccc
            border-radius: 5px
            //width: 200px
            background-color: #007bff
            color: white
            cursor: pointer

            &:hover
                background-color: #0056b3


</style>