<template>
    <main class="form-register text-center">
        <form @submit.prevent="submit">
            <br>
            <h1 class="h3 mb-3 fw-normal text-primary"><b>SIGN UP</b></h1>
            <input v-model="name" class="form-control" placeholder="Name" required>
            <br>
            <input v-model="email" class="form-control" placeholder="Email" required>
            <br>
            <input v-model="password" class="form-control" placeholder="Password" required>
            <input v-model="passwordConfirm" class="form-control" placeholder="Password Confirm" required>
            <br>
            <button class="w-50 btn btn-lg btn-primary" type="submit">REGISTER</button>
        </form>
    </main>
</template>

<script>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

export default {
    name: "Register",
    setup() {
        const name = ref('');
        const email = ref('');
        const password = ref('');
        const passwordConfirm = ref('');
        const router = useRouter();

        const submit = async () => {
            await axios.post('register', {
                name: name.value,
                email: email.value,
                password: password.value,
                password_confirm: passwordConfirm.value,
            })

            await router.push('/login')
        }

        return {
            name,
            email,
            password,
            passwordConfirm,
            submit
        }
    }
}
</script>

<style>
.form-register {
    width: 100%;
    max-width: 330px;
    padding: 15px;
    margin: auto;
}

.form-register .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
    font-size: 16px;
}

.form-register .form-control:focus {
    z-index: 2;
}
</style>
