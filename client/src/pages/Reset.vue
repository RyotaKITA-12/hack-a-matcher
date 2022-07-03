<template>
    <main class="form-reset text-center">
        <form @submit.prevent="submit">
            <br>
            <h1 class="h3 mb-3 fw-normal">Please new password.</h1>
            <br>
            <input v-model="password" type="password" class="form-control" placeholder="Password" required>
            <input v-model="passwordConfirm" type="password" class="form-control" placeholder="Password Confirm"
                required>
            <br>
            <button class="w-50 btn btn-lg btn-primary" type="submit">RESET</button>
        </form>
    </main>
</template>

<script>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'

export default {
    name: "Reset",
    setup() {
        const password = ref('')
        const passwordConfirm = ref('')
        const route = useRoute()
        const router = useRouter()

        const submit = async () => {
            await axios.post('password/reset', {
                token: route.params.token,
                password: password.value,
                password_confirm: passwordConfirm.value
            })

            await router.push('/login')
        }

        return {
            password,
            passwordConfirm,
            submit
        }
    }
}
</script>

<style>
.form-reset {
    width: 100%;
    max-width: 330px;
    padding: 15px;
    margin: auto;
}

.form-reset .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
    font-size: 16px;
}

.form-reset .form-control:focus {
    z-index: 2;
}

.form-reset input[type="email"] {
    border-top-left-radius: 0;
    border-top-right-radius: 0;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
}

.form-reset input[type="password"] {
    border-top-left-radius: 0;
    border-top-right-radius: 0;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
}
</style>
