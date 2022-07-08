<template>
    <main class="form-profile text-center">
        <br>
        <h1 class="h1 mb-3 fw-normal"><b>Home</b></h1>
        <hr>
        <br>
        <div class="profile-container">
            <h2 class="h6"><b>個人情報</b></h2>
            <div class="container bg-gray rounded">
                <div class="row">
                    <p></p>
                    <p class="text-secondary col-4">Eメール</p>
                    <p class="text-dark col-8">{{ email }}</p>
                </div>
                <div class="row">
                    <p class="text-secondary col-4">生年月日</p>
                    <p class="text-dark col-8">{{ birth }}</p>
                </div>
                <div class="row">
                    <p class="text-secondary col-4">性別</p>
                    <p class="text-dark col-8">{{ gender }}</p>
                </div>
                <div class="row">
                    <p class="text-secondary col-4">住所</p>
                    <p class="text-dark col-8">{{ address }}</p>
                </div>
            </div>
        </div>
    </main>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from "vuex"
import moment from "moment"

export default {
    name: "profile",
    setup() {
        const email = ref('none')
        const birth = ref('none')
        const gender = ref('none')
        const address = ref('none')
        const store = useStore()

        onMounted(async () => {
            try {
                const { data } = await axios.get('profile')
                email.value = `${data.email}`
                birth.value = `${data.birthday}`
                birth.value = moment(birth.value).format("YYYY年MM月DD日")
                gender.value = `${data.gender}`
                address.value = `${data.address}`
                await store.dispatch('setAuth', true)
            } catch (e) {
                await store.dispatch('setAuth', false)
            }
        })

        return {
            email,
            birth,
            gender,
            address,
        }
    }
}
</script>

<style>
.form-profile {
    width: 100%;
    max-width: 500px;
    padding: 15px;
    margin: auto;
}

.form-profile .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
    font-size: 16px;
}

.form-profile .form-control:focus {
    z-index: 2;
}

/* .profile-container div {       */
/*     background-color: #F0F0F5; */
/* }                              */

.container p {
    display: inline;
    text-align: left;
}

.profile-container h2 {
    text-align: left;
}

.profile-rect svg {
    margin-left: 5px;
    margin-right: 5px;
}

.rect-left {
    float: left;
}

.bg-gray {
    background-color: #F5F5FF;
}
</style>

