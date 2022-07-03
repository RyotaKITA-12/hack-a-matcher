<template>
    <main class="form-forgot text-center">
        <form @submit.prevent="forgot">
            <br>
            <h1 class="h3 mb-3 fw-normal">Please e-mail address.</h1>
            <br>
            <input v-model="email" type="email" class="form-control mb-3" placeholder="Email" required>
            <br>
            <button class="w-50 btn btn-lg btn-primary" type="submit">SEND</button>
            <br>
            <br>
            <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
                {{ notify.message }}
            </div>
        </form>
    </main>
</template>
<script>
import { reactive, ref } from 'vue'
import axios from 'axios'

export default {
    name: "Forgot",
    setup() {
        const email = ref('')
        const notify = reactive({
            cls: '',
            message: ''
        })
        const forgot = () => {
            try {
                axios.post('password/forgot', {
                    email: email.value
                })
                notify.cls = 'success'
                notify.message = 'Email was sent.'
            } catch (e) {
                notify.cls = 'danger'
                notify.message = 'Email does not exit.'
            }
        }

        return {
            email,
            notify,
            forgot
        }
    }
}
</script>

<style>
.form-forgot {
    width: 100%;
    max-width: 330px;
    padding: 15px;
    margin: auto;
}

.form-forgot .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
    font-size: 16px;
}

.form-forgot .form-control:focus {
    z-index: 2;
}

.form-forgot input[type="email"] {
    border-top-right-radius: 0;
    border-top-left-radius: 0;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
}
</style>
