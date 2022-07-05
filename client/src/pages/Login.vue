<template>
    <main class="form-login text-center">
        <form @submit.prevent="login">
            <br>
            <h1 class="h1 mb-3 fw-normal text-primary"><b>ログイン</b></h1>
            <hr>
            <br>
            <div class="input-container">
                <label for="email" class="text-secondary"><b>Eメール</b></label>
                <input id="email" v-model="email" class="form-control" placeholder="Email" required>
            </div>
            <br>
            <br>
            <div class="input-container">
                <label for="password" class="text-secondary"><b>パスワード</b></label>
                <input id="password" v-model="password" type="password" class="form-control" placeholder="Password" required>
            </div>
            <br>
            <br>
            <br>
            <button class="w-50 btn btn-lg btn-primary" type="submit">LOG IN</button>
            <br>
            <br>
            <div class="mb-2">
                <router-link to="/forgot">パスワードを忘れた</router-link>
            </div>
        </form>
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
</style>
