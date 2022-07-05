<template>
    <main class="form-login text-center">
        <br>
        <h1 class="h1 mb-3 fw-normal text-secondary"><b>プロフィール</b></h1>
        <hr>
        <br>
        <div class="profile-container">
            <h2 class="h6 text-dark"><b>基本情報</b></h2>
            <div class="container bg-light rounded">
                <div class="row">
                    <p></p>
                    <p class="text-secondary col-4">Eメール</p>
                    <p class="text-dark col-8">test@example.com</p>
                </div>
                <div class="row">
                    <p class="text-secondary col-4">生年月日</p>
                    <p class="text-dark col-8">2001年12月12日</p>
                </div>
                <div class="row">
                    <p class="text-secondary col-4">性別</p>
                    <p class="text-dark col-8">男性</p>
                </div>
            </div>
        </div>
        <br>
        <div class="profile-container">
            <h2 class="h6 text-dark"><b>属性</b></h2>
            <div class="container  bg-light rounded">
                <div class="row">
                    <p></p>
                    <p class="text-secondary col-4">役割</p>
                </div>
                <div class="profile-rect">
                    <svg class="rounded-pill" width="100" height="35">
                        <rect width="100%" height="100%" fill="#CCC"></rect>
                        <text x="50%" y="50%" fill="#000" text-anchor="middle" dominant-baseline="central">エンジニア</text>
                    </svg>
                </div>
                <br>
                <br>
                <div class="row">
                    <p class="text-secondary col-4">生年月日</p>
                    <p class="text-dark col-8">2001年12月12日</p>
                </div>
                <div class="row">
                    <p class="text-secondary col-4">性別</p>
                    <p class="text-dark col-8">男性</p>
                </div>
            </div>
        </div>
    </main>
</template>

<script>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

export default {
    name: "Login",
    setup() {
        const email = ref('')
        const password = ref('')
        const router = useRouter()

        const login = async () => {
            await axios.post('login', {
                email: email.value,
                password: password.value
            })

            await router.push('/')
        }

        return {
            email,
            password,
            login
        }
    }
}
</script>

<style>
.form-login {
    width: 100%;
    max-width: 330px;
    padding: 15px;
    margin: auto;
}

.form-login .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
    font-size: 16px;
}

.form-login .form-control:focus {
    z-index: 2;
}

/* .profile-container div {             */
/*     /* background-color: #F0F0F0; */
/* }                                    */

.container p {
    display: inline;
    text-align: left;
}

.profile-container h2 {
    text-align: left;
}

/* .profile-left svg {     */
/*     margin-right: auto; */
/* }                       */

.profile-rect div {
    margin-right: auto;
    /* text-align: left; */
}
</style>
