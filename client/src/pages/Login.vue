<template>
    <main class="form-login text-center">
        <form @submit.prevent="login">
            <br>
            <h1 class="h3 mb-3 fw-normal text-primary"><b>SIGN IN</b></h1>
            <input v-model="email" class="form-control" placeholder="Email" required>
            <br>
            <input v-model="password" class="form-control" placeholder="Password" required>
            <br>
            <button class="w-50 btn btn-lg btn-primary" type="submit">LOG IN</button>
            <br>
            <br>
            <div class="mb-2">
                <router-link to="/forgot">Forgot Password</router-link>
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

.form-login input[type="email"] {
    margin-bottom: -1px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
}

.form-login input[type="password"] {
    margin-bottom: 10px;
    border-top-left-radius: 0;
    border-top-right-radius: 0;
}
</style>
