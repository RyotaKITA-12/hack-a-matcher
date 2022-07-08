<template>
    <main class="form-profile text-center">
        <br>
        <h1 class="h1 mb-3 fw-normal"><b>プロフィール</b></h1>
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
        <br>
        <div class="profile-container">
            <h2 class="h6"><b>リンク</b></h2>
            <div class="container bg-gray rounded">
                <div class="row">
                    <p></p>
                    <p class="text-secondary col-4">Portfolio</p>
                    <p class="text-dark col-8"><a href="https://kt12.jp">kt12.jp</a></p>
                </div>
                <div class="row">
                    <p></p>
                    <p class="text-secondary col-4">Twitter</p>
                    <p class="text-dark col-8"><a href="https://twitter.com/dhs_1212">@dhs_1212</a></p>
                </div>
                <div class="row">
                    <p></p>
                    <p class="text-secondary col-4">Github</p>
                    <p class="text-dark col-8"><a href="https://github.com/RyotaKITA-12">RyotaKITA-12</a></p>
                </div>
            </div>
        </div>
        <br>
        <div class="profile-container">
            <h2 class="h6"><b>スキル</b></h2>
            <div class="container  bg-gray rounded">
                <div class="row">
                    <p></p>
                    <p class="text-secondary col-4">役割</p>
                </div>
                <div class="profile-rect">
                    <svg class="rounded-pill rect-left" width="100" height="35">
                        <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                        <text x="50%" y="50%" fill="#334" text-anchor="middle" dominant-baseline="central">エンジニア</text>
                    </svg>
                </div>
                <br>
                <br>
                <hr>
                <br>
                <div class="row">
                    <p class="text-secondary col-4">使用経験</p>
                    <p class="text-dark col-8">授業</p>
                </div>
                <div class="row">
                    <p class="text-secondary col-4">言語</p>
                    <div class="profile-rect">
                        <svg class="rounded-pill rect-left" width="100" height="35">
                            <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                            <text x="50%" y="50%" fill="#334" text-anchor="middle"
                                dominant-baseline="central">Pyhon</text>
                        </svg>
                    </div>
                </div>
                <br>
                <hr>
                <br>
                <div class="row">
                    <p class="text-secondary col-4">使用経験</p>
                    <p class="text-dark col-8">趣味</p>
                </div>
                <div class="row">
                    <p class="text-secondary col-4">言語</p>
                    <div class="profile-rect">
                        <svg class="rounded-pill rect-left" width="130" height="35">
                            <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                            <text x="50%" y="50%" fill="#334" text-anchor="middle"
                                dominant-baseline="central">Python</text>
                        </svg>
                        <svg class="rounded-pill rect-left" width="130" height="35">
                            <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                            <text x="50%" y="50%" fill="#334" text-anchor="middle" dominant-baseline="central">Go</text>
                        </svg>
                        <svg class="rounded-pill rect-left" width="130" height="35">
                            <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                            <text x="50%" y="50%" fill="#334" text-anchor="middle"
                                dominant-baseline="central">JavaScript</text>
                        </svg>
                    </div>
                </div>
                <br>
                <br>
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
