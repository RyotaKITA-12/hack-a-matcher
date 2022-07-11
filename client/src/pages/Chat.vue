<template>
    <main class="form-profile text-center">
        <br>
        <h1 class="h1 mb-3 fw-normal"><b>チャット</b></h1>
        <hr>
        <br>
        <div class="overflow-scroll" style="height:750px;">
            <div class="profile-container">
                <h2 class="h6"><b>ryota-k</b></h2>
                <div class="container bg-gray rounded" style="text-align: left;">
                    <br>
                    <p class="text-dark">8月のハッカソンのグループ、応募ありがとうございます！</p>
                    <br>
                    <br>
                </div>
                <br>
                <h2 class="h6" style="text-align: right"><b>山田太郎</b></h2>
                <div class="container bg-primary rounded" style="text-align: left;">
                    <br>
                    <p class="text-white">よろしくお願いします！</p>
                    <br>
                    <br>
                </div>
                <br>
                <h2 class="h6"><b>ryota-k</b></h2>
                <div class="container bg-gray rounded" style="text-align: left;">
                    <br>
                    <p class="text-dark">こちらこそ！</p>
                    <br>
                    <br>
                </div>
                <br>
                <h2 class="h6" style="text-align: right"><b>山田太郎</b></h2>
                <div class="container bg-primary rounded" style="text-align: left;">
                    <br>
                    <p class="text-white">はい！</p>
                    <br>
                    <br>
                </div>
            </div>
            <form action="" class="form">
                <textarea></textarea>
                <button type="submit" class="btn btn-primary send-button">送信</button>
            </form>
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
        const portfolio = ref('none')
        const twitter = ref('none')
        const github = ref('none')
        const store = useStore()

        onMounted(async () => {
            try {
                const p = await axios.get('profile')
                const l = await axios.get('link')
                email.value = `${p.data.email}`
                birth.value = `${p.data.birthday}`
                birth.value = moment(birth.value).format("YYYY年MM月DD日")
                gender.value = `${p.data.gender}`
                address.value = `${p.data.address}`
                portfolio.value = `${l.data.portfolio}`
                twitter.value = `${l.data.twitter}`
                github.value = `${l.data.github}`
                await store.dispatch('setAuth', true)
            } catch (e) {
                await store.dispatch('setAuth', false)
            }
        })

        var showEditProfile = ref(false)
        const openEditProfile = () => {
            showEditProfile.value = true
        }
        const closeEditProfile = () => {
            showEditProfile.value = false
        }

        const submit = async () => {
            const { data } = await axios.get('user')
            console.log(data)
            await axios.post('register/profile', {
                user_id: `${data.ID}`,
                email: `${data.email}`,
                birthday: birth.value,
                gender: gender.value,
                address: address.value,
            })
            showEditProfile.value = false
        }

        return {
            email,
            birth,
            gender,
            address,
            portfolio,
            twitter,
            github,
            submit,
            showEditProfile,
            openEditProfile,
            closeEditProfile
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
    margin: 5px;
}

.rect-left {
    float: left;
}

.bg-gray {
    background-color: #F5F5FF;
}

a:hover {
    cursor: pointer;
}

.send-button {
    margin: 10px;
    height: 4em;
}

.form {
    position: fixed;
    display: flex;
    justify-content: center;
    align-items: center;
    bottom: 30px;
    height: 80px;
    width: 450px;
    background: #f5f5f5;
}

.form textarea {
    border: 1px solid #ccc;
    border-radius: 2px;
    height: 4em;
    width: calc(100% - 6em);
    resize: none;
}
</style>
