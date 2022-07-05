<template>
    <main class="form-register text-center">
        <form @submit.prevent="submit">
            <br>
            <h1 class="h1 mb-3 fw-normal text-success"><b>アカウントを作成</b></h1>
            <hr>
            <br>
            <div class="input-container">
                <label for="name" class="text-secondary"><b>ユーザ名</b></label>
                <input id="name" v-model="name" class="form-control" placeholder="Name" required>
            </div>
            <br>
            <br>
            <div class="input-container">
                <label for="email" class="text-secondary"><b>Eメール</b></label>
                <input id="email" v-model="email" class="form-control" placeholder="Email" required>
            </div>
            <br>
            <br>
            <div class="input-container">
                <label for="password" class="text-secondary"><b>パスワード</b></label>
                <input id="password" v-model="password" type="password" class="form-control" placeholder="Password"
                    required>
            </div>
            <br>
            <br>
            <div class="input-container">
                <label for="passwordConfirm" class="text-secondary"><b>パスワード(確認用)</b></label>
                <input id="passwordConfirm" v-model="passwordConfirm" type="password" class="form-control"
                    placeholder="Password Confirm" required>
            </div>
            <br>
            <br>
            <br>
            <button class="w-50 btn btn-lg btn-success" type="submit">REGISTER</button>
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

.input-container {
    display: inline-block;
    position: relative;
    right: -75px
}

.input-container label {
    position: absolute;
    left: -150px;
    top: 50%;
    transform: translateY(-50%);
}
</style>
